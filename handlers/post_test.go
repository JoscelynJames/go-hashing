package handlers

import (
	"github.com/joscelynjames/go-hashing/hashing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPostHandler(t *testing.T) {
	data := url.Values{}
	data.Set("password", "password")

	r, err := http.NewRequest("POST", "https://test.com", strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	MakePosty(w, r)

	response := w.Result()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	expect := hashing.HashItOut("password")
	if string(body) != expect {
		t.Errorf("expected body to be password but body was '%s'", string(body))
	}

}
