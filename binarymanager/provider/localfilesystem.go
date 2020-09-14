package provider

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type LocalFileSystem struct {
	ID               string
	StorageDirectory string
}

func (p LocalFileSystem) GetID() string {
	return p.ID
}

func (p LocalFileSystem) Initialize() error {
	if p.ID == "" {
		return errors.New("ID must be set")
	}
	if p.StorageDirectory == "" {
		return errors.New("StorageDirectory must be set")
	}

	log.Println("Local file system provider", p.ID, "loaded at", p.StorageDirectory)
	return nil
}

func (p LocalFileSystem) FileExists(path string) bool {
	var fullPath = p.GetFullFilePath(path)

	_, err := os.Stat(fullPath)
	return !os.IsNotExist(err)
}

func (p LocalFileSystem) GetFileInfo(path string) (string, error) {
	return "NOT IMPLEMENTED", nil
}

func (p LocalFileSystem) CreateUploadHandle() UploadFile {
	uploadFile := &LocalFileSystemUploadFile{
		LocalFileSystem: p,
	}
	uploadFile.Initialize()

	return uploadFile
}

func (p LocalFileSystem) MoveFile(oldPath string, newPath string) error {
	log.Println(oldPath, newPath)

	// create the base directory if it doesn't already exist
	baseDirectory := filepath.Dir(newPath)
	os.MkdirAll(baseDirectory, os.ModePerm)

	err := os.Rename(oldPath, newPath)
	return err
}

func (p LocalFileSystem) Download(path string) ([]byte, error) {
	return []byte{}, nil
}

func (p LocalFileSystem) Delete(path string) error {
	return nil
}

func (p LocalFileSystem) GetFullFilePath(path string) string {
	var fullPath = filepath.Join(p.StorageDirectory, "files", path)

	var fullAbsPath, _ = filepath.Abs(fullPath) //probably should use this error
	return fullAbsPath
}

// should probably be cleaned up
func (p LocalFileSystem) GetInProgressFilePath(path string) string {
	var fullPath = filepath.Join(p.StorageDirectory, "inprogress", path)

	var fullAbsPath, _ = filepath.Abs(fullPath) //probably should use this error
	return fullAbsPath
}