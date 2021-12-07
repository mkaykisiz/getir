package app

import (
	"encoding/json"
	"fmt"
	"getir/inmemory"
	"getir/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var storeKey = "active-tabs"
var storeValue = "getir"

func TestApp_InMemoryRetrieveHandler(t *testing.T) {
	t.Parallel()

	im := inmemory.New()
	app := App{inMemoryDB: im}
	app.inMemoryDB.Write(storeKey, storeValue)
	r, _ := http.NewRequest("GET", fmt.Sprintf("/in-memory?key=%s", storeKey), nil)
	w := httptest.NewRecorder()

	app.InMemoryRetrieveHandler(w, r)

	if http.StatusOK != w.Code {
		t.Errorf("Status Failed.  Got: %d, Want: %d", w.Code, http.StatusOK)
	}
}

func TestApp_InMemoryCreateHandler(t *testing.T) {
	t.Parallel()

	im := inmemory.New()
	app := App{inMemoryDB: im}
	app.inMemoryDB.Write(storeKey, storeValue)

	dataModel := models.Entity{Key: storeKey, Value: storeValue}
	data, _ := json.Marshal(dataModel)

	r, _ := http.NewRequest("POST", "/in-memory", strings.NewReader(string(data)))
	w := httptest.NewRecorder()

	app.InMemoryCreateHandler(w, r)

	if http.StatusCreated != w.Code {
		t.Errorf("Status Failed.  Got: %d, Want: %d", w.Code, http.StatusCreated)
	}
}
