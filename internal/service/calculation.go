package service

import (
	"context"
	"fmt"

	"sto-calculator/internal/models"
)

type DB interface {
	ListCalculations(ctx context.Context) ([]models.Calculation, error)
	GetCalculation(ctx context.Context, calculationID int) (models.Calculation, error)
	GetCalculationParameters(ctx context.Context, calculationID int) ([]models.Parameter, error)
	GetUnit(ctx context.Context, unitID int) (models.Unit, error)
	GetUnitsByType(ctx context.Context, unitType string) ([]models.Unit, error)
}

type FunctionExecutor interface {
	Exist(fn string) bool
	Execute(fn string, parameters map[string]float64) (float64, error)
}

type Service struct {
	db       DB
	executor FunctionExecutor
}

func NewService(db DB, executor FunctionExecutor) *Service {
	return &Service{
		db:       db,
		executor: executor,
	}
}

func (s *Service) Init(ctx context.Context) error {
	calculations, err := s.db.ListCalculations(ctx)
	if err != nil {
		return err
	}

	notExist := make([]string, 0, len(calculations))
	for _, calculation := range calculations {
		if !s.executor.Exist(calculation.Function) {
			notExist = append(notExist, calculation.Function)
		}
	}

	if len(notExist) > 0 {
		return fmt.Errorf("functions %v needed to be implemented", notExist)
	}

	return nil
}

func (s *Service) ListCalculations(ctx context.Context) ([]models.Calculation, error) {
	return s.db.ListCalculations(ctx)
}

func (s *Service) GetCalculation(ctx context.Context, calculationID int) (models.Calculation, error) {
	return s.db.GetCalculation(ctx, calculationID)
}

func (s *Service) GetCalculationParameters(ctx context.Context, calculationID int) ([]models.Parameter, error) {
	return s.db.GetCalculationParameters(ctx, calculationID)
}

func (s *Service) ExecuteCalculation(ctx context.Context, calculationID int, parameters map[string]float64) (float64, error) {
	calculation, err := s.db.GetCalculation(ctx, calculationID)
	if err != nil {
		return 0, err
	}

	if err = s.checkParameters(ctx, calculation, parameters); err != nil {
		return 0, err
	}

	return s.executor.Execute(calculation.Function, parameters)
}

func (s *Service) GetUnit(ctx context.Context, unitID int) (models.Unit, error) {
	return s.db.GetUnit(ctx, unitID)
}

func (s *Service) GetUnitsByType(ctx context.Context, unitType string) ([]models.Unit, error) {
	return s.db.GetUnitsByType(ctx, unitType)
}

func (s *Service) checkParameters(ctx context.Context, calculation models.Calculation, parameters map[string]float64) error {
	requiredParams, err := s.db.GetCalculationParameters(ctx, calculation.ID)
	if err != nil {
		return err
	}

	for _, param := range requiredParams {
		if _, ok := parameters[param.Symbol]; !ok {
			return fmt.Errorf("parameter %s is required", param.Symbol)
		}
	}

	return nil
}
