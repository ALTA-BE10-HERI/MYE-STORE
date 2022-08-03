package delivery

import "project3/group3/domain"

type InsertFormat struct {
	Qty       int `json:"qty" form:"qty"`
	Total     int `json:"total" form:"total"`
	IDProduct int
}

func (i *InsertFormat) ToModel() domain.Orders {
	return domain.Orders{
		Qty:   i.Qty,
		Total: i.Total,
		Product: domain.ProductBuyyed{
			ID:           i.IDProduct,
			ProductName:  i.ToModel().Product.ProductImage,
			ProductImage: i.ToModel().Product.ProductImage,
		},
	}
}
