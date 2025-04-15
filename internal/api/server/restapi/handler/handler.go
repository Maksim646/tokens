package handler

import (
	_ "context"
	_ "fmt"

	"net/http"
	"strings"

	"encoding/json"

	_ "github.com/Maksim646/Tokens/internal/api/definition"
	_ "github.com/Maksim646/Tokens/internal/model"
	_ "github.com/Maksim646/Tokens/pkg/jsonwebtoken"
	"go.uber.org/zap"

	"github.com/Maksim646/Tokens/internal/api/server/restapi"
	"github.com/Maksim646/Tokens/internal/api/server/restapi/api"
	"github.com/go-openapi/loads"
)

type Handler struct {
	router       http.Handler
	HashSalt     string
	jwtSigninKey string
}

func New(

	version string,
	HashSalt string,
	jwtSigninKey string,
) *Handler {

	withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", version)
	swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	if err != nil {
		panic(err)
	}

	h := &Handler{

		HashSalt:     HashSalt,
		jwtSigninKey: jwtSigninKey,
	}

	zap.L().Error("server http handler request")
	router := api.NewTokensBackendServiceAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof
	// router.BearerAuth = h.ValidateHeader

	h.router = router.Serve(nil)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Received HTTP request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	if h.router == nil {
		zap.L().Error("h.router is nil")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	zap.L().Info("h.router is not nil, processing request")
	h.router.ServeHTTP(w, r)
}

// func (h *Handler) ValidateHeader(bearerHeader string) (*definition.Principal, error) {
// 	ctx := context.Background()

// 	fmt.Println("1")
// 	bearerToken := strings.TrimPrefix(bearerHeader, "Bearer ")
// 	userID, roleID, err := jsonwebtoken.ParseToken(bearerToken, h.jwtSigninKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if roleID == 0 {
// 		_, err = h.userUsecase.GetUserByID(ctx, userID)
// 		if err != nil {
// 			return nil, err
// 		}

// 	} else {
// 		_, err = h.adminUsecase.GetAdminByID(ctx, userID)
// 		if err != nil {
// 			return nil, err
// 		}

// 	}

// 	return &definition.Principal{ID: userID, Role: roleID}, nil
// }
