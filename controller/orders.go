package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (ctr *Controller) Orders(c echo.Context) error {
	userID := getEchoParamInt(c, "user")

	bodyOrderNumber, _ := ioutil.ReadAll(c.Request().Body)
	orderNumber, _ := strconv.Atoi(string(bodyOrderNumber))

	createdOrder, err := ctr.services.Order.Create(ctr.ctx, userID, orderNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("could not create an order: %w", err))
	}

	return c.JSON(http.StatusOK, createdOrder)
}

func (ctr *Controller) OrdersList(c echo.Context) error {
	userID := getEchoParamInt(c, "user")
	orders, err := ctr.services.Order.List(ctr.ctx, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("could not create an order: %w", err))
	}

	return c.JSON(http.StatusOK, orders)
}
