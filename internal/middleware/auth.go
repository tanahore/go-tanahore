package middleware

import (
	"fmt"
	"strings"

	"tanahore/internal/pkg/jwt"
	res "tanahore/internal/pkg/responses"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc, condition string) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header["Authorization"]

		if len(authorization) <= 0 {
			return res.StatusUnauthorized(c, "missing or malformed jwt.", fmt.Errorf("unauthorized"))
		}

		userToken := strings.Split(authorization[0], " ")

		if len(userToken) <= 1 {
			return res.StatusBadRequest(c, "invalid token.", fmt.Errorf("invalid token"))
		}

		if userToken[0] != "Bearer" {
			return res.StatusBadRequest(c, "invalid token type.", fmt.Errorf("invalid token type"))
		}

		data, err := jwt.ExtractToken(userToken[1])

		if err != nil {
			return res.StatusBadRequest(c, "invalid token!", err)
		}

		if condition == "all" {
			if data.RoleName != "admin" && data.RoleName != "user" && data.RoleName != "student" {
				return res.StatusForbidden(c, "you dont have access!", fmt.Errorf("access forbidden!"))
			}
		}

		if condition == "admin" {
			if data.RoleName != "admin" {
				return res.StatusForbidden(c, "you dont have access!", fmt.Errorf("access forbidden!"))
			}
		}

		if condition == "user" {
			if data.RoleName != "user" && data.RoleName != "admin" {
				return res.StatusForbidden(c, "you dont have access!", fmt.Errorf("access forbidden!"))
			}
		}

		c.Set("user_id", data.ID)
		c.Set("role_name", data.RoleName)
		c.Set("username", data.Username)
		c.Set("create_user", data.CreatedAt.Unix())

		return next(c)
	}
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return AuthMiddleware(next, "admin")
}

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return AuthMiddleware(next, "user")
}
func AllMiddleare(next echo.HandlerFunc) echo.HandlerFunc {
	return AuthMiddleware(next, "all")
}
