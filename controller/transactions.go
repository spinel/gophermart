package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spinel/gophermart/model"
)

func (ctr *Controller) Balance(c echo.Context) error {
	userID := getEchoParamInt(c, "user")

	balance, err := ctr.services.Transaction.Balance(ctr.ctx, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("could not get balance: %w", err))
	}

	return c.JSON(http.StatusOK, balance)
}

func (ctr *Controller) Withdraw(c echo.Context) error {
	userID := getEchoParamInt(c, "user")

	transactionRequest := new(model.TransactionRequest)
	if err := c.Bind(transactionRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("could not get body: %w", err))
	}

	err := ctr.services.Transaction.Withdraw(ctr.ctx, userID, transactionRequest.Order, transactionRequest.Sum)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("could not get balance: %w", err))
	}

	return c.JSON(http.StatusOK, transactionRequest)
}
