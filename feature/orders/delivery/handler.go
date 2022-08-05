package delivery

import (
	"fmt"
	"net/http"
	"project3/group3/domain"
	"strconv"

	_middleware "project3/group3/feature/common"
	_helper "project3/group3/helper"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUseCase domain.OrderUseCase
}

func New(ou domain.OrderUseCase) domain.OrderHandler {
	return &OrderHandler{
		orderUseCase: ou,
	}
}

func (oh *OrderHandler) ConfirmStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idOrder, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		row, err := oh.orderUseCase.ConfirmData(idOrder, idFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (oh *OrderHandler) CancelStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idOrder, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		row, err := oh.orderUseCase.ConfirmData(idOrder, idFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (oh *OrderHandler) GetAllData() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitcnv, _ := strconv.Atoi(limit)
		offsetcnv, _ := strconv.Atoi(offset)
		result, err := oh.orderUseCase.GetAllData(limitcnv, offsetcnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed get all data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(result)))
	}

}

func (oh *OrderHandler) InsertOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		var insertFormat OrderFormat
		idUser, _ := _middleware.ExtractData(c)
		err := c.Bind(&insertFormat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind insert"))
		}
		fmt.Println("id product: ", insertFormat.IdProduct)
		fmt.Println("id user: ", idUser)
		product, err := oh.orderUseCase.GetProduct(int(insertFormat.IdProduct))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("cant parse data"))
		}
		user, err := oh.orderUseCase.GetUser(idUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("cant parse data"))
		}
		fmt.Println("product", product)
		fmt.Println("user:", user)
		var Total int
		Name := user.Name
		Email := user.Email
		qty := insertFormat.Stock
		price := product.TotalPrice
		for i := 0; i < qty; i++ {
			Total += price
		}
		fmt.Println("nama: ", Name)
		fmt.Println("email: ", Email)
		fmt.Println("total:", Total)
		data := domain.Order{}
		data.ProductID = insertFormat.IdProduct
		data.UserID = idUser
		data.Status = "Pending"
		data.Stock = qty
		fmt.Println(data)
		res := _helper.Payment(Name, Email, Total)
		datas, _ := oh.orderUseCase.CreateData(data)
		fmt.Println("hasil dari data:", datas)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"message":  "success checkout",
			"payments": res,
		})
	}
}
