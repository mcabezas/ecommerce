package usecases

import (
	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/mcabezas/ecommerce/internal/money"
)

type CreateCheckout interface {
	create(watchIDs []models.WatchID) (money.Money, error)
}

func NewCreateCheckout(repository infrastructure.WatchCatalogRepository) CreateCheckout {
	return &createCheckout{
		repository: repository,
	}
}

type createCheckout struct {
	repository infrastructure.WatchCatalogRepository
}

func (c *createCheckout) create(watchIDs []models.WatchID) (money.Money, error) {
	watchIDsSet := removeDuplicatedIDs(watchIDs)
	watchesCatalog, err := c.repository.GetWatchesCatalogue(watchIDsSet)
	if err != nil {
		return money.Money{}, err
	}

	purchaseOrder := watchIDsToPurchaseOrder(watchIDs, watchesCatalog)
	price := c.calculateBasePrice(purchaseOrder)
	return money.Money{Amount: price}, nil
}

func (c *createCheckout) calculateBasePrice(purchaseOrder models.PurchaseOrder) float64 {
	var price float64
	for _, watch := range purchaseOrder.Items {
		price += watch.UnitPrice * float64(watch.Qty)
	}
	return price
}
