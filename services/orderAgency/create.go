package services

import (
	"context"
	"test-go/models"
)

type OrderAgencyService struct{}

func (orderAgencyService *OrderAgencyService) Create(ctx context.Context) (models.OrderAgency, error) {
	orderAgency := models.OrderAgency{
		OrderAgencyId:   "test1",
		OrderAgencyName: "test1Name",
	}

	return orderAgency, nil
}
