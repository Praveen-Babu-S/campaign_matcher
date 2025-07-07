package campaignsfetcher

import (
	"campaigns/models"
	"context"
)

type InMemoryDataFetcher struct{}

// NewInMemoryDataFetcher creates a new InMemoryDataFetcher.
func NewInMemoryDataFetcher() *InMemoryDataFetcher {
	return &InMemoryDataFetcher{}
}

// FetchCampaigns retrieves campaign data from a memory store.
func (f *InMemoryDataFetcher) FetchCampaigns(ctx context.Context) (map[string]models.Campaign, error) {

	campaigns := map[string]models.Campaign{
		"spotify": {
			CampaignId:   "spotify",
			CampaignName: "Spotify - Music for everyone",
			Creatives: []models.Creative{
				{
					ImageUrl: "https://somelink",
				},
			},
			CTA:            "Download",
			CampaignStatus: "ACTIVE",
		},
		"duolingo": {
			CampaignId:   "duolingo",
			CampaignName: "Duolingo: Best way to learn",
			Creatives: []models.Creative{
				{
					ImageUrl: "https://somelink2",
				},
			},
			CTA:            "Install",
			CampaignStatus: "INATCIVE",
		},
		"subwaysurfer": {
			CampaignId:   "subwaysurfer",
			CampaignName: "Subway Surfer",
			Creatives: []models.Creative{
				{
					ImageUrl: "https://somelink3",
				},
			},
			CTA:            "Play",
			CampaignStatus: "ATCIVE",
		},
		"inactive_campaign": {
			CampaignId:   "inactive_campaign",
			CampaignName: "Inactive Test Campaign",
			Creatives: []models.Creative{
				{
					ImageUrl: "https://somelink4",
				},
			},
			CTA:            "Buy Now",
			CampaignStatus: "INATCIVE",
		},
	}
	return campaigns, nil
}

// FetchTargetingRules retrieves targeting rule data from data store.
func (f *InMemoryDataFetcher) FetchTargetingRules(ctx context.Context) (map[string]models.TargetingRule, error) {
	rules := map[string]models.TargetingRule{
		"spotify": {
			CampaignID:       "spotify",
			IncludeCountries: []string{"US", "Canada"},
		},
		"duolingo": {
			CampaignID:       "duolingo",
			IncludeOS:        []string{"Android", "iOS"},
			ExcludeCountries: []string{"US"},
		},
		"subwaysurfer": {
			CampaignID:    "subwaysurfer",
			IncludeOS:     []string{"Android"},
			IncludeAppIDs: []string{"com.gametion.ludokinggame"},
		},
	}
	return rules, nil
}
