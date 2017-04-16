package data

type Address struct {
	AddressLine1, AddressLine2 string
	PostCode                   string
	Suburb                     string
	State                      string
}

type Image []byte

type Property struct {
	Price       string
	Description string
	Address     Address
	Images      []Image
	Auction     bool
	Type        string
}