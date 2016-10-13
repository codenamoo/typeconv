# typeconv
go type convertor

### functions
* func InterfaceMapToStringMap(src *map[interface{}]interface{}) map[string]interface{}
* func StringToMap(input string, out *map[string]interface{}) error
* func InterfaceToString(input interface{}) string
* func InterfaceToInt(input interface{}) (int, error)
* func InterfaceToInt64(input interface{}) (int, error)
* func InterfaceToBool(input interface{}) (bool, error)

### test

  $ go test
