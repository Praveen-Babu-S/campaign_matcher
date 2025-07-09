package models

type CampaignStatus int

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
