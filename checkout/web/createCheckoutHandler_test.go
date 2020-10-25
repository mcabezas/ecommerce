package web

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/mcabezas/ecommerce/checkout/usecases"
	"github.com/mcabezas/ecommerce/internal/commons/logs"
	"github.com/stretchr/testify/assert"
)

func TestCanCalculatePrice(t *testing.T) {
	// setup
	logs.InitDefault()
	handler := NewCreateCheckoutHandler(newWatchCatalogueRepository(), usecases.NewDiscountFinder(newComboRepository()))
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.CreateCheckoutHandler(w, r)
	}))
	response, err := server.Client().Post(server.URL, "", strings.NewReader(`["001"]`))
	// when
	assert.Nil(t, err)
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, `{"price":100}`, string(body))
}

func TestGivenInvalidIDCannotCalculatePrice(t *testing.T) {
	// setup
	logs.InitDefault()
	handler := NewCreateCheckoutHandler(newWatchCatalogueRepository(), usecases.NewDiscountFinder(newComboRepository()))
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.CreateCheckoutHandler(w, r)
	}))
	response, err := server.Client().Post(server.URL, "", strings.NewReader(`["wrongID"]`))
	// when
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func newComboRepository() infrastructure.ComboDiscountRepository {
	return infrastructure.NewInMemoryComboDiscountRepository(
		[]models.ComboDiscount{
			models.NewComboDiscount("001", 3, 200),
			models.NewComboDiscount("002", 2, 120),
		},
	)
}

func newWatchCatalogueRepository() infrastructure.WatchCatalogRepository {
	return infrastructure.NewInMemoryWatchCatalogueRepository(
		[]models.WatchCatalogueItem{
			{ID: "001", Name: "Rolex", UnitPrice: 100},
			{ID: "002", Name: "Michael Kors", UnitPrice: 80},
			{ID: "003", Name: "Swatch", UnitPrice: 50},
			{ID: "004", Name: "Casio", UnitPrice: 30},
		},
	)
}
