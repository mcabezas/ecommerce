package models

// WatchID models a watch id type
type WatchID string

// WatchCatalogueItem models a watch catalogue entity
type WatchCatalogueItem struct {
	ID        WatchID
	Name      string
	UnitPrice float64
}

// PurchaseOrder models a purchase order that will be use by checkout
type PurchaseOrder struct {
	Items map[WatchID]PurchaseOrderItem
}

// PurchaseOrderItem models an item of PurchaseOrder
type PurchaseOrderItem struct {
	WatchID
	Qty       int
	UnitPrice float64
}
