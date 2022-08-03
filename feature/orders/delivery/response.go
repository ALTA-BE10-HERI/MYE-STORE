package delivery

import "project3/group3/domain"

type Orders struct {
	ID            int     `json:"id"`
	ProductBuyyed Product `json:"product"`
	Qty           int     `json:"qty"`
	Total         int     `json:"total"`
}

type Product struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Stock        int    `json:"stock"`
	Price        int    `json:"price"`
}

func FromModel(data domain.Orders) Orders {
	return Orders{
		ID:    data.ID,
		Qty:   data.Qty,
		Total: data.Total,
		ProductBuyyed: Product{
			ID:           data.Product.ID,
			ProductName:  data.Product.ProductName,
			ProductImage: data.Product.ProductImage,
			Stock:        data.Product.Stock,
			Price:        data.Product.Price,
		},
	}
}

func FromModelList(data []domain.Orders) []Orders {
	result := []Orders{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}
