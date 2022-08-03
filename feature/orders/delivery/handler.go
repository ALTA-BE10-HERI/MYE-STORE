package delivery

import (
	"net/http"
	"project3/group3/domain"
	"strconv"

	// _middleware "project3/group3/feature/common"
	_helper "project3/group3/helper"

	"github.com/labstack/echo/v4"
)

type ordersHandler struct {
	ordersUseCase domain.OrdersUseCase
}

func New(os domain.OrdersUseCase) domain.OrdersHandler {
	return &ordersHandler{
		ordersUseCase: os,
	}

}
func (oh *ordersHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitint, _ := strconv.Atoi(limit)
		offsetint, _ := strconv.Atoi(offset)
		res, err := oh.ordersUseCase.GetAllData(limitint, offsetint)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("canot to read all data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(res)))

	}
}
