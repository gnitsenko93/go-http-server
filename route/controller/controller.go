package controller

import (
	"log"
	"net/http"
)

type Controller struct {
	Logger *log.Logger
}

func (c Controller) Handle(res http.ResponseWriter, req *http.Request) {
	c.Logger.Println("Start of inbound request handling")
	c.Logger.Println("End of inbound request handling")
}
