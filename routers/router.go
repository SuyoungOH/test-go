package routers

import (
	"fmt"
	"net/http"
)

// 커스텀 타입 정의
type httpMethod string
type urlPattern string

type routerRules struct {
	methods map[httpMethod]http.Handler
}

type router struct {
	routes map[urlPattern]routerRules
}

// 요청이 들어오면 여기를 탐
func (router *router) ServeHTTP(resw http.ResponseWriter, req *http.Request) {
	fmt.Println("router ServeHTTP")
	rules, isExists := router.routes[urlPattern(req.URL.Path)]
	if !isExists {
		http.NotFound(resw, req)
		return
	}

	handler, isExists := rules.methods[httpMethod(req.Method)]
	if !isExists {
		// TODO: Method Not Allowed 응답 필요
		http.NotFound(resw, req)
		return
	}

	handler.ServeHTTP(resw, req)
}

// 라우터 핸들러 초기화: router map 에 없으면 핸들러 넣어줌
func (router *router) HandleFunc(method httpMethod, pattern urlPattern, fn http.HandlerFunc) {
	fmt.Println("router HandleFunc")

	rules, isExists := router.routes[pattern]
	if !isExists {
		rules = routerRules{methods: make(map[httpMethod]http.Handler)}
		router.routes[pattern] = rules
	}

	rules.methods[method] = http.HandlerFunc(fn)
}

func NewRouter() *router {
	// make(): map 초기화
	return &router{routes: make(map[urlPattern]routerRules)}
}

func ApiRouter() *router {
	// mux := &http.ServeMux{}
	// mux.HandleFunc("GET /test", GetApiHandler)
	// mux.HandleFunc("POST /test", PostApiHandler)
	// mux.HandleFunc("GET /", RootHandler)
	// return mux

	mux := NewRouter()
	mux.HandleFunc("GET", "/test", GetApiHandler)
	mux.HandleFunc("POST", "/test", PostApiHandler)
	mux.HandleFunc("GET", "/", RootHandler)
	return mux
}
