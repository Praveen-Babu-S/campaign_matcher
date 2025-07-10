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
	Cache cache.ICacher
}

// NewCampaignMapper creates a new CampaignMapper.
func NewCampaignMapper(cache cache.ICacher) ICampaignMapper {
	return &CampaignMapper{
		Cache: cache,
	}
}

func (m *CampaignMapper) GetTargetedCampaigns(ctx context.Context, req models.DeliveryRequest) []models.DeliveryResponse {
	matchedCampaigns := []models.DeliveryResponse{}

	activeCampaigns := m.Cache.GetActiveCampaignIds()
	app := strings.ToLower(req.AppID)
	os := strings.ToLower(req.OS)
	country := strings.ToLower(req.Country)
	for _, campaignID := range activeCampaigns {
		campaign, f1 := m.Cache.GetCampaign(campaignID)
		rule, f2 := m.Cache.GetRule(campaignID)
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

		//check if campaign has exclude params matches with req params
		if len(rule.ExcludedAppIdsMap) > 0 {
			if _, ok := rule.ExcludedAppIdsMap[app]; ok {
				continue
			}
		}

		if len(rule.ExcludedCountriesMap) > 0 {
			if _, ok := rule.ExcludedCountriesMap[country]; ok {
				continue
			}
		}

		if len(rule.ExcludedOsMap) > 0 {
			if _, ok := rule.ExcludedOsMap[os]; ok {
				continue
			}
		}
		include := true
		if len(rule.IncludedAppIdsMap) > 0 {
			if _, ok := rule.IncludedAppIdsMap[app]; !ok {
				include = false
			}
		}

		if len(rule.IncludedCountiesMap) > 0 {
			if _, ok := rule.IncludedCountiesMap[country]; !ok {
				include = false
			}
		}

		if len(rule.IncludedOsMap) > 0 {
			if _, ok := rule.IncludedOsMap[os]; !ok {
				include = false
			}
		}
		if include {
			matchedCampaigns = append(matchedCampaigns, models.DeliveryResponse{
				CampaignId: campaign.CampaignId,
				Creatives:  campaign.Creatives,
				CTA:        campaign.CTA,
			})
		}
	}
	return matchedCampaigns
}
