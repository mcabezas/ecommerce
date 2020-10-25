package main

import (
	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/mcabezas/ecommerce/checkout/usecases"
	logs "github.com/mcabezas/ecommerce/internal/commons/logs"
)

func main() {
	logs.InitDefault()
	watchCatalogueRepository := newWatchCatalogueRepository()
	comboRepository := newComboRepository()
	routes := Routes(watchCatalogueRepository, usecases.NewDiscountFinder(comboRepository))
	srv := newServer("9092", routes)
	srv.Start()
}

func newComboRepository() infrastructure.ComboDiscountRepository {
	return infrastructure.NewInMemoryComboDiscountRepository(
		[]models.ComboDiscount{
			models.NewComboDiscount("001", 3, 200),
			models.NewComboDiscount("002", 2, 120),
		},
	)
}

func newWatchCatalogueRepository() infrastructure.WatchCatalogRepository {
	return infrastructure.NewInMemoryWatchCatalogueRepository(
		[]models.WatchCatalogueItem{
			{ID: "001", Name: "Rolex", UnitPrice: 100},
			{ID: "002", Name: "Michael Kors", UnitPrice: 80},
			{ID: "003", Name: "Swatch", UnitPrice: 50},
			{ID: "004", Name: "Casio", UnitPrice: 30},
		},
	)
}
