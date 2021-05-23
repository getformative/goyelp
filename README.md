[![Coverage Status](https://coveralls.io/repos/github/getformative/goyelp/badge.svg?branch=main)](https://coveralls.io/github/getformative/goyelp?branch=main)

[![Test Yelp SDK](https://github.com/getformative/goyelp/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/getformative/goyelp/actions/workflows/ci-cd.yml)

# Go Yelp

Go Yelp is a simple API wrapper for the [Yelp v3 API](https://www.yelp.com/developers/documentation/v3).

## Installation

To install Go Yelp, simply do
```
go get github.com/getformative/goyelp
```

## Usage

### Common Example

``` go
sdk, err := yelpsdk.NewYelpSDK("https://api.yelp.com/v3", os.Getenv("YELP_KEY"))
if err != nil {
  panic(err)
}
results, err := sdk.BusinessSearch(
  yelpsdk.YelpBusinessSearchParameters{
    Latitude:  40.7608,
    Longitude: -111.8910,
    Radius:    2000,
  },
)
if err != nil {
  panic(err)
}
print(results)
```