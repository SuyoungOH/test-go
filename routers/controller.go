package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-go/models"
	services "test-go/services/order"
)

var orderService = &services.OrderService{}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	msg := "test root api message. url path=" + req.URL.Path
	// TODO: Service 호출

	fmt.Println(msg)
	w.Write([]byte(msg))
}

// TODO: func (서비스 정의 필요) GetApiHandler() {}
func GetApiHandler(w http.ResponseWriter, req *http.Request) {

	msg := "test api message. url path=" + req.URL.Path

	fmt.Println(msg)
	w.Write([]byte(msg))
}

// TODO: func (서비스 정의 필요) PostApiHandler() {}
func PostApiHandler(w http.ResponseWriter, r *http.Request) {
	msg := "test post api message. url path=" + r.URL.Path

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
