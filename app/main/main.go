package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pobyzaarif/cake_store/app/main/controller"
	"github.com/pobyzaarif/cake_store/app/main/router"
	"github.com/pobyzaarif/cake_store/business/cake"
	"github.com/pobyzaarif/cake_store/config"
	cakeModule "github.com/pobyzaarif/cake_store/modules/cake"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	goLoggerAppName "github.com/pobyzaarif/go-logger/appname"
	goLogger "github.com/pobyzaarif/go-logger/logger"
	goLoggerEchoMiddlerware "github.com/pobyzaarif/go-logger/rest/framework/echo/v4/middleware"
)

var logger = goLogger.NewLog("MAIN")

func main() {

	conf := config.GetAPPConfig()
	db := conf.GetDatabaseConnection()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(goLoggerEchoMiddlerware.ServiceRequestTime)
	e.Use(goLoggerEchoMiddlerware.ServiceTrackerID)
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: goLoggerEchoMiddlerware.APILogHandler,
		Skipper: middleware.DefaultSkipper,
	}))

	e.Use(goLoggerEchoMiddlerware.Recover())

	cakeRepo := cakeModule.RepositoryFactory(db)
	cakeService := cake.NewService(cakeRepo)

	controllerAPP := controller.NewController(
		cakeService,
	)

	router.RegisterPath(
		e,
		conf,
		controllerAPP,
	)

	address := "0.0.0.0:" + conf.AppMainPort
	go func() {
		if err := e.Start(address); err != http.ErrServerClosed {
			logger.Fatal("failed on http server " + conf.AppMainPort)
		}
	}()

	logger.SetTrackerID("main")
	logger.Info(goLoggerAppName.GetAPPName() + " service running in " + address)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal(fmt.Sprintf("failed to shutting down echo server %v", err))
	} else {
		logger.Info("successfully shutting down echo server")
	}
}
