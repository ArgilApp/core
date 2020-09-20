package provider

import (
	"cloud.google.com/go/storage"
)

type GoogleCloudStorageDownloadFile struct {
	GoogleCloudStorage GoogleCloudStorage
	Reader             *storage.Reader
}

func (f *GoogleCloudStorageDownloadFile) Initialize(path string) {
	fullPath := f.GoogleCloudStorage.GetFullFilePath(path)
	// reader, err := os.Open(fullPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	reader, _ := f.GoogleCloudStorage.Client.Bucket(f.GoogleCloudStorage.Bucket).Object(fullPath).NewReader(*f.GoogleCloudStorage.Context)
	f.Reader = reader
}

func (f *GoogleCloudStorageDownloadFile) Read(buffer []byte) (int, error) {
	n, err := f.Reader.Read(buffer)
	return n, err
}

func (f *GoogleCloudStorageDownloadFile) Cleanup() {
	f.Reader.Close()
}
