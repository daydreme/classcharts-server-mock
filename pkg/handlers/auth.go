package handlers

import (
	"encoding/json"
	"net/http"
)

func HasDOBHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := map[string]interface{}{}

	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.Write(jsonData)
}
