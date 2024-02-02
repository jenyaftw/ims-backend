package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jenyaftw/scaffold-go/internal/adapters/config"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http/handlers"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

func AuthMiddleware(next http.Handler) http.Handler {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				handlers.HandleError(w, domain.ErrMissingAuthHeader)
				return
			}

			splitToken := strings.Split(tokenString, " ")
			if splitToken[0] != "Bearer" {
				handlers.HandleError(w, domain.ErrInvalidTokenType)
				return
			}

			token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(cfg.Jwt.Secret), nil
			})
			if err != nil {
				handlers.HandleError(w, domain.ErrInvalidAuthToken)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				ctx := context.WithValue(r.Context(), "user", claims["sub"])
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				handlers.HandleError(w, domain.ErrInvalidAuthToken)
			}
		},
	)
}
