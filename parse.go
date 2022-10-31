package ezenv

import (
	"fmt"
	"github.com/gobeam/stringy"
	"log"
	"os"
	"reflect"
	"strings"
)

func Provider[T any, CastAs any]() func() T {
	fullTypeName := fmt.Sprintf("%T", *new(T))

	parts := strings.SplitAfter(fullTypeName, ".")
	envVarNameCamel := parts[len(parts)-1]
	str := stringy.New(envVarNameCamel)
	snakeStr := str.SnakeCase("?", "")
	envVarName := snakeStr.ToUpper()

	value := os.Getenv(envVarName)
	if value == "" {
		log.Fatalf("Var %s not present.", envVarName)
	}

	return func() T {
		output := new(T)
		v := reflect.ValueOf(output).Elem()
		v.SetString(value)

		return *output
	}
}
