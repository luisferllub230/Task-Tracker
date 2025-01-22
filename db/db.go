package db

import (
	"encoding/json"
	"os"
)

var path = "db/task.json"

func Connect(path string) ([]byte, error) {
	var data []byte
	data, err := os.ReadFile(path)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func FindAll(entity interface{}) ([]interface{}, error) {
	data, err := Connect(path)

	if err != nil {
		return nil, err
	}

	jsonErr := json.Unmarshal(data, &entity)

	if jsonErr != nil {
		return nil, jsonErr
	}

	return entity.([]interface{}), nil

}

func Save(entities []interface{}) error {
	data, err := json.Marshal(entities)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
