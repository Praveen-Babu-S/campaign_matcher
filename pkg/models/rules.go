package models

type TargetingRule struct {
	CampaignID string `json:"campaignId"`

	IncludeAppIDs []string `json:"includeAppIDs"`
	ExcludeAppIDs []string `json:"excludeAppIDs"`

	IncludeCountries []string `json:"includeCountries"`
	ExcludeCountries []string `json:"excludeCountries"`

	IncludeOS []string `json:"includeOs"`
	ExcludeOS []string `json:"excludeOs"`
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
