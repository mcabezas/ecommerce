package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/usecases"
	"github.com/mcabezas/ecommerce/checkout/web"
)

func Routes(repository infrastructure.WatchCatalogRepository, discountService usecases.DiscountFinder) *chi.Mux {
	mux := chi.NewMux()
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
	)
	mux.Mount("/", checkout(repository, discountService))
	return mux
}

func checkout(repository infrastructure.WatchCatalogRepository, discountService usecases.DiscountFinder) *chi.Mux {
	handler := web.NewCreateCheckoutHandler(repository, discountService)
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Post("/checkout", handler.CreateCheckoutHandler)
	})
	return router
}
