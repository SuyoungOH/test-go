package routers

import (
	routerCommon "test-go/routers/common"
)

func OrderAgencyApiRouter(mux *routerCommon.Router) {
	mux.HandleFunc("POST", "/order-agency/test", PostOrderAgencyHandler)
}
