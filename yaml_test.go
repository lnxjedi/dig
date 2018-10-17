package dig

import (
	"testing"
)

func TestYamlGetIdxString(t *testing.T) {
	sv := yamlData.Get("Strings", 0)
	if _, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	}
}

func TestYamlSetIdxString(t *testing.T) {
	err := yamlData.Set("Strings", 0, "ONE")
	if err != nil {
		t.Errorf("Error setting value: %v", err)
	}
	sv := yamlData.Get("Strings", 0)
	if str, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	} else {
		if str != "ONE" {
			t.Errorf("Wrong value retrieving 'Strings', 0; got '%s', wanted: 'ONE'", str)
		}
	}
}

func TestYamlGetMapString(t *testing.T) {
	sv := yamlData.Get("Structs", 0, "Name")
	if str, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	} else if str != "Bob" {
		t.Errorf("Wrong value for Structs/0/Name; want: 'Bob', got: '%s'", str)
	}
}

func TestYamlGetFloat(t *testing.T) {
	iv := yamlData.Get("Floats", 1)
	if floatval, ok := iv.(float64); !ok {
		t.Errorf("Type assertion failed; wanted type 'int', got '%v', type '%T'", iv, iv)
	} else if floatval != 2 {
		t.Errorf("Wrong value retrieving Ints/1; wanted: 2, got: %f", floatval)
	}
}

func TestYamlPathGet(t *testing.T) {
	iv := yamlData.PathGet("Floats/1")
	if floatval, ok := iv.(float64); !ok {
		t.Errorf("Type assertion failed; wanted type 'int', got '%v', type '%T'", iv, iv)
	} else if floatval != 2 {
		t.Errorf("Wrong value retrieving Ints/1; wanted: 2, got: %f", floatval)
	}
}

func TestYamlPathSet(t *testing.T) {
	err := yamlData.PathSet("Strings/0", "ONE")
	if err != nil {
		t.Errorf("Error setting value: %v", err)
	}
	sv := yamlData.PathGet("Strings/0")
	if str, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	} else {
		if str != "ONE" {
			t.Errorf("Wrong value retrieving 'Strings', 0; got '%s', wanted: 'ONE'", str)
		}
	}
}
