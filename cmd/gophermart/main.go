package main

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/spinel/gophermart/config"
	"github.com/spinel/gophermart/controller"
	"github.com/spinel/gophermart/logger"
	"github.com/spinel/gophermart/route"
	"github.com/spinel/gophermart/service"
	"github.com/spinel/gophermart/store"
)

func main() {
	l := logger.Get()
	if err := run(l); err != nil {
		l.Fatal().Msgf("init error, %s", err.Error())
	}
}

func run(l *logger.Logger) error {
	ctx := context.Background()

	// config
	cfg := config.Get()

	store, err := store.New(cfg)
	if err != nil {
		return errors.Wrap(err, "store")
	}

	// init service manager
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		return errors.Wrap(err, "manager")
	}
	c := controller.New(ctx, serviceManager, l)

	// init routes
	r := route.New(ctx, c)
	r.InitRoutes()

	// server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	err = r.Echo.StartServer(s)

	return errors.Wrap(err, "server")
}
