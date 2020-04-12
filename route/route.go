package route

import (
	"go-http-server/route/controller"
	"log"
	"net/http"
)

type Route struct {
	Logger     *log.Logger
	controller controller.Controller
}

func New(logger *log.Logger) *Route {
	var controller = controller.Controller{Logger: logger}

	return &Route{Logger: logger, controller: controller}
}

func (r Route) Route(res http.ResponseWriter, req *http.Request) {
	r.Logger.Print("Start of inbound request routing")

	r.controller.Handle(res, req)

	r.Logger.Print("End of inbound request routing")
}
