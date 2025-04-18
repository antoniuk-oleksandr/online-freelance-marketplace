package utils

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func AddServerURLToFiles[T any](data T) T {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr && val.Kind() != reflect.Slice {
		return data
	}

	if val.Kind() == reflect.Slice {
		return handleSlice[T](val)
	}

	elem := val.Elem()

	switch elem.Kind() {
	case reflect.String:
		return handleString[T](elem, data)
	case reflect.Ptr:
		return handlePointer[T](elem)
	case reflect.Slice:
		return handleSlicePointer[T](elem)
	case reflect.Struct:
		return handleStruct[T](elem, data)
	default:
		return data
	}
}

func handleSlicePointer[T any](elem reflect.Value) T {
	newSlice := reflect.MakeSlice(elem.Type(), elem.Len(), elem.Cap())

	for i := 0; i < elem.Len(); i++ {

		elemValue := elem.Index(i)

		if !elemValue.CanAddr() {
			return elem.Interface().(T)
		}

		modifiedValue := AddServerURLToFiles(elemValue.Addr().Interface())

		if reflect.ValueOf(modifiedValue).Kind() == reflect.Ptr {
			modifiedValue = reflect.ValueOf(modifiedValue).Elem().Interface() // Dereference
		}

		newSlice.Index(i).Set(reflect.ValueOf(modifiedValue))
	}

	result := reflect.New(newSlice.Type())
	result.Elem().Set(newSlice)

	return result.Interface().(T)
}

func handleSlice[T any](val reflect.Value) T {
	newSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())

	for i := 0; i < val.Len(); i++ {
		elemValue := val.Index(i)

		if !elemValue.CanAddr() {
			return val.Interface().(T)
		}

		modifiedValue := AddServerURLToFiles(elemValue.Addr().Interface())

		if reflect.ValueOf(modifiedValue).Kind() != reflect.Ptr {
			modifiedValue = reflect.ValueOf(modifiedValue).Addr().Interface()
		}

		newSlice.Index(i).Set(reflect.ValueOf(modifiedValue))
	}

	return newSlice.Interface().(T)
}

func handleString[T any](elem reflect.Value, data T) T {
	if !checkIfStringAFile(elem.String()) {
		return data
	}

	host := os.Getenv("FILE_SERVER_HOST")
	link := fmt.Sprintf("%s/", host)

	elem.SetString(link + elem.String())

	return data
}

func handlePointer[T any](elem reflect.Value) T {
	nestedElem := elem.Elem()

	if !nestedElem.CanAddr() {
		return elem.Interface().(T)
	}

	return AddServerURLToFiles(nestedElem.Addr().Interface().(T))
}

func handleStruct[T any](elem reflect.Value, data T) T {
	if !elem.CanSet() {
		return data
	}

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		if field.CanSet() && field.CanAddr() {
			AddServerURLToFiles(field.Addr().Interface())
		}
	}

	return data
}

func checkIfStringAFile(stringValue string) bool {
	toExclude := []string{"react.js"}

	if Contains(toExclude, strings.ToLower(stringValue)) {
		return false
	}

	re := regexp.MustCompile(`\.[a-zA-Z0-9]+$`)
	return re.MatchString(stringValue)
}
