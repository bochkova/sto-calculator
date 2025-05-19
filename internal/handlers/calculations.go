package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "sto-calculator/pkg/logging"

	"github.com/go-chi/chi/v5"
)

func (a *API) registerCalculation(router chi.Router) {
	router.Route("/calculations", func(router chi.Router) {
		router.Get("/", a.listCalculations)
		router.Get("/{calculationID}", a.getCalculation)
		router.Get("/{calculationID}/parameters", a.getCalculationParameters)
		router.Post("/{calculationID}/execute", a.executeCalculation)
	})
}

func (a *API) listCalculations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	calculations, err := a.service.ListCalculations(ctx)
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("get calculations list error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(calculations); err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("encode error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) getCalculation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	calculationID, err := strconv.Atoi(chi.URLParam(r, "calculationID"))
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("calculationID must be an integer: %v", err)
		http.Error(w, "calculationID must be an integer", http.StatusBadRequest)
	}

	calculation, err := a.service.GetCalculation(ctx, calculationID)
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("get calculation error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(calculation); err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("encode error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) getCalculationParameters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	calculationID, err := strconv.Atoi(chi.URLParam(r, "calculationID"))
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("calculationID must be an integer: %v", err)
		http.Error(w, "calculationID must be an integer", http.StatusBadRequest)
	}

	parameters, err := a.service.GetCalculationParameters(ctx, calculationID)
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("get calculation parameters error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(parameters); err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("encode error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) executeCalculation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	calculationID, err := strconv.Atoi(chi.URLParam(r, "calculationID"))
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("calculationID must be an integer: %v", err)
		http.Error(w, "calculationID must be an integer", http.StatusBadRequest)
	}

	// ExecuteCalculationRequest represents the parameters needed to execute a calculation
	type ExecuteCalculationRequest struct {
		Parameters map[string]float64 `json:"parameters"`
	}

	var req ExecuteCalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("failed to parse body: %v", err)
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	result, err := a.service.ExecuteCalculation(ctx, calculationID, req.Parameters)
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("execute calculation error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("encode error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
