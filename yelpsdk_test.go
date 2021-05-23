package yelpsdk_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/getformative/yelpsdk"
)

type YelpSDKTestCase struct{}

func TestYelpSearch(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			data, err := ioutil.ReadFile("./sample_response.json")
			fmt.Println(string(data))
			if err != nil {
				t.Fatal(err)
			}
			w.Write(data)
		}),
	)
	defer server.Close()
	sdk, err := yelpsdk.NewYelpSDK(server.URL, "test-api-key")
	if err != nil {
		t.Fatal(err)
	}
	_, err = sdk.BusinessSearch(yelpsdk.YelpBusinessSearchParameters{
		Latitude:  46.0,
		Longitude: -110.9,
		Radius:    39000,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestBusinessSearchIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	sdk, err := yelpsdk.NewYelpSDK("https://api.yelp.com/v3", os.Getenv("YELP_KEY"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sdk)
	results, err := sdk.BusinessSearch(
		yelpsdk.YelpBusinessSearchParameters{
			Latitude:  40.7608,
			Longitude: -111.8910,
			Radius:    2000,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Yelp SDK search results: %v", results)
}
