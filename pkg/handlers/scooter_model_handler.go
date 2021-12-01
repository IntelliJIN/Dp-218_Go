package handlers

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pkg/services"
	"encoding/json"
	"net/http"
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
	_,err = sm.scooterModelService.CreateScooterModel(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterModelHandler) GetScooterModels(w http.ResponseWriter, r *http.Request) {
	var models []entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&models)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = sm.scooterModelService.GetScooterModels()
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (sm ScooterModelHandler) GetScooterModelByID(w http.ResponseWriter, r *http.Request) {
	var model entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = sm.scooterModelService.GetScooterModelByID(model.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterModelHandler) EditScooterModel(w http.ResponseWriter, r *http.Request) {
	var model entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = sm.scooterModelService.EditScooterModel(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterModelHandler) DeleteScooterModel(w http.ResponseWriter, r *http.Request) {
	var model entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = sm.scooterModelService.DeleteScooterModel(model.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}
