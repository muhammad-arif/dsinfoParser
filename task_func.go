package dsinfoParsingLibrary

import (
	"encoding/json"
	"fmt"
)

func NewTasksParser() *Task {
	return &Task{}
}
func (u Task) GetAllJSON(i *json.RawMessage) string {
	var DressedJson = dressing(i)
	return DressedJson
}
func (u Task) GetAll(i *json.RawMessage) []Task {
	var DressedJson = dressing(i)
	// Converting to byte array to prepare for json.Unmarshal as it takes only `[]byte`
	data := []byte(DressedJson)
	//
	var simple []Task
	err := json.Unmarshal(data, &simple)
	if err != nil {
		fmt.Errorf("[task_func.go.GetAll()] : cannot unmarshal %s", err)
	}
	return simple
}
