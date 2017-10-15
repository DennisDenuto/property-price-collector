package data

import (
	"time"
)

type LonLat struct {
	Lon string
	Lat string
}

type Address struct {
	AddressLine1 string
	PostCode     string
	Suburb       string
	State        string
	LonLat       LonLat
}

type Image []byte

type Property struct {
	DateListed    time.Time
	Size          string
	AgeOfProperty string
	Price         string
	Description   string
	Address       Address
	Images        []Image
	Auction       bool
	Type          string
	NumBeds       string
	NumBaths      string
	NumCars       string
}

type PropertyHistoryData struct {
	Type          string
	DateSold      time.Time
	Price         string
	Size          string
	AgeOfProperty string
	PriceHistory []struct {
		Price    string
		DateSole time.Time
	}
	Description string
	Address     Address
	NumBeds     string
	NumBaths    string
	NumCars     string
}
