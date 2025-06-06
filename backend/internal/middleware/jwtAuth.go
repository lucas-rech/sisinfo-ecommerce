package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuth is a middleware function that checks for a valid JWT token in the request.
func JWTAuthAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := ValidateAdminRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only Administrator is allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

// check for valid customer token
func JWTAuthCustomer() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := ValidateCustomerRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only registered Customers are allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

