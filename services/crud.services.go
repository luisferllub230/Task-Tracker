package services

import "github.com/luisferllub230/task_tracker/db"

func Create(model interface{}) (interface{}, error) {
	return model, nil
}

func Update(model interface{}) (interface{}, error) {
	return model, nil
}

func Delete(model interface{}) (interface{}, error) {
	return model, nil
}

func Read(model interface{}) (interface{}, error) {
	return model, nil
}

func List(entity interface{}) ([]interface{}, error) {
	dataList, err := db.FindAll(entity)
	if err != nil {
		return nil, err
	}
	return dataList, nil
}
