package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func BodyToStruct(req *http.Request) (result map[string]interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	return result
}
