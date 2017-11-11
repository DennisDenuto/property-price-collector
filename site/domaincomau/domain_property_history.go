package domaincomau

import "time"

type DomainComAuPropertyWrapper struct {
	AdvertBannerObject     interface{} `json:"advertBannerObject"`
	BuildingProfileURLSlug string      `json:"buildingProfileUrlSlug"`
	EndpointHost           string      `json:"endpointHost"`
	NoEstimateInfo         string      `json:"noEstimateInfo"`
	PostcodeRegex          struct {
	} `json:"postcodeRegex"`
	Property struct {
		Address string `json:"address"`
		ID      string `json:"id"`
	} `json:"property"`
	PropertyObject DomainComAuPropertyHistory `json:"propertyObject"`
	ShowBuildingProfileLink bool   `json:"showBuildingProfileLink"`
	ShowStreetProfileLink   bool   `json:"showStreetProfileLink"`
	StickyHeaderEnabled     bool   `json:"stickyHeaderEnabled"`
	StreetProfileURLSlug    string `json:"streetProfileUrlSlug"`
	TermsInput              string `json:"termsInput"`
	TypeaheadHintCSS        string `json:"typeaheadHintCss"`
	TypeaheadMenuCSS        string `json:"typeaheadMenuCss"`
	TypeaheadSuggestCSS     string `json:"typeaheadSuggestCss"`
	User                    struct {
		IsAuthenticated bool   `json:"IsAuthenticated"`
		FamilyName      string `json:"familyName"`
		FullName        string `json:"fullName"`
		GivenName       string `json:"givenName"`
	} `json:"user"`
	Valuation struct {
		Date            string  `json:"date"`
		LowerPrice      int     `json:"lowerPrice"`
		MidPrice        int     `json:"midPrice"`
		PriceConfidence string  `json:"priceConfidence"`
		RentPerWeek     int     `json:"rentPerWeek"`
		RentYield       float64 `json:"rentYield"`
		UpperPrice      int     `json:"upperPrice"`
	} `json:"valuation"`
}

