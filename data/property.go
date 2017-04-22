package data

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
	Price       string
	Description string
	Address     Address
	Images      []Image
	Auction     bool
	Type        string
	NumBeds     string
	NumBaths    string
	NumCars     string
}
