package delivery

import (
	"log"
	"net/http"
	"project3/group3/domain"
	_middleware "project3/group3/feature/common"
	_helper "project3/group3/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUseCase domain.ProductUseCase
}

func New(ps domain.ProductUseCase) domain.ProductHandler {
	return &productHandler{
		productUseCase: ps,
	}
}

func (ph *productHandler) InsertProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertProductFormat
		idFromToken, _ := _middleware.ExtractData(c)
		if idFromToken != 1 && idFromToken != 2 {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("your account not admin"))
		}

		err := c.Bind(&tmp)
		tmp.UserID = idFromToken
		if err != nil {
			log.Println("cannot parse data", err)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data"))
		}

		dataProduct := tmp.ToModel()
		result, errCreate := ph.productUseCase.InsertProduct(dataProduct)
		if errCreate != nil {
			log.Println("ini data product : ", dataProduct)
			log.Println("err : ", errCreate)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to post product"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModel(result)))
	}
}

func (ph *productHandler) DeleteProductHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idproduct := c.Param("id")
		id, _ := strconv.Atoi(idproduct)
		if id == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
		}
		row, errDel := ph.productUseCase.DeleteProduct(id)

		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to delete data product"))
		}

		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data product"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (ph *productHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitint, _ := strconv.Atoi(limit)
		offsetint, _ := strconv.Atoi(offset)
		res, err := ph.productUseCase.GetAllData(limitint, offsetint)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("canot to read all data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(res)))

	}
}
func (ph *productHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idProduct, _ := strconv.Atoi(id)
		res, err := ph.productUseCase.GetProductById(idProduct)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get detail roduct"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success ", FromModel(res)))
	}

}
