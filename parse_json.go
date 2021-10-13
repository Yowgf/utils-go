package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const jsonschemaFile = "pleasePutAnActualJsonFileName"

func main() {
	body, err := ioutil.ReadFile(jsonschemaFile)
	if err != nil {
		panic(err)
	}
	// It's best to unmarshall to an actual struct with fields marked by
	// `json:...`. The type below is just for testing purposes.
	var topJs map[string]interface{}
	if json.Unmarshal(body, &topJs) == nil {
		fmt.Println("Valid json!")
	} else {
		panic(fmt.Errorf("Invalid json!"))
	}
}
