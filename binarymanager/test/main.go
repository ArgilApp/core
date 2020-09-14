package main

import (
	"log"
	"os"

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

	stream, err := os.Open("./tmp/example_file")
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	hashes, err := binarymanager.Upload(stream)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(hashes.SHA256)

	var exists = binarymanager.HashExists(hashes.SHA256)

	if exists {
		log.Println("Found hash", hashes.SHA256)
	} else {
		log.Fatalln("Could not find hash", hashes.SHA256)
	}
}
