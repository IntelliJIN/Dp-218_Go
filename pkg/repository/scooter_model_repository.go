package repository

import (
	"Dp-218_Go/entities"
	"Dp-218_Go/pgdb"
	"context"
	"fmt"
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
	CreateScooterModel(scooterModel *entities.ScooterModel)
	GetScooterModels() (*[]entities.ScooterModel, error)
	UpdateScooterModel(scooterModel *entities.ScooterModel) (int, error)
	DeleteScooterModel(id int) (error, error)
	GetScooterModelById(modelId int) (*entities.ScooterModel, error)
}

func (sm ScooterModelRepository)CreateScooterModel(model *entities.ScooterModel){
	err, _ := sm.db.Exec(context.Background(), `INSERT INTO scooter_models (payment_type_id, model_name, max_weight, speed) VALUES ($1, $2, $3, $4)`,
		&model.PaymentTypeId, &model.ModelName, &model.MaxWeight, &model.Speed)
	if err != nil {
		fmt.Println(err)
	}
}

func (sm ScooterModelRepository)GetScooterModels() (*[]entities.ScooterModel, error) {
	var models []entities.ScooterModel
	rows, err := sm.db.Query(context.Background(),`SELECT id, payment_type_id, model_name, max_weight, speed FROM scooter_models ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}

	model := entities.ScooterModel{}
	for rows.Next() {
		err = rows.Scan(&model.Id,&model.PaymentTypeId, &model.ModelName, &model.MaxWeight, &model.Speed)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return &models, nil
}

func (sm ScooterModelRepository)UpdateScooterModel(model *entities.ScooterModel) (int, error) {
	res, err := sm.db.Exec(context.Background(), `UPDATE scooter_models SET payment_type_id=$1, model_name=$2, max_weight=$3, speed=$4 WHERE id=$5`,
		&model.PaymentTypeId, &model.ModelName, &model.MaxWeight, &model.Speed, &model.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := res.RowsAffected()
	return int(rowsAffected), nil
}

func (sm ScooterModelRepository)DeleteScooterModel(id int)  (error,error) {
/*	_, err := sm.db.Exec(context.Background(), `DELETE * FROM scooters WHERE model_id=$1`, id)
	if err != nil {
		return  nil, err
	}

 */

	_, err := sm.db.Exec(context.Background(), `DELETE FROM scooter_models WHERE id=$1`, id)
	if err != nil {
		return  err,nil
	}

	return  nil,nil
}

func (sm ScooterModelRepository) GetScooterModelById(id int) (*entities.ScooterModel, error) {
	var scooterModel entities.ScooterModel
	rows, err := sm.db.Query(context.Background(),`SELECT id, payment_type_id, model_name, max_weight, speed FROM scooter_models  WHERE id=$1`, id)
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


