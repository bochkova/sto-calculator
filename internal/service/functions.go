package service

import (
	"fmt"
	"reflect"

	log "sto-calculator/pkg/logging"
)

type Functions struct{}

func NewFunctions() *Functions {
	return &Functions{}
}

func (f *Functions) Exist(fnName string) bool {
	val := reflect.ValueOf(f)
	method := val.MethodByName(fnName)
	log.WithFields(log.Fields{
		"function_name":     fnName,
		"is_valid":          method.IsValid(),
		"available_methods": val.NumMethod(),
	}).Debug("Checking function existence")
	return method.IsValid()
}

func (f *Functions) Execute(fnName string, params map[string]float64) (float64, error) {
	fn := reflect.ValueOf(f).MethodByName(fnName)
	if !fn.IsValid() {
		return 0, fmt.Errorf("calculation function %s not found", fnName)
	}

	args := []reflect.Value{
		reflect.ValueOf(params),
	}

	return fn.Call(args)[0].Float(), nil
}

func (f *Functions) EmptyingOfEquipment(params map[string]float64) float64 {
	return params["K"] * params["V"] * (params["Pн"]/(params["Tн"]*params["Zн"]) - params["Pк"]/(params["Tк"]*params["Zк"])) * params["n"]
}

func (f *Functions) BlowingOfEquipment(params map[string]float64) float64 {
	return params["K"] * params["V"] * params["Pср"] / (params["T"] * params["Z"]) * params["nпр"]
}

func (f *Functions) AfterRepairOfEquipment(params map[string]float64) float64 {
	return params["K"] * params["V"] * params["P"] / (params["T"] * params["Z"]) * (params["b"] + 1) * params["nв"]
}

func (f *Functions) BeforeRepairOfAdsorbers(params map[string]float64) float64 {
	return (params["K"]*params["Pн"]/(params["Tн"]*params["Zн"])*(params["Vад"]-(params["k"]*params["Gад"])/params["pн.ад"]) + (params["a"]*params["Gад"]*params["v"])/params["M"]) * (params["n"] + params["m"])
}

func (f *Functions) AfterRepairOfAdsorbers(params map[string]float64) float64 {
	return params["K"] * params["P"] / (params["T"] * params["Z"]) * (params["Vад"] - (params["k"]*params["Gад"])/params["pн.ад"]) * (params["n"] + params["m"]) * (params["b"] + 1)
}

func (f *Functions) TitrometricAnalyses(params map[string]float64) float64 {
	return params["Vп.п.р"] * params["nан.т"]
}

func (f *Functions) IndividualTestsOfEquipment(params map[string]float64) float64 {
	return params["Vисп"] * params["tисп"]
}

func (f *Functions) DischargeOfReservoirWater(params map[string]float64) float64 {
	return params["Vж"] * params["r"] * params["t"]
}

func (f *Functions) PurgingOfWell(params map[string]float64) float64 {
	return params["q"] * params["t"] * params["n"]
}

func (f *Functions) SubcriticalExpiration(params map[string]float64) float64 {
	return 1121.67 * params["F"] * params["P"] * params["tнкр"]
}

func (f *Functions) CriticalExpiration(params map[string]float64) float64 {
	return 3018.31 * params["F"] * params["P"] * params["tкр"]
}

func (f *Functions) OperationOfFlareSystem(params map[string]float64) float64 {
	return 3600*params["w"]*params["F"]*params["tф"] + params["V"]*params["nг"]*params["tг"]
}

func (f *Functions) TreatmentOfGasCondensate(params map[string]float64) float64 {
	return params["Gк"] * params["N"]
}

func (f *Functions) DegassingOfLiquids(params map[string]float64) float64 {
	return params["Vж"] * params["r"] * params["t"]
}
