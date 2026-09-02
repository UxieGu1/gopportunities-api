package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret      []byte
	tokenExpiresIn time.Duration
)

func Configure() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "super-secret-key-change-me"
	}
	jwtSecret = []byte(secret)

	expiresInStr := os.Getenv("JWT_EXPIRES_IN")
	if expiresInStr == "" {
		tokenExpiresIn = time.Hour * 24
	} else {
		hours, err := strconv.Atoi(expiresInStr)
		if err != nil {
			tokenExpiresIn = time.Hour * 24
		} else {
			tokenExpiresIn = time.Duration(hours) * time.Hour
		}
	}
}

func init() {
	Configure()
}

type JwtCustomClaims struct {
	UserID uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, role string) (string, error) {
	if userID == 0 {
		return "", errors.New("userID inválido")
	}
	if role == "" {
		return "", errors.New("role inválido")
	}

	claims := &JwtCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gopportunities-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar token: %v", err)
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*JwtCustomClaims, error) {
	if tokenString == "" {
		return nil, errors.New("token vazio")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do token: %v", err)
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
