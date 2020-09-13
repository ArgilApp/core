package binarymanager

import (
	"testing"

	"github.com/argilapp/core/binarymanager"
	"github.com/argilapp/core/binarymanager/provider"
)

func TestAddProvider(t *testing.T) {
	var localfs provider.Provider
	localfs = provider.LocalFileSystem{
		ID:               "localfs-test",
		StorageDirectory: "/var/argil/storage",
	}
	binarymanager.AddProvider(localfs)

	for _, a := range binarymanager.ListProviders() {
		if a == localfs {
			return
		}
	}
	t.Errorf("No maching provider")
}
