package handlers

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pkg/services"
	"encoding/json"
	"net/http"
)

type ScooterHandler struct {
	scooterService services.ScooterServiceI
}

func NewScooterHandler(scooterService services.ScooterServiceI) *ScooterHandler {
	return &ScooterHandler{
		scooterService: scooterService,
	}
}

type ScooterHandlerI interface {
	CreateScooter(w http.ResponseWriter, r *http.Request)
	GetScooters(w http.ResponseWriter, r *http.Request)
	UpdateScooterSerial(w http.ResponseWriter, r *http.Request)
	DeleteScooter(w http.ResponseWriter, r *http.Request)
	GetScooterByModelId(w http.ResponseWriter, r *http.Request)
	GetScooterByModelName(w http.ResponseWriter, r *http.Request)
	GetScooterByID(w http.ResponseWriter, r *http.Request)
}

func (sm ScooterHandler) CreateScooter(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = sm.scooterService.CreateScooter(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) GetScooters(w http.ResponseWriter, r *http.Request) {
	var scooters []entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = sm.scooterService.GetScooters()
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) UpdateScooterSerial(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = sm.scooterService.UpdateScooterSerial(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) DeleteScooter(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = sm.scooterService.DeleteScooter(scooter.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) GetScooterByModelId(w http.ResponseWriter, r *http.Request) {
	var scooters []entities.Scooter
	var model entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&scooters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = sm.scooterService.GetScooterByModelId(model.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) GetScooterByModelName(w http.ResponseWriter, r *http.Request) {
	var scooters []entities.Scooter
	var model entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&scooters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = sm.scooterService.GetScooterByModelName(model.ModelName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) GetScooterById(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = sm.scooterService.GetScooterByID(scooter.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

