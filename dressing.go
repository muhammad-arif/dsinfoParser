package dsinfoparser

import (
	"encoding/json"
	"fmt"
)

func dressing(data *json.RawMessage) string {
	var nodesDetails []interface{}
	err := json.Unmarshal(*data, &nodesDetails)
	if err != nil {
		fmt.Errorf("[ucpnodestxt.go.dressing()] : cannot unmarshal %s", err)
	}
	valStr := fmt.Sprint(nodesDetails)
	// Trimming Prefix and Suffix for extra `[`,`]`
	last := len(valStr) - 1
	valStr = valStr[:last]
	valStr = valStr[1:]
	return valStr
}
