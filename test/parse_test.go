package test

import (
	"github.com/JamesArthurHolland/ezenv"
	"os"
	"testing"
)

type DbUrl string

func TestParse(t *testing.T) {
	os.Setenv("DB_URL", "localhost")
	output := ezenv.Provider[DbUrl, string]()

	dbUrl := output()

	if dbUrl != "localhost" {
		t.Error("Should equal localhost")
	}
}
