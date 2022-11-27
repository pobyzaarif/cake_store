package router

import (
	"github.com/labstack/echo/v4"

	"github.com/pobyzaarif/cake_store/app/main/controller"
	"github.com/pobyzaarif/cake_store/config"
)

var apiVersion1 = "api/v1"

func RegisterPath(
	e *echo.Echo,
	appConfig *config.Config,
	controller *controller.Controller,
) {
	// public
	e.GET(apiVersion1+"/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	// cake
	cake := e.Group(apiVersion1 + "/cakes")
	cake.POST("", controller.CakeCreate)
	cake.GET("", controller.CakeFindAll)
	cake.GET("/:id", controller.CakeFindByID)
	cake.PUT("", controller.CakeUpdate)
	cake.PATCH("/:id", controller.CakeUpdate)
	cake.DELETE("/:id", controller.CakeDelete)
}
