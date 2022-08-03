package data

import (
	// "errors"
	// "log"
	"project3/group3/domain"

	"gorm.io/gorm"
)

type ordersData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrdersData {
	return &ordersData{
		db: db,
	}
}

func (od *ordersData) GetOrders(idOrders int) (result domain.Orders, err error) {
	var tmp Orders
	res := od.db.Where("id = ?", idOrders).First(&tmp)
	if res.Error != nil {
		return domain.Orders{}, res.Error
	}
	return tmp.ToDomain(), nil
}

func (od *ordersData) SelectData(limit, offset int) (data []domain.Orders, err error) {
	dataOrders := []Orders{}
	res := od.db.Limit(limit).Offset(offset).Find(&dataOrders)
	if res.Error != nil {
		return []domain.Orders{}, nil
	}
	return ParseToArrOrders(dataOrders), nil
}
