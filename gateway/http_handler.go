package main

import (
	"log"
	"net/http"
)

type handler struct {
	//gateway
}

func NewHandler() *handler {
	return &handler{}
}
func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
	log.Println("Hello World!")

}

func (h *handler) HandleCreateOrder(writer http.ResponseWriter, request *http.Request) {}
