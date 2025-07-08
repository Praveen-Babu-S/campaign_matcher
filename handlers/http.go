package handlers

import (
	"campaigns/pkg/delivery"
	"encoding/json"
	"log"
	"net/http"
)

type CampaignsHTTPHandler struct {
	deliveryService delivery.DeliveryService
}

func NewCampaignsHTTPHandler(service delivery.DeliveryService) *CampaignsHTTPHandler {
	return &CampaignsHTTPHandler{deliveryService: service}
}

func (h *CampaignsHTTPHandler) FetchCampaigns(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, err := fetchAndValidateReqParams(r.URL.Query())
	if err != nil {
		log.Printf("query param validation failed:%s", err.Error())
		h.writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	campaigns, err := h.deliveryService.GetCampaigns(ctx, *req)
	if err != nil {
		log.Printf("error getting campaigns: %s", err.Error())
		h.writeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if len(campaigns) == 0 {
		w.WriteHeader(http.StatusNoContent) // HTTP 204
		return
	}
	log.Println("successfully fetched campaigns")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(campaigns); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *CampaignsHTTPHandler) writeError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
