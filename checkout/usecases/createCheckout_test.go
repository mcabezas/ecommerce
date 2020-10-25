package usecases

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateSimpleCheckout(t *testing.T) {
	checkout := buildCreateCheckout(EmptyWatchCatalogue())
	_, err := checkout.create([]models.WatchID{
		"001",
		"002",
		"001",
		"004",
		"003",
	})
	assert.Nil(t, err)
}

func TestCanCreateCheckoutGivenEmptyWatchList(t *testing.T) {
	checkout := buildCreateCheckout(EmptyWatchCatalogue())
	_, err := checkout.create([]models.WatchID{})
	assert.Nil(t, err)
}

func TestCanCreateCheckoutWithoutDiscounts(t *testing.T) {
	checkout := buildCreateCheckout(
		[]models.WatchCatalogueItem{
			{ID: "001", Name: "Rolex", UnitPrice: 100},
			{ID: "002", Name: "Michael Kors", UnitPrice: 80},
			{ID: "003", Name: "Swatch", UnitPrice: 50},
			{ID: "004", Name: "Casio", UnitPrice: 30},
		},
	)
	price, err := checkout.create([]models.WatchID{
		"001",
		"002",
		"001",
		"004",
		"003",
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(360), price.Amount)
}

func buildCreateCheckout(watchCatalog []models.WatchCatalogueItem) CreateCheckout {
	return NewCreateCheckout(
		infrastructure.NewInMemoryWatchCatalogueRepository(
			watchCatalog,
		),
	)
}

func EmptyWatchCatalogue() []models.WatchCatalogueItem {
	return []models.WatchCatalogueItem{}
}
