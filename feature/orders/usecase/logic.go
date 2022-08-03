package usecase

import (
	"project3/group3/domain"
)

type ordersUseCase struct {
	ordersData domain.OrdersData
}

func New(od domain.OrdersData) domain.OrdersUseCase {
	return &ordersUseCase{
		ordersData: od,
	}
}

func (od *ordersUseCase) GetAllData(limit, offset int) (data []domain.Orders, err error) {
	res, err := od.ordersData.SelectData(limit, offset)
	return res, err
}