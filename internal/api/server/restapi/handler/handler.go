package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/loads"
	"go.uber.org/zap"

	"github.com/Maksim646/tokens/internal/api/definition"
	"github.com/Maksim646/tokens/internal/api/server/restapi"
	"github.com/Maksim646/tokens/internal/api/server/restapi/api"
	"github.com/Maksim646/tokens/internal/model"
	"github.com/Maksim646/tokens/pkg/jsonwebtoken"
)

type contextKey string

const (
	contextKeyUserID    contextKey = "userID"
	contextKeyIP        contextKey = "ip"
	contextKeyRefreshID contextKey = "refreshID"
)

type Handler struct {
	userUsecase         model.IUserUsecase
	refreshTokenUsecase model.IRefreshTokenUsecase

	router             http.Handler
	jwtSigninKey       string
	accessTokenTTL     time.Duration
	refreshTokenTTL    time.Duration
	refreshTokenLength int
}

func New(
	userUsecase model.IUserUsecase,
	refreshTokenUsecase model.IRefreshTokenUsecase,

	version string,
	jwtSigninKey string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
	refreshTokenLength int,
) *Handler {

	swaggerContent := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", version)
	swagger, err := loads.Analyzed(json.RawMessage(swaggerContent), "")
	if err != nil {
		panic("failed to parse Swagger JSON: " + err.Error())
	}

	h := &Handler{
		userUsecase:         userUsecase,
		refreshTokenUsecase: refreshTokenUsecase,

		jwtSigninKey:       jwtSigninKey,
		accessTokenTTL:     accessTokenTTL,
		refreshTokenTTL:    refreshTokenTTL,
		refreshTokenLength: refreshTokenLength,
	}

	zap.L().Info("Initializing API routes")

	apiRouter := api.NewTokensBackendServiceAPI(swagger)
	apiRouter.UseSwaggerUI()
	apiRouter.Logger = zap.S().Infof

	// AUTH
	apiRouter.GetAuthTokenHandler = api.GetAuthTokenHandlerFunc(h.GetAccessRefreshTokensHandler)
	apiRouter.PostAuthRefreshHandler = api.PostAuthRefreshHandlerFunc(h.UpdateAccessRefreshTokensHandler)

	h.router = apiRouter.Serve(nil)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) ValidateAccessTokenOnly(bearerHeader string) (*definition.Principal, error) {
	bearerToken := extractToken(bearerHeader)
	userID, ip, refreshID, err := jsonwebtoken.ParseToken(bearerToken, h.jwtSigninKey)
	if err != nil {
		return nil, err
	}

	return &definition.Principal{
		ID:        userID,
		IP:        ip,
		RefreshID: refreshID,
	}, nil
}

func (h *Handler) ValidateExpiredOrValidToken(bearerHeader string) (*definition.Principal, error) {
	bearerToken := extractToken(bearerHeader)

	userID, ip, refreshID, err := jsonwebtoken.ParseToken(bearerToken, h.jwtSigninKey)
	if err == nil {
		return &definition.Principal{ID: userID, IP: ip, RefreshID: refreshID}, nil
	}

	userID, ip, refreshID, err = jsonwebtoken.ParseTokenWithOutClaims(bearerToken, h.jwtSigninKey)
	if err != nil {
		return nil, err
	}

	return &definition.Principal{ID: userID, IP: ip, RefreshID: refreshID}, nil
}

func extractToken(header string) string {
	return strings.TrimPrefix(header, "Bearer ")
}
