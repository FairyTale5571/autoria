package autoria

import (
	"fmt"
	"net/http"
)

// GetStates returns all states of Ukraine
func (s *service) GetStates() (States, error) {
	var states States
	if err := s.request(http.MethodGet, "/auto/states", nil, nil, &states); err != nil {
		return nil, err
	}
	return states, nil
}

// GetCities returns all cities for given stateID
func (s *service) GetCities(stateID int) (States, error) {
	var cities States
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/states/%d/cities", stateID), nil, nil, &cities); err != nil {
		return nil, err
	}
	return cities, nil
}

// GetCountries returns all countries
func (s *service) GetCountries() (CountryTypes, error) {
	var countries CountryTypes
	if err := s.request(http.MethodGet, "/auto/countries", nil, nil, &countries); err != nil {
		return nil, err
	}
	return countries, nil
}
