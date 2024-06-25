package services

import (
	"context"
	"fmt"
	"test-go/models"
)

type OrderService struct {
}

// TODO: func (DB ?) Create(data input body)
func (service *OrderService) Create(ctx context.Context) (models.Order, error) {
	fmt.Println("services order create", ctx)

	order := models.Order{
		OrderId:       1,
		OrderAgencyId: "test1",
	}

	return order, nil
}
