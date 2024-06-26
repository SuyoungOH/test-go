package models

type Order struct {
	OrderId       int64  `json:"orderId"`
	OrderAgencyId string `json:"orderAgencyId"`
	Email         string `json:"email"`
}
