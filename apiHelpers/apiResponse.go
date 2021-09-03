package apiHelpers

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(rw http.ResponseWriter, code int, errorMessage string) {
	RespondWithJSON(rw, code, map[string]string{
		"error": errorMessage,
	})
}

func RespondWithJSON(rw http.ResponseWriter, code int, responseBody interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)

	json.NewEncoder(rw).Encode(responseBody)
}
