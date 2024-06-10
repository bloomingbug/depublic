package middleware

import (
	"net/http"

	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	SecretKey string
}

type MiddlewareInterface interface {
	For(role string) echo.MiddlewareFunc
}

func NewMiddleware(secretKey string) MiddlewareInterface {
	return &Middleware{SecretKey: secretKey}
}

func (m *Middleware) For(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Guest role doesn't require JWT
			if role == "guest" {
				user := c.Get("user")
				if user != nil {
					return c.JSON(http.StatusForbidden, map[string]string{"message": "anda sudah login"})
				}
				return next(c)
			}

			// Set up the JWT middleware
			config := echojwt.Config{
				NewClaimsFunc: func(c echo.Context) jwt.Claims {
					return new(jwt_token.JwtCustomClaims)
				},
				SigningKey: []byte(m.SecretKey),
				ErrorHandler: func(c echo.Context, err error) error {
					return c.JSON(http.StatusUnauthorized, map[string]string{"message": "anda harus login untuk mengakses resource ini"})
				},
			}

			// Apply the JWT middleware
			if err := echojwt.WithConfig(config)(func(c echo.Context) error {
				user := c.Get("user")
				if user == nil {
					return c.JSON(http.StatusUnauthorized, map[string]string{"message": "anda harus login untuk mengakses resource ini"})
				}

				token := user.(*jwt.Token)
				claims := token.Claims.(*jwt_token.JwtCustomClaims)

				switch role {
				case "admin":
					if claims.Role != "admin" {
						return c.JSON(http.StatusForbidden, map[string]string{"message": "anda tidak memiliki akses"})
					}
				case "buyer":
					if claims.Role != "buyer" {
						return c.JSON(http.StatusForbidden, map[string]string{"message": "anda tidak memiliki akses"})
					}
				case "login":
					// User is already authenticated
				default:
					return c.JSON(http.StatusForbidden, map[string]string{"message": "invalid role"})
				}

				return next(c)
			})(c); err != nil {
				return err
			}

			return nil
		}
	}
}
