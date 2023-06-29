package autoria

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCategories returns all categories
// Легкові, грузові, причепи і т.д. в массиві
func (s *service) GetCategories() (Categories, error) {
	var categories Categories
	if err := s.request(http.MethodGet, "/auto/categories", nil, nil, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

// GetBodyStyles returns all body styles for given parentID
// Example: GetBodyStyles(1) returns all body styles for "Легкові"
func (s *service) GetBodyStyles(parentID int) ([]BaseWithParentID, error) {
	var bodyStyles []BaseWithParentID
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/bodystyles", parentID), nil, nil, &bodyStyles); err != nil {
		return nil, err
	}
	return bodyStyles, nil
}

type ObjectOrArray any

// GetBodyStylesWithGroups returns all body styles for given parentID grouped by type
// Example: GetBodyStylesWithGroups(1) returns all body styles for "Легкові" grouped by type
func (s *service) GetBodyStylesWithGroups(parentID int) ([][]BaseWithParentID, error) {
	var raw json.RawMessage
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/bodystyles/_group", parentID), nil, nil, &raw); err != nil {
		return nil, err
	}

	var data []ObjectOrArray
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	var bodyStyles [][]BaseWithParentID
	for _, item := range data {
		switch value := item.(type) {
		case []interface{}:
			var group []BaseWithParentID
			for _, obj := range value {
				// Вызываем преобразование map в Object структуру
				objectData, _ := json.Marshal(obj)
				var object BaseWithParentID
				if err := json.Unmarshal(objectData, &object); err != nil {
					return nil, err
				}
				group = append(group, object)
			}
			bodyStyles = append(bodyStyles, group)
		case map[string]interface{}:
			// Преобразование map в Object структуру
			objectData, _ := json.Marshal(value)
			var object BaseWithParentID
			if err := json.Unmarshal(objectData, &object); err != nil {
				return nil, err
			}
			bodyStyles = append(bodyStyles, []BaseWithParentID{object})
		}
	}

	return bodyStyles, nil
}

// GetMarksByCategory returns all marks for given categoryID
// Example: GetMarksByCategory(1) returns all marks for "Легкові"
func (s *service) GetMarksByCategory(categoryID int) (Marks, error) {
	var marks Marks
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/marks", categoryID), nil, nil, &marks); err != nil {
		return nil, err
	}
	return marks, nil
}

// GetModelsByCategoryAndMarkID returns all models for given categoryID and markID
// Example: GetModelsByCategoryAndMarkID(1, 1) returns all models for "Легкові" and "Acura"
func (s *service) GetModelsByCategoryAndMarkID(categoryID, markID int) (Models, error) {
	var models Models
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/marks/%d/models", categoryID, markID), nil, nil, &models); err != nil {
		return nil, err
	}
	return models, nil
}

// GetModelsByCategoryAndMarkIDWithGroups returns all models for given categoryID and markID grouped by type
// Example: GetModelsByCategoryAndMarkIDWithGroups(1, 1) returns all models for "Легкові" and "Acura" grouped by type
func (s *service) GetModelsByCategoryAndMarkIDWithGroups(categoryID, markID int) (Models, error) {
	var raw json.RawMessage
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/marks/%d/models/_group", categoryID, markID), nil, nil, &raw); err != nil {
		return nil, err
	}

	var data []ObjectOrArray
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	var bodyStyles Models
	for _, item := range data {
		switch value := item.(type) {
		case []interface{}:
			var group Models
			for _, obj := range value {
				// Вызываем преобразование map в Object структуру
				objectData, _ := json.Marshal(obj)
				var object Base
				if err := json.Unmarshal(objectData, &object); err != nil {
					return nil, err
				}
				group = append(group, object)
			}
			bodyStyles = append(bodyStyles, group...)
		case map[string]interface{}:
			// Преобразование map в Object структуру
			objectData, _ := json.Marshal(value)
			var object Base
			if err := json.Unmarshal(objectData, &object); err != nil {
				return nil, err
			}
			bodyStyles = append(bodyStyles, object)
		}
	}

	return bodyStyles, nil
}

func (s *service) GetGenerationsByModelID(modelID int) ([]Generations, error) {
	var generations []Generations
	if err := s.request(http.MethodGet, fmt.Sprintf("/generations/by/models/%d/generations", modelID), nil, nil, &generations); err != nil {
		return nil, err
	}
	return generations, nil
}

// GetModificationsByGenerationID returns all modifications for given generationID
// Example: GetModificationsByGenerationID(1) де 1 - це id покоління з слайсу GetGenerationsByModelID
func (s *service) GetModificationsByGenerationID(generationID int) (Modifications, error) {
	var modifications Modifications
	if err := s.request(http.MethodGet, fmt.Sprintf("/modifications/by/generation/%d/modifications", generationID), nil, nil, &modifications); err != nil {
		return nil, err
	}
	return modifications, nil
}

// GetDriverTypes returns all driver types
// Типи приводу
func (s *service) GetDriverTypes(categoryID int) (DriverTypes, error) {
	var driverTypes DriverTypes
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/driverTypes", categoryID), nil, nil, &driverTypes); err != nil {
		return nil, err
	}
	return driverTypes, nil
}

// GetFuelTypes returns all fuel types
// Типи палива
func (s *service) GetFuelTypes() (FuelTypes, error) {
	var fuelTypes FuelTypes
	if err := s.request(http.MethodGet, "/auto/type", nil, nil, &fuelTypes); err != nil {
		return nil, err
	}
	return fuelTypes, nil
}

// GetGearboxTypes returns all gearbox types
// Типи коробки передач
func (s *service) GetGearboxTypes(categoryID int) (GearboxesTypes, error) {
	var gearboxTypes GearboxesTypes
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/gearboxes", categoryID), nil, nil, &gearboxTypes); err != nil {
		return nil, err
	}
	return gearboxTypes, nil
}

// GetOptions returns all options
// Опції
func (s *service) GetOptions(categoryID int) (OptionTypes, error) {
	var options OptionTypes
	if err := s.request(http.MethodGet, fmt.Sprintf("/auto/categories/%d/options", categoryID), nil, nil, &options); err != nil {
		return nil, err
	}
	return options, nil
}

// GetColors returns all colors
// Кольори
func (s *service) GetColors() (ColorTypes, error) {
	var colors ColorTypes
	if err := s.request(http.MethodGet, "/auto/colors", nil, nil, &colors); err != nil {
		return nil, err
	}
	return colors, nil
}
