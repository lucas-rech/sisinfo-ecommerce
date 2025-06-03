package middleware

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/backend/utils"
)



var secretKey = []byte(utils.GetEnv("JWT_SECRET"))

func GenerateJWT(user dto.UserResponse) (string, error) {
	hoursToExpire, err := strconv.ParseInt(utils.GetEnv("JWT_EXPIRATION_TIME"), 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid JWT expiration time: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
	jwt.MapClaims{
		"username": user.Name,
		"role":    user.Role,
		"iat":  time.Now().Unix(),	
		"exp": time.Now().Add(time.Hour * time.Duration(hoursToExpire)).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func ValidateJWT (context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return fmt.Errorf("invalid token claims")
}

func ValidateCustomerRoleJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token claims")
	}

	role, exists := claims["role"]
	if !exists || role != "CUSTOMER" {
		return fmt.Errorf("unauthorized access: customer role required")
	}

	return nil
}

func ValidateAdminRoleJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token claims")
	}

	role, exists := claims["role"]
	if !exists || role != "ADMIN" {
		return fmt.Errorf("unauthorized access: admin role required")
	}

	return nil
}
