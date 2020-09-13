package provider

import (
	"errors"
	"log"
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
	log.Println(p.getFullFilePath(path))
	return false
}

func (p LocalFileSystem) GetFileInfo(path string) (error, string) {
	return nil, "NOT IMPLEMENTED"
}

func (p LocalFileSystem) Upload(path string) (error, string) {
	return nil, "NOT IMPLEMENTED"
}

func (p LocalFileSystem) Download(path string) (error, string) {
	return nil, "NOT IMPLEMENTED"
}

func (p LocalFileSystem) Delete(path string) error {
	return nil
}

func (p LocalFileSystem) getFullFilePath(path string) string {
	var fullPath = filepath.Join(p.StorageDirectory, path)

	var fullAbsPath, _ = filepath.Abs(fullPath) //probably should use this error
	return fullAbsPath
}
