package typeconv_test

import (
	"testing"

	"github.com/codenamoo/typeconv"
)

func TestMain(t *testing.T) {
	m := map[string]interface{}{}
	string_result := "string test"
	int_result := 10
	bool_result := true
	sbool_result := true

	m["string"] = string_result
	m["int"] = int_result
	m["not_int"] = string_result
	m["boolean"] = bool_result
	m["sboolean"] = "true"

	str := typeconv.InterfaceToString(m["string"])
	if str != string_result {
		t.Error("string convert failed")
	}

	i, err := typeconv.InterfaceToInt(m["int"])
	if err != nil {
		t.Error("int convert error")
	}
	if i != int_result {
		t.Error("int convert failed")
	}

	b, err := typeconv.InterfaceToBool(m["boolean"])
	if err != nil {
		t.Error("bool convert error")
	}
	if b != bool_result {
		t.Error("bool convert failed")
	}

	b, err = typeconv.InterfaceToBool(m["sboolean"])
	if err != nil {
		t.Error("bool convert error")
	}
	if b != sbool_result {
		t.Error("bool convert failed")
	}
}
