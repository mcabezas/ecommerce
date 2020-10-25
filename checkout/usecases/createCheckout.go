package usecases

import (
	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/mcabezas/ecommerce/internal/money"
)

type CreateCheckout interface {
	create(watchIDs []models.WatchID) (money.Money, error)
}

func NewCreateCheckout(repository infrastructure.WatchCatalogRepository, discountFinder DiscountFinder) CreateCheckout {
	return &createCheckout{
		repository:     repository,
		discountFinder: discountFinder,
	}
}

type createCheckout struct {
	repository     infrastructure.WatchCatalogRepository
	discountFinder DiscountFinder
}

func (c *createCheckout) create(watchIDs []models.WatchID) (money.Money, error) {
	watchIDsSet := removeDuplicatedIDs(watchIDs)
	watchesCatalog, err := c.repository.GetWatchesCatalogue(watchIDsSet)
	if err != nil {
		return money.Money{}, err
	}

	purchaseOrder := watchIDsToPurchaseOrder(watchIDs, watchesCatalog)
	price := calculateBasePrice(purchaseOrder)

	discounts, err := c.discountFinder.findComboDiscounts(watchIDsSet)
	if err != nil {
		return money.Money{}, err
	}
	discountedAmount, err := calculateDiscountedPrice(purchaseOrder, discounts)
	if err != nil {
		return money.Money{}, err
	}
	return price.Sub(discountedAmount)
}

func calculateBasePrice(purchaseOrder models.PurchaseOrder) money.Money {
	var price float64
	for _, watch := range purchaseOrder.Items {
		price += watch.UnitPrice * float64(watch.Qty)
	}
	return money.Money{Amount: price}
}

func calculateDiscountedPrice(purchaseOrder models.PurchaseOrder, discounts []models.Discount) (money.Money, error) {
	var discountedPrice money.Money
	for _, discount := range discounts {
		calculatedDiscount, err := discount.CalculateDiscount(purchaseOrder)
		if err != nil {
			return money.Money{}, err
		}
		discountedPrice, err = discountedPrice.Plus(calculatedDiscount)
		if err != nil {
			return money.Money{}, err
		}
	}
	return discountedPrice, nil
}
