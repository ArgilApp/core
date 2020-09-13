package main

import (
	"log"

	"github.com/argilapp/core/binarymanager"
	"github.com/argilapp/core/binarymanager/provider"
)

func main() {
	var localfs provider.Provider
	localfs = provider.LocalFileSystem{
		ID:               "localfs-test",
		StorageDirectory: "./tmp/localfs",
	}
	binarymanager.AddProvider(localfs)

	// h := sha256.New()
	// h.Write([]byte("hello world\n"))
	// fmt.Printf("%x", h.Sum(nil))

	log.Println(binarymanager.HashExists("a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447"))
}
