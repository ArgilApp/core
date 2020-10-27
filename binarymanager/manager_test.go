package binarymanager

import (
	"os"
	"testing"

	"github.com/argilapp/core/binarymanager/provider"
)

func CreateLocalFSProvider() {
	var localfs provider.Provider
	localfs = &provider.LocalFileSystem{
		ID:               "localfs-test",
		StorageDirectory: "./tmp/localfs",
	}
	AddProvider(localfs)
}

func TestAddProvider(t *testing.T) {
	CreateLocalFSProvider()
	for _, a := range ListProviders() {
		if a.GetID() == "localfs-test" {
			return
		}
	}
	t.Errorf("No maching provider")
}

func TestUpload(t *testing.T) {
	CreateLocalFSProvider()
	temporaryFile, _ := os.Create("./tmp/test-file-1")
	Upload(temporaryFile)

	//not done
}
