package services

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pkg/repository"
)

type ScooterServiceI interface {
	CreateScooter(scooter *entities.Scooter) (int, error)
	GetScooters() (*[]entities.Scooter, error)
	UpdateScooterSerial(scooter *entities.Scooter) (int, error)
	DeleteScooter(id int) (int, error)
	GetScooterByModelId(ide int) (*[]entities.Scooter, error)
	GetScooterByModelName(name string) (*[]entities.Scooter, error)
	GetScooterByID(scooterID int) (*entities.Scooter, error)
}

func NewScooterService(scooterRepository repository.ScooterRepositoryI) *ScooterService {
	return &ScooterService{
		scooterRepository,
	}
}

type ScooterService struct {
	scooterRepository repository.ScooterRepositoryI
}

func (s ScooterService) CreateScooter(scooter *entities.Scooter) (int, error) {
	lastID, err := s.scooterRepository.CreateScooter(scooter)
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

func (s ScooterService) GetScooters() (*[]entities.Scooter, error) {
	scooters, err := s.scooterRepository.GetAllScooters()
	if err != nil {
		return nil, err
	}
	return scooters, nil
}

func (s ScooterService) UpdateScooterSerial(scooter *entities.Scooter) (int, error) {
	rowsAffected, err := s.scooterRepository.UpdateScooterSerial(scooter)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (s ScooterService) DeleteScooter(scooterID int) (int, error) {
	rowsAffected, err := s.scooterRepository.DeleteScooter(scooterID)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (s ScooterService) GetScooterByModelId(id int) (*[]entities.Scooter, error) {
	scooter, err := s.scooterRepository.GetScooterByModelId(id)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}

func (s ScooterService) GetScooterByModelName(modelName string) (*[]entities.Scooter, error) {
	scooter, err := s.scooterRepository.GetScooterByModelName(modelName)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}

func (s ScooterService) GetScooterByID(scooterID int) (*entities.Scooter, error) {
	scooter, err := s.scooterRepository.GetScooterByID(scooterID)
	if err != nil {
		return nil, err
	}
	return scooter, nil
}













