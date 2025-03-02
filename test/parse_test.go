package test

import (
	"github.com/JamesArthurHolland/ezenv/v2"
	"os"
	"testing"
)

type DbUrl string
type NumbersList []int
type StringList []string
type SingleBool bool

func TestParseSingleStringEnvVar(t *testing.T) {
	os.Setenv("DB_URL", "localhost")
	dbUrl, err := ezenv.Provider[DbUrl]()

	if err != nil {
		t.Error("Should not error")
	}

	if dbUrl != "localhost" {
		t.Error("Should equal localhost")
	}
}

func TestParseIntArrayEnvVar(t *testing.T) {
	os.Setenv("NUMBERS_LIST", "1;2;3")
	parts, err := ezenv.SliceProvider[NumbersList]()
	if err != nil {
		t.Error("Should not error")
	}

	if parts[0] != 1 || parts[1] != 2 || parts[2] != 3 {
		t.Error("parts slice elements should be {1, 2, 3}")
	}
}

func TestParseStringArrayEnvVar(t *testing.T) {
	os.Setenv("STRING_LIST", "Alice;Bob;Charlie")

	parts, err := ezenv.SliceProvider[StringList]()
	if err != nil {
		t.Error("Should not error")
	}

	if parts[0] != "Alice" || parts[1] != "Bob" || parts[2] != "Charlie" {
		t.Error("parts slice elements should be {Alice, Bob, Charlie}")
	}
}

func TestParseSingleBoolEnvVar(t *testing.T) {
	os.Setenv("SINGLE_BOOL", "true")
	output, err := ezenv.Provider[SingleBool]()
	if err != nil {
		t.Error("Should not error")
	}

	if output != true {
		t.Error("Should equal true")
	}
}
