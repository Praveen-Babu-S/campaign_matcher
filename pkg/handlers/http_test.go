package handlers

import (
	"campaigns/mocks" // Adjust to your generated mock path
	"campaigns/pkg/models"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestDeliverCampaigns(t *testing.T) {

	sampleReq := models.DeliveryRequest{
		AppID:   "com.test.app",
		Country: "US",
		OS:      "Android",
	}
	campaign := models.DeliveryResponse{
		CampaignId: "camp1",
		Creatives: []models.Creative{
			{ImageUrl: "url"},
		},
		CTA: "Download",
	}

	tests := []struct {
		name           string
		method         string
		queryParams    url.Values
		mockSetup      func(m *mocks.MockICampaignMapper)
		wantStatusCode int
		wantBody       string
		wantHeader     http.Header
	}{
		{
			name:   "invalid HTTP Method - POST",
			method: http.MethodPost,
			queryParams: url.Values{
				"app":     []string{"com.test.app"},
				"country": []string{"US"},
				"os":      []string{"Android"},
			},
			mockSetup:      func(m *mocks.MockICampaignMapper) {},
			wantStatusCode: http.StatusMethodNotAllowed,
			wantBody:       "Method not allowed\n",
			wantHeader: http.Header{
				"Content-Type": {"text/plain; charset=utf-8"},
			},
		},
		{
			name:   "missing 'app' query parameter",
			method: http.MethodGet,
			queryParams: url.Values{
				"country": []string{"US"},
				"os":      []string{"Android"},
			},
			mockSetup:      func(m *mocks.MockICampaignMapper) {},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"error":"missing app param"}` + "\n",
			wantHeader: http.Header{
				"Content-Type": {"application/json"},
			},
		},
		{
			name:   "Missing 'country' query parameter",
			method: http.MethodGet,
			queryParams: url.Values{
				"app": []string{"com.test.app"},
				"os":  []string{"Android"},
			},
			mockSetup:      func(m *mocks.MockICampaignMapper) {},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"error":"missing country param"}` + "\n",
			wantHeader: http.Header{
				"Content-Type": {"application/json"},
			},
		},
		{
			name:   "No campaigns found (204 No Content)",
			method: http.MethodGet,
			queryParams: url.Values{
				"app":     []string{sampleReq.AppID},
				"country": []string{sampleReq.Country},
				"os":      []string{sampleReq.OS},
			},
			mockSetup: func(m *mocks.MockICampaignMapper) {
				m.EXPECT().GetTargetedCampaigns(gomock.Any(), sampleReq).Return([]models.DeliveryResponse{}).Times(1)
			},
			wantStatusCode: http.StatusNoContent,
			wantBody:       "",
			wantHeader:     http.Header{},
		},
		{
			name:   "campaigns found (200 OK)",
			method: http.MethodGet,
			queryParams: url.Values{
				"app":     []string{sampleReq.AppID},
				"country": []string{sampleReq.Country},
				"os":      []string{sampleReq.OS},
			},
			mockSetup: func(m *mocks.MockICampaignMapper) {
				m.EXPECT().GetTargetedCampaigns(gomock.Any(), sampleReq).Return([]models.DeliveryResponse{campaign}).Times(1)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockMapper := mocks.NewMockICampaignMapper(ctrl)
			tt.mockSetup(mockMapper)
			handler := NewCampaignsHTTPHandler(mockMapper)

			req := httptest.NewRequest(tt.method, "/v1/delivery?"+tt.queryParams.Encode(), nil)
			ctx := context.Background()
			req = req.WithContext(ctx)

			// rr is  to capture the HTTP response
			rr := httptest.NewRecorder()

			handler.DeliverCampaigns(rr, req)

			if status := rr.Code; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatusCode)
			}

			for k, v := range tt.wantHeader {
				if gotHeader := rr.Header().Get(k); len(v) > 0 && gotHeader != v[0] {
					t.Errorf("handler returned wrong header %s: got %q want %q", k, gotHeader, v[0])
				}
			}
		})
	}
}
