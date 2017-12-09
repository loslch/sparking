package main

import (
	"fmt"
	"log"
)

func main() {
	if isExistParkingJson() == false {
		// Get ShinHan Parking Page File
		pageFilePath, err := getPageFilePath()
		if err != nil {
			log.Fatal(err)
		}

		// Generate Parking Spaces
		parkingSpaces, err := generateParkingSpace(pageFilePath)
		if err != nil {
			log.Fatal(err)
		}

		// Store Parking Spaces
		if err := storeParkingJson(parkingSpaces); err != nil {
			log.Fatal(err)
		}
	}

	// Read Parking Spaces
	json, err := readParkingJson()
	if err != nil {
		log.Fatal(err)
	}

	// Print Parking Spaces
	fmt.Println(string(json))
}
