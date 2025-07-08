package handlers

import (
	"campaigns/models"
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
		return nil, errors.New("missing appId")
	}
	if country == "" {
		return nil, errors.New("missing country code")
	}
	if os == "" {
		return nil, errors.New("missing os code")
	}
	return req, nil
}
