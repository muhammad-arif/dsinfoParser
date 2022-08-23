package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var file = "/home/arif/Documents/go/dsinfo_ddl/jsonRaw/352/dsinfo.json"
	//var file = "/home/arif/Documents/go/dsinfo.json"
	dsinfoJson, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("[ReadFile] cannot Read File %s", err)
	}
	var core CoreDsinfo
	err = json.Unmarshal(dsinfoJson, &core)
	if err != nil {
		fmt.Errorf("[Unmarshal-Core] Cannot unmarshal %s", err)
	}

}
