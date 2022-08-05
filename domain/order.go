package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID         int
	Stock      int
	Status     string
	TotalPrice int
	UserID     int
	ProductID  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    ProductOrder
	User       UserOrder
}

type ProductOrder struct {
	ID           int
	ProductName  string
	Stock        int
	Price        int
	ProductImage string
	UserID       int
	User         UserCart
}
type UserOrder struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Role    string
	Address string
}
type OrderDetail struct {
	Name       string
	Email      string
	Qty        int
	TotalPrice int
}

//logic
type OrderUseCase interface {
	ConfirmData(idOrder, idFromToken int) (row int, err error)
	CancelData(idOrder, idFromToken int) (row int, err error)
	GetAllData(limit, offset int) (data []Order, err error)
	GetProduct(idProduct int) (data OrderDetail, err error)
	CreateData(data Order) (row int, err error)
	GetUser(idUser int) (OrderDetail, error)
}

//query
type OrderData interface {
	ConfirmOrder(idOrder, idFromToken int) (row int, err error)
	CancelOrder(idOrder, idFromToken int) (row int, err error)
	SelectData(limit, offset int) (data []Order, err error)
	GetUser(idUser int) (data OrderDetail, err error)
	GetProduct(idProduct int) (data OrderDetail, err error)
	InsertData(data Order) (row int, err error)
}

//handler
type OrderHandler interface {
	ConfirmStatus() echo.HandlerFunc
	CancelStatus() echo.HandlerFunc
	GetAllData() echo.HandlerFunc
	InsertOrder() echo.HandlerFunc
}
