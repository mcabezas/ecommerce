package models

import "github.com/mcabezas/ecommerce/internal/money"

type Discount interface {
	Name() string
	CalculateDiscount(purchaseOrder PurchaseOrder) (money.Money, error)
}

type ComboDiscount struct {
	watchID    WatchID
	units      int
	fixedPrice float64
}

func (ud *ComboDiscount) WatchID() WatchID {
	return ud.watchID
}

func NewComboDiscount(watchID WatchID, units int, fixedPrice float64) ComboDiscount {
	return ComboDiscount{
		watchID:    watchID,
		units:      units,
		fixedPrice: fixedPrice,
	}
}
func (ud *ComboDiscount) CalculateDiscount(purchaseOrder PurchaseOrder) (money.Money, error) {
	return money.Money{Amount: 0}, nil
}

func (ud *ComboDiscount) Name() string {
	return "Combo discount"
}
