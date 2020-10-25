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
	if ud.units == 0 {
		return money.Money{Amount: 0}, nil
	}
	originalPrice := float64(purchaseOrder.Items[ud.watchID].Qty) * purchaseOrder.Items[ud.watchID].UnitPrice

	combos := purchaseOrder.Items[ud.watchID].Qty / ud.units
	outOfCombos := purchaseOrder.Items[ud.watchID].Qty % ud.units
	discountedPrice := float64(combos)*ud.fixedPrice +
		float64(outOfCombos)*purchaseOrder.Items[ud.watchID].UnitPrice

	return money.Money{Amount: originalPrice - discountedPrice}, nil
}

func (ud *ComboDiscount) Name() string {
	return "Combo discount"
}
