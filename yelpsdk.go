package yelpsdk

import "errors"

// See: https://www.yelp.com/developers/documentation/v3/business_search

// NewYelpSDK constructs a yelp sdk and returns a pointer to the SDK or returns an error
func NewYelpSDK(baseURL string, APIKey string) (*YelpSDK, error) {
	if APIKey == "" {
		return nil, errors.New("api key must not be empty")
	}
	if baseURL == "" {
		return nil, errors.New("base url must not be empty")
	}
	return &YelpSDK{
		BaseURL: baseURL,
		APIKey:  APIKey,
	}, nil
}

// YelpSDK provides an easy interface for interacting with the yelp places api
type YelpSDK struct {
	BaseURL string
	APIKey  string
}
