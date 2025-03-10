package ezenv

import (
	"errors"
	"fmt"
	"github.com/gobeam/stringy"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func getEnvVarName[T any](fullTypeName string) string {
	log.Printf("Ezenv: Full type name: %s", fullTypeName)
	parts := strings.SplitAfter(fullTypeName, ".")
	envVarNameCamel := parts[len(parts)-1]
	str := stringy.New(envVarNameCamel)
	snakeStr := str.SnakeCase("?", "")
	return snakeStr.ToUpper()
}

func checkErr(err error, fullTypeName string) {
	if err != nil {
		log.Fatalf("Error for %s = %s", fullTypeName, err)
	}
}

func Provider[T any]() (T, error) {
	fullTypeName := fmt.Sprintf("%T", *new(T))

	envVarName := getEnvVarName[T](fullTypeName)

	value := os.Getenv(envVarName)
	if value == "" {
		return *new(T), errors.New(fmt.Sprintf("FATAL ERROR: Var %s not present.", envVarName))
	}

	output := new(T)
	v := reflect.ValueOf(output)
	e := v.Elem()

	switch e.Kind() {
	case reflect.String:
		e.SetString(value)
	case reflect.Int32:
		intValue, err := strconv.ParseInt(value, 10, 32)
		checkErr(err, fullTypeName)
		e.SetInt(intValue)
	case reflect.Int, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		checkErr(err, fullTypeName)
		e.SetInt(intValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		checkErr(err, fullTypeName)
		e.SetBool(boolValue)
	default:
		log.Fatalf("Not a bool, string or int %s", fullTypeName)
	}

	return *output, nil
}

func SliceProvider[S ~[]T, T any]() (S, error) {
	fullTypeName := fmt.Sprintf("%T", *new(S))

	envVarName := getEnvVarName[S](fullTypeName)

	value := os.Getenv(envVarName)
	if value == "" {
		return nil, errors.New(fmt.Sprintf("FATAL ERROR: Slice var %s not present.", envVarName))
	}

	tNew := new(T)
	v := reflect.ValueOf(tNew)

	parts := strings.Split(value, ";")
	outSlice := make([]T, len(parts))

	switch v.Elem().Kind() {
	case reflect.Int, reflect.Int64:
		for i := 0; i < len(parts); i++ {
			newElement := new(T)
			intValue, err := strconv.ParseInt(parts[i], 10, 64)
			checkErr(err, fullTypeName)
			reflect.ValueOf(newElement).Elem().SetInt(intValue)
			outSlice[i] = *newElement
		}
	case reflect.String:
		for i := 0; i < len(parts); i++ {
			newElement := new(T)
			reflect.ValueOf(newElement).Elem().SetString(parts[i])
			outSlice[i] = *newElement
		}
	case reflect.Bool:
		for i := 0; i < len(parts); i++ {
			newElement := new(T)
			boolValue, err := strconv.ParseBool(parts[i])
			checkErr(err, fullTypeName)
			reflect.ValueOf(newElement).Elem().SetBool(boolValue)
			outSlice[i] = *newElement
		}
	default:
		log.Println(reflect.TypeOf(tNew).Elem().Kind())
		log.Fatalf("Not a bool, string or int %s", fullTypeName)
	}

	return outSlice, nil
}
