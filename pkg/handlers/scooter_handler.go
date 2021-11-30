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
	GetScooterById(w http.ResponseWriter, r *http.Request)
	GetScooterByEmail(w http.ResponseWriter, r *http.Request)
	EditScooter(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetScooterStations(w http.ResponseWriter, r *http.Request)
}

func (s ScooterHandler) GetScooterStations(w http.ResponseWriter, r *http.Request) {
	var stations []entities.ScooterStation
	err := json.NewDecoder(r.Body).Decode(&stations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = s.scooterService.ShowScooterStation()
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s ScooterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = s.scooterService.CreateScooter(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s ScooterHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var scooters []entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = s.scooterService.GetScooters()
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s ScooterHandler) GetById(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = s.scooterService.GetScooterByID(scooter.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s ScooterHandler) GetByModelName(w http.ResponseWriter, r *http.Request) {
	var scooters []entities.Scooter
	var model entities.ScooterModel
	err := json.NewDecoder(r.Body).Decode(&scooters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_, err = s.scooterService.GetScooterByModelName(model.ModelName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s ScooterHandler) EditInfo(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = s.scooterService.EditScooter(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s ScooterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	_,err = s.scooterService.DeleteScooter(scooter.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}
