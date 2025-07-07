package campaignsfetcher

import (
	"campaigns/models"
	"context"
)

type DataFetcher interface {
	FetchCampaigns(ctx context.Context) (map[string]models.Campaign, error)
	FetchTargetingRules(ctx context.Context) (map[string]models.TargetingRule, error)
}

type CampaignMatcher interface {
	Match(req models.DeliveryRequest, campaigns map[string]models.Campaign, rules map[string]models.TargetingRule) []models.DeliveryResponse
}

type DeliveryService interface {
	GetCampaigns(ctx context.Context, req models.DeliveryRequest) ([]models.DeliveryResponse, error)
}
