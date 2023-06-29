package autoria

// Base is a base struct for all categories
type Base struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type BaseWithParentID struct {
	Base
	ParentID int `json:"parent_id"`
}

// Categories тип з категоріями авто
// Categories is a slice of Base
type Categories []Base

// Marks тип з марками авто
// Example: Marks{{Name: "Audi", Value: 1}, {Name: "BMW", Value: 2}}
type Marks []Base

// Models тип з моделями авто
// Example: Models{{Name: "A4", Value: 1}, {Name: "A6", Value: 2}}
type Models []Base

// Modifications тип з модифікаціями авто
// Example: Modifications{{Name: "1.8", Value: 1}, {Name: "2.0", Value: 2}}
type Modifications []Base

// DriverTypes тип з модифікаціями авто
// Example: DriverTypes{{Name: "1.8", Value: 1}, {Name: "2.0", Value: 2}}
type DriverTypes []Base

// FuelTypes тип з модифікаціями авто
// Example: FuelTypes{{Name: "1.8", Value: 1}, {Name: "2.0", Value: 2}}
type FuelTypes []Base

// GearboxesTypes тип з модифікаціями авто
// Example: GearboxesTypes{{Name: "1.8", Value: 1}, {Name: "2.0", Value: 2}}
type GearboxesTypes []Base

// OptionTypes тип з модифікаціями авто
// Example: OptionTypes{{Name: "1.8", Value: 1}, {Name: "2.0", Value: 2}}
type OptionTypes []Base

// ColorTypes тип з модифікаціями авто
// Example: ColorTypes{{Name: "1.8", Value: 1}, {Name: "2.0", Value: 2}}
type ColorTypes []Base

// CountryTypes тип з модифікаціями авто
// Example: CountryTypes{{Name: "1.8", Value: 1}, {Name: "2.0", Value: 2}}
type CountryTypes []Base

// States області України
// Example: States{{Name: "Вінницька", Value: 1}, {Name: "Волинська", Value: 18}}
type States []Base

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RestError struct {
	ErrorBody `json:"error"`
}

func (e RestError) Error() string {
	return e.Message
}

type Generation struct {
	GenerationId int    `json:"generationId"`
	Name         string `json:"name"`
	YearFrom     int    `json:"yearFrom"`
	YearTo       int    `json:"yearTo"`
	ModelId      int    `json:"modelId"`
	Eng          string `json:"eng"`
}

type Generations struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Generations []Generation `json:"generations"`
}
