package mapper

import (
	"campaigns/pkg/cache"
	"campaigns/pkg/models"
	"context"
	"strings"
)

type ICampaignMapper interface {
	GetTargetedCampaigns(ctx context.Context, req models.DeliveryRequest) []models.DeliveryResponse
}

type CampaignMapper struct {
	Cache *cache.CampsignCacher
}

// NewCampaignMapper creates a new CampaignMapper.
func NewCampaignMapper(cache *cache.CampsignCacher) ICampaignMapper {
	return &CampaignMapper{
		Cache: cache,
	}
}

func (m *CampaignMapper) GetTargetedCampaigns(ctx context.Context, req models.DeliveryRequest) []models.DeliveryResponse {
	matchedCampaigns := []models.DeliveryResponse{}

	activeCampaigns := m.Cache.Cache.ActiveCampaignIds
	for _, campaignID := range activeCampaigns {
		campaign, f1 := m.Cache.GetCampaign(campaignID)
		rule, f2 := m.Cache.GetRules(campaignID)
		if !f1 {
			// campaign not found
			continue
		}
		if !f2 {
			// campaign without any rules can be served for all
			matchedCampaigns = append(matchedCampaigns, models.DeliveryResponse{
				CampaignId: campaign.CampaignId,
				Creatives:  campaign.Creatives,
				CTA:        campaign.CTA,
			})
			continue
		}

		// check if campaign has exclude params matches with req params
		for _, id := range rule.ExcludeAppIDs {
			if strings.EqualFold(id, req.AppID) {
				continue
			}
		}
		for _, id := range rule.ExcludeCountries {
			if strings.EqualFold(id, req.Country) {
				continue
			}
		}
		for _, id := range rule.ExcludeOS {
			if strings.EqualFold(id, req.OS) {
				continue
			}
		}
		includeCountry, includeOS, includeAppId := true, true, true
		for _, id := range rule.IncludeAppIDs {
			includeAppId = false
			if strings.EqualFold(id, req.AppID) {
				includeAppId = true
				break
			}
		}

		for _, id := range rule.IncludeCountries {
			includeCountry = false
			if strings.EqualFold(id, req.Country) {
				includeCountry = true
				break
			}
		}

		for _, id := range rule.IncludeOS {
			includeOS = false
			if strings.EqualFold(id, req.OS) {
				includeOS = true
				break
			}
		}
		if includeCountry && includeOS && includeAppId {
			matchedCampaigns = append(matchedCampaigns, models.DeliveryResponse{
				CampaignId: campaign.CampaignId,
				Creatives:  campaign.Creatives,
				CTA:        campaign.CTA,
			})
		}
	}

	return matchedCampaigns
}
