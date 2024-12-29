package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type TransferHandler struct {
	svc ports.TransferService
}

func NewTransferHandler(
	svc ports.TransferService,
) TransferHandler {
	return TransferHandler{
		svc: svc,
	}
}

func (h TransferHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

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

	transfers, err := h.svc.GetTransfers(ctx, offsetInt, limitInt)
	if err != nil {
		HandleError(w, err)
		return
	}

	transferResponses := make([]transferResponse, 0, len(transfers))
	for _, transfer := range transfers {
		transferResponses = append(transferResponses, newTransferResponse(transfer))
	}

	res, err := json.Marshal(transferResponses)
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

type TransferRequest struct {
	ItemID      string `json:"item_id" validate:"required"`
	ToSectionID string `json:"to_section_id" validate:"required"`
	Quantity    uint64 `json:"quantity" validate:"required"`
}

func (h TransferHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var transferRequest TransferRequest
	err := json.NewDecoder(r.Body).Decode(&transferRequest)
	if err != nil {
		HandleError(w, err)
		return
	}

	sectionUuid, err := uuid.Parse(transferRequest.ToSectionID)
	if err != nil {
		HandleError(w, err)
		return
	}

	itemUuid, err := uuid.Parse(transferRequest.ItemID)
	if err != nil {
		HandleError(w, err)
		return
	}

	request := domain.TransferRequest{
		Item: domain.Item{
			ID: itemUuid,
		},
		ToSectionID: sectionUuid,
		Quantity:    transferRequest.Quantity,
	}

	transfer, err := h.svc.CreateTransfer(ctx, request)
	if err != nil {
		HandleError(w, err)
		return
	}

	res, err := json.Marshal(newTransferResponse(transfer))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h TransferHandler) ProcessTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	transferIdString := chi.URLParam(r, "id")
	transferId, err := uuid.Parse(transferIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	transfer, err := h.svc.GetTransferById(ctx, transferId)
	if err != nil {
		HandleError(w, err)
		return
	}

	err = h.svc.ProcessTransfer(ctx, transfer.ID)
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write([]byte("ok"))
}
