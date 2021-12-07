package app

import (
	"encoding/json"
	"fmt"
	"getir/helper"
	"getir/models"
	"log"
	"net/http"
)

// Handler is root url resource.
func (app *App) Handler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w, "Server is Up and Running...")
}

// RecordListHandler is Mongo db records handler.
func (app *App) RecordListHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	filter := new(models.RecordFilter)
	err := decoder.Decode(filter)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, models.RecordsResponse{Code: 1, Msg: "Decode Error"})
		return
	}

	result := models.GetRecordsByFilter(app.mongoDB, *filter)
	status := http.StatusOK
	if result.Code != 0 {
		status = http.StatusBadRequest
	}
	helper.RespondWithJSON(w, status, result)
}

// InMemoryHandler is in memory request handler.
func (app *App) InMemoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		log.Println("[InMemoryHandler] Received Post request")
		app.InMemoryCreateHandler(w, r)
	case http.MethodGet:
		log.Println("[InMemoryHandler] Received Get request")
		app.InMemoryRetrieveHandler(w, r)
	}
}

// InMemoryCreateHandler for store new key and value.
func (app *App) InMemoryCreateHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	entity := new(models.Entity)
	err := decoder.Decode(entity)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"err": err.Error()})
		return
	}

	app.inMemoryDB.Write(entity.Key, entity.Value)
	helper.RespondWithJSON(w, http.StatusCreated, entity)

}

// InMemoryRetrieveHandler is get value by key.
func (app *App) InMemoryRetrieveHandler(w http.ResponseWriter, r *http.Request) {

	// Get Key from url querystring
	key := r.URL.Query().Get("key")
	if key == "" { // not parameter
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"err": "key required"})
		return
	}

	val, err := app.inMemoryDB.Read(key)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"err": err.Error()})
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, models.Entity{Key: key, Value: val})
}
