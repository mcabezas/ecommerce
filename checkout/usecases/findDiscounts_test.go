package usecases

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/mcabezas/ecommerce/internal/money"
	"github.com/stretchr/testify/assert"
)

func TestGivenAEmptyListCanFindAUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
	})
	discounts, err := finder.findComboDiscounts([]models.WatchID{})
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 0)
}

func TestCanFindSingleUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
		models.NewComboDiscount("003", 0, 120),
	})
	discounts, err := finder.findComboDiscounts(
		[]models.WatchID{"002"},
	)
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 1)
}

func TestCanFindMultipleUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
		models.NewComboDiscount("003", 0, 120),
	})
	discounts, err := finder.findComboDiscounts(
		[]models.WatchID{"002", "003"},
	)
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 2)
}

func TestCanCalculateZeroDiscount(t *testing.T) {
	// given
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
		models.NewComboDiscount("003", 0, 120),
	})
	watchIDs := []models.WatchID{"002", "003"}
	watchIDsSet := removeDuplicatedIDs(watchIDs)
	discounts, err := finder.findComboDiscounts(watchIDsSet)

	purchaseOrder := watchIDsToPurchaseOrder(watchIDs, map[models.WatchID]models.WatchCatalogueItem{
		"002": {ID: "002", Name: "Random2", UnitPrice: 80},
		"003": {ID: "003", Name: "Random3", UnitPrice: 50},
	})
	// when
	var totalDiscount money.Money
	for _, discount := range discounts {
		calculatedDiscount, err := discount.CalculateDiscount(purchaseOrder)
		assert.Nil(t, err)
		totalDiscount.Amount += calculatedDiscount.Amount
	}
	// then
	assert.Nil(t, err)
	assert.True(t, totalDiscount.Amount == float64(0))
}

func TestCanCalculateADiscountDiscount(t *testing.T) {
	// given
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 3, 200),
		models.NewComboDiscount("002", 3, 120),
		models.NewComboDiscount("003", 0, 120),
	})
	watchIDs := []models.WatchID{"002", "003", "002", "001", "002"}
	watchIDsSet := removeDuplicatedIDs(watchIDs)
	discounts, err := finder.findComboDiscounts(watchIDsSet)

	purchaseOrder := watchIDsToPurchaseOrder(watchIDs, map[models.WatchID]models.WatchCatalogueItem{
		"001": {ID: "001", Name: "Random1", UnitPrice: 80},
		"002": {ID: "002", Name: "Random2", UnitPrice: 80},
		"003": {ID: "003", Name: "Random3", UnitPrice: 50},
	})
	// when
	var totalDiscount money.Money
	for _, discount := range discounts {
		calculatedDiscount, err := discount.CalculateDiscount(purchaseOrder)
		assert.Nil(t, err)
		totalDiscount.Amount += calculatedDiscount.Amount
	}
	// then
	assert.Nil(t, err)
	assert.True(t, totalDiscount.Amount == float64(120))
}

func TestCanCalculateMultipleDiscountDiscount(t *testing.T) {
	// given
	finder := buildDiscountFinder([]models.ComboDiscount{
		models.NewComboDiscount("001", 4, 200),
		models.NewComboDiscount("002", 3, 120),
		models.NewComboDiscount("003", 0, 120),
	})
	// There are:
	// 4 watchID "002" => 120 discount
	// 2 watchID "003  => 0 discount"
	// 5 watchID "001" => 200 discount
	watchIDs := []models.WatchID{"002", "002", "003", "003", "002", "001", "002", "001", "001", "001", "001"}
	watchIDsSet := removeDuplicatedIDs(watchIDs)
	discounts, err := finder.findComboDiscounts(watchIDsSet)

	purchaseOrder := watchIDsToPurchaseOrder(watchIDs, map[models.WatchID]models.WatchCatalogueItem{
		"001": {ID: "001", Name: "Random1", UnitPrice: 100},
		"002": {ID: "002", Name: "Random2", UnitPrice: 80},
		"003": {ID: "003", Name: "Random3", UnitPrice: 50},
	})
	// when
	var totalDiscount money.Money
	for _, discount := range discounts {
		calculatedDiscount, err := discount.CalculateDiscount(purchaseOrder)
		assert.Nil(t, err)
		totalDiscount.Amount += calculatedDiscount.Amount
	}
	// then
	assert.Nil(t, err)
	assert.True(t, totalDiscount.Amount == float64(320))
}

func buildDiscountFinder(comboDiscounts []models.ComboDiscount) DiscountFinder {
	return NewDiscountFinder(
		infrastructure.NewInMemoryComboDiscountRepository(comboDiscounts),
	)
}
