package interfaceToMap

import "reflect"

// InterfaceToMap will return a map if the underlying interface is already a map or a struct (returns a nil map if neither)
func InterfaceToMap(i interface{}) map[string]interface{} {
	rv := reflect.ValueOf(i)

	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.Elem().CanInterface() {
			i = rv.Elem().Interface()
		}

		// assign value for default rv.Kind() switch below
		rv = rv.Elem()
	}

	switch i.(type) {
	case map[string]interface{}:
		return i.(map[string]interface{})
	default:
		switch rv.Kind() {
		case reflect.Struct:
			// insert logic to unmarshal the struct to map[string]interface{} here and return it
			return nil
		default:
			return nil
		}
	}
}
