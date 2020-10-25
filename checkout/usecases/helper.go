package usecases

import "github.com/mcabezas/ecommerce/checkout/models"

func removeDuplicatedIDs(src []models.WatchID) []models.WatchID {
	// Put into a map to remove duplicates
	watchIDs := make(map[models.WatchID]bool, len(src))
	for _, id := range src {
		watchIDs[id] = true
	}
	// Converts map to slice
	var result = make([]models.WatchID, len(watchIDs))
	var i int
	for id := range watchIDs {
		result[i] = id
		i++
	}
	return result
}

func watchIDsToPurchaseOrder(watchIDs []models.WatchID, catalogue map[models.WatchID]models.WatchCatalogueItem) models.PurchaseOrder {
	watchListItems := make(map[models.WatchID]models.PurchaseOrderItem, len(watchIDs))
	for _, watchID := range watchIDs {
		if item, ok := watchListItems[watchID]; ok {
			item.Qty += 1
			watchListItems[watchID] = item
			continue
		}
		watchListItems[watchID] = models.PurchaseOrderItem{
			WatchID:   watchID,
			Qty:       1,
			UnitPrice: catalogue[watchID].UnitPrice,
		}
	}
	return models.PurchaseOrder{Items: watchListItems}
}
