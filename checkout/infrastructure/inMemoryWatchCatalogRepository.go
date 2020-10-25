package infrastructure

import (
	"sync"

	"github.com/mcabezas/ecommerce/checkout/models"
)

type inMemoryWatchCatalogue struct {
	watchCatalog *sync.Map
}

// NewInMemoryWatchCatalogueRepository ...
func NewInMemoryWatchCatalogueRepository(watches []models.WatchCatalogueItem) WatchCatalogRepository {
	watchCatalogue := &sync.Map{}
	for _, watch := range watches {
		watchCatalogue.Store(watch.ID, watch)
	}
	return &inMemoryWatchCatalogue{watchCatalog: watchCatalogue}
}

// GetWatchesCatalogue retrieves a map containing the requested models.WatchCatalogueItem
func (m *inMemoryWatchCatalogue) GetWatchesCatalogue(watchIDs []models.WatchID) (map[models.WatchID]models.WatchCatalogueItem, error) {
	var result = make(map[models.WatchID]models.WatchCatalogueItem, len(watchIDs))
	for _, watchID := range watchIDs {
		if watch, ok := m.watchCatalog.Load(watchID); ok {
			result[watchID] = watch.(models.WatchCatalogueItem)
		}
	}
	return result, nil
}
