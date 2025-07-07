package campaignmapper

import (
	"campaigns/models"
	"strings"
)

type SimpleCampaignMatcher struct{}

// NewSimpleCampaignMatcher creates a new SimpleCampaignMatcher.
func NewSimpleCampaignMatcher() *SimpleCampaignMatcher {
	return &SimpleCampaignMatcher{}
}

// Match performs the campaign matching operation.
func (m *SimpleCampaignMatcher) Match(req models.DeliveryRequest, campaigns map[string]models.Campaign, rules map[string]models.TargetingRule) []models.DeliveryResponse {
	var matchedCampaigns []models.DeliveryResponse

	for campaignID, campaign := range campaigns {
		if campaign.CampaignStatus != "ACTIVE" {
			continue
		}

		rule, exists := rules[campaignID]
		if !exists {
			// If no targeting rule exists, assume it can be served everywhere
			matchedCampaigns = append(matchedCampaigns, models.DeliveryResponse{
				CampaignId: campaign.CampaignId,
				Creatives:  campaign.Creatives,
				CTA:        campaign.CTA,
			})
			continue
		}

		// Evaluate App ID
		if len(rule.IncludeAppIDs) > 0 {
			found := false
			for _, id := range rule.IncludeAppIDs {
				if strings.EqualFold(id, req.AppID) {
					found = true
					break
				}
			}
			if !found {
				continue // App ID not included
			}
		}
		if len(rule.ExcludeAppIDs) > 0 {
			for _, id := range rule.ExcludeAppIDs {
				if strings.EqualFold(id, req.AppID) {
					continue // App ID excluded
				}
			}
		}

		// Evaluate Country
		if len(rule.IncludeCountries) > 0 {
			found := false
			for _, country := range rule.IncludeCountries {
				if strings.EqualFold(country, req.Country) {
					found = true
					break
				}
			}
			if !found {
				continue // Country not included
			}
		}
		if len(rule.ExcludeCountries) > 0 {
			excluded := false
			for _, country := range rule.ExcludeCountries {
				if strings.EqualFold(country, req.Country) {
					excluded = true
					break
				}
			}
			if excluded {
				continue // Country excluded
			}
		}

		// Evaluate OS
		if len(rule.IncludeOS) > 0 {
			found := false
			for _, os := range rule.IncludeOS {
				if strings.EqualFold(os, req.OS) {
					found = true
					break
				}
			}
			if !found {
				continue // OS not included
			}
		}
		if len(rule.ExcludeOS) > 0 {
			for _, os := range rule.ExcludeOS {
				if strings.EqualFold(os, req.OS) {
					continue // OS excluded
				}
			}
		}

		// If all checks pass, add to matched campaigns
		matchedCampaigns = append(matchedCampaigns, models.DeliveryResponse{
			CampaignId: campaign.CampaignId,
			Creatives:  campaign.Creatives,
			CTA:        campaign.CTA,
		})
	}

	return matchedCampaigns
}
