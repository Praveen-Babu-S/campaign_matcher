package internal

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartCampaignsServer() {
	// Initialize components
	dataFetcher := NewInMemoryDataFetcher()
	campaignMatcher := NewSimpleCampaignMatcher()
	deliveryService := NewDeliveryService(dataFetcher, campaignMatcher)
	deliveryHandler := NewDeliveryHTTPHandler(deliveryService)

	// Set up HTTP server
	router := mux.NewRouter()
	router.HandleFunc("/v1/delivery", deliveryHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Starting campaigns service on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
	}
}
