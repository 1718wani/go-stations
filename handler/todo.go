package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// ServeHTTP handles the HTTP requests for TODOHandler.
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleCreate(w, r)
	case http.MethodPut:
		h.handleUpdate(w, r)
	case http.MethodGet:
        h.handleRead(w, r) 
	case http.MethodDelete:
		h.handleDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleCreate handles the endpoint that creates the TODO.
func (h *TODOHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTODORequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Subject == "" {
		http.Error(w, "Subject is required", http.StatusBadRequest)
		return
	}

	todo, err := h.svc.CreateTODO(r.Context(), req.Subject, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := model.CreateTODOResponse{TODO: *todo}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

// handleUpdate handles the endpoint that updates the TODO.
func (h *TODOHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateTODORequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Subject == "" {
		http.Error(w, "Subject is required", http.StatusBadRequest)
		return
	}

	todo, err := h.svc.UpdateTODO(r.Context(), req.ID, req.Subject, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := model.UpdateTODOResponse{TODO: *todo}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

// handleRead handles the endpoint that reads the TODOs.
func (h *TODOHandler) handleRead(w http.ResponseWriter, r *http.Request) {
	prevID, _ := strconv.ParseInt(r.URL.Query().Get("prev_id"), 10, 64)
	size, _ := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)

	req := &model.ReadTODORequest{
		PrevID: int(prevID),
		Size:   int(size),
	}

	todos, err := h.svc.ReadTODO(r.Context(), int64(req.PrevID), int64(req.Size))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert []*model.TODO to []model.TODO
	todoValues := make([]model.TODO, len(todos))
	for i, todo := range todos {
		todoValues[i] = *todo
	}

	res := &model.ReadTODOResponse{TODOs: todoValues}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

// handleDelete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	var req model.DeleteTODORequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(req.IDs) == 0 {
		http.Error(w, "IDs are required", http.StatusBadRequest)
		return
	}

	// Convert []int to []int64
	ids := make([]int64, len(req.IDs))
	for i, id := range req.IDs {
		ids[i] = int64(id)
	}

	err := h.svc.DeleteTODO(r.Context(), ids)
	if err != nil {
		if _, ok := err.(*model.ErrNotFound); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	

	res := model.DeleteTODOResponse{}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
