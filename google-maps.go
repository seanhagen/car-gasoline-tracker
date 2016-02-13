package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const GoogleMapsApi = "https://maps.googleapis.com"
const GoogleMapsPath = "/maps/api/geocode/json"

type (
	Response struct {
		Status  string   `json:"status"`
		Results []Result `json:"results"`
	}

	Result struct {
		Types             []string           `json:"types"`
		FormattedAddress  string             `json:"formatted_address"`
		AddressComponents []AddressComponent `json:"address_components"`
		Geometry          GeometryData       `json:"geometry"`
	}

	Address struct {
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Address string  `json:"address"`
	}

	AddressComponent struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	}

	GeometryData struct {
		Location     LatLng `json:"location"`
		LocationType string `json:"location_type"`
		Viewport     struct {
			Southwest LatLng `json:"southwest"`
			Northeast LatLng `json:"northeast"`
		} `json:"viewport"`
		Bounds struct {
			Southwest LatLng `json:"southwest"`
			Northeast LatLng `json:"northeast"`
		} `json:"bounds"`
	}

	LatLng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
)

func getAddressGeo(address string) *Address {
	var Url *url.URL
	Url, err := url.Parse(GoogleMapsApi)
	if err != nil {
		panic("boom")
	}

	Url.Path += GoogleMapsPath

	parameters := url.Values{}
	parameters.Add("key", *googleMapsApiKey)
	parameters.Add("address", address)
	Url.RawQuery = parameters.Encode()

	Uri := Url.String()

	resp, err := http.Get(Uri)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var g = new(Response)
	err = json.NewDecoder(resp.Body).Decode(g)

	if err != nil {
		panic(err)
	}

	return &Address{
		Lat:     g.Results[0].Geometry.Location.Lat,
		Lng:     g.Results[0].Geometry.Location.Lng,
		Address: g.Results[0].FormattedAddress,
	}
}
