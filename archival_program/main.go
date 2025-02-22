package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mholt/archiver"
)

type Manifest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func main() {
	fmt.Println("This is an archival program.")

	// Read JSON Manifest
	jsonFile, json_err := os.Open("data.json")
	if json_err != nil {
		fmt.Println(json_err)
	}

	defer jsonFile.Close()
	jsonBytes, _ := io.ReadAll(jsonFile)

	var manifest Manifest
	json.Unmarshal(jsonBytes, &manifest)

	// Debug Print Statement
	fmt.Println("Data is at: " + manifest.Source)

	outputArchive := manifest.Destination

	err := archiver.Archive([]string{manifest.Source}, outputArchive)
	if err != nil {
		log.Fatalf("Failed to create archive: %v", err)
	}

	fmt.Println("Archive created successfully:", outputArchive)

}
