package repository

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pgdb"
	"context"
	_ "github.com/lib/pq"
)

type ScooterModelRepository struct {
	db *pgdb.PgDB
}

func NewScooterModelRepository(db *pgdb.PgDB) *ScooterModelRepository {
	return &ScooterModelRepository{
		db: db,
	}
}

type ScooterModelRepositoryI interface {
	CreateScooterModel(scooter *entities.ScooterModel) (int, error)
	GetScooterModels() (*[]entities.ScooterModel, error)
	UpdateScooterModel(scooter *entities.ScooterModel) (int, error)
	DeleteScooterModel(id int) (int, error)
	GetScooterModelById(modelId int) (*entities.ScooterModel, error)
}

func (sm ScooterModelRepository)CreateScooterModel(model *entities.ScooterModel) (int, error) {
	res, err := sm.db.Exec(context.Background(),"INSERT INTO scooter_models (payment_type_id,  model_name, max_weight, speed) VALUES ($1, $2, $3, $4)",
		 &model.PaymentTypeId, &model.ModelName, &model.MaxWeight, &model.Speed)
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

func (sm ScooterModelRepository)GetScooterModels() (*[]entities.ScooterModel, error) {
	var models []entities.ScooterModel
	rows, err := sm.db.Query(context.Background(),"SELECT * FROM scooter_models")

	if err != nil {
		return nil, err
	}
	model := entities.ScooterModel{}
	for rows.Next() {
		err = rows.Scan(&model.Id, &model.ModelName, &model.MaxWeight, &model.Speed)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return &models, nil
}

func (sm ScooterModelRepository)UpdateScooterModel(model *entities.ScooterModel) (int, error) {
	res, err := sm.db.Exec(context.Background(), "UPDATE scooter_models SET payment_type_id=$1, model_name=$2, max_weight=$3, speed WHERE id=$4",
		&model.PaymentTypeId, &model.ModelName, &model.MaxWeight, &model.Speed, &model.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := res.RowsAffected()
	return int(rowsAffected), nil
}

func (sm ScooterModelRepository)DeleteScooterModel(id int) (int, error) {
	_, err := sm.db.Exec(context.Background(), "`DELETE * FROM scooter WHERE model_id=$1", id)
	if err != nil {
		return 0, err
	}

	res1, err := sm.db.Exec(context.Background(), "`DELETE FROM scooter_models WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	rowsAffected := res1.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsAffected), nil
}

func (sm ScooterModelRepository) GetScooterModelById(modelId int) (*entities.ScooterModel, error) {
	var scooterModel entities.ScooterModel
	rows, err := sm.db.Query(context.Background(),"SELECT * FROM scooter_models  WHERE id=$1", modelId)
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


