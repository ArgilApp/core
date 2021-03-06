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

func (p *LocalFileSystem) GetID() string {
	return p.ID
}

func (p *LocalFileSystem) Initialize() error {
	if p.ID == "" {
		return errors.New("ID must be set")
	}
	if p.StorageDirectory == "" {
		return errors.New("StorageDirectory must be set")
	}

	log.Println("Local file system provider", p.ID, "loaded at", p.StorageDirectory)
	return nil
}

func (p *LocalFileSystem) FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (p *LocalFileSystem) GetFileInfo(path string) (FileInfo, error) {
	fileInfo := FileInfo{}
	return fileInfo, errors.New("File info not implemented yet")
}

func (p *LocalFileSystem) CreateUploadHandle() UploadFile {
	uploadFile := &LocalFileSystemUploadFile{
		LocalFileSystem: *p,
	}
	uploadFile.Initialize()

	return uploadFile
}

func (p *LocalFileSystem) MoveFile(oldPath string, newPath string) error {
	// create the base directory if it doesn't already exist
	baseDirectory := filepath.Dir(newPath)
	os.MkdirAll(baseDirectory, os.ModePerm)

	err := os.Rename(oldPath, newPath)
	return err
}

func (p *LocalFileSystem) CreateDownloadHandle(path string) DownloadFile {
	downloadFile := &LocalFileSystemDownloadFile{
		LocalFileSystem: *p,
	}
	downloadFile.Initialize(path)

	return downloadFile
}

func (p *LocalFileSystem) GetDirectDownloadLink(path string) (string, error) {
	return "N/A", errors.New("Local file system does not support direct download links")
}

func (p *LocalFileSystem) Delete(path string) error {
	err := os.RemoveAll(path) // we should also cleanup all the empty directories that could be left behind here
	return err
}

func (p *LocalFileSystem) GetFullFilePath(path string) string {
	fullPath := filepath.Join(p.StorageDirectory, "files", path)

	fullAbsPath, _ := filepath.Abs(fullPath) //probably should use this error
	return fullAbsPath
}

// should probably be cleaned up
func (p *LocalFileSystem) GetInProgressFilePath(path string) string {
	fullPath := filepath.Join(p.StorageDirectory, "inprogress", path)

	fullAbsPath, _ := filepath.Abs(fullPath) //probably should use this error
	return fullAbsPath
}

func (p *LocalFileSystem) SupportedDownloadAccessTypes() []AccessType {
	return []AccessType{
		Streamable,
	}
}
