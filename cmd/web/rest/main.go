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
	srv := newServer("9090", routes)
	srv.Start()
}

func newComboRepository() infrastructure.ComboDiscountRepository {
	return infrastructure.NewInMemoryComboDiscountRepository(
		[]models.ComboDiscount{},
	)
}

func newWatchCatalogueRepository() infrastructure.WatchCatalogRepository {
	return infrastructure.NewInMemoryWatchCatalogueRepository(
		[]models.WatchCatalogueItem{},
	)
}
