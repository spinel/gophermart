package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/spinel/gophermart/model"

	"github.com/labstack/echo/v4"
)

// RegisterUser creates new user
func (ctr *Controller) RegisterUser(c echo.Context) error {
	var userRegisterForm model.UserRegisterForm
	err := c.Bind(&userRegisterForm)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("could not decode user data"))
	}

	err = c.Validate(&userRegisterForm)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	createdUser, err := ctr.services.User.Create(ctr.ctx, userRegisterForm)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("could not create an user"))
	}

	if createdUser == nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("could not create an user"))
	}

	ctr.logger.Debug().Msgf("Created user '%d'", createdUser)

	return c.JSON(http.StatusOK, createdUser)
}

// UserLogin handle user sign in
func (ctr *Controller) UserLogin(c echo.Context) error {
	var userRegisterForm model.UserRegisterForm
	if err := c.Bind(&userRegisterForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("could not decode credentials data"))
	}

	if err := c.Validate(&userRegisterForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := ctr.services.User.Login(ctr.ctx, userRegisterForm)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("user could not login"))
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("permission denied"))
	}

	// create new token
	sessionToken := createSessionToken()

	// Set the token in the cache, along with the user whom it represents
	ctr.services.Memory.Add(ctr.ctx, sessionToken, user.ID)

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	http.SetCookie(c.Response().Writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(1200 * time.Second),
	})

	return c.JSON(http.StatusOK, user)
}

func createSessionToken() string {
	// Create a new random session token
	return uuid.NewString()
}
