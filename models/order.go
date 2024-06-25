package models

type Order struct {
	OrderId       int64
	OrderAgencyId string
	Email         string "json:email"
}
