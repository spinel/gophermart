package controller

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/spinel/gophermart/controller/middleware"
	"github.com/spinel/gophermart/logger"
	"github.com/spinel/gophermart/service"
)

// Controller all controllers
type Controller struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger

	Auth *middleware.Auth
}

// New controller
func New(ctx context.Context, services *service.Manager, logger *logger.Logger) *Controller {
	auth := middleware.NewAuth(services)
	return &Controller{
		ctx:      ctx,
		services: services,
		logger:   logger,

		Auth: auth,
	}
}

func getEchoParamInt(c echo.Context, key string) int {
	echoParam := c.Get(key)
	value, ok := echoParam.(int)
	if ok {
		return value
	}
	return 0
}
