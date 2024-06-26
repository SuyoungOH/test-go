package routers

import (
	// "fmt"
	// "net/http"
	routerCommon "test-go/routers/common"
	routerOrder "test-go/routers/order"
	routerOrderAgency "test-go/routers/orderAgency"
)

func ApiRouter() *routerCommon.Router {
	// mux := &http.ServeMux{}
	// mux.HandleFunc("GET /test", GetApiHandler)
	// mux.HandleFunc("POST /test", PostApiHandler)
	// mux.HandleFunc("GET /", RootHandler)
	// return mux

	mux := routerCommon.NewRouter()

	routerOrder.OrderApiRouter(mux)

	routerOrderAgency.OrderAgencyApiRouter(mux)

	mux.HandleFunc("GET", "/", RootHandler)
	return mux
}
