package main

import (
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
)

func TestShinHanPageConnection(t *testing.T) {
	resp, err := http.Get(TargetURL)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	text := string(data)

	if !strings.Contains(text, "tblNfud") {
		t.Fatal("Cannot find 'tblNfud' in this page.")
	}
}
