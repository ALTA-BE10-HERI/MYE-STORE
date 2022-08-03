package factory

import (
	// pd "project3/group3/feature/product/data"
	// productDelivery "project3/group3/feature/product/delivery"
	// ps "project3/group3/feature/product/usecase"

	ud "project3/group3/feature/users/data"
	userDelivery "project3/group3/feature/users/delivery"
	us "project3/group3/feature/users/usecase"

	pd "project3/group3/feature/products/data"
	productDelivery "project3/group3/feature/products/delivery"
	ps "project3/group3/feature/products/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := ud.New(db)
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)

	productData := pd.New(db)
	productCase := ps.New(productData)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteProduct(e, productHandler)

}
