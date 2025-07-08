package delivery

import (
	"campaigns/models"
	"campaigns/pkg/mapper"
	"campaigns/pkg/rules"
	"context"
	"log"
	"time"
)

type DeliveryService interface {
	GetCampaigns(ctx context.Context, req models.DeliveryRequest) ([]models.DeliveryResponse, error)
}

// DeliveryServiceImpl implements the DeliveryService interface.
type DeliveryServiceImpl struct {
	dataFetcher     rules.DataFetcher
	campaignMatcher mapper.CampaignMapper
	campaignStore   *models.CampaignStore // In-memory store
}

// NewDeliveryService creates a new DeliveryServiceImpl.
func NewDeliveryService(dataFetcher rules.DataFetcher, campaignMatcher mapper.CampaignMapper) *DeliveryServiceImpl {
	ds := &DeliveryServiceImpl{
		dataFetcher:     dataFetcher,
		campaignMatcher: campaignMatcher,
		campaignStore:   models.NewCampaignStore(),
	}
	// Start a goroutine to periodically refresh campaign data
	go ds.refreshCampaignData()
	return ds
}

// refreshCampaignData periodically fetches updated campaign and rule data from the data source.
func (ds *DeliveryServiceImpl) refreshCampaignData() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5-second timeout for data fetch
		campaigns, err := ds.dataFetcher.FetchCampaigns(ctx)
		if err != nil {
			log.Printf("Error refreshing campaigns: %v", err)
			cancel()
			continue
		}

		rules, err := ds.dataFetcher.FetchTargetingRules(ctx)
		if err != nil {
			log.Printf("Error refreshing targeting rules: %v", err)
			cancel()
			continue
		}
		cancel()

		ds.campaignStore.UpdateCampaigns(campaigns, rules)
		log.Println("Campaign data refreshed successfully.")
	}
}

// GetCampaigns retrieves campaigns matching the delivery request.
func (ds *DeliveryServiceImpl) GetCampaigns(ctx context.Context, req models.DeliveryRequest) ([]models.DeliveryResponse, error) {
	campaigns, rules := ds.campaignStore.GetCampaignsAndRules()

	matched := ds.campaignMatcher.Match(req, campaigns, rules)
	return matched, nil
}
