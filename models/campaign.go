package models

import "sync"

type CampaignStatus int

const (
	CampaignStatusActive CampaignStatus = iota
	CampaignStatusInActive
)

var (
	CampaignStatusMap = map[CampaignStatus]string{
		CampaignStatusActive:   "ATCIVE",
		CampaignStatusInActive: "INATCIVE",
	}
)

type Campaign struct {
	CampaignId     string     `json:"campaignId"`
	CampaignName   string     `json:"campaignName"`
	Creatives      []Creative `json:"creatives"`
	CTA            string     `json:"cta"`
	CampaignStatus string     `json:"campaignStatus"`
}

type Creative struct {
	ImageUrl string `json:"imageUrl"`
}

type CampaignStore struct {
	mu             sync.RWMutex // For concurrent access
	campaigns      map[string]Campaign
	targetingRules map[string]TargetingRule // Key is CampaignID
}

func NewCampaignStore() *CampaignStore {
	return &CampaignStore{
		campaigns:      make(map[string]Campaign),
		targetingRules: make(map[string]TargetingRule),
	}
}

func (cs *CampaignStore) UpdateCampaigns(campaigns map[string]Campaign, rules map[string]TargetingRule) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.campaigns = campaigns
	cs.targetingRules = rules
}

// GetCampaignsAndRules retrieves the current campaigns and rules.
func (cs *CampaignStore) GetCampaignsAndRules() (map[string]Campaign, map[string]TargetingRule) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	campaignsCopy := make(map[string]Campaign)
	for k, v := range cs.campaigns {
		campaignsCopy[k] = v
	}
	rulesCopy := make(map[string]TargetingRule)
	for k, v := range cs.targetingRules {
		rulesCopy[k] = v
	}
	return campaignsCopy, rulesCopy
}
