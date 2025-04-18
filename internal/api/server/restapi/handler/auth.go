package handler

import (
	"fmt"
	"net"
	"time"

	"github.com/Maksim646/tokens/internal/api/definition"
	"github.com/Maksim646/tokens/internal/api/server/restapi/api"
	"github.com/Maksim646/tokens/internal/model"
	"github.com/Maksim646/tokens/pkg/jsonwebtoken"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) GetAccessRefreshTokensHandler(req api.GetAuthTokenParams) middleware.Responder {
	zap.L().Info("get access and refresh tokens request")
	ctx := req.HTTPRequest.Context()

	ip, _, err := net.SplitHostPort(req.HTTPRequest.RemoteAddr)
	if err != nil {
		zap.L().Warn("failed to parse IP from RemoteAddr", zap.Error(err))
		ip = req.HTTPRequest.RemoteAddr
	}

	user, err := h.userUsecase.GetUserByID(ctx, req.UserID)
	if err != nil {
		zap.L().Error(fmt.Sprintf("user not found, user id: %s", req.UserID), zap.Error(err))
		return api.NewGetAuthTokenBadRequest().WithPayload(&definition.Error{
			Message: &model.UserNotFound,
		})
	}

	refreshToken, err := h.refreshTokenUsecase.GetRefreshTokenByUserID(ctx, user.ID)
	if err == nil {
		if refreshToken.ExpiredAt.UTC().After(time.Now().UTC()) {
			zap.L().Info("refresh token is valid")
			return api.NewGetAuthTokenBadRequest().WithPayload(&definition.Error{
				Message: &model.TokensAlreadyExists,
			})
		}

		if refreshToken.ExpiredAt.UTC().Before(time.Now().UTC()) {
			err = h.refreshTokenUsecase.DeleteRefreshTokenByID(ctx, refreshToken.ID)
			if err != nil {
				zap.L().Error("error delete refresh token", zap.Error(err))
			}
		}

		refreshToken, err = h.refreshTokenUsecase.GetRefreshTokenByUserID(ctx, user.ID)
		if err == nil && refreshToken.ExpiredAt.UTC().After(time.Now().UTC()) {
			zap.L().Info("refresh token is valid")
			return api.NewGetAuthTokenBadRequest().WithPayload(&definition.Error{
				Message: &model.TokensAlreadyExists,
			})
		}
	}

	refreshID := uuid.NewString()

	accessToken, err := jsonwebtoken.GenerateAccessToken(user.ID, ip, refreshID, h.jwtSigninKey, h.accessTokenTTL)
	if err != nil {
		zap.L().Error("error generate access token", zap.Error(err))
		return api.NewGetAuthTokenInternalServerError()
	}

	refreshTokenStr, err := jsonwebtoken.GenerateRefreshToken(user.ID, ip, refreshID, h.jwtSigninKey, h.refreshTokenLength)
	if err != nil {
		zap.L().Error("error generate refresh token", zap.Error(err))
		return api.NewGetAuthTokenInternalServerError()
	}

	refreshTokenHashed, err := bcrypt.GenerateFromPassword([]byte(refreshTokenStr), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("error hash refresh token", zap.Error(err))
		return api.NewGetAuthTokenInternalServerError()
	}

	refreshToken = model.RefreshToken{
		TokenHash: string(refreshTokenHashed),
		ID:        refreshID,
		UserID:    user.ID,
		ExpiredAt: time.Now().UTC().Add(h.refreshTokenTTL),
	}

	err = h.refreshTokenUsecase.CreateRefreshToken(ctx, refreshToken)
	if err != nil {
		zap.L().Error("error save refresh token", zap.Error(err))
		return api.NewGetAuthTokenInternalServerError()
	}

	return api.NewGetAuthTokenOK().WithPayload(&definition.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
	})
}

func (h *Handler) UpdateAccessRefreshTokensHandler(req api.PostAuthRefreshParams) middleware.Responder {
	zap.L().Info("update access and refresh tokens request")
	ctx := req.HTTPRequest.Context()

	ip, _, err := net.SplitHostPort(req.HTTPRequest.RemoteAddr)
	if err != nil {
		zap.L().Warn("failed to parse IP from RemoteAddr", zap.Error(err))
		ip = req.HTTPRequest.RemoteAddr
	}

	userID, ok := req.HTTPRequest.Context().Value(contextKeyUserID).(string)
	if !ok {
		zap.L().Info("error fetch user ID from context")
		return api.NewPostAuthRefreshBadRequest()
	}

	currentIP, ok := ctx.Value(contextKeyIP).(string)
	if !ok {
		zap.L().Info("error fetch user IP from context")
		return api.NewPostAuthRefreshBadRequest()
	}

	currentRefreshID, ok := ctx.Value(contextKeyRefreshID).(string)
	if !ok {
		zap.L().Info("error fetch refresh token ID from context")
		return api.NewPostAuthRefreshBadRequest()
	}

	user, err := h.userUsecase.GetUserByID(ctx, userID)
	if err != nil {
		zap.L().Error(fmt.Sprintf("user not found, user id: %s", userID), zap.Error(err))
		return api.NewPostAuthRefreshBadRequest().WithPayload(&definition.Error{
			Message: &model.UserNotFound,
		})
	}

	refreshToken, err := h.refreshTokenUsecase.GetRefreshTokenByUserID(ctx, user.ID)
	if err == nil && refreshToken.ExpiredAt.UTC().Before(time.Now().UTC()) {
		zap.L().Info("refresh token was expired")
		return api.NewPostAuthRefreshBadRequest().WithPayload(&definition.Error{
			Message: &model.RefreshTokenExpired,
		})
	}

	if currentIP != ip {
		zap.L().Warn("IP address mismatch detected")

		fmt.Println("old ip: ", currentIP, "new ip: ", ip)
		h.SendEmailWarning(user.Email, user.ID, currentIP, ip)
	}

	if currentRefreshID != refreshToken.ID {
		zap.L().Info("tokens have different id")
		return api.NewPostAuthRefreshConflict().WithPayload(&definition.Error{
			Message: &model.TokensDoNotMatch,
		})
	}

	accessToken, err := jsonwebtoken.GenerateAccessToken(user.ID, currentIP, currentRefreshID, h.jwtSigninKey, h.accessTokenTTL)
	if err != nil {
		zap.L().Error("error generate access token", zap.Error(err))
		return api.NewPostAuthRefreshInternalServerError()
	}

	return api.NewPostAuthRefreshOK().WithPayload(&definition.AccessTokenBody{
		AccessToken: accessToken,
	})
}
