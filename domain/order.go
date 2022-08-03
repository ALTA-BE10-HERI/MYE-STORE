package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Orders struct {
	ID        int
	ProductID int
	Qty       int
	Total     int
	CreatedAt time.Time
	UpdatedAt time.Time
	Product   ProductBuyyed
}

type ProductBuyyed struct {
	ID           int
	ProductName  string
	Stock        int
	Price        int
	ProductImage string
}

// logic
type OrdersUseCase interface {
	GetAllData(limit, offset int) (data []Orders, err error)
}

// query
type OrdersData interface {
	GetOrders(idOrders int) (result Orders, err error) //buat ambil nama id
	SelectData(limit, offset int) (data []Orders, err error)

}


// handler
type OrdersHandler interface {
	GetAll() echo.HandlerFunc
}
