package typeconv

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
)

func InterfaceMapToStringMap(src *map[interface{}]interface{}) map[string]interface{} {
	dst := map[string]interface{}{}

	for k, v := range *src {
		if v1, ok := v.(map[interface{}]interface{}); ok {
			dst[k.(string)] = InterfaceMapToStringMap(&v1)
		} else {

			dst[k.(string)] = v
		}
	}

	return dst
}

func StringToMap(input string, out *map[string]interface{}) error {
	b := []byte(input)
	err := json.Unmarshal(b, out)
	return err
}

func ByteArrayToMap(input []byte, out *map[string]interface{}) error {
	err := json.Unmarshal(input, out)
	return err
}

func InterfaceToString(input interface{}) string {
	if input == nil {
		return ""
	}
	if str, ok := input.(string); ok {
		return str
	} else {
		str = ""
		if i, ok := input.(int); ok {
			str = strconv.Itoa(i)
			return str
		} else {
			f := input.(float64)
			c := math.Ceil(f)
			if c == f {
				i := int(f)
				str = strconv.Itoa(i)
			} else {
				str = strconv.FormatFloat(f, 'f', 1, 64)
			}

			return str
		}
	}
}

func InterfaceToInt(input interface{}) (int, error) {
	if _, ok := input.(string); ok {
		if i, err := strconv.Atoi(InterfaceToString(input)); err == nil {
			return i, nil
		}
		return 0, errors.New("Not a number")
	} else {
		if i, ok := input.(int); ok {
			return i, nil
		} else {
			f := input.(float64)
			c := math.Ceil(f)
			if c == f {
				i := int(f)
				return i, nil
			} else {
				return int(f), nil
			}
		}
	}
}

func InterfaceToInt64(input interface{}) (int64, error) {
	if _, ok := input.(string); ok {
		if i, err := strconv.ParseInt(InterfaceToString(input), 10, 64); err == nil {
			return i, nil
		}
		return 0, errors.New("Not a number")
	} else {
		if i, ok := input.(int64); ok {
			return i, nil
		} else {
			f := input.(float64)
			c := math.Ceil(f)
			if c == f {
				i := int64(f)
				return i, nil
			} else {
				return int64(f), nil
			}
		}
	}
}

func InterfaceToBool(input interface{}) (bool, error) {
	if b, ok := input.(bool); ok {
		return b, nil
	} else {
		s := InterfaceToString(input)
		if s == "true" {
			return true, nil
		} else if s == "false" {
			return false, nil
		}

		return false, errors.New("Not a boolean")
	}
}

func main() {
	return
}
