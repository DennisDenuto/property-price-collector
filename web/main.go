package main

import (
	"net/http"
	"github.com/DennisDenuto/property-price-collector/site/domaincomau"
	"github.com/DennisDenuto/property-price-collector/data"
	"fmt"
	"googlemaps.github.io/maps"
	"github.com/Sirupsen/logrus"
	"context"
	"os"
	"github.com/pkg/errors"
	"math"
	"strings"
)

func main() {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("PLACES_KEY")))
	if err != nil {
		logrus.Errorf("new client fatal error: %s", err)
		panic(err)

	}

	err = http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.URL.RawQuery)
		propertyUrl := r.URL.Path[1:]
		resp, err := http.DefaultClient.Get(propertyUrl)
		domaincomau.GetDomainComAuPropertyDetailWrapper(resp, err, func(wrapper *data.DomainComAuPropertyDetailWrapper) {

			latLng, err := getLatLon(c, wrapper.Page.PageInfo.Property.Address)
			if err != nil {
				logrus.Errorf("geocode fatal error: %s", err)
				w.WriteHeader(500)
				return
			}

			stations, err := getNearbyTrainStations(c, latLng)
			if err != nil {
				logrus.Errorf("nearby search fatal error: %s", err)
				w.WriteHeader(500)
				return
			}

			property := wrapper.Page.PageInfo.Property
			w.Write([]byte(fmt.Sprintf(`Address,	Suburb,	Land Size,	Price,	Train Stations, Num Beds, Num Baths, Num Garage, url
"%v", "%v", "%v", "%v", "%v", "%v", "%v", "%v", "%v"`,
				property.Address, property.Suburb, property.Landsize, property.Price, stations, property.Bedrooms, property.Bathrooms, property.Parking, propertyUrl)))
		})

		if err != nil {
			w.WriteHeader(500)
		}
	}))

	if err != nil {
		panic(err.Error())
	}
}

func getNearbyTrainStations(client *maps.Client, lng maps.LatLng) (string, error) {
	request := &maps.NearbySearchRequest{
		Location: &lng,
		RankBy:   maps.RankByDistance,
		Type:     maps.PlaceTypeTrainStation,
	}

	nearbySearch, err := client.NearbySearch(context.Background(), request)
	if err != nil {
		logrus.Errorf("nearby search fatal error: %s", err)
		return "", err
	}

	var nearbyStations []string
	for _, result := range nearbySearch.Results {
		distance := Distance(result.Geometry.Location.Lat, result.Geometry.Location.Lng, lng.Lat, lng.Lng)
		station := fmt.Sprintf("%v distance: %.2f metres", result.Name, distance)
		nearbyStations = append(nearbyStations, station)
	}

	return strings.Join(nearbyStations, ","), nil
}

func getLatLon(client *maps.Client, address string) (maps.LatLng, error) {
	reverseGeocode, err := client.Geocode(context.Background(), &maps.GeocodingRequest{
		Address: address,
	})
	if err != nil {
		logrus.Errorf("geocode fatal error: %s", err)
		return maps.LatLng{}, err
	}

	if len(reverseGeocode) <= 0 {
		return maps.LatLng{}, errors.New("could not get latlon for address")
	}

	return reverseGeocode[0].Geometry.Location, nil
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
