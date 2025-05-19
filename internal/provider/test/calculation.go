package test

import (
	"context"

	"sto-calculator/internal/models"
)

type Test struct{}

func (t *Test) ListCalculations(ctx context.Context) ([]models.Calculation, error) {
	return []models.Calculation{
		{
			ID:      1,
			Name:    "Тестовый расчет",
			Formula: "P = Q * t",
		},
	}, nil
}

func (t *Test) GetCalculation(ctx context.Context, calculationID string) (models.Calculation, error) {
	return models.Calculation{
		ID:      1,
		Name:    "Тестовый расчет",
		Formula: "P = Q * t",
	}, nil
}

func (t *Test) GetCalculationParameters(ctx context.Context, calculationID string) ([]models.Parameter, error) {
	return []models.Parameter{
		{
			ID:         1,
			Name:       "Q",
			IsConstant: false,
			UnitID:     1,
		},
		{
			ID:         2,
			Name:       "t",
			IsConstant: false,
			UnitID:     1,
		},
	}, nil
}

func (t *Test) GetUnitsByType(ctx context.Context, unitType string) ([]models.Unit, error) {
	return []models.Unit{
		{
			ID:     1,
			Name:   "Килограмм",
			Symbol: "кг",
			Type:   "mass",
		},
	}, nil
}
