package handlers

import (
	"campaigns/pkg/models"
	"errors"
	"net/url"
	"testing"
)

func Test_fetchAndValidateReqParams(t *testing.T) {
	tests := []struct {
		name        string
		queryParams url.Values
		wantReq     *models.DeliveryRequest
		wantErr     error
	}{
		{
			name: "Success - All parameters present",
			queryParams: url.Values{
				"app":     []string{"com.example.app"},
				"country": []string{"US"},
				"os":      []string{"Android"},
			},
			wantReq: &models.DeliveryRequest{
				AppID:   "com.example.app",
				Country: "US",
				OS:      "Android",
			},
			wantErr: nil,
		},
		{
			name: "Error - Missing app param",
			queryParams: url.Values{
				"country": []string{"US"},
				"os":      []string{"Android"},
			},
			wantReq: nil,
			wantErr: errors.New("missing app param"),
		},
		{
			name: "Error - Missing country param",
			queryParams: url.Values{
				"app": []string{"com.example.app"},
				"os":  []string{"Android"},
			},
			wantReq: nil,
			wantErr: errors.New("missing country param"),
		},
		{
			name: "Error - Missing os param",
			queryParams: url.Values{
				"app":     []string{"com.example.app"},
				"country": []string{"US"},
			},
			wantReq: nil,
			wantErr: errors.New("missing os param"),
		},
		{
			name:        "Error - All parameters missing",
			queryParams: url.Values{}, // Empty query params
			wantReq:     nil,
			wantErr:     errors.New("missing app param"), // Should return error for 'app' first
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReq, gotErr := fetchAndValidateReqParams(tt.queryParams)

			// Check for error
			if (gotErr != nil) != (tt.wantErr != nil) {
				t.Errorf("fetchAndValidateReqParams() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error() {
				t.Errorf("fetchAndValidateReqParams() error message = %q, wantErr message %q", gotErr.Error(), tt.wantErr.Error())
				return
			}

			// Check for returned request struct
			if gotReq == nil && tt.wantReq != nil {
				t.Errorf("fetchAndValidateReqParams() gotReq = %v, wantReq %v", gotReq, tt.wantReq)
				return
			}
			if gotReq != nil && tt.wantReq == nil {
				t.Errorf("fetchAndValidateReqParams() gotReq = %v, wantReq %v", gotReq, tt.wantReq)
				return
			}
			if gotReq != nil && tt.wantReq != nil {
				if *gotReq != *tt.wantReq { // Compare struct values
					t.Errorf("fetchAndValidateReqParams() gotReq = %+v, wantReq %+v", *gotReq, *tt.wantReq)
				}
			}
		})
	}
}
