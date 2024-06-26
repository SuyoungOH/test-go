package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-go/models"
	orderServices "test-go/services/order"
)

var orderService = &orderServices.OrderService{}

// TODO: func (서비스 정의 필요) GetApiHandler() {}
func GetOrderHandler(w http.ResponseWriter, req *http.Request) {

	msg := "test get order message. url path=" + req.URL.Path

	fmt.Println(msg)
	w.Write([]byte(msg))
}

// TODO: func (서비스 정의 필요) PostApiHandler() {}
func PostOrderHandler(w http.ResponseWriter, r *http.Request) {
	msg := "test post order message. url path=" + r.URL.Path

	fmt.Println(msg)
	o := models.Order{}
	_ = json.NewDecoder(r.Body).Decode(&o)
	fmt.Println("email=" + o.Email)

	order, err := orderService.Create(r.Context())
	if err != nil {
		w.Write([]byte("err...."))
	}

	convertedOrder, _ := json.Marshal(order)
	w.Header().Add("content-type", "application/json")
	w.Write(convertedOrder)
}
