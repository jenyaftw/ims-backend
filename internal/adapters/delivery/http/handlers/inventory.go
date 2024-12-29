package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
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

	inventories, err := h.svc.ListInventories(ctx, 10000, 0)
	if err != nil {
		HandleError(w, err)
		return
	}

	fmt.Println("Inventories:", len(inventories))

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

type InventoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h InventoryHandler) CreateInventory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req InventoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HandleError(w, err)
		return
	}

	newInventory := domain.Inventory{
		Name:        req.Name,
		Description: req.Description,
	}

	inventory, err := h.svc.CreateInventory(ctx, newInventory)
	if err != nil {
		HandleError(w, err)
		return
	}

	res, err := json.Marshal(newInventoryResponse(inventory))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	inventoryIdString := chi.URLParam(r, "id")
	inventoryId, err := uuid.Parse(inventoryIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventory, err := h.svc.GetInventoryById(ctx, inventoryId)
	if err != nil {
		HandleError(w, err)
		return
	}

	res, err := json.Marshal(newInventoryResponse(inventory))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h InventoryHandler) GetInventorySections(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	inventoryIdString := chi.URLParam(r, "id")
	inventoryId, err := uuid.Parse(inventoryIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventory, err := h.svc.GetInventoryById(ctx, inventoryId)
	if err != nil {
		HandleError(w, err)
		return
	}

	res, err := json.Marshal(newSectionResponses(inventory.Sections))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h InventoryHandler) CreateInventorySection(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	inventoryIdString := chi.URLParam(r, "id")
	inventoryId, err := uuid.Parse(inventoryIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventory, err := h.svc.GetInventoryById(ctx, inventoryId)
	if err != nil {
		HandleError(w, err)
		return
	}

	var req InventoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HandleError(w, err)
		return
	}

	newSection := domain.Section{
		Name:        req.Name,
		Description: req.Description,
	}

	section, err := h.svc.CreateInventorySection(ctx, inventory, newSection)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventory.Sections = append(inventory.Sections, section)

	res, err := json.Marshal(newInventoryResponse(inventory))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h InventoryHandler) GetInventoryItemBySKU(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	skuString := r.URL.Query().Get("sku")
	if skuString != "" {
		item, err := h.svc.GetInventoryItemBySKU(ctx, skuString)
		if err != nil {
			HandleError(w, err)
			return
		}

		res, err := json.Marshal(newItemResponse(item))
		if err != nil {
			HandleError(w, err)
			return
		}

		w.Write(res)
		return
	}
}

func (h InventoryHandler) GetInventoryItems(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	sectionIdString := r.URL.Query().Get("section_id")
	var sectionId *uuid.UUID
	if sectionIdString != "" {
		sectionIdParsed, err := uuid.Parse(sectionIdString)
		if err != nil {
			HandleError(w, err)
			return
		}
		sectionId = &sectionIdParsed
	}

	// limit offset default limit 100 offset 0

	offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "100"
	}
	offsetInt, err := strconv.ParseUint(offset, 10, 32)
	if err != nil {
		HandleError(w, err)
		return
	}
	limitInt, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventoryIdString := chi.URLParam(r, "id")
	inventoryId, err := uuid.Parse(inventoryIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventory, err := h.svc.GetInventoryById(ctx, inventoryId)
	if err != nil {
		HandleError(w, err)
		return
	}

	var items []domain.Item
	if sectionId != nil {
		items, err = h.svc.GetInventoryItemsBySection(ctx, inventory.ID, *sectionId, offsetInt, limitInt)
	} else {
		items, err = h.svc.GetInventoryItems(ctx, inventory.ID, offsetInt, limitInt)
	}
	if err != nil {
		HandleError(w, err)
		return
	}

	itemResponses := make([]itemResponse, 0, len(items))
	for _, item := range items {
		itemResponses = append(itemResponses, newItemResponse(item))
	}

	res, err := json.Marshal(itemResponses)
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

type ItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	SectionID   string `json:"section_id"`
}

func (h InventoryHandler) CreateInventoryItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	inventoryIdString := chi.URLParam(r, "id")
	inventoryId, err := uuid.Parse(inventoryIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	inventory, err := h.svc.GetInventoryById(ctx, inventoryId)
	if err != nil {
		HandleError(w, err)
		return
	}

	var req ItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HandleError(w, err)
		return
	}

	sectionId, err := uuid.Parse(req.SectionID)
	if err != nil {
		HandleError(w, err)
		return
	}

	foundSection := false
	for _, section := range inventory.Sections {
		if section.ID == sectionId {
			foundSection = true
			break
		}
	}
	if !foundSection {
		HandleError(w, fmt.Errorf("section not found"))
		return
	}

	newItem := domain.Item{
		Name:        req.Name,
		Description: req.Description,
		Quantity:    req.Quantity,
	}

	item, err := h.svc.CreateInventoryItem(ctx, inventory, domain.Section{ID: sectionId}, newItem)
	if err != nil {
		HandleError(w, err)
		return
	}

	res, err := json.Marshal(newItemResponse(item))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h InventoryHandler) DeleteInventoryItem(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	itemIdString := chi.URLParam(r, "itemID")
	itemId, err := uuid.Parse(itemIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	err = h.svc.DeleteInventoryItem(ctx, itemId)
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write([]byte("ok"))
}
