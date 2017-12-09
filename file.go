package main

import (
	"time"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
)

const (
	TargetURL = "https://www.shinhancard.com/conts/person/card_info/premium/platinum/1197988_12791.jsp"
	DataDir   = "./data/"
)

var (
	now         = time.Now().Format(time.RFC3339)
	parkingPage = DataDir + now[0:10]
	parkingJson = DataDir + now[0:10] + ".json"
)

func getPageFilePath() (string, error) {
	if _, err := os.Stat(parkingPage); err != nil {
		if os.IsNotExist(err) {
			resp, err := http.Get(TargetURL)
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}

			if err := ioutil.WriteFile(parkingPage, body, 0644); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	return parkingPage, nil
}

func storeParkingJson(parkingSpaces []ParkingSpace) error {
	if _, err := os.Stat(parkingJson); err != nil {
		if os.IsNotExist(err) {
			jsonData, err := MarshalJSON(parkingSpaces)
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(parkingJson, jsonData, 0644); err != nil {
				return err
			}
		}
	}
	return nil
}

func readParkingJson() ([]byte, error) {
	if _, err := os.Stat(parkingJson); err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
	}

	jsonFile, err := ioutil.ReadFile(parkingJson)
	if err != nil {
		return nil, err
	}

	return jsonFile, nil
}

func isExistParkingJson() bool {
	if _, err := os.Stat(parkingJson); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
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
