package models

import "strings"

type TargetingRule struct {
	CampaignID string `json:"campaignId"`

	IncludeAppIDs []string `json:"includeAppIDs"`
	ExcludeAppIDs []string `json:"excludeAppIDs"`

	IncludeCountries []string `json:"includeCountries"`
	ExcludeCountries []string `json:"excludeCountries"`

	IncludeOS []string `json:"includeOs"`
	ExcludeOS []string `json:"excludeOs"`
}

type ProcessedRule struct {
	CampaignId           string
	IncludedAppIdsMap    map[string]struct{}
	IncludedOsMap        map[string]struct{}
	IncludedCountiesMap  map[string]struct{}
	ExcludedAppIdsMap    map[string]struct{}
	ExcludedCountriesMap map[string]struct{}
	ExcludedOsMap        map[string]struct{}
}

type DeliveryRequest struct {
	AppID   string `json:"app,omitempty"`
	Country string `json:"country,omitempty"`
	OS      string `json:"os,omitempty"`
}

type DeliveryResponse struct {
	CampaignId string     `json:"campaignId,omitempty"`
	Creatives  []Creative `json:"creatives,omitempty"`
	CTA        string     `json:"cta,omitempty"`
}

func NewProcessedRule(r *TargetingRule) *ProcessedRule {
	rule := &ProcessedRule{
		IncludedAppIdsMap:    make(map[string]struct{}),
		ExcludedAppIdsMap:    make(map[string]struct{}),
		IncludedOsMap:        make(map[string]struct{}),
		ExcludedOsMap:        make(map[string]struct{}),
		IncludedCountiesMap:  make(map[string]struct{}),
		ExcludedCountriesMap: make(map[string]struct{}),
	}
	rule.CampaignId = r.CampaignID
	for _, id := range r.IncludeAppIDs {
		rule.IncludedAppIdsMap[strings.ToLower(id)] = struct{}{}
	}
	for _, id := range r.IncludeOS {
		rule.IncludedOsMap[strings.ToLower(id)] = struct{}{}
	}
	for _, id := range r.IncludeCountries {
		rule.IncludedCountiesMap[strings.ToLower(id)] = struct{}{}
	}

	for _, id := range r.ExcludeAppIDs {
		rule.ExcludedAppIdsMap[strings.ToLower(id)] = struct{}{}
	}
	for _, id := range r.ExcludeOS {
		rule.ExcludedOsMap[strings.ToLower(id)] = struct{}{}
	}
	for _, id := range r.ExcludeCountries {
		rule.ExcludedCountriesMap[strings.ToLower(id)] = struct{}{}
	}

	return rule
}
