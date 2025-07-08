package rules

import (
	"campaigns/models"
	"context"
)

type DataFetcher interface {
	FetchCampaigns(ctx context.Context) (map[string]models.Campaign, error)
	FetchTargetingRules(ctx context.Context) (map[string]models.TargetingRule, error)
}
