package mapper

import (
	"campaigns/mocks" // Adjust this import path to your generated mock file
	"campaigns/pkg/models"
	"context"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

var (
	testCampaigns = []models.Campaign{
		{
			CampaignId:   "spotify",
			CampaignName: "Spotify - Music for everyone",
			Creatives: []models.Creative{
				{ImageUrl: "https://somelink"},
			},
			CTA:            "Download",
			CampaignStatus: "ACTIVE",
		},
		{
			CampaignId:   "duolingo",
			CampaignName: "Duolingo: Best way to learn",
			Creatives: []models.Creative{
				{ImageUrl: "https://somelink2"},
			},
			CTA:            "Install",
			CampaignStatus: "ACTIVE",
		},
		{
			CampaignId:   "subwaysurfer",
			CampaignName: "Subway Surfer",
			Creatives: []models.Creative{
				{ImageUrl: "https://somelink3"},
			},
			CTA:            "Play",
			CampaignStatus: "ACTIVE",
		},
		{
			CampaignId:   "testcampaign",
			CampaignName: "Campaign with no rules",
			Creatives: []models.Creative{
				{ImageUrl: "https://no.rule.link"},
			},
			CTA:            "Learn More",
			CampaignStatus: "ACTIVE",
		},
		{
			CampaignId:   "multi_incl_camp",
			CampaignName: "camp_incl",
			Creatives:    []models.Creative{{ImageUrl: "link"}},
			CTA:          "cta", CampaignStatus: "ACTIVE",
		},
	}

	testRules = []models.TargetingRule{
		{
			CampaignID:       "spotify",
			IncludeCountries: []string{"US", "Canada"},
		},
		{
			CampaignID:       "duolingo",
			IncludeOS:        []string{"Android", "iOS"},
			ExcludeCountries: []string{"US"},
		},
		{
			CampaignID:    "subwaysurfer",
			IncludeOS:     []string{"Android"},
			IncludeAppIDs: []string{"com.gametion.ludokinggame"},
		},
		{
			CampaignID:       "multi_incl_camp",
			IncludeAppIDs:    []string{"my.app"},
			IncludeCountries: []string{"US"},
			IncludeOS:        []string{"Android"},
		},
	}
)

func TestGetTargetedCampaigns(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		req          models.DeliveryRequest
		mocks        func(*mocks.MockICacher)
		wantResponse []models.DeliveryResponse
	}{
		{
			name: "should return empty campaign list",
			req:  models.DeliveryRequest{AppID: "test", Country: "test", OS: "test"},
			mocks: func(m *mocks.MockICacher) {
				m.EXPECT().GetActiveCampaignIds().Return([]string{}).Times(1)
			},
			wantResponse: []models.DeliveryResponse{},
		},
		{
			name: "should return campaigns with no rules)",
			req:  models.DeliveryRequest{AppID: "some.app", Country: "US", OS: "iOS"},
			mocks: func(m *mocks.MockICacher) {
				m.EXPECT().GetActiveCampaignIds().Return([]string{"testcampaign"}).Times(1)
				m.EXPECT().GetCampaign("testcampaign").Return(testCampaigns[3], true).Times(1)
				m.EXPECT().GetRule("testcampaign").Return(models.ProcessedRule{}, false).Times(1)
			},
			wantResponse: []models.DeliveryResponse{
				{CampaignId: testCampaigns[3].CampaignId, Creatives: testCampaigns[3].Creatives, CTA: testCampaigns[3].CTA},
			},
		},
		{
			name: "should return campaign-duolingo",
			req:  models.DeliveryRequest{AppID: "com.abc.xyz", Country: "IND", OS: "iOS"},
			mocks: func(m *mocks.MockICacher) {
				m.EXPECT().GetActiveCampaignIds().Return([]string{testCampaigns[1].CampaignId}).Times(1)
				m.EXPECT().GetCampaign(testCampaigns[1].CampaignId).Return(testCampaigns[1], true).Times(1)
				m.EXPECT().GetRule(testCampaigns[1].CampaignId).Return(*models.NewProcessedRule(&testRules[1]), true).Times(1)
			},
			wantResponse: []models.DeliveryResponse{
				{CampaignId: testCampaigns[1].CampaignId, Creatives: testCampaigns[1].Creatives, CTA: testCampaigns[1].CTA},
			},
		},
		{
			name: "should return all mapper campaigns",
			req:  models.DeliveryRequest{AppID: "com.gametion.ludokinggame", Country: "US", OS: "Android"},
			mocks: func(m *mocks.MockICacher) {
				activeIDs := []string{testCampaigns[0].CampaignId, testCampaigns[2].CampaignId}
				m.EXPECT().GetActiveCampaignIds().Return(activeIDs).Times(1)
				// mocks for spotify campaign
				m.EXPECT().GetCampaign(testCampaigns[0].CampaignId).Return(testCampaigns[0], true).Times(1)
				m.EXPECT().GetRule(testCampaigns[0].CampaignId).Return(*models.NewProcessedRule(&testRules[0]), true).Times(1)

				// mock for subwaysurfer campaign
				m.EXPECT().GetCampaign(testCampaigns[2].CampaignId).Return(testCampaigns[2], true).Times(1)
				m.EXPECT().GetRule(testCampaigns[2].CampaignId).Return(*models.NewProcessedRule(&testRules[2]), true).Times(1)
			},
			wantResponse: []models.DeliveryResponse{
				{CampaignId: testCampaigns[0].CampaignId, Creatives: testCampaigns[0].Creatives, CTA: testCampaigns[0].CTA},
				{CampaignId: testCampaigns[2].CampaignId, Creatives: testCampaigns[2].Creatives, CTA: testCampaigns[2].CTA},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCacher := mocks.NewMockICacher(ctrl)
			tt.mocks(mockCacher)

			mapper := NewCampaignMapper(mockCacher)
			gotResponse := mapper.GetTargetedCampaigns(context.Background(), tt.req)

			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GetTargetedCampaigns() got = %+v, want %+v", gotResponse, tt.wantResponse)
			}
		})
	}
}
