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

