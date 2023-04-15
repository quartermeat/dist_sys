package data

import (
	"encoding/json"
	"os"
)

type GenericMessage map[string]interface{}

func LoadMessageFromJSONFile(filename string) (*GenericMessage, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var message GenericMessage
	err = decoder.Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
