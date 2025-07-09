package repository

import (
	"campaigns/pkg/models"
	"context"
	"encoding/json"
	"fmt"
)

type Repository interface {
	FetchCampaigns(ctx context.Context) (map[string]models.Campaign, error)
	FetchTargetingRules(ctx context.Context) (map[string]models.TargetingRule, error)
}

type DataStore struct{}

func NewDataStore() Repository {
	return &DataStore{}
}

func (d *DataStore) FetchCampaigns(ctx context.Context) (map[string]models.Campaign, error) {
	campaigns := []models.Campaign{}
	err := json.Unmarshal([]byte(JsonCampaigns), &campaigns)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	campaignsMap := make(map[string]models.Campaign)
	for i := range campaigns {
		campaignsMap[campaigns[i].CampaignId] = campaigns[i]
	}
	return campaignsMap, nil
}

func (d *DataStore) FetchTargetingRules(ctx context.Context) (map[string]models.TargetingRule, error) {
	rules := []models.TargetingRule{}
	err := json.Unmarshal([]byte(JsonRules), &rules)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
	}
	rulesMap := make(map[string]models.TargetingRule)
	for _, r := range rules {
		rulesMap[r.CampaignID] = r
	}
	return rulesMap, nil
}
