package utils

import (
	"net/http"
	"encoding/json"
	"log"
)

func BodyToStruct(req *http.Request) map[string]interface{} {
	decoder := json.NewDecoder(req.Body)
	var result map[string]interface{}
	err := decoder.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	return result
}