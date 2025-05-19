package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "sto-calculator/pkg/logging"

	"github.com/go-chi/chi/v5"
)

func (a *API) registerUnits(router chi.Router) {
	router.Route("/unit", func(router chi.Router) {
		router.Get("/{unitID}", a.getUnit)
	})
	//router.Route("/units", func(router chi.Router) {
	//	router.Get("/{unitType}", a.getUnitsByType)
	//})
}

func (a *API) getUnit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	unitID, err := strconv.Atoi(chi.URLParam(r, "unitID"))
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("unitID must be an integer: %v", err)
		http.Error(w, "unitID must be an integer", http.StatusBadRequest)
	}

	unit, err := a.service.GetUnit(ctx, unitID)
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("get unit by id error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(unit); err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("encode error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) getUnitsByType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	unitType := chi.URLParam(r, "unitType")

	units, err := a.service.GetUnitsByType(ctx, unitType)
	if err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("get units by type error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(units); err != nil {
		log.GetLoggerFromCtx(ctx).Errorf("encode error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
