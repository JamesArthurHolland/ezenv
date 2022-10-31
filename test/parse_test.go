package test

import (
	"github.com/JamesArthurHolland/ezenv"
	"log"
	"os"
	"testing"
)

type DbUrl string

func TestParse(t *testing.T) {
	os.Setenv("DB_URL", "localhost")
	output := ezenv.Provider[DbUrl, string]()

	log.Println(output())
}
