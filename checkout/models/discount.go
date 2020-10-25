package models

import "github.com/mcabezas/ecommerce/internal/money"

type Discount interface {
	Name() string
	CalculateDiscount(purchaseOrder PurchaseOrder) (money.Money, error)
}

type FixedDiscount struct {
	watchID    WatchID
	units      int
	fixedPrice float64
}

func (ud *FixedDiscount) WatchID() WatchID {
	return ud.watchID
}

func NewFixedDiscount(watchID WatchID, units int, fixedPrice float64) FixedDiscount {
	return FixedDiscount{
		watchID:    watchID,
		units:      units,
		fixedPrice: fixedPrice,
	}
}
func (ud *FixedDiscount) CalculateDiscount(purchaseOrder PurchaseOrder) (money.Money, error) {
	return money.Money{Amount: 0}, nil
}

func (ud *FixedDiscount) Name() string {
	return "Fixed discount"
}
