package dsinfoParsingLibrary

import (
	"encoding/json"
	"fmt"
)

func NewServiceParser() *Service {
	return &Service{}
}
func (c Service) GetAll(i *json.RawMessage) Service {
	var data Service
	err := json.Unmarshal(*i, &data)
	if err != nil {
		fmt.Errorf("[service_func.go.GetAll()] : cannot unmarshal %s", err)
	}
	return data

}
