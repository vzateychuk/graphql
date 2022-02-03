package meta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var store = make(map[string]Metadata)

//Helper function to import json from file to map
func InitMetaStore(fileName string) error {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile: %w", err)
	}
	err = json.Unmarshal(content, &store)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)

	}
	return nil
}
