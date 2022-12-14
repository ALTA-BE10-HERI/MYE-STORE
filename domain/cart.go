package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Cart struct {
	ID         int
	Stock      int
	Status     string
	TotalPrice int
	UserID     int
	ProductID  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    ProductCart
	User       UserCart
}

type ProductCart struct {
	ID           int
	ProductName  string
	Stock        int
	Price        int
	ProductImage string
	UserID       int
	User         UserCart
}
type UserCart struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Role    string
	Address string
}

//logic
type CartUseCase interface {
	GetAllData(limit, offset, idFromToken int) (data []Cart, err error)
	CreateData(data Cart) (row int, err error)
	UpdateData(stock, idCart, idFromToken int) (row int, err error)
	DeleteData(idProd, idFromToken int) (row int, err error)
}

//usecase
type ChartData interface {
	InsertData(data Cart) (row int, err error)
	SelectData(limit, offset, idFromToken int) (data []Cart, err error)
	CheckCart(idProd, idFromToken int) (isExist bool, idCart, stock int, err error)
	UpdateDataDB(stock, idCart, idFromToken int) (row int, err error)
	DeleteDataDB(idProd, idFromToken int) (row int, err error)
}

type CartHandler interface {
	PostCart() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}
