package dsinfoParsingLibrary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var file = "/home/arif/Documents/go/dsinfo_ddl/jsonRaw/352/dsinfo.json"
	//var file = "/home/arif/Documents/go/dsinfo.json
	dsinfoJson, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("cannot Read File %s", err)
	}
	var core CoreDsinfo
	err = json.Unmarshal(dsinfoJson, &core)
	if err != nil {
		fmt.Errorf("Cannot unmarshal %s", err)
	}
	nodeInspect := NewNodeParser().GetAll(&core.UcpNodesTxt)

	for k, _ := range nodeInspect {
		fmt.Printf("Hostname: %s\n\tEngine Version: %s\n\tRole: %s\n", nodeInspect[k].Description.Hostname, nodeInspect[k].Description.Engine.EngineVersion, nodeInspect[k].Spec.Role)
	}
}
