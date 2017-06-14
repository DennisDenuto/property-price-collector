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
	Property
	DateSold time.Time
}

//TODO consider:
// Infrastructure: water, sewerage, power, gas and telco – all publically accessible and highly relevant.
// Upcoming developments planned future developments are going to play a big part in predicting the hot spots of the future, enabling algorithms to catch gentrification before it happens
