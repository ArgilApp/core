package provider

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type LocalFileSystemUploadFile struct {
	LocalFileSystem LocalFileSystem
	Identifier      string
	TemporaryFile   *os.File
}

func (f *LocalFileSystemUploadFile) Initialize() {
	f.Identifier = uuid.New().String()

	temporaryPath := f.LocalFileSystem.GetInProgressFilePath(f.Identifier)

	// create the base directory if it doesn't already exist
	baseDirectory := filepath.Dir(temporaryPath)
	os.MkdirAll(baseDirectory, os.ModePerm)
	file, _ := os.Create(temporaryPath)

	f.TemporaryFile = file
}

func (f *LocalFileSystemUploadFile) GetIdentifier() string {
	return f.Identifier
}

func (f *LocalFileSystemUploadFile) Write(bytes []byte) error {
	_, err := f.TemporaryFile.Write(bytes)
	return err
}

func (f *LocalFileSystemUploadFile) Cleanup() {
	f.TemporaryFile.Sync()
	f.TemporaryFile.Close()
}
