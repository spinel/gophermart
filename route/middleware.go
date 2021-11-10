package route

import (
	"github.com/spinel/gophermart/pkg/validator"

	libError "github.com/spinel/gophermart/pkg/error"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
)

func (r *Route) middleware() {
	r.Echo.Validator = validator.NewValidator()
	r.Echo.HTTPErrorHandler = libError.Error

	r.Echo.Use(middleware.Logger())
	r.Echo.Use(middleware.Recover())
	r.Echo.Use(session.Middleware(sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))))
}
