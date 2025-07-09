package cache

import (
	"campaigns/pkg/models"
	"campaigns/pkg/repository"
	"context"
	"log"
	"sync"
	"time"
)

type ICacher interface {
	GetCampaign(id string) (models.Campaign, bool)
	GetRules(id string) (models.TargetingRule, bool)
	GetActiveCampaignIds() []string
}

type CampaignStore struct {
	Mu                sync.RWMutex // For concurrent access
	Campaigns         map[string]models.Campaign
	TargetingRules    map[string]models.TargetingRule // Key is CampaignID
	ActiveCampaignIds []string
}

type CampsignCacher struct {
	Cache *CampaignStore
	repo  repository.Repository
}

func NewCampaignStore(repo repository.Repository) ICacher {
	cache := &CampsignCacher{
		Cache: &CampaignStore{
			Campaigns:         make(map[string]models.Campaign),
			TargetingRules:    make(map[string]models.TargetingRule),
			ActiveCampaignIds: make([]string, 0),
		},
		repo: repo,
	}
	ctx := context.Background()
	cache.RefreshCache(ctx)
	go cache.refreshTask(ctx, 30*time.Second)
	return cache
}

// RefreshCache reloads data from the data store into the cache.
func (c *CampsignCacher) RefreshCache(ctx context.Context) {
	log.Println("Refreshing cache...")
	newCampaigns := make(map[string]models.Campaign)
	newRules := make(map[string]models.TargetingRule)

	// fetching from Data Store
	dbCampaigns, err := c.repo.FetchCampaigns(ctx)
	if err != nil {
		log.Printf("error at fetchign db campaigns:%s\n", err.Error())
		return
	}
	dbRules, err := c.repo.FetchTargetingRules(ctx)
	if err != nil {
		log.Printf("error at fetchign db rules:%s\n", err.Error())
		return
	}

	for _, campaign := range dbCampaigns {
		if campaign.CampaignStatus == "ACTIVE" {
			c.Cache.ActiveCampaignIds = append(c.Cache.ActiveCampaignIds, campaign.CampaignId)
		}
		newCampaigns[campaign.CampaignId] = campaign
	}

	for _, rule := range dbRules {
		newRules[rule.CampaignID] = rule
	}

	// Safely update the cache
	c.Cache.Mu.RLock()
	c.Cache.Campaigns = newCampaigns
	c.Cache.TargetingRules = newRules
	c.Cache.Mu.RUnlock()
	log.Println("Cache refreshed successfully.")
}

// startRefreshRoutine periodically refreshes the cache.
func (c *CampsignCacher) refreshTask(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.RefreshCache(ctx)
	}
}

func (c *CampsignCacher) GetCampaign(id string) (models.Campaign, bool) {
	c.Cache.Mu.RLock()
	defer c.Cache.Mu.RUnlock()
	camp, ok := c.Cache.Campaigns[id]
	return camp, ok
}

func (c *CampsignCacher) GetRules(id string) (models.TargetingRule, bool) {
	c.Cache.Mu.RLock()
	defer c.Cache.Mu.RUnlock()
	rule, ok := c.Cache.TargetingRules[id]
	return rule, ok
}

func (c *CampsignCacher) GetActiveCampaignIds() []string {
	return c.Cache.ActiveCampaignIds
}
