package infrastructure

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/stretchr/testify/assert"
)

func TestCanGetAWatchCatalogue(t *testing.T) {
	repository := NewInMemoryWatchCatalogueRepository([]models.WatchCatalogueItem{
		{"100", "random", 10},
		{"200", "random2", 20},
		{"300", "random3", 30},
		{"400", "random4", 40},
		{"500", "random5", 50},
	})
	discounts, err := repository.GetWatchesCatalogue([]models.WatchID{"100"})
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 1)
}

func TestCanGetMorThanOneWatchCatalogue(t *testing.T) {
	repository := NewInMemoryWatchCatalogueRepository([]models.WatchCatalogueItem{
		{"100", "random", 10},
		{"200", "random2", 20},
		{"300", "random3", 30},
		{"400", "random4", 40},
		{"500", "random5", 50},
	})
	watches, err := repository.GetWatchesCatalogue([]models.WatchID{"100", "200"})
	assert.Nil(t, err)
	assert.True(t, len(watches)==2)
}

func TestWhenGetInvalidIDThenFails(t *testing.T) {
	repository := NewInMemoryWatchCatalogueRepository([]models.WatchCatalogueItem{
		{"100", "random", 10},
		{"200", "random2", 20},
		{"300", "random3", 30},
		{"400", "random4", 40},
		{"500", "random5", 50},
	})
	_, err := repository.GetWatchesCatalogue([]models.WatchID{"fail", "200"})
	assert.NotNil(t, err)
}