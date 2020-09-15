package provider

import (
	"log"
	"os"
)

type LocalFileSystemDownloadFile struct {
	LocalFileSystem LocalFileSystem
	Reader          *os.File
}

func (f *LocalFileSystemDownloadFile) Initialize(path string) {
	fullPath := f.LocalFileSystem.GetFullFilePath(path)
	reader, err := os.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	f.Reader = reader
}

func (f *LocalFileSystemDownloadFile) Read(buffer []byte) (int, error) {
	n, err := f.Reader.Read(buffer)
	return n, err
}

func (f *LocalFileSystemDownloadFile) Cleanup() {
	f.Reader.Close()
}
