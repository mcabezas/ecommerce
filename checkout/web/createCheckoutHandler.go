package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/mcabezas/ecommerce/checkout/infrastructure"
	"github.com/mcabezas/ecommerce/checkout/models"
	"github.com/mcabezas/ecommerce/checkout/usecases"
	"github.com/mcabezas/ecommerce/internal/commons/logs"
	"go.uber.org/zap"
)

type Handler struct {
	CreateCheckout usecases.CreateCheckout
}

func NewCreateCheckoutHandler(repository infrastructure.WatchCatalogRepository, discountService usecases.DiscountFinder) *Handler {
	return &Handler{CreateCheckout: usecases.NewCreateCheckout(repository, discountService)}
}

func (h *Handler) CreateCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	watchIDs, err := parseUpdateBody(r)
	if err != nil {
		logs.Sugar().Info("there was an issue reading body", zap.String("requestID", middleware.GetReqID(r.Context())))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	price, err := h.CreateCheckout.Create(watchIDs)
	if err != nil {
		logs.Log().Info("there was an issue creating checkout", zap.String("requestID", middleware.GetReqID(r.Context())))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	marshal, _ := json.Marshal(CheckoutResponse{Price: price.Amount})
	_, _ = w.Write(marshal)
}

func parseUpdateBody(r *http.Request) ([]models.WatchID, error) {
	var watchIDs []models.WatchID
	err := json.NewDecoder(r.Body).Decode(&watchIDs)
	if err != nil {
		return nil, err
	}
	return watchIDs, nil
}

type CheckoutResponse struct {
	Price float64 `json:"price"`
}
