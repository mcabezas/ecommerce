package usecases

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/stretchr/testify/assert"
)

func TestGivenAEmptyListCanFindAUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.FixedDiscount{
		models.NewFixedDiscount("001", 3, 200),
		models.NewFixedDiscount("002", 3, 120),
	})
	discounts, err := finder.findFixedDiscounts([]models.WatchID{})
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 0)
}

func TestCanFindSingleUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.FixedDiscount{
		models.NewFixedDiscount("001", 3, 200),
		models.NewFixedDiscount("002", 3, 120),
		models.NewFixedDiscount("003", 0, 120),
	})
	discounts, err := finder.findFixedDiscounts(
		[]models.WatchID{"002"},
	)
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 1)
}

func TestCanFindMultipleUnitDiscount(t *testing.T) {
	finder := buildDiscountFinder([]models.FixedDiscount{
		models.NewFixedDiscount("001", 3, 200),
		models.NewFixedDiscount("002", 3, 120),
		models.NewFixedDiscount("003", 0, 120),
	})
	discounts, err := finder.findFixedDiscounts(
		[]models.WatchID{"002", "003"},
	)
	assert.Nil(t, err)
	assert.True(t, len(discounts) == 2)
}

func buildDiscountFinder(fixedDiscounts []models.FixedDiscount) DiscountFinder {
	return NewDiscountFinder(
		infrastructure.NewInMemoryFixedDiscountRepository(fixedDiscounts),
	)
}
