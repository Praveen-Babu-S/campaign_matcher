package internal

import (
	"campaigns/handlers"
	"campaigns/pkg/delivery"
	"campaigns/pkg/mapper"
	"campaigns/pkg/rules"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartCampaignsServer() {
	dataFetcher := rules.NewInMemoryDataFetcher()
	campaignMapper := mapper.NewSimpleCampaignMapper()
	deliveryService := delivery.NewDeliveryService(dataFetcher, campaignMapper)
	campaignsHandler := handlers.NewCampaignsHTTPHandler(deliveryService)

	// Set up HTTP server
	router := mux.NewRouter()
	router.HandleFunc("/v1/delivery", campaignsHandler.FetchCampaigns).Methods(http.MethodGet)

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
