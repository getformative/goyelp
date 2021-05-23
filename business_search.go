package yelpsdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

func (y *YelpSDK) businessSearchEndpoint() string {
	return fmt.Sprintf("%v/businesses/search", y.BaseURL)
}

func (y *YelpSDK) searchParamsToQueryString(params *YelpBusinessSearchParameters) (string, error) {
	qs, err := query.Values(params)
	if err != nil {
		return "", err
	}
	return qs.Encode(), nil
}

func (y *YelpSDK) getYelpClient() *http.Client {
	client := &http.Client{}
	return client
}

// BusinessSearch executes a search against the yelp api
func (y *YelpSDK) BusinessSearch(parameters YelpBusinessSearchParameters) (*YelpBusinessSearchResult, error) {
	if ok := parameters.Validate(); !ok {
		return nil, errors.New("search parameters are invalid")
	}
	endpoint := y.businessSearchEndpoint()
	qs, err := y.searchParamsToQueryString(&parameters)
	if err != nil {
		return nil, err
	}
	client := y.getYelpClient()
	url := fmt.Sprintf("%v?%v", endpoint, qs)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("there was an error during the http request")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", y.APIKey))
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("something went wrong while making the request")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var yelpResponse YelpBusinessSearchResult

	err = json.Unmarshal(body, &yelpResponse)
	if err != nil {
		return nil, err
	}
	return &yelpResponse, nil
}

// YelpBusinessSearchParameters are the parameters
// available when executing a business search against
// the Yelp Fusion API
type YelpBusinessSearchParameters struct {
	Term       string   `url:"term,omitempty"`
	Location   string   `url:"location,omitempty"`
	Latitude   float64  `url:"latitude,omitempty"`
	Longitude  float64  `url:"longitude,omitempty"`
	Radius     int      `url:"radius,omitempty"`
	Categories []string `url:"categories,omitempty"`
	Locale     string   `url:"locale,omitempty"`
	Limit      int      `url:"limit,omitempty"`
	Offset     int      `url:"offset,omitempty"`
	SortBy     string   `url:"sort_by,omitempty"`
	Price      []string `url:"price,omitempty"`
	OpenNow    bool     `url:"open_now,omitempty"`
	OpenAt     string   `url:"open_at,omitempty"`
	Attributes []string `url:"attributes,omitempty"`
}

// Validate ensures the YelpBusinessSearchParameters are in a valid state
func (y *YelpBusinessSearchParameters) Validate() bool {
	if y.Location == "" && (y.Latitude == 0.0 && y.Longitude == 0.0) {
		return false
	}
	if y.Radius == 0 || y.Radius > 40000 {
		return false
	}
	return true
}

// YelpBusinessSearchResult is the result that returns from the Yelp Fusion 3 API
type YelpBusinessSearchResult struct {
	Total      int        `json:"total"`
	Businesses []Business `json:"businesses"`
	Region     Region     `json:"region"`
}

func (y *YelpBusinessSearchResult) String() string {
	businessNames := make([]string, len(y.Businesses))
	for i := range y.Businesses {
		businessNames[i] = y.Businesses[i].Name
	}
	return fmt.Sprintf("Yelp Business Search Result: %v businesses: %v",
		len(y.Businesses),
		strings.Join(businessNames, ", "),
	)
}

// Business is a yelp business
type Business struct {
	Rating       int          `json:"rating"`
	Price        string       `json:"price"`
	Phone        string       `json:"phone"`
	ID           string       `json:"id"`
	Alias        string       `json:"alias"`
	IsClosed     bool         `json:"is_closed"`
	Categories   []Categories `json:"categories"`
	ReviewCount  int          `json:"review_count"`
	Name         string       `json:"name"`
	URL          string       `json:"url"`
	Coordinates  Coordinates  `json:"coordinates"`
	ImageURL     string       `json:"image_url"`
	Location     Location     `json:"location"`
	Distance     float64      `json:"distance"`
	Transactions []string     `json:"transactions"`
}

// Categories is the category of business defined by the Yelp API
type Categories struct {
	Alias string `json:"alias"`
	Title string `json:"title"`
}

// Coordinates represent the lat/lng location of the business
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Location represents the human-readable location of the business
type Location struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	State    string `json:"state"`
	Address1 string `json:"address1"`
	ZipCode  string `json:"zip_code"`
}

// Center is the geographical center of the search in lat/lng coordinates
type Center struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Region is a simple wrap for the center
type Region struct {
	Center Center `json:"center"`
}
