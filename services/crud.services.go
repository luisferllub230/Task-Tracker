package services

import (
	"reflect"

	"github.com/luisferllub230/task_tracker/db"
)

func Create(model interface{}) (interface{}, error) {
	data, err := db.FindAll(model)
	if err != nil {
		return nil, err
	}

	data = append(data, model)
	db.Save(data)
	return model, nil
}

func Update(model interface{}) (interface{}, error) {
	data, err := db.FindAll(model)
	if err != nil {
		return nil, err
	}

	var modelValues = reflect.ValueOf(model)
	modelId := modelValues.FieldByName("Id").Int()
	id := float64(modelId)

	for i, v := range data {
		dataModelId := v.(map[string]interface{})["id"].(float64)
		if id == dataModelId {
			data = append(data[:i], data[i+1:]...)
			data = append(data, model)
			break
		}
	}
	db.Save(data)
	return model, nil
}

func Delete(model interface{}) (interface{}, error) {
	return model, nil
}

func Read(model interface{}) (interface{}, error) {
	data, err := db.FindAll(model)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		var modelValues = reflect.ValueOf(model)
		modelId := modelValues.FieldByName("Id").Int()
		id := float64(modelId)
		dataModelId := v.(map[string]interface{})["id"].(float64)
		if id == dataModelId {
			return v, nil
		}
	}
	return nil, nil
}

func List(entity interface{}) ([]interface{}, error) {
	dataList, err := db.FindAll(entity)
	if err != nil {
		return nil, err
	}
	return dataList, nil
}
