package usecases

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateSimpleCheckout(t *testing.T) {
	checkout := buildCreateCheckout(EmptyWatchCatalogue(), buildDiscountFinder([]models.ComboDiscount{}))
	_, err := checkout.Create([]models.WatchID{
		"001",
		"002",
		"001",
		"004",
		"003",
	})
	assert.Nil(t, err)
}

func TestCanCreateCheckoutGivenEmptyWatchList(t *testing.T) {
	checkout := buildCreateCheckout(EmptyWatchCatalogue(), buildDiscountFinder([]models.ComboDiscount{}))
	_, err := checkout.Create([]models.WatchID{})
	assert.Nil(t, err)
}

func TestCanCreateCheckoutWithoutDiscounts(t *testing.T) {
	checkout := buildCreateCheckout(
		[]models.WatchCatalogueItem{
			{ID: "001", Name: "Rolex", UnitPrice: 100},
			{ID: "002", Name: "Michael Kors", UnitPrice: 80},
			{ID: "003", Name: "Swatch", UnitPrice: 50},
			{ID: "004", Name: "Casio", UnitPrice: 30},
		}, buildDiscountFinder([]models.ComboDiscount{}),
	)
	price, err := checkout.Create([]models.WatchID{
		"001",
		"002",
		"001",
		"004",
		"003",
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(360), price.Amount)
}

func TestCanCreateCheckoutWithDiscounts(t *testing.T) {
	checkout := buildCreateCheckout(
		[]models.WatchCatalogueItem{
			{ID: "001", Name: "Rolex", UnitPrice: 100},
			{ID: "002", Name: "Michael Kors", UnitPrice: 80},
			{ID: "003", Name: "Swatch", UnitPrice: 50},
			{ID: "004", Name: "Casio", UnitPrice: 30},
		}, buildDiscountFinder([]models.ComboDiscount{
			models.NewComboDiscount("001", 3, 200),
			models.NewComboDiscount("002", 2, 120),
		}),
	)
	price, err := checkout.Create([]models.WatchID{
		"001",
		"002",
		"002",
		"001",
		"004",
		"003",
		"001",
		"001",
	})
	assert.Nil(t, err)
	assert.Equal(t, float64(500), price.Amount)
}

func buildCreateCheckout(watchCatalog []models.WatchCatalogueItem, discountFinder DiscountFinder) CreateCheckout {
	return NewCreateCheckout(
		infrastructure.NewInMemoryWatchCatalogueRepository(
			watchCatalog,
		),
		discountFinder,
	)
}

func EmptyWatchCatalogue() []models.WatchCatalogueItem {
	return []models.WatchCatalogueItem{}
}
