package repository

import (
	"Dp-218_Go/entities"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jszwec/csvutil"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

type FileRepository struct {
	db *pgx.Conn
}

func NewFileRepository(db *pgx.Conn) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

type FileRepositoryI interface {
	CreateTempFile(file multipart.File)string
	ConvertToStruct(path string)[]entities.ScooterUploaded
	InsertScooterModelData(scooters []entities.ScooterUploaded)error
	InsertScooterData(scooters []entities.ScooterUploaded)error
}

func (f FileRepository) CreateTempFile(file multipart.File)string{
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile, err := ioutil.TempFile("./../../internal/temp_files", "upload-*.сsv")
	if err != nil {
		fmt.Println(err)
	}
//	defer tempFile.Close()
	tempFile.Write(fileBytes)
	return tempFile.Name()
}

func (f FileRepository) ConvertToStruct(path string)[]entities.ScooterUploaded {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'

	scooterHeader, _ := csvutil.Header(entities.ScooterUploaded{}, "csv")
	dec, _ := csvutil.NewDecoder(reader, scooterHeader...)

	var fileData []entities.ScooterUploaded
	for {
		var s entities.ScooterUploaded
		if err := dec.Decode(&s); err == io.EOF {
			break
		}
		fileData = append(fileData, s)
	}
	return fileData
}

func (f FileRepository) InsertScooterModelData(scooterModels []entities.ScooterUploaded)error{
	valueStrings := make([]string, 0, len(scooterModels))
	model := make([]interface{}, 0, len(scooterModels) * 4)

	for i, scooterModel := range scooterModels {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		model = append(model, scooterModel.PaymentTypeId)
		model = append(model, scooterModel.ModelName)
		model = append(model, scooterModel.MaxWeight)
		model = append(model, scooterModel.Speed)
	}

	stmt := fmt.Sprintf("INSERT INTO scooter_models (payment_type_id, model_name, max_weight, speed) VALUES %s ON CONFLICT (model_name) DO NOTHING;", strings.Join(valueStrings, ","))
	if _, err := f.db.Exec(context.Background(),stmt, model...)
		err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return err
	}
	return nil
}

func (f FileRepository) InsertScooterData(scooters []entities.ScooterUploaded)error{
	valueStrings := make([]string, 0, len(scooters))
	scooterInfo := make([]interface{}, 0, len(scooters) * 3)

		for i, scooter := range scooters {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
			scooterInfo = append(scooterInfo, scooter.ModelId)
			scooterInfo = append(scooterInfo, scooter.OwnerId)
			scooterInfo = append(scooterInfo, scooter.SerialNumber)
		}

		stmt1 := fmt.Sprintf("INSERT INTO scooters (model_id, owner_id, serial_number) VALUES %s ON CONFLICT (serial_number) DO NOTHING;", strings.Join(valueStrings, ","))
		if _, err := f.db.Exec(context.Background(),stmt1, scooterInfo...)
			err != nil {
			fmt.Println("Unable to insert due to: ", err)
		}
	return nil
}

