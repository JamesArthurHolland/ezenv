package test

import (
	"github.com/JamesArthurHolland/ezenv"
	"os"
	"testing"
)

type DbUrl string
type NumbersList []int
type StringList []string

func TestParseSingleStringEnvVar(t *testing.T) {
	os.Setenv("DB_URL", "localhost")
	output := ezenv.Provider[DbUrl]()

	dbUrl := output()

	if dbUrl != "localhost" {
		t.Error("Should equal localhost")
	}
}

func TestParseIntArrayEnvVar(t *testing.T) {
	os.Setenv("NUMBERS_LIST", "1;2;3")
	output := ezenv.SliceProvider[NumbersList]()

	parts := output()

	if parts[0] != 1 || parts[1] != 2 || parts[2] != 3 {
		t.Error("parts slice elements should be {1, 2, 3}")
	}
}

func TestParseStringArrayEnvVar(t *testing.T) {
	os.Setenv("STRING_LIST", "Alice;Bob;Charlie")

	parts := ezenv.SliceProvider[StringList]()()

	if parts[0] != "Alice" || parts[1] != "Bob" || parts[2] != "Charlie" {
		t.Error("parts slice elements should be {1, 2, 3}")
	}
}
