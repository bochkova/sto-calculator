package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"sto-calculator/internal/models"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d *DB) ListCalculations(ctx context.Context) ([]models.Calculation, error) {
	var calculations []models.Calculation
	err := d.db.SelectContext(ctx, &calculations, `
		SELECT id, name, formula, function
		FROM calculations
	`)
	if err != nil {
		return nil, err
	}
	return calculations, nil
}

func (d *DB) GetCalculation(ctx context.Context, calculationID int) (models.Calculation, error) {
	var calc models.Calculation
	err := d.db.GetContext(ctx, &calc, `
		SELECT id, name, formula, function
		FROM calculations 
		WHERE id = $1
	`, calculationID)
	return calc, err
}

func (d *DB) GetCalculationParameters(ctx context.Context, calculationID int) ([]models.Parameter, error) {
	var parameters []models.Parameter
	err := d.db.SelectContext(ctx, &parameters, `
		SELECT id, calculation_id, symbol, description, is_constant, value, unit_id
		FROM parameters
		WHERE calculation_id = $1
	`, calculationID)
	if err != nil {
		return nil, err
	}
	return parameters, nil
}

func (d *DB) GetUnit(ctx context.Context, unitID int) (models.Unit, error) {
	var unit models.Unit
	err := d.db.GetContext(ctx, &unit, `
		SELECT id, name, symbol, type
		FROM units
		WHERE id = $1
	`, unitID)
	return unit, err
}

func (d *DB) GetUnitsByType(ctx context.Context, unitType string) ([]models.Unit, error) {
	var units []models.Unit
	err := d.db.SelectContext(ctx, &units, `
		SELECT id, name, symbol, type
		FROM units
		WHERE type = $1
	`, unitType)
	if err != nil {
		return nil, err
	}
	return units, nil
}
