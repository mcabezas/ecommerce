package infrastructure

import "github.com/mcabezas/ecommerce/checkout/models"

type WatchCatalogRepository interface {
	GetWatchesCatalogue(watchIDs []models.WatchID) (map[models.WatchID]models.WatchCatalogueItem, error)
}

