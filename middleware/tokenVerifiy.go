package middleware

import (
	"Portofolio/SistemPerpustakaan/utils"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// TokenVerifyMiddleware to verify token
func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (i interface{}, err error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There Was An Error %w", err)
				}
				return []byte(os.Getenv("SECRET")), nil
			})

			if err != nil {
				fmt.Println("There are error with : %w", err.Error())
				utils.HandleError(resp, http.StatusUnauthorized, "Ooppss , something when wrong")
				return
			}

			if token.Valid {
				next.ServeHTTP(resp, req)
			} else {
				utils.HandleError(resp, http.StatusUnauthorized, "Invalid Token")
				return
			}

		} else {
			utils.HandleError(resp, http.StatusUnauthorized, "Invalid Token")
			return
		}
	})
}
