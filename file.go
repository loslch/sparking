package main

import (
	"time"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
	"os/exec"
	"errors"
)

const (
	targetURL = "https://www.shinhancard.com/conts/person/card_info/premium/platinum/1197988_12791.jsp"
	dataDir   = "./data/"
)

var (
	now         = time.Now().Format(time.RFC3339)
	parkingPage = dataDir + now[0:10] + ".jsp"
	parkingData = dataDir + now[0:10] + ".dat"
	parkingJson = dataDir + now[0:10] + ".json"
)

func isExist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

func DownloadParkingPage() error {
	if isExist(parkingPage) {
		return nil
	}

	resp, err := http.Get(targetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(parkingPage, body, 0644); err != nil {
		return err
	}

	return nil
}

func TransformPageToData() error {
	if isExist(parkingData) {
		return nil
	}

	cmd := exec.Command("./transform.sh", parkingPage, parkingData)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ReadParkingData() (*os.File, error) {
	if !isExist(parkingData) {
		return nil, errors.New("file does not exist")
	}

	file, err := os.Open(parkingData)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func StoreParkingJson(parkingSpaces []ParkingSpace) error {
	if isExist(parkingJson) {
		return nil
	}

	jsonData, err := MarshalJSON(parkingSpaces)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(parkingJson, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

func ReadParkingJson() ([]byte, error) {
	if !isExist(parkingJson) {
		return nil, errors.New("file does not exist")
	}

	jsonFile, err := ioutil.ReadFile(parkingJson)
	if err != nil {
		return nil, err
	}

	return jsonFile, nil
}

// How to stop json.Marshal from escaping < and >?
// href: https://stackoverflow.com/a/28596225
func MarshalJSON(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
