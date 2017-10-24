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

type Suburb struct {
	Name  string
	State string
}

func (pss *PostcodeSuburbStore) GetSuburb(postcode int) ([]Suburb, bool) {
	var suburbsForPostcode []Suburb
	suburbs, found := pss.suburbs.Suburbs[int64(postcode)]

	if !found {
		return nil, false
	}
	for _, suburb := range suburbs {

		suburbsForPostcode = append(suburbsForPostcode, Suburb{
			Name:  suburb.Name,
			State: suburb.State.Code,
		})
	}

	return suburbsForPostcode, true
}