type DomainComAuPropertyHistory struct {
	Address           string `json:"address"`
	AddressCoordinate struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"addressCoordinate"`
	AddressID int           `json:"addressId"`
	Adverts   []interface{} `json:"adverts"`
	AreaSize  int           `json:"areaSize"`
	Bathrooms int           `json:"bathrooms"`
	Bedrooms  int           `json:"bedrooms"`
	Cadastre  struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"cadastre"`
	CadastreType string    `json:"cadastreType"`
	CarSpaces    int       `json:"carSpaces"`
	Created      time.Time `json:"created"`
	Features     []string  `json:"features"`
	FlatNumber   string    `json:"flatNumber"`
	History      struct {
		Rentals []interface{} `json:"rentals"`
		Sales   []struct {
			Agency           string  `json:"agency,omitempty"`
			ApmAgencyID      int     `json:"apmAgencyId,omitempty"`
			Date             string  `json:"date"`
			DaysOnMarket     float64 `json:"daysOnMarket,omitempty"`
			DocumentedAsSold bool    `json:"documentedAsSold"`
			Price            int     `json:"price"`
			ReportedAsSold   bool    `json:"reportedAsSold"`
			SuppressDetails  bool    `json:"suppressDetails"`
			SuppressPrice    bool    `json:"suppressPrice"`
			Type             string  `json:"type"`
			First            struct {
				AdvertisedDate  string `json:"advertisedDate"`
				Agency          string `json:"agency"`
				ApmAgencyID     int    `json:"apmAgencyId"`
				Source          string `json:"source"`
				SuppressDetails bool   `json:"suppressDetails"`
				SuppressPrice   bool   `json:"suppressPrice"`
				Type            string `json:"type"`
			} `json:"first"`
			ID   int64 `json:"id"`
			Last struct {
				AdvertisedDate  string `json:"advertisedDate"`
				AdvertisedPrice int    `json:"advertisedPrice"`
				Agency          string `json:"agency"`
				ApmAgencyID     int    `json:"apmAgencyId"`
				Source          string `json:"source"`
				SuppressDetails bool   `json:"suppressDetails"`
				SuppressPrice   bool   `json:"suppressPrice"`
				Type            string `json:"type"`
			} `json:"last"`
			PropertyType string `json:"propertyType"`
		} `json:"sales"`
		Valuations []struct {
			Confidence string `json:"confidence"`
			Date       string `json:"date"`
			LowerPrice int    `json:"lowerPrice"`
			UpperPrice int    `json:"upperPrice"`
		} `json:"valuations"`
	} `json:"history"`
	ID               string        `json:"id"`
	IsResidential    bool          `json:"isResidential"`
	LegacyPropertyID int           `json:"legacyPropertyId"`
	LotNumber        string        `json:"lotNumber"`
	OnMarketTypes    []interface{} `json:"onMarketTypes"`
	PdsID            int           `json:"pdsId"`
	Photos           []struct {
		AdvertID  int    `json:"advertId"`
		Date      string `json:"date"`
		FullURL   string `json:"fullUrl"`
		Rank      int    `json:"rank"`
		ImageType string `json:"imageType,omitempty"`
	} `json:"photos"`
	PlanNumber         string `json:"planNumber"`
	Postcode           string `json:"postcode"`
	PropertyCategory   string `json:"propertyCategory"`
	PropertyCategoryID int    `json:"propertyCategoryId"`
	PropertyType       string `json:"propertyType"`
	PropertyTypeID     int    `json:"propertyTypeId"`
	SectionNumber      string `json:"sectionNumber"`
	State              string `json:"state"`
	StreetAddress      string `json:"streetAddress"`
	StreetName         string `json:"streetName"`
	StreetNumber       string `json:"streetNumber"`
	StreetType         string `json:"streetType"`
	StreetTypeLong     string `json:"streetTypeLong"`
	Suburb             string `json:"suburb"`
	SuburbID           int    `json:"suburbId"`
	Timeline           []struct {
		SaleMetadata struct {
			IsSold      bool `json:"isSold"`
			MarketSteps []struct {
				Agency struct {
					Name string `json:"name"`
				} `json:"agency"`
				Date string `json:"date"`
				Type string `json:"type"`
			} `json:"marketSteps"`
		} `json:"saleMetadata"`
		Agency struct {
			Name string `json:"name"`
		} `json:"agency,omitempty"`
		Category         string  `json:"category"`
		DaysOnMarket     float64 `json:"daysOnMarket,omitempty"`
		EventDate        string  `json:"eventDate"`
		EventPrice       int     `json:"eventPrice"`
		PriceDescription string  `json:"priceDescription"`
		IsMajorEvent     bool    `json:"isMajorEvent"`
		AdvertID         int     `json:"advertId,omitempty"`
	} `json:"timeline"`
	Updated      time.Time `json:"updated"`
	URLSlug      string    `json:"urlSlug"`
	URLSlugShort string    `json:"urlSlugShort"`
	Valuation    struct {
		Date            string  `json:"date"`
		LowerPrice      int     `json:"lowerPrice"`
		PriceConfidence string  `json:"priceConfidence"`
		RentPerWeek     int     `json:"rentPerWeek"`
		RentYield       float64 `json:"rentYield"`
		UpperPrice      int     `json:"upperPrice"`
		MidPrice        int     `json:"midPrice"`
	} `json:"valuation"`
	Zone             string `json:"zone"`
	DefaultPhotoURL  string `json:"defaultPhotoUrl"`
	Occupancy        string `json:"occupancy"`
	LastSaleActivity struct {
		Date             string `json:"date"`
		DocumentedAsSold bool   `json:"documentedAsSold"`
		Price            int    `json:"price"`
		ReportedAsSold   bool   `json:"reportedAsSold"`
		Type             string `json:"type"`
		SuppressPrice    bool   `json:"suppressPrice"`
		SuppressDetails  string `json:"suppressDetails"`
		CapitalGrowth    struct {
			AnnualGrowth                float64 `json:"annualGrowth"`
			PreviousSoldPrice           int     `json:"previousSoldPrice"`
			PreviousSoldDate            string  `json:"previousSoldDate"`
			SuppressedPreviousSoldPrice bool    `json:"suppressedPreviousSoldPrice"`
		} `json:"capitalGrowth"`
		ActivityID int64 `json:"activityId"`
	} `json:"lastSaleActivity"`
	FlatNumberSortKey   string `json:"flatNumberSortKey"`
	StreetNumberSortKey string `json:"streetNumberSortKey"`
}