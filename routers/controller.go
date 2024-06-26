package routers

import (
	"fmt"
	"net/http"
)

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
