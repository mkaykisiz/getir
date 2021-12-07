package helper

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON returns json output.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(response)
		if err != nil {
			panic(err)
		}
	}
}
