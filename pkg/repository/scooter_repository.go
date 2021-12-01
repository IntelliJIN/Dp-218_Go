package repository

import (
	"Dp-218_Go/entities"
	"context"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type ScooterRepository struct {
	db *pgx.Conn
}


func NewScooterRepository(db *pgx.Conn) *ScooterRepository {
	return &ScooterRepository{
		db: db,
	}
}

type ScooterRepositoryI interface {
	CreateScooter(scooter *entities.Scooter) (int, error)
	GetAllScooters() (*[]entities.Scooter, error)
	UpdateScooterSerial(scooter *entities.Scooter) (int, error)
	DeleteScooter(id int) (int, error)
	GetScooterByModelId(id int) (*[]entities.Scooter, error)
	GetScooterByModelName(name string) (*[]entities.Scooter, error)
	GetScooterByID(id int) (*entities.Scooter, error)
}

func (sm ScooterRepository) CreateScooter(scooter *entities.Scooter) (int, error) {
	res, err := sm.db.Exec(context.Background(),"INSERT INTO scooters (id, entities, brand, max_distance, capacity, max_weight) VALUES ($1, $2, $3, $4, $5, $6)",
		0, &scooter.Id, &scooter.ModelId, &scooter.OwnerId, &scooter.SerialNumber)
	if err != nil {
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	lastID := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (sm ScooterRepository) GetAllScooters() (*[]entities.Scooter, error) {
	var scooters []entities.Scooter
	rows, err := sm.db.Query(context.Background(),"SELECT * FROM scooters")

	if err != nil {
		return nil, err
	}
	scooter := entities.Scooter{}
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.ModelId, &scooter.OwnerId, &scooter.SerialNumber)
		if err != nil {
			return nil, err
		}
		scooters = append(scooters, scooter)
	}
	return &scooters, nil
}

func (sm ScooterRepository) UpdateScooterSerial(scooter *entities.Scooter) (int, error) {
	res, err := sm.db.Exec(context.Background(), "UPDATE scooters SET serial_number=$1 WHERE id=$2",
		&scooter.SerialNumber, &scooter.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := res.RowsAffected()
	return int(rowsAffected), nil
}

func (sm ScooterRepository) DeleteScooter(id int) (int, error) {
	res, err := sm.db.Exec(context.Background(), "`DELETE FROM scooters WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsAffected := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (sm ScooterRepository) GetScooterByModelId(modelId int) (*[]entities.Scooter, error) {
	var scooters []entities.Scooter

	var scooter entities.Scooter
	rows, err := sm.db.Query(context.Background(), "SELECT * FROM scooters WHERE model_id=$1", modelId)
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.ModelId, &scooter.OwnerId, &scooter.SerialNumber)
		if err != nil {
			return nil, err
		}
		scooters = append(scooters, scooter)
	}

	return &scooters, nil
}

func (sm ScooterRepository) GetScooterByModelName(name string) (*[]entities.Scooter, error) {
	var scooterModel entities.ScooterModel
	var scooters []entities.Scooter
	row := sm.db.QueryRow(context.Background(), "SELECT id FROM scooter_models WHERE model_name=$1", name )

	_ = row.Scan(&scooterModel.Id)

	var scooter entities.Scooter
	rows, err := sm.db.Query(context.Background(), "SELECT * FROM scooters WHERE model_id=$1", &scooterModel.Id)
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.ModelId, &scooter.OwnerId, &scooter.SerialNumber)
		if err != nil {
			return nil, err
		}
		scooters = append(scooters, scooter)
	}

	return &scooters, nil
}

func (sm ScooterRepository) GetScooterByID(id int) (*entities.Scooter, error) {
	scooter := entities.Scooter{}
	rows, err := sm.db.Query(context.Background(), "SELECT * FROM scooters WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.OwnerId, &scooter.ModelId, &scooter.SerialNumber)
		if err != nil {
			return nil, err
		}
	}

	return &scooter, nil
}

