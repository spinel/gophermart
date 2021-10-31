package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spinel/gophermart/service"
)

type (
	Auth struct {
		services *service.Manager
	}
)

// NewAuth create new auth middleware.
func NewAuth(services *service.Manager) *Auth {
	return &Auth{
		services: services,
	}
}

// Process is the middleware auth api.
func (s *Auth) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		openURLs := map[string]bool{
			"/api/user/register": true,
			"/api/user/login":    true,
		}

		if !openURLs[c.Request().RequestURI] {
			ctx := c.Request().Context()
			// We can obtain the session token from the requests cookies, which come with every request
			cookie, err := c.Request().Cookie("session_token")
			if err != nil {
				if err == http.ErrNoCookie {
					// If the cookie is not set, return an unauthorized status
					return echo.ErrUnauthorized
				}
				// For any other type of error, return a bad request status
				return echo.ErrBadRequest
			}
			sessionToken := cookie.Value

			// We then get the name of the user from our cache, where we set the session token
			cacheResponse, err := s.services.Cache.Do(ctx, "GET", sessionToken)
			if err != nil {
				// If there is an error fetching from cache, return an internal server error status
				return echo.ErrInternalServerError
			}

			if cacheResponse == nil {
				// If the session token is not present in cache, return an unauthorized error
				return echo.ErrUnauthorized
			}

			c.Set("user", cacheResponse)
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
