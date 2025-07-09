package handlers

import (
	"campaigns/pkg/mapper"
	"encoding/json"
	"log"
	"net/http"
)

type CampaignsHTTPHandler struct {
	mapper mapper.ICampaignMapper
}

func NewCampaignsHTTPHandler(mapper mapper.ICampaignMapper) *CampaignsHTTPHandler {
	return &CampaignsHTTPHandler{
		mapper: mapper,
	}
}

func (h *CampaignsHTTPHandler) DeliverCampaigns(w http.ResponseWriter, r *http.Request) {
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
	mappedCampaigns := h.mapper.GetTargetedCampaigns(ctx, *req)

	if len(mappedCampaigns) == 0 {
		log.Println("no mapped campaings found for given req!")
		w.WriteHeader(http.StatusNoContent) // HTTP 204
		return
	}

	log.Println("successfully fetched campaigns")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(mappedCampaigns); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *CampaignsHTTPHandler) writeError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
