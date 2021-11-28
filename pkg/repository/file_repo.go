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
	InsertToDb(scooters []entities.ScooterUploaded) error
	CreateTempFile(file multipart.File)string
	ConvertToStruct(path string)[]entities.ScooterUploaded
}


func (f FileRepository) CreateTempFile(file multipart.File)string{
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile, err := ioutil.TempFile("./../internal/temp_files", "upload-*.сsv")
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

	var scooters []entities.ScooterUploaded
	for {
		var s entities.ScooterUploaded
		if err := dec.Decode(&s); err == io.EOF {
			break
		}
		scooters = append(scooters, s)
	}
	return scooters
}

func (f FileRepository) InsertToDb(scooters []entities.ScooterUploaded) error {
	valueStrings := make([]string, 0, len(scooters))
	valueArgs := make([]interface{}, 0, len(scooters) * 4)
	for i, scooter := range scooters {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d", i*4+1, i*4+2, i*4+3, i*4+4))
		valueArgs = append(valueArgs, scooter.ModelName)
		valueArgs = append(valueArgs, scooter.MaxWeight)
		valueArgs = append(valueArgs, scooter.PaymentType)
		valueArgs = append(valueArgs, scooter.Speed)
	}

	stmt := fmt.Sprintf("INSERT INTO scooters(model_name, max_weight, payment_type, speed) VALUES %s", strings.Join(valueStrings, ","))
	if _, err := f.db.Exec(context.Background(),stmt, valueArgs...)
		err != nil {
		fmt.Println("Unable to insert due to: ", err)
		return err
	}
	return nil
}
