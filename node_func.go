package dsinfoParsingLibrary

import (
	"encoding/json"
	"fmt"
	swarmTypes "github.com/docker/docker/api/types/swarm"
)

type Node swarmTypes.Node

func NewNodeParser() *Node {
	return &Node{}
}

func (u Node) GetAllJSON(i *json.RawMessage) string {
	var DressedJson = dressing(i)
	return DressedJson
}

func (u Node) GetAll(i *json.RawMessage) []Node {
	var DressedJson = dressing(i)
	// Converting to byte array to prepare for json.Unmarshal as it takes only `[]byte`
	data := []byte(DressedJson)
	//
	var simple []Node
	err := json.Unmarshal(data, &simple)
	if err != nil {
		fmt.Errorf("[ucpnodestxt.go.GetAll()] : cannot unmarshal %s", err)
	}
	return simple
}
