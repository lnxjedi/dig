package dig

import (
	"testing"
)

func TestJSONGetIdxString(t *testing.T) {
	sv := jsonData.Get("Strings", 0)
	if _, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	}
}

func TestJSONSetIdxString(t *testing.T) {
	err := jsonData.Set("Strings", 0, "ONE")
	if err != nil {
		t.Errorf("Error setting value: %v", err)
	}
	sv := jsonData.Get("Strings", 0)
	if str, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	} else {
		if str != "ONE" {
			t.Errorf("Wrong value retrieving 'Strings', 0; got '%s', wanted: 'ONE'", str)
		}
	}
}

func TestJSONGetMapString(t *testing.T) {
	sv := jsonData.Get("Structs", 0, "Name")
	if str, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	} else if str != "Bob" {
		t.Errorf("Wrong value for Structs/0/Name; want: 'Bob', got: '%s'", str)
	}
}

func TestJSONGetFloat(t *testing.T) {
	iv := jsonData.Get("Floats", 1)
	if floatval, ok := iv.(float64); !ok {
		t.Errorf("Type assertion failed; wanted type 'int', got '%v', type '%T'", iv, iv)
	} else if floatval != 2 {
		t.Errorf("Wrong value retrieving Ints/1; wanted: 2, got: %f", floatval)
	}
}

func TestJSONPathGet(t *testing.T) {
	iv := jsonData.PathGet("Floats/1")
	if floatval, ok := iv.(float64); !ok {
		t.Errorf("Type assertion failed; wanted type 'int', got '%v', type '%T'", iv, iv)
	} else if floatval != 2 {
		t.Errorf("Wrong value retrieving Ints/1; wanted: 2, got: %f", floatval)
	}
}

func TestJSONPathSet(t *testing.T) {
	err := jsonData.PathSet("Strings/0", "ONE")
	if err != nil {
		t.Errorf("Error setting value: %v", err)
	}
	sv := jsonData.PathGet("Strings/0")
	if str, ok := sv.(string); !ok {
		t.Errorf("Type assertion failed; wanted type 'string', got '%v', type '%T'", sv, sv)
	} else {
		if str != "ONE" {
			t.Errorf("Wrong value retrieving 'Strings', 0; got '%s', wanted: 'ONE'", str)
		}
	}
}
