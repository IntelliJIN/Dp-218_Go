package services

import (
	"Dp-218_Go/pkg/repository"
	"mime/multipart"
)

type FileServiceI interface {
	InsertScootersToDb(file multipart.File) string
}

func NewFileService(fileRepository repository.FileRepositoryI) *FileService {
	return &FileService{
		fileRepository,
	}
}

type FileService struct {
	fileRepository repository.FileRepositoryI
}

func (f FileService)InsertScootersToDb(file multipart.File)string{
	tempFilePath := f.fileRepository.CreateTempFile(file)
	uploadModel := f.fileRepository.ConvertToStruct(tempFilePath)

	f.fileRepository.InsertScooterData(uploadModel)

	return tempFilePath
}