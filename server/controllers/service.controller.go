package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"server/models"
	"server/storage/servicestore"
	"time"
)

type ServiceHandler struct {
	store servicestore.ServiceStore
}

func NewServiceHandler(store *servicestore.MongoStore) *ServiceHandler {
	return &ServiceHandler{store: store}
}

func (h *ServiceHandler) CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var service models.CreateService
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newService := models.CreateService{
		ServiceId:   primitive.NewObjectID(),
		ServiceName: service.ServiceName,
		CreatedAt:   time.Now()}

	err := h.store.CreateService(r.Context(), (*models.DBService)(&newService))
	if err != nil {
		http.Error(w, "Failed to create service", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Service created successfully"}); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
	}
}

func (h *ServiceHandler) GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	services, err := h.store.GetAllService(context.Background())
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(services); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ServiceHandler) UpdateServiceHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var service models.UpdateService
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if service.ServiceId == primitive.NilObjectID {
		http.Error(w, "Service ID is required", http.StatusBadRequest)
		return
	}

	err := h.store.UpdateService(context.Background(), &service)
	if err != nil {
		http.Error(w, "Failed to update service", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Service updated successfully"}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
