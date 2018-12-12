package annotation

import (
	"errors"
	"reflect"
	"strings"
)

func Parse(tagName string, value interface{}, tag interface{}) error {

	v := reflect.ValueOf(tag)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("annotation: attempt to load into an invalid pointer")
	}
	v = v.Elem()

	isSlice := v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8

	if isSlice {
		return errors.New("annotation: attempt to load into an invalid pointer")
	}

	fields := make([]string, v.Elem().NumField())

	for i := 0; i < v.Elem().NumField(); i++ {
		field := v.Elem().Type().Field(i)
		fields[i] = field.Name
	}

	return load(tagName, fields, value, tag)

}

// Load loads any value from sql.Rows
func load(tagName string, fields []string, value interface{}, tag interface{}) error {
	v := reflect.ValueOf(tag)

	v = v.Elem()

	v2 := reflect.ValueOf(value)

	for i := 0; i < v2.Elem().NumField(); i++ {
		infoTag := v2.Elem().Type().Field(i).Tag.Get(tagName)

		var tags []string
		if infoTag != "" {
			infoTag = strings.Replace(infoTag, ", ", ",", -1)
			infoTag = strings.Replace(infoTag, "= ", "=", -1)
			infoTag = strings.Replace(infoTag, " =", "=", -1)

			tags = strings.Split(infoTag, ",")
		}

		for _, _tag := range tags {
			if strings.Contains(_tag, "=") {
				tmp := strings.Split(_tag, "=")

				for iv := 0; iv < v.Elem().NumField(); iv++ {
					name := v.Elem().Type().Field(iv).Name
					if strings.ToUpper(name) == strings.ToUpper(tmp[0]) {
						v.Elem().Field(iv).SetString(tmp[1])
						break
					}
				}
			}
		}
	}

	return nil
}
