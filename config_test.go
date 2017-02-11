package config

import (
	_ "fmt"
	"os"
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	rice "github.com/GeertJohan/go.rice"
)

var (
	box           = rice.MustFindBox("_fixtures")
	exampleConfig = box.MustString("config_example.yml")
)

func TestMain(m *testing.M) {
	ConfigFileName = filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "config_example.yml")
	Init()
	Debug()
	os.Exit(m.Run())
}