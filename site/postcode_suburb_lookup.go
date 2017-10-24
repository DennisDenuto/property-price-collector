package site

import "github.com/tomjowitt/ozdata/ozdata"

type PostcodeSuburbStore struct {
	suburbs ozdata.Suburbs
}

func (pss *PostcodeSuburbStore) Load() error {
	suburbs, err := ozdata.LoadSuburbs()
	if err != nil {
		return err
	}

	pss.suburbs = suburbs
	return nil
}

func (pss *PostcodeSuburbStore) GetSuburb(postcode int) ([]string, bool) {
	var suburbsForPostcode []string
	suburbs, found := pss.suburbs.Suburbs[int64(postcode)]

	if !found {
		return nil, false
	}
	for _, suburb := range suburbs {
		suburbsForPostcode = append(suburbsForPostcode, suburb.Name)
	}

	return suburbsForPostcode, true
}
