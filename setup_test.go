package dig

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/ghodss/yaml"
)

var yamlData, jsonData InterfaceMap

func TestMain(m *testing.M) {
	rawJSONData := []byte(`{
		"Floats": [ 1, 2, 3, 4 ],
		"Strings": [
			"one",
			"two",
			"three"
		],
		"Structs": [
			{
				"Name": "Bob",
				"Relation": "Uncle"
			},
			{
				"Name": "Martha",
				"Relation": "Aunt"
			}
		]
	}`)
	rawYamlData := []byte(`
Floats: [ 1, 2, 3, 4 ]
Strings:
- one
- two
- three
Structs:
- Name: Bob
  Relation: Uncle
- Name: Martha
  Relation: Aunt
`)
	jsonData = New()
	yamlData = New()
	if err := json.Unmarshal(rawJSONData, &jsonData); err != nil {
		log.Fatalf("json.Unmarshal failed: %v", err)
	}
	if err := yaml.Unmarshal(rawYamlData, &yamlData); err != nil {
		log.Fatalf("yaml.Unmarshal failed: %v", err)
	}
	os.Exit(m.Run())
}
