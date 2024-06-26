package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"test-go/models"
	orderAgencyServices "test-go/services/orderAgency"
)

var orderAgencyService = &orderAgencyServices.OrderAgencyService{}

func PostOrderAgencyHandler(w http.ResponseWriter, r *http.Request) {
	msg := "test post order agency message. url paht=" + r.URL.Path
	fmt.Println(msg)

	oa := models.OrderAgency{}
	_ = json.NewDecoder(r.Body).Decode(&oa)

	orderAgency, err := orderAgencyService.Create(r.Context())
	if err != nil {
		w.Write([]byte("err...."))
	}

	convertedOrderAgency, _ := json.Marshal(orderAgency)
	w.Header().Add("content-type", "application/json")
	w.Write(convertedOrderAgency)
}
