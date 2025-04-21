package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var expectedAudience = "https://apigomedium-production.up.railway.app" // Cambia esto por tu audiencia esperada
var expectedIssuer = "https://dev-arcw7okepellj1bi.us.auth0.com/"      // Cambia esto por tu emisor esperado

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token no proporcionado", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, JWKS.Keyfunc)
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "No se pudieron extraer los claims del token", http.StatusUnauthorized)
			return
		}

		if claims["aud"] != expectedAudience || claims["iss"] != expectedIssuer {
			http.Error(w, "Audiencia o emisor inválido", http.StatusUnauthorized)
			return

		}

		// ✅ Token válido, continuar
		next.ServeHTTP(w, r)
	})
}
