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
	ShowScooterStation() (*[]entities.ScooterStation, error)
	Create(scooter *entities.Scooter) (int, error)
	GetAll() (*[]entities.Scooter, error)
	Update(scooter *entities.Scooter) (int, error)
	Delete(id int) (int, error)
	GetModelById(modelId int) (*entities.ScooterModel, error)
	GetOwnerIdByName(userName string) (*int, error)
	GetPaymentTypeIdByName(PaymentType string) (*int, error)
	GetScooterByModelName(modelName string) (*[]entities.Scooter, error)
	GetScooterByID(id int) (*entities.Scooter, error)
}

func (s ScooterRepository) ShowScooterStation() (*[]entities.ScooterStation, error){
	var stations []entities.ScooterStation
	rows, err := s.db.Query(context.Background(),"SELECT * FROM scooter_stations")

	if err != nil {
		return nil, err
	}
	station := entities.ScooterStation{}
	for rows.Next() {
		err = rows.Scan(&station.Id, &station.Name, &station.LocationId, &station.IsActive)
		if err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	return &stations, nil
}

func (s ScooterRepository) GetModelById(modelId int) (*entities.ScooterModel, error) {
	var scooterModel entities.ScooterModel
	rows, err := s.db.Query(context.Background(),"SELECT * FROM scooter_models  WHERE id=$1", modelId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&scooterModel.Id, &scooterModel.PaymentTypeId, &scooterModel.ModelName, &scooterModel.MaxWeight, &scooterModel.Speed)
		if err != nil {
			return nil, err
		}
	}

	return &scooterModel, nil
}

func (s ScooterRepository) GetOwnerIdByName(userName string) (*int, error) {
	var user entities.User
	rows, err := s.db.Query(context.Background(),"SELECT id FROM user WHERE name=$1", userName)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&user.Id)
		if err != nil {
			panic(err)
		}
	}

	return &user.Id, nil
}

func (s ScooterRepository) GetPaymentTypeIdByName(PaymentType string) (*int, error) {
	var paymentType entities.PaymentType
	rows, err := s.db.Query(context.Background(),"SELECT id FROM payment_types WHERE name=$1", PaymentType)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&paymentType.Id)
		if err != nil {
			panic(err)
		}
	}

	return &paymentType.Id, nil
}

func (s ScooterRepository) GetAll() (*[]entities.Scooter, error) {
	var scooters []entities.Scooter
	rows, err := s.db.Query(context.Background(),"SELECT * FROM scooters")

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

func (s ScooterRepository) Create(scooter *entities.Scooter) (int, error) {
	res, err := s.db.Exec(context.Background(),"INSERT INTO scooters (id, entities, brand, max_distance, capacity, max_weight) VALUES ($1, $2, $3, $4, $5, $6)",
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

func (s ScooterRepository) GetScooterByModelName(modelName string) (*[]entities.Scooter, error) {
	var scooterModel entities.ScooterModel
	var scooters []entities.Scooter
	row := s.db.QueryRow(context.Background(), "SELECT id FROM scooter_models WHERE model_name=$1", modelName )

	_ = row.Scan(&scooterModel.Id)

	var scooter entities.Scooter
	rows, err := s.db.Query(context.Background(), "SELECT * FROM scooters WHERE model_id=$1", &scooterModel.Id)
	for rows.Next() {
		err = rows.Scan(&scooter.Id, &scooter.ModelId, &scooter.OwnerId, &scooter.SerialNumber)
		if err != nil {
			return nil, err
		}
		scooters = append(scooters, scooter)
	}

	return &scooters, nil
}

func (s ScooterRepository) GetScooterByID(id int) (*entities.Scooter, error) {
	scooter := entities.Scooter{}
	rows, err := s.db.Query(context.Background(), "SELECT * FROM scooters WHERE id=$1", id)
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

func (s ScooterRepository) Update(scooter *entities.Scooter) (int, error) {
	res, err := s.db.Exec(context.Background(), "UPDATE scooters SET serial_number=$1 WHERE id=$2",
		&scooter.SerialNumber, &scooter.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := res.RowsAffected()
	return int(rowsAffected), nil
}

func (s ScooterRepository) Delete(id int) (int, error) {
	res, err := s.db.Exec(context.Background(), "`DELETE FROM scooters WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsAffected := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}
