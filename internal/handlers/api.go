package handlers

import (
	"context"

	"sto-calculator/internal/models"

	"github.com/go-chi/chi/v5"
)

type Service interface {
	ListCalculations(ctx context.Context) ([]models.Calculation, error)
	GetCalculation(ctx context.Context, calculationID int) (models.Calculation, error)
	GetCalculationParameters(ctx context.Context, calculationID int) ([]models.Parameter, error)
	ExecuteCalculation(ctx context.Context, calculationID int, parameters map[string]float64) (float64, error)
	GetUnit(ctx context.Context, unitID int) (models.Unit, error)
	GetUnitsByType(ctx context.Context, unitType string) ([]models.Unit, error)
}

type API struct {
	service Service
}

func NewHandlers(service Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) RegisterHandlers(router chi.Router) {
	a.registerCalculation(router)
	a.registerUnits(router)
}
