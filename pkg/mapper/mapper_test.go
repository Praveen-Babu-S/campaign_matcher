package mapper

import (
	"campaigns/mocks" // Adjust this import path to your generated mock file
	"campaigns/pkg/models"
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestGetTargetedCampaigns(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	spotifyCampaign := models.Campaign{
		CampaignId:   "spotify",
		CampaignName: "Spotify - Music for everyone",
		Creatives: []models.Creative{
			{ImageUrl: "https://somelink"},
		},
		CTA:            "Download",
		CampaignStatus: "ACTIVE",
	}
	duolingoCampaign := models.Campaign{
		CampaignId:   "duolingo",
		CampaignName: "Duolingo: Best way to learn",
		Creatives: []models.Creative{
			{ImageUrl: "https://somelink2"},
		},
		CTA:            "Install",
		CampaignStatus: "ACTIVE",
	}
	subwaySurferCampaign := models.Campaign{
		CampaignId:   "subwaysurfer",
		CampaignName: "Subway Surfer",
		Creatives: []models.Creative{
			{ImageUrl: "https://somelink3"},
		},
		CTA:            "Play",
		CampaignStatus: "ACTIVE",
	}

	spotifyRule := models.TargetingRule{
		CampaignID:       "spotify",
		IncludeCountries: []string{"US", "Canada"},
	}
	duolingoRule := models.TargetingRule{
		CampaignID:       "duolingo",
		IncludeOS:        []string{"Android", "iOS"},
		ExcludeCountries: []string{"US"},
	}
	subwaySurferRule := models.TargetingRule{
		CampaignID:    "subwaysurfer",
		IncludeOS:     []string{"Android"},
		IncludeAppIDs: []string{"com.gametion.ludokinggame"},
	}
	noRuleCampaignID := "no_rule_campaign"
	noRuleCampaign := models.Campaign{
		CampaignId:   noRuleCampaignID,
		CampaignName: "Campaign with no rules",
		Creatives: []models.Creative{
			{ImageUrl: "https://no.rule.link"},
		},
		CTA:            "Learn More",
		CampaignStatus: "ACTIVE",
	}

	tests := []struct {
		name         string
		req          models.DeliveryRequest
		mockSetup    func(*mocks.MockICacher)
		wantResponse []models.DeliveryResponse
	}{
		{
			name: "should return empty campaign list",
			req:  models.DeliveryRequest{AppID: "test", Country: "test", OS: "test"},
			mockSetup: func(m *mocks.MockICacher) {
				m.EXPECT().GetActiveCampaignIds().Return([]string{}).Times(1)
			},
			wantResponse: []models.DeliveryResponse{},
		},
		{
			name: "should return campaigns with no rules)",
			req:  models.DeliveryRequest{AppID: "some.app", Country: "US", OS: "iOS"},
			mockSetup: func(m *mocks.MockICacher) {
				m.EXPECT().GetActiveCampaignIds().Return([]string{noRuleCampaignID}).Times(1)
				m.EXPECT().GetCampaign(noRuleCampaignID).Return(noRuleCampaign, true).Times(1)
				m.EXPECT().GetRules(noRuleCampaignID).Return(models.TargetingRule{}, false).Times(1)
			},
			wantResponse: []models.DeliveryResponse{
				{CampaignId: noRuleCampaign.CampaignId, Creatives: noRuleCampaign.Creatives, CTA: noRuleCampaign.CTA},
			},
		},
		{
			name: "should return campaign-duolingo",
			req:  models.DeliveryRequest{AppID: "com.abc.xyz", Country: "IND", OS: "iOS"},
			mockSetup: func(m *mocks.MockICacher) {
				m.EXPECT().GetActiveCampaignIds().Return([]string{duolingoCampaign.CampaignId}).Times(1)
				m.EXPECT().GetCampaign(duolingoCampaign.CampaignId).Return(duolingoCampaign, true).Times(1)
				m.EXPECT().GetRules(duolingoCampaign.CampaignId).Return(duolingoRule, true).Times(1)
			},
			wantResponse: []models.DeliveryResponse{
				{CampaignId: duolingoCampaign.CampaignId, Creatives: duolingoCampaign.Creatives, CTA: duolingoCampaign.CTA},
			},
		},
		{
			name: "should return all mapper campaigns",
			req:  models.DeliveryRequest{AppID: "com.gametion.ludokinggame", Country: "US", OS: "Android"},
			mockSetup: func(m *mocks.MockICacher) {
				activeIDs := []string{spotifyCampaign.CampaignId, subwaySurferCampaign.CampaignId}
				m.EXPECT().GetActiveCampaignIds().Return(activeIDs).Times(1)
				// mocks for spotify campaign
				m.EXPECT().GetCampaign(spotifyCampaign.CampaignId).Return(spotifyCampaign, true).Times(1)
				m.EXPECT().GetRules(spotifyCampaign.CampaignId).Return(spotifyRule, true).Times(1)

				// mock for subwaysurfer campaign
				m.EXPECT().GetCampaign(subwaySurferCampaign.CampaignId).Return(subwaySurferCampaign, true).Times(1)
				m.EXPECT().GetRules(subwaySurferCampaign.CampaignId).Return(subwaySurferRule, true).Times(1)
			},
			wantResponse: []models.DeliveryResponse{
				{CampaignId: spotifyCampaign.CampaignId, Creatives: spotifyCampaign.Creatives, CTA: spotifyCampaign.CTA},
				{CampaignId: subwaySurferCampaign.CampaignId, Creatives: subwaySurferCampaign.Creatives, CTA: subwaySurferCampaign.CTA},
			},
		},
		{
			name: "should return campaigns with multiple Include rules, all matching",
			req:  models.DeliveryRequest{AppID: "my.app", Country: "US", OS: "Android"},
			mockSetup: func(m *mocks.MockICacher) {
				campaign := models.Campaign{CampaignId: "multi_incl_camp", Creatives: []models.Creative{{ImageUrl: "link"}}, CTA: "cta", CampaignStatus: "ACTIVE"}
				rule := models.TargetingRule{
					CampaignID:       "multi_incl_camp",
					IncludeAppIDs:    []string{"my.app"},
					IncludeCountries: []string{"US"},
					IncludeOS:        []string{"Android"},
				}
				m.EXPECT().GetActiveCampaignIds().Return([]string{campaign.CampaignId}).Times(1)
				m.EXPECT().GetCampaign(campaign.CampaignId).Return(campaign, true).Times(1)
				m.EXPECT().GetRules(campaign.CampaignId).Return(rule, true).Times(1)
			},
			wantResponse: []models.DeliveryResponse{
				{CampaignId: "multi_incl_camp", Creatives: []models.Creative{{ImageUrl: "link"}}, CTA: "cta"},
			},
		},
		{
			name: "should return campaigns with multiple Include rules, one not matching",
			req:  models.DeliveryRequest{AppID: "my.app", Country: "CA", OS: "Android"}, // Country mismatch
			mockSetup: func(m *mocks.MockICacher) {
				campaign := models.Campaign{CampaignId: "multi_incl_camp", Creatives: []models.Creative{{ImageUrl: "link"}}, CTA: "cta", CampaignStatus: "ACTIVE"}
				rule := models.TargetingRule{
					CampaignID:       "multi_incl_camp",
					IncludeAppIDs:    []string{"my.app"},
					IncludeCountries: []string{"US"},
					IncludeOS:        []string{"Android"},
				}
				m.EXPECT().GetActiveCampaignIds().Return([]string{campaign.CampaignId}).Times(1)
				m.EXPECT().GetCampaign(campaign.CampaignId).Return(campaign, true).Times(1)
				m.EXPECT().GetRules(campaign.CampaignId).Return(rule, true).Times(1)
			},
			wantResponse: []models.DeliveryResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCacher := mocks.NewMockICacher(ctrl)
			tt.mockSetup(mockCacher)

			mapper := NewCampaignMapper(mockCacher)
			gotResponse := mapper.GetTargetedCampaigns(context.Background(), tt.req)

			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GetTargetedCampaigns() got = %+v, want %+v", gotResponse, tt.wantResponse)
			}
		})
	}
}
