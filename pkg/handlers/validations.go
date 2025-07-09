package handlers

import (
	"campaigns/pkg/models"
	"errors"
	"net/url"
)

func fetchAndValidateReqParams(query url.Values) (*models.DeliveryRequest, error) {

	appID := query.Get("app")
	country := query.Get("country")
	os := query.Get("os")
	req := &models.DeliveryRequest{
		AppID:   appID,
		Country: country,
		OS:      os,
	}

	if appID == "" {
		return nil, errors.New("missing app param")
	}
	if country == "" {
		return nil, errors.New("missing country param")
	}
	if os == "" {
		return nil, errors.New("missing os param")
	}
	return req, nil
}
