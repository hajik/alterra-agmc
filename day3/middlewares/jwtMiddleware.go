package middlewares

import (
	"mvcapp/constants"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		strconv.Itoa(userId),
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(constants.SECRET_JWT))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ExtractTOkenUserId(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(*JwtCustomClaims)
		userId := claims.UserID
		return userId
	}
	return ""
}

func JWTConfig() middleware.JWTConfig {
	// Configure middleware with the custom claims type
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(constants.SECRET_JWT),
	}
}
