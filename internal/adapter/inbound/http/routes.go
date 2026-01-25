package http

import (
	"kasir-api/internal/adapter/inbound/http/handler"
	"net/http"
)

type Router struct {
	produkHandler   *handler.ProdukHandler
	categoryHandler *handler.CategoryHandler
}

func NewRouter(produkHandler *handler.ProdukHandler, categoryHandler *handler.CategoryHandler) *Router {
	return &Router{
		produkHandler:   produkHandler,
		categoryHandler: categoryHandler,
	}
}

func (r *Router) RegisterRoutes(mux *http.ServeMux) {
	r.produkHandler.RegisterRoutes(mux)
	r.categoryHandler.RegisterRoutes(mux)
}
