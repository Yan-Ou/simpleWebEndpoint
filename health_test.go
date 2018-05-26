package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	h := health()
	h.ServeHTTP(w, req)

	status := "OK"
	diskspace := "used: 54% threshold: 1% location: /var/log"

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	info := struct {
		Status string `json:"status"`
		Errors struct {
			Diskspace string `json:"diskspace"`
		} `json:"errors"`
	}{}

	err = json.Unmarshal(data, &info)
	if err != nil {
		t.Fatal(err)
	}
	if info.Status != status {
		t.Errorf("Status is wrong. Have: %s, want: %s", info.Status, status)
	}
	if info.Errors.Diskspace != diskspace {
		t.Errorf("Diskspace info is wrong. Have: %s, want: %s", info.Errors.Diskspace, diskspace)
	}
}
