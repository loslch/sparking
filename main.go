package main

import (
	"fmt"
	"log"
)

func main() {
	// Download Parking Information Page
	if err := DownloadParkingPage(); err != nil {
		log.Fatal(err)
	}

	// Transform Parking Page into simple text
	if err := TransformPageToData(); err != nil {
		log.Fatal(err)
	}

	// Read Transformed Parking Data
	dataReader, err := ReadParkingData()
	if err != nil {
		log.Fatal(err)
	}
	defer dataReader.Close()

	// Generate Parking Spaces Object
	parkingSpaces, err := ParseParkingData(dataReader)
	if err != nil {
		log.Fatal(err)
	}

	// Store Parking Spaces Json
	if err := StoreParkingJson(parkingSpaces); err != nil {
		log.Fatal(err)
	}

	// Read Parking Spaces Json
	json, err := ReadParkingJson()
	if err != nil {
		log.Fatal(err)
	}

	// Print Parking Spaces
	fmt.Print(string(json))
}
