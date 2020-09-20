package provider

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type GoogleCloudStorageUploadFile struct {
	GoogleCloudStorage *GoogleCloudStorage
	Identifier         string
	TemporaryFile      *storage.Writer
	Cancel             context.CancelFunc
}

func (f *GoogleCloudStorageUploadFile) Initialize() {
	f.Identifier = uuid.New().String()

	temporaryPath := f.GoogleCloudStorage.GetInProgressFilePath(f.Identifier)

	// create the base directory if it doesn't already exist
	// baseDirectory := filepath.Dir(temporaryPath)
	// os.MkdirAll(baseDirectory, os.ModePerm)
	// file, _ := os.Create(temporaryPath)

	// ctx, cancel := context.WithTimeout(*f.GoogleCloudStorage.Context, time.Second*50)
	// defer cancel()

	writer := f.GoogleCloudStorage.Client.Bucket(f.GoogleCloudStorage.Bucket).Object(temporaryPath).NewWriter(*f.GoogleCloudStorage.Context)

	f.TemporaryFile = writer
}

func (f *GoogleCloudStorageUploadFile) GetIdentifier() string {
	return f.Identifier
}

func (f *GoogleCloudStorageUploadFile) Write(bytes []byte) error {
	_, err := f.TemporaryFile.Write(bytes)
	return err
}

func (f *GoogleCloudStorageUploadFile) Cleanup() {
	f.TemporaryFile.Close()
}
