package route

import (
	"context"

	"github.com/spinel/gophermart/controller"

	"github.com/labstack/echo/v4"
)

type Route struct {
	ctx        context.Context
	controller *controller.Controller
	Echo       *echo.Echo
}

// New Route
func New(ctx context.Context, controller *controller.Controller) *Route {
	e := echo.New()
	return &Route{
		ctx:        ctx,
		controller: controller,
		Echo:       e,
	}
}

func (r *Route) InitRoutes() {
	r.middleware()
	r.initUser()
}
