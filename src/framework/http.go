package framework

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type HTTPSystem struct {
	Port   string
	Router *routing.Router
}

func (h *HTTPSystem) init() {
	h.Router = routing.New()
	h.routeSysRoute()
}

func (h *HTTPSystem) listen() {
	fmt.Printf("\033[1;34m%s\033[0m", "[HTTP] ")
	fmt.Printf("Listening on port %+v\n", h.Port)
	panic(fasthttp.ListenAndServe(":"+h.Port, h.Router.HandleRequest))
}

func (h *HTTPSystem) routeSysRoute() {
	h.Router.Get("/_/sys", func(c *routing.Context) error {
		fmt.Fprintf(c, "HTTP System functional!")
		return nil
	})

}
