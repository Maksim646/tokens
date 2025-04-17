package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/Maksim646/tokens/database/postgresql"
	"github.com/Maksim646/tokens/internal/api/server/restapi/handler"
	"github.com/Maksim646/tokens/pkg/logger"
	"github.com/justinas/alice"

	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	_userRepo "github.com/Maksim646/tokens/internal/domain/user/repository/postgresql"
	_userUsecase "github.com/Maksim646/tokens/internal/domain/user/usecase"

	_refreshTokenRepo "github.com/Maksim646/tokens/internal/domain/refresh_token/repository/postgresql"
	_refreshTokenUsecase "github.com/Maksim646/tokens/internal/domain/refresh_token/usecase"
)

const (
	httpVersion = "development"
)

var config struct {
	Addr               string        `envconfig:"ADDR" default:"8000"`
	LogLevel           string        `envconfig:"LOG_LEVEL" default:"DEBUG"`
	MigrationsDir      string        `envconfig:"MIGRATIONS_DIR" default:"../../database/postgresql/migrations"`
	PostgresURI        string        `envconfig:"POSTGRES_URI" default:"postgres://postgres:tokens@localhost:5448/tokens_db?sslmode=disable"`
	JWTSigninKey       string        `envconfig:"JWT_SIGNIN_KEY" default:"MaximAdamov2002"`
	AccessTokenTTL     time.Duration `envconfig:"ACCESS_TOKEN_TTL" default:"15m"`
	RefreshTokenTTL    time.Duration `envconfig:"REFRESH_TOKEN_TTL" default:"30m"`
	RefreshTokenLength int           `envconfig:"REFRESH_TOKEN_LENGTH" default:"32"`
}

func main() {
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("env processing failed: %v", err)
	}

	if err := logger.BuildLogger(config.LogLevel); err != nil {
		log.Fatal("cannot build logger: ", err)
	}
	zap.L().Sync()

	zap.L().Info("PostgresURI: ", zap.String("uri", config.PostgresURI))
	zap.L().Info("MigrationsDir: ", zap.String("dir", config.MigrationsDir))

	time.Sleep(3 * time.Second)

	migrator := postgresql.NewMigrator(config.PostgresURI, config.MigrationsDir)
	if err := migrator.Apply(); err != nil {
		log.Fatal("cannot apply migrations: ", err)
	}

	sqlxConn, err := sqlx.Connect("postgres", config.PostgresURI)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}
	defer sqlxConn.Close()

	sqlxConn.SetMaxOpenConns(100)
	sqlxConn.SetMaxIdleConns(100)
	sqlxConn.SetConnMaxLifetime(5 * time.Minute)

	sqalxConn, err := sqalx.New(sqlxConn)
	if err != nil {
		log.Fatal("cannot connect to postgres db: ", err)
	}
	defer sqalxConn.Close()

	zap.L().Info("Database manage was process successfully")

	userRepo := _userRepo.New(sqalxConn)
	userUsecase := _userUsecase.New(userRepo)

	refreshTokenRepo := _refreshTokenRepo.New(sqalxConn)
	refreshTokenUsecase := _refreshTokenUsecase.New(refreshTokenRepo)

	appHandler := handler.New(
		userUsecase,
		refreshTokenUsecase,

		httpVersion,
		config.JWTSigninKey,
		config.AccessTokenTTL,
		config.RefreshTokenTTL,
		config.RefreshTokenLength,
	)

	chain := alice.New(appHandler.WsMiddleware).Then(appHandler)
	if chain == nil {
		fmt.Println(chain)
	}
	server := http.Server{
		Handler: chain,
		Addr:    ":" + config.Addr,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	zap.L().Info("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	zap.L().Info("Server exiting")
}
