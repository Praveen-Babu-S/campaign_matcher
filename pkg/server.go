package campaigns

import (
	"campaigns/pkg/cache"
	"campaigns/pkg/handlers"
	"campaigns/pkg/mapper"
	"campaigns/pkg/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartCampaignsServer(port string) {
	// data store layer
	repo := repository.NewDataStore()
	// cache layer for quick data access
	cache := cache.NewCampaignStore(repo)
	// campaign mapper to get relevant campaigns for req
	mapper := mapper.NewCampaignMapper(cache)
	// http handler layer
	campaignsHandler := handlers.NewCampaignsHTTPHandler(mapper)

	log.Println("repos initialisations are completed....")
	router := mux.NewRouter()
	router.HandleFunc("/v1/delivery", campaignsHandler.DeliverCampaigns).Methods(http.MethodGet)

	log.Printf("starting campaigns service on %s....", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}
