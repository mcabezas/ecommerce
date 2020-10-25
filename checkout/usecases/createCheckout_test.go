package usecases

import (
	"testing"

	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/stretchr/testify/assert"

)

func TestCanCreateSimpleCheckout(t *testing.T) {
	checkout := NewCreateCheckout()
	_, err := checkout.create([]models.WatchID{
		"001",
		"002",
		"001",
		"004",
		"003",
	})
	assert.Nil(t, err)
}

func TestCanCreateCheckoutGivenEmptyWatchList(t *testing.T) {
	checkout := NewCreateCheckout()
	_, err := checkout.create([]models.WatchID{})
	assert.Nil(t, err)
}
