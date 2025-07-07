package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type DeliveryHTTPHandler struct {
	deliveryService DeliveryService
}

func NewDeliveryHTTPHandler(service DeliveryService) *DeliveryHTTPHandler {
	return &DeliveryHTTPHandler{deliveryService: service}
}

func (h *DeliveryHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	appID := r.URL.Query().Get("app")
	country := r.URL.Query().Get("country")
	os := r.URL.Query().Get("os")

	// Validate required parameters
	if appID == "" {
		h.writeErrorResponse(w, "missing app param", http.StatusBadRequest)
		return
	}
	if country == "" {
		h.writeErrorResponse(w, "missing country param", http.StatusBadRequest)
		return
	}
	if os == "" {
		h.writeErrorResponse(w, "missing os param", http.StatusBadRequest)
		return
	}

	req := DeliveryRequest{
		AppID:   appID,
		Country: country,
		OS:      os,
	}

	ctx := r.Context() // Use request context for cancellation/timeouts

	campaigns, err := h.deliveryService.GetCampaigns(ctx, req)
	if err != nil {
		log.Printf("Error getting campaigns: %v", err)
		h.writeErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if len(campaigns) == 0 {
		w.WriteHeader(http.StatusNoContent) // HTTP 204
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // HTTP 200
	if err := json.NewEncoder(w).Encode(campaigns); err != nil {
		log.Printf("Error encoding response: %v", err)
		// Best effort to write error, but response body might already be partially written
	}
}

// writeErrorResponse writes an error message with the given status code.
func (h *DeliveryHTTPHandler) writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
