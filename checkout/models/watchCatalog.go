package models

type WatchID string

type WatchCatalogueItem struct {
	ID        WatchID
	Name      string
	UnitPrice float64
}
