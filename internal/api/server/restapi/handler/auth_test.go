package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/Maksim646/tokens/internal/api/client/operations"
	"github.com/Maksim646/tokens/internal/api/definition"
	"github.com/Maksim646/tokens/internal/api/server/restapi/api"
	"github.com/Maksim646/tokens/internal/model"
)

func (s *Suite) TestGetAccessRefreshTokens() {
	ctx := context.Background()

	user1ID := "41d862b0-81ee-4167-b19c-111e8126cc32"
	_, err := s.handler.userUsecase.CreateUserByEmail(ctx, user1ID, "user1@example.ru")
	s.Require().NoError(err)

	refreshTokenValid := model.RefreshToken{
		TokenHash: "111",
		ID:        "41d862b0-81ee-4167-b19c-111e8126cc34",
		UserID:    "41d862b0-81ee-4167-b19c-111e8126cc32",
		ExpiredAt: time.Now().UTC().Add(1 * time.Minute),
	}

	err = s.handler.refreshTokenUsecase.CreateRefreshToken(ctx, refreshTokenValid)
	s.Require().NoError(err)

	user2ID := "41d862b0-81ee-4167-b19c-111e8126cc33"
	_, err = s.handler.userUsecase.CreateUserByEmail(ctx, user2ID, "user2@example.ru")
	s.Require().NoError(err)

	refreshTokenNoValid := model.RefreshToken{
		TokenHash: "222",
		ID:        "41d862b0-81ee-4167-b19c-111e8126cc35",
		UserID:    "41d862b0-81ee-4167-b19c-111e8126cc33",
		ExpiredAt: time.Now().UTC().Add(1 * time.Microsecond),
	}

	err = s.handler.refreshTokenUsecase.CreateRefreshToken(ctx, refreshTokenNoValid)
	s.Require().NoError(err)
	time.Sleep(1 * time.Second)

	user3ID := "123e4567-e89b-12d3-a456-426614174000"
	_, err = s.handler.userUsecase.CreateUserByEmail(ctx, user3ID, "user3@example.ru")
	s.Require().NoError(err)

	testCases := []struct {
		name          string
		userID        string
		expectedError bool
	}{
		{
			name:          "RefreshTokenValid",
			userID:        user1ID,
			expectedError: true,
		},
		{
			name:          "RefreshTokenExpired",
			userID:        user2ID,
			expectedError: false,
		},
		{
			name:          "UserExistsWithoutToken",
			userID:        user3ID,
			expectedError: false,
		},
		{
			name:          "InvalidUUID",
			userID:        "not-a-uuid",
			expectedError: true,
		},
		{
			name:          "UserDoesNotExist",
			userID:        "11111111-1111-1111-1111-111111111111",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			params := operations.NewGetAuthTokenParams().
				WithUserID(tc.userID)

			resp, err := s.api.GetAuthToken(params)

			if tc.expectedError {
				s.Error(err)
				s.Nil(resp)
			} else {
				s.NoError(err)
				s.NotNil(resp)
				s.NotEmpty(resp.Payload.AccessToken)
				s.NotEmpty(resp.Payload.RefreshToken)
			}
		})
	}
}

func (s *Suite) TestUpdateAccessTokenByRefreshToken() {
	ctx := context.Background()

	user1ID := "41d862b0-81ee-4167-b19c-111e8126cc32"
	_, err := s.handler.userUsecase.CreateUserByEmail(ctx, user1ID, "user1@example.ru")
	s.Require().NoError(err)

	user1TokenHash := "111"
	user1RefreshTokenID := "41d862b0-81ee-4167-b19c-111e8126cc34"
	refreshTokenValid := model.RefreshToken{
		TokenHash: user1TokenHash,
		ID:        user1RefreshTokenID,
		UserID:    user1ID,
		ExpiredAt: time.Now().UTC().Add(1 * time.Minute),
	}
	err = s.handler.refreshTokenUsecase.CreateRefreshToken(ctx, refreshTokenValid)
	s.Require().NoError(err)

	user2ID := "41d862b0-81ee-4167-b19c-111e8126cc33"
	_, err = s.handler.userUsecase.CreateUserByEmail(ctx, user2ID, "user2@example.ru")
	s.Require().NoError(err)

	user2TokenHash := "222"
	refreshTokenExpired := model.RefreshToken{
		TokenHash: user2TokenHash,
		ID:        "41d862b0-81ee-4167-b19c-111e8126cc35",
		UserID:    user2ID,
		ExpiredAt: time.Now().UTC().Add(1 * time.Microsecond),
	}
	err = s.handler.refreshTokenUsecase.CreateRefreshToken(ctx, refreshTokenExpired)
	s.Require().NoError(err)
	time.Sleep(1 * time.Second)

	user3ID := "123e4567-e89b-12d3-a456-426614174000"
	_, err = s.handler.userUsecase.CreateUserByEmail(ctx, user3ID, "user3@example.ru")
	s.Require().NoError(err)

	user3TokenHash := "333"
	refreshTokenNoValid := model.RefreshToken{
		TokenHash: user3TokenHash,
		ID:        "41d862b0-81ee-4167-b19c-111e8126cc35",
		UserID:    user2ID,
		ExpiredAt: time.Now().UTC().Add(1 * time.Microsecond),
	}
	err = s.handler.refreshTokenUsecase.CreateRefreshToken(ctx, refreshTokenNoValid)
	s.Require().NoError(err)

	testCases := []struct {
		name           string
		userID         string
		refreshTokenID string
		refreshToken   string
		expectedError  bool
		expectedStatus int
	}{
		{
			name:           "RefreshTokenValid",
			userID:         user1ID,
			refreshTokenID: refreshTokenValid.ID,
			refreshToken:   user1TokenHash,
			expectedError:  false,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "RefreshTokenExpired",
			userID:         user2ID,
			refreshTokenID: refreshTokenExpired.ID,
			refreshToken:   user2TokenHash,
			expectedError:  true,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "InvalidRefreshToken",
			userID:         user3ID,
			refreshTokenID: "invalid token id",
			refreshToken:   "invalid-token",
			expectedError:  true,
			expectedStatus: http.StatusConflict,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ctx := context.Background()
			ctx = context.WithValue(ctx, contextKeyUserID, tc.userID)
			ctx = context.WithValue(ctx, contextKeyIP, "127.0.0.1")
			ctx = context.WithValue(ctx, contextKeyRefreshID, tc.refreshTokenID)

			req := httptest.NewRequest(http.MethodPost, "/auth/refresh", nil)
			req = req.WithContext(ctx)
			req.RemoteAddr = "127.0.0.1:12345"

			params := api.PostAuthRefreshParams{
				HTTPRequest: req,
				RefreshTokenBody: &definition.RefreshTokenBody{
					RefreshToken: tc.refreshToken,
				},
			}

			resp := s.handler.UpdateAccessRefreshTokensHandler(params)

			if tc.expectedError {
				if tc.expectedStatus == 400 {
					s.IsType(api.NewPostAuthRefreshBadRequest(), resp)
				} else {
					s.IsType(api.NewPostAuthRefreshConflict(), resp)
				}
			} else {
				result, ok := resp.(*api.PostAuthRefreshOK)
				s.True(ok)
				s.NotEmpty(result.Payload.AccessToken)
			}
		})
	}
}
