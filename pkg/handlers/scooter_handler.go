package handlers

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pkg/services"
	"encoding/json"
	"net/http"
	"strconv"
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
	sm.scooterService.CreateScooter(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) GetScooters(w http.ResponseWriter, r *http.Request) {
	p, err := sm.scooterService.GetScooters()
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var resp []entities.Scooter
	for _, x := range *p {
		resp = append(
			resp,
			entities.Scooter{Id: x.Id, ModelId: x.ModelId, OwnerId: x.OwnerId, SerialNumber: x.SerialNumber},
		)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (sm ScooterHandler) UpdateScooterSerial(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Scooter
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := sm.scooterService.UpdateScooterSerial(&entities.Scooter{Id: scooter.Id,
		SerialNumber: scooter.SerialNumber})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "nothing was changed", http.StatusNotModified)
		return
	}
}

func (sm ScooterHandler) DeleteScooter(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	_, err := sm.scooterService.DeleteScooter(ID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("scooter successfully deleted"))
}

func (sm ScooterHandler) GetScooterByModelId(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.Atoi(r.URL.Query().Get("model_id"))

	_, err := sm.scooterService.GetScooterByModelId(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) GetScooterByModelName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("model_name")

	_, err := sm.scooterService.GetScooterByModelName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sm ScooterHandler) GetScooterById(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	_, err := sm.scooterService.GetScooterByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

