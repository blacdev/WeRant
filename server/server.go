package server

import (
	"github.com/blacdev/werant/controller"
	"github.com/blacdev/werant/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Start(addr string, cts *controller.Container, sc *service.Container) error {
	e := echo.New()

	//todo: register logging middleware

	buildRoutes(e, cts)

	if err := e.Start(addr); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func buildRoutes(e *echo.Echo, cts *controller.Container) {
	e.GET("/health", controller.Health)
	e.POST("/user/register", controller.UserRegistration)
	e.POST("/user/login", controller.UserLogin)
}
