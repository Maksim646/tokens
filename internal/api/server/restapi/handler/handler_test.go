package handler

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_httpTransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Maksim646/tokens/database/postgresql/pgtest"
	"github.com/Maksim646/tokens/internal/api/client"
	"github.com/Maksim646/tokens/internal/api/client/operations"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	_userRepo "github.com/Maksim646/tokens/internal/domain/user/repository/postgresql"
	_userUsecase "github.com/Maksim646/tokens/internal/domain/user/usecase"

	_refreshTokenRepo "github.com/Maksim646/tokens/internal/domain/refresh_token/repository/postgresql"
	_refreshTokenUsecase "github.com/Maksim646/tokens/internal/domain/refresh_token/usecase"
)

const (
	jwtSigninKey       = "MaximAdamov2002"
	accessTokenTTL     = time.Duration(15 * time.Minute)
	refreshTokenTTL    = time.Duration(30 * time.Minute)
	refreshTokenLength = 32
)

type Suite struct {
	pgtest.Suite
	ctx     context.Context
	cancel  context.CancelFunc
	api     operations.ClientService
	server  *httptest.Server
	handler *Handler
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	zap.ReplaceGlobals(zaptest.NewLogger(s.T(), zaptest.WrapOptions(zap.AddCaller())))
	s.Suite.SetupSuite("./../../../..")

	userRepo := _userRepo.New(s.DB())
	userUsecase := _userUsecase.New(userRepo)

	refreshTokenRepo := _refreshTokenRepo.New(s.DB())
	refreshTokenUsecase := _refreshTokenUsecase.New(refreshTokenRepo)

	s.handler = New(
		userUsecase,
		refreshTokenUsecase,

		"test",
		jwtSigninKey,
		accessTokenTTL,
		refreshTokenTTL,
		refreshTokenLength,
	)

	s.server = httptest.NewServer(s.handler)
	transport := _httpTransport.New(strings.TrimPrefix(s.server.URL, "http://"), "", nil)
	s.api = client.New(transport, strfmt.Default).Operations
}

func (s *Suite) SetupTest() {
	zap.ReplaceGlobals(zaptest.NewLogger(s.T(), zaptest.WrapOptions(zap.AddCaller())))
	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.Suite.SetupTest()

	userRepo := _userRepo.New(s.DB())
	userUsecase := _userUsecase.New(userRepo)

	refreshTokenRepo := _refreshTokenRepo.New(s.DB())
	refreshTokenUsecase := _refreshTokenUsecase.New(refreshTokenRepo)

	s.handler.userUsecase = userUsecase
	s.handler.refreshTokenUsecase = refreshTokenUsecase

}

func (s *Suite) TearDownTest() {
	zap.ReplaceGlobals(zap.NewNop())
	s.cancel()
	s.Suite.TearDownTest()
}

func (s *Suite) TearDownSuite() {
	s.Suite.TearDownSuite()
	s.server.Close()
}
