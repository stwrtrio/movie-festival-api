package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	JTI    string `json:"jti"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(userID, email string) (string, error) {
	expiry := LoadJWTExpiry()
	expirationTime := time.Now().Add(expiry)
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		JTI: GenerateUUID(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secretKey))
}

func ValidateJWTToken(tokenString string) (*Claims, error) {
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}

func LoadJWTExpiry() time.Duration {
	expiryStr := os.Getenv("JWT_EXPIRY")
	expiry, err := time.ParseDuration(expiryStr)
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRY in .env: %v", err)
	}
	return expiry
}
