package jwt

import (
	"fmt"
	"os"
	"time"

	"tanahore/internal/model/web"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(tokenString string) (*web.AuthResponse, error) {

	type CustomClaims struct {
		ID        uint      `json:"id"`
		Email     string    `json:"email"`
		Username  string    `json:"username"`
		RoleName  string    `json:"role_name"`
		CreatedAt time.Time `json:"created_at"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		extractedToken := &web.AuthResponse{
			ID:        claims.ID,
			Username:  claims.Username,
			Email:     claims.Email,
			RoleName:  web.Role(claims.RoleName),
			CreatedAt: claims.CreatedAt,
		}
		return extractedToken, nil
	}

	return nil, fmt.Errorf("invalid token")

}
