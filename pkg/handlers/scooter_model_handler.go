package handlers

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pkg/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ScooterModelHandler struct {
	scooterModelService services.ScooterModelServiceI
}

func NewScooterModelHandler(scooterModelService services.ScooterModelServiceI) *ScooterModelHandler {
	return &ScooterModelHandler{
		scooterModelService: scooterModelService,
	}
}

type ScooterModelHandlerI interface {
	CreateScooterModel(w http.ResponseWriter, r *http.Request)
	GetScooterModels(w http.ResponseWriter, r *http.Request)
	GetScooterModelByID(w http.ResponseWriter, r *http.Request)
	EditScooterModel(w http.ResponseWriter, r *http.Request)
	DeleteScooterModel(w http.ResponseWriter, r *http.Request)
}

func (sm ScooterModelHandler) CreateScooterModel(w http.ResponseWriter, r *http.Request) {
	var model entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	 sm.scooterModelService.CreateScooterModel(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterModelHandler) GetScooterModels(w http.ResponseWriter, r *http.Request) {
	p, err := sm.scooterModelService.GetScooterModels()
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []entities.ScooterModel
	for _, x := range *p {
		resp = append(
			resp,
			entities.ScooterModel{Id: x.Id, PaymentTypeId: x.PaymentTypeId, ModelName: x.ModelName, MaxWeight: x.Speed},
		)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (sm ScooterModelHandler) GetScooterModelByID(w http.ResponseWriter, r *http.Request) {
	modelID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	models, err := sm.scooterModelService.GetScooterModelByID(modelID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models)
}

func (sm ScooterModelHandler) EditScooterModel(w http.ResponseWriter, r *http.Request) {
	updateModelRequest := new(entities.ScooterModel)
	err := json.NewDecoder(r.Body).Decode(&updateModelRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := sm.scooterModelService.EditScooterModel(&entities.ScooterModel{Id: updateModelRequest.Id,
		PaymentTypeId: updateModelRequest.PaymentTypeId,
		ModelName: updateModelRequest.ModelName, MaxWeight: updateModelRequest.MaxWeight,Speed: updateModelRequest.Speed})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "nothing was changed", http.StatusNotModified)
		return
	}
}

func (sm ScooterModelHandler) DeleteScooterModel(w http.ResponseWriter, r *http.Request) {
	modelID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := sm.scooterModelService.DeleteScooterModel(modelID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("model successfully deleted"))
}
