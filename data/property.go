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
	PriceHistory  []struct {
		Price    string
		DateSole time.Time
	}
	Description string
	Address     Address
	NumBeds     string
	NumBaths    string
	NumCars     string
}

type DomainComAuPropertyListWrapper struct {
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
	PropertyObject          DomainComAuPropertyHistory `json:"propertyObject"`
	ShowBuildingProfileLink bool                       `json:"showBuildingProfileLink"`
	ShowStreetProfileLink   bool                       `json:"showStreetProfileLink"`
	StickyHeaderEnabled     bool                       `json:"stickyHeaderEnabled"`
	StreetProfileURLSlug    string                     `json:"streetProfileUrlSlug"`
	TermsInput              string                     `json:"termsInput"`
	TypeaheadHintCSS        string                     `json:"typeaheadHintCss"`
	TypeaheadMenuCSS        string                     `json:"typeaheadMenuCss"`
	TypeaheadSuggestCSS     string                     `json:"typeaheadSuggestCss"`
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

type DomainComAuPropertyListRentWrapper struct {
	Page struct {
		PageInfo struct {
			PageID   string `json:"pageId"`
			PageName string `json:"pageName"`
			Search   struct {
				AgencyIds             string `json:"agencyIds"`
				BathroomsFrom         string `json:"bathroomsFrom"`
				BathroomsTo           string `json:"bathroomsTo"`
				BedroomsFrom          string `json:"bedroomsFrom"`
				BedroomsTo            string `json:"bedroomsTo"`
				CarSpaces             string `json:"carSpaces"`
				GeoType               string `json:"geoType"`
				MapSearch             bool   `json:"mapSearch"`
				MedianPrice           int    `json:"medianPrice"`
				Postcode              string `json:"postcode"`
				PrimaryPropertyType   string `json:"primaryPropertyType"`
				SecondaryPropertyType string `json:"secondaryPropertyType"`
				ResultsPages          int    `json:"resultsPages"`
				ResultsRecords        string `json:"resultsRecords"`
				SearchArea            string `json:"searchArea"`
				SearchDepth           int    `json:"searchDepth"`
				SearchLocationCat     string `json:"searchLocationCat"`
				SearchRegion          string `json:"searchRegion"`
				SearchResultCount     int    `json:"searchResultCount"`
				SearchSuburb          string `json:"searchSuburb"`
				SearchTerm            string `json:"searchTerm"`
				SearchTypeView        string `json:"searchTypeView"`
				SortBy                string `json:"sortBy"`
				State                 string `json:"state"`
				SuburbID              string `json:"suburbId"`
				SurroundingSuburbs    string `json:"surroundingSuburbs"`
			} `json:"search"`
			Brand         string `json:"brand"`
			Generator     string `json:"generator"`
			SysEnv        string `json:"sysEnv"`
			IsEmbeddedApp bool   `json:"isEmbeddedApp"`
		} `json:"pageInfo"`
		Category struct {
			PrimaryCategory string `json:"primaryCategory"`
			SubCategory1    string `json:"subCategory1"`
			PageType        string `json:"pageType"`
		} `json:"category"`
		AbTest5 string `json:"abTest5"`
	} `json:"page"`
	Titan struct {
		AdZone      string `json:"adZone"`
		AdKeyValues struct {
			Cat         string   `json:"cat"`
			Cat1        string   `json:"cat1"`
			Ctype       string   `json:"ctype"`
			E           string   `json:"e"`
			Locstate    string   `json:"locstate"`
			Locarea     string   `json:"locarea"`
			Locsuburb   []string `json:"locsuburb"`
			Locpostcode []string `json:"locpostcode"`
			Usertype    string   `json:"usertype"`
			MedianPrice int      `json:"medianPrice"`
			Bedexact    []string `json:"bedexact"`
		} `json:"adKeyValues"`
		AdSite  string        `json:"adSite"`
		AdSlots []interface{} `json:"adSlots"`
	} `json:"titan"`
	User struct {
		MembershipType  string `json:"membershipType"`
		SessionToken    string `json:"sessionToken"`
		MembershipState string `json:"membershipState"`
		IPAddress       string `json:"ipAddress"`
	} `json:"user"`
	Version string        `json:"version"`
	Events  []interface{} `json:"events"`
}

type DomainComAuPropertyDetailWrapper struct {
	Page struct {
		PageInfo struct {
			Author    string `json:"author"`
			Brand     string `json:"brand"`
			Generator string `json:"generator"`
			PageID    string `json:"pageId"`
			PageName  string `json:"pageName"`
			SysEnv    string `json:"sysEnv"`
			Property  struct {
				Address               string `json:"address"`
				Agency                string `json:"agency"`
				AgentNames            string `json:"agentNames"`
				AdType                string `json:"adType"`
				Bedrooms              string `json:"bedrooms"`
				Bathrooms             string `json:"bathrooms"`
				FloorPlansCount       string `json:"floorPlansCount"`
				Landsize              string `json:"landsize"`
				Buildingsize          string `json:"buildingsize"`
				Parking               string `json:"parking"`
				PhotoCount            string `json:"photoCount"`
				Postcode              string `json:"postcode"`
				Price                 string `json:"price"`
				PrimaryPropertyType   string `json:"primaryPropertyType"`
				SecondaryPropertyType string `json:"secondaryPropertyType"`
				State                 string `json:"state"`
				Suburb                string `json:"suburb"`
				VideoCount            string `json:"videoCount"`
				PropertyID            string `json:"propertyId"`
				AgencyID              string `json:"agencyId"`
			} `json:"property"`
			SuburbID string `json:"suburbId"`
		} `json:"pageInfo"`
		Category struct {
			PrimaryCategory string `json:"primaryCategory"`
			SubCategory1    string `json:"subCategory1"`
			PageType        string `json:"pageType"`
		} `json:"category"`
	} `json:"page"`
	Titan struct {
		AdSite      string `json:"adSite"`
		AdZone      string `json:"adZone"`
		AdKeyValues struct {
			Cat         string   `json:"cat"`
			Cat1        string   `json:"cat1"`
			Ctype       string   `json:"ctype"`
			Locstate    string   `json:"locstate"`
			Locarea     string   `json:"locarea"`
			Locsuburb   []string `json:"locsuburb"`
			Locpostcode []string `json:"locpostcode"`
			Usertype    string   `json:"usertype"`
			MedianPrice string   `json:"medianPrice"`
			E           string   `json:"e"`
			Proptypes   []string `json:"proptypes"`
		} `json:"adKeyValues"`
		AdSlots      []string `json:"adSlots"`
		ContentWidth string   `json:"contentWidth"`
	} `json:"titan"`
	User struct {
		Profile struct {
			ProfileInfo struct {
			} `json:"profileInfo"`
		} `json:"profile"`
		MembershipType  string `json:"membershipType"`
		MembershipState string `json:"membershipState"`
		IPAddress       string `json:"ipAddress"`
	} `json:"user"`
	Version string        `json:"version"`
	Events  []interface{} `json:"events"`
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
