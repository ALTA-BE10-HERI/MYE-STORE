package data

import (
	"project3/group3/domain"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Qty       int `json:"qty" form:"qty"`
	Total     int `json:"total" form:"total"`
	ProductID int
	Product   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type Product struct {
	gorm.Model
	ProductName  string   `json:"product_name" form:"product_name"`
	ProductImage string   `form:"product_image"`
	Stock        int      `json:"stock" form:"stock"`
	Price        int      `json:"price" form:"price"`
	Orders       []Orders `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (o *Orders) ToDomain() domain.Orders {
	return domain.Orders{
		ID:    int(o.ID),
		Qty:   o.Qty,
		Total: o.Total,
		Product: domain.ProductBuyyed{
			ID:           int(o.ProductID),
			ProductName:  o.ProductName,
			ProductImage: o.ProductImage,
			Stock:        o.Stock,
			Price:        o.Price,
		},
	}
}

func ParseToArrOrders(arr []Orders) []domain.Orders {
	var res []domain.Orders
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func FromModel(data domain.Orders) Orders {
	var res Orders
	res.Qty = data.Qty
	res.Total = data.Total
	res.ProductID = data.Product.ID
	return res
}
