package services

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pkg/repository"
)

type ScooterModelServiceI interface {
	CreateScooterModel(scooterModel *entities.ScooterModel) (int, error)
	GetScooterModels() (*[]entities.ScooterModel, error)
	GetScooterModelByID(modelID int) (*entities.ScooterModel, error)
	EditScooterModel(model *entities.ScooterModel) (int, error)
	DeleteScooterModel(id int) (int, error)
}

func NewScooterModelService(scooterModelRepository repository.ScooterModelRepositoryI) *ScooterModelService {
	return &ScooterModelService{
		scooterModelRepository,
	}
}

type ScooterModelService struct {
	scooterModelRepository repository.ScooterModelRepositoryI
}

func (sm ScooterModelService) CreateScooterModel(scooterModel *entities.ScooterModel) (int, error) {
	lastID, err := sm.scooterModelRepository.CreateScooterModel(scooterModel)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (sm ScooterModelService) GetScooterModels() (*[]entities.ScooterModel, error) {
	scooters, err := sm.scooterModelRepository.GetScooterModels()
	if err != nil {
		return nil, err
	}
	return scooters, nil
}

func (sm ScooterModelService) EditScooterModel(model *entities.ScooterModel) (int, error) {
	rowsAffected, err := sm.scooterModelRepository.UpdateScooterModel(model)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (sm ScooterModelService) DeleteScooterModel(scooterID int) (int, error) {
	rowsAffected, err := sm.scooterModelRepository.DeleteScooterModel(scooterID)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (sm ScooterModelService) GetScooterModelByID(modelID int) (*entities.ScooterModel, error) {
	scooter, err := sm.scooterModelRepository.GetScooterModelById(modelID)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}