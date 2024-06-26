package routers

import (
	routerCommon "test-go/routers/common"
)

func OrderApiRouter(mux *routerCommon.Router) {
	mux.HandleFunc("GET", "/order/test", GetOrderHandler)
	mux.HandleFunc("POST", "/order/test", PostOrderHandler)
}
