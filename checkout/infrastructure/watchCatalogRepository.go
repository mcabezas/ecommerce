package infrastructure

import "github.com/mcabezas/ecommerce/checkout/models"

// WatchCatalogRepository is a dedicated repository interface for WatchCatalogue
type WatchCatalogRepository interface {
	GetWatchesCatalogue(watchIDs []models.WatchID) (map[models.WatchID]models.WatchCatalogueItem, error)
}

