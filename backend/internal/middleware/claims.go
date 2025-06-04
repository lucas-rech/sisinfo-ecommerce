package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractClaims() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Extrai a claim de cart_ID do token JWT do usuário logado
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if cartID, ok := claims["cart_Id"].(float64); ok {
				c.Set("cart_ID", uint(cartID))
			}
			if email, ok := claims["email"].(string); ok {
				c.Set("email", email)
			}
		}

		c.Next()
	}
}