package provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"cloud.google.com/go/storage"
)

type GoogleCloudStorage struct {
	ID      string
	Bucket  string
	Context *context.Context
	Client  *storage.Client
}

func (p *GoogleCloudStorage) GetID() string {
	return p.ID
}

func (p *GoogleCloudStorage) Initialize() error {
	if p.ID == "" {
		return errors.New("ID must be set")
	}
	if p.Bucket == "" {
		return errors.New("Bucket must be set")
	}

	log.Println("Google Cloud Storage system provider", p.ID, "loaded at", p.Bucket)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}

	p.Client = client
	p.Context = &ctx

	return nil
}

func (p *GoogleCloudStorage) FileExists(path string) bool {
	obj := p.Client.Bucket(p.Bucket).Object(path)

	_, err := obj.Attrs(*p.Context)

	return !(err == storage.ErrObjectNotExist)
}

func (p *GoogleCloudStorage) GetFileInfo(path string) (FileInfo, error) {
	fileInfo := FileInfo{}
	return fileInfo, errors.New("File info not implemented yet")
}

func (p *GoogleCloudStorage) CreateUploadHandle() UploadFile {
	uploadFile := &GoogleCloudStorageUploadFile{
		GoogleCloudStorage: p,
	}
	uploadFile.Initialize()

	return uploadFile
}

func (p *GoogleCloudStorage) MoveFile(oldPath string, newPath string) error {
	src := p.Client.Bucket(p.Bucket).Object(oldPath)
	dst := p.Client.Bucket(p.Bucket).Object(newPath)

	if _, err := dst.CopierFrom(src).Run(*p.Context); err != nil {
		return fmt.Errorf("Failed copying %q to %q: %v", oldPath, newPath, err)
	}
	if err := src.Delete(*p.Context); err != nil {
		return fmt.Errorf("Failed deleting %q: %v", oldPath, err)
	}

	return nil
}

func (p *GoogleCloudStorage) CreateDownloadHandle(path string) DownloadFile {
	downloadFile := &GoogleCloudStorageDownloadFile{
		GoogleCloudStorage: *p,
	}
	downloadFile.Initialize(path)

	return downloadFile
}

func (p *GoogleCloudStorage) Delete(path string) error {
	obj := p.Client.Bucket(p.Bucket).Object(path)

	err := obj.Delete(*p.Context)

	return err
}

func (p *GoogleCloudStorage) GetFullFilePath(path string) string {
	fullPath := filepath.Join("files", path)

	return fullPath
}

// should probably be cleaned up
func (p *GoogleCloudStorage) GetInProgressFilePath(path string) string {
	fullPath := filepath.Join("inprogress", path)

	return fullPath
}

func (p *GoogleCloudStorage) SupportedDownloadAccessTypes() []AccessType {
	return []AccessType{
		DirectLink, Streamable,
	}
}
