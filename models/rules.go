package models

type TargetingRule struct {
	CampaignID    string
	IncludeAppIDs []string
	ExcludeAppIDs []string

	IncludeCountries []string
	ExcludeCountries []string

	IncludeOS []string
	ExcludeOS []string
}

type DeliveryRequest struct {
	AppID   string `json:"appId,omitempty"`
	Country string `json:"country,omitempty"`
	OS      string `json:"os,omitempty"`
}

type DeliveryResponse struct {
	CampaignId string     `json:"campaignId,omitempty"`
	Creatives  []Creative `json:"creatives,omitempty"`
	CTA        string     `json:"cta,omitempty"`
}
