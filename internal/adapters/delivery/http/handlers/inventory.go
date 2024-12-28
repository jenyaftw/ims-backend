package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type InventoryHandler struct {
	svc ports.InventoryService
}

func NewInventoryHandler(
	svc ports.InventoryService,
) InventoryHandler {
	return InventoryHandler{
		svc: svc,
	}
}

func (h InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	inventories, err := h.svc.ListInventories(ctx, 0, 10)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventoryResponse := make([]inventoryResponse, 0, len(inventories))
	for _, inventory := range inventories {
		inventoryResponse = append(inventoryResponse, newInventoryResponse(inventory))
	}

	res, err := json.Marshal(inventoryResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(res)
}
