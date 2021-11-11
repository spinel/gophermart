package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spinel/gophermart/pkg/types"
)

func (ctr *Controller) Orders(c echo.Context) error {
	userID := getEchoParamInt(c, "user")

	bodyOrderNumber, _ := ioutil.ReadAll(c.Request().Body)
	orderNumber := string(bodyOrderNumber)

	createdOrder, err := ctr.services.Order.Create(ctr.ctx, userID, orderNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("could not create an order: %w", err))
	}

	if createdOrder == nil {
		return c.JSON(http.StatusNoContent, "")
	}

	return c.JSON(http.StatusAccepted, createdOrder)
}

func (ctr *Controller) OrdersList(c echo.Context) error {
	userID := getEchoParamInt(c, "user")
	orders, err := ctr.services.Order.List(ctr.ctx, userID)

	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create an order"))
		}
	}

	return c.JSON(http.StatusOK, orders)
}
