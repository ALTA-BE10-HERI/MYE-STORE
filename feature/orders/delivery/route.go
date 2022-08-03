package delivery

import (
	"project3/group3/domain"
	// _middleware "project3/group3/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteOrders(e *echo.Echo, do domain.OrdersHandler) {
	// e.GET("/products/:id", do.GetById(), _middleware.JWTMiddleware())                 //user & admin
	// e.GET("/products", dp.GetAll())
}
