package models

// Calculation represents a type of calculation
type Calculation struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Formula  string `json:"formula" db:"formula"`
	Function string `json:"-" db:"function"`
}

// Parameter represents a calculation parameter
type Parameter struct {
	ID            int      `json:"id" db:"id"`
	CalculationID int      `json:"calculation_id" db:"calculation_id"`
	Symbol        string   `json:"symbol" db:"symbol"`
	Description   string   `json:"description" db:"description"`
	IsConstant    bool     `json:"is_constant" db:"is_constant"`
	Value         *float64 `json:"value" db:"value"`
	UnitID        int      `json:"unit_id" db:"unit_id"`
}

// Unit represents a unit of measurement
type Unit struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Symbol string `json:"symbol" db:"symbol"`
	Type   string `json:"type" db:"type"`
}
