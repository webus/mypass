package base

import (
	"testing"
	"os"
	"log"
)

func TestInitConfiguration(t *testing.T) {
	err := os.Setenv("CONFIG_PATH","/tmp")
	if err != nil {
		log.Fatal(err)
	}

	cfg := MyPassConfiguration{}
	// delete file if exists
	_ = os.Remove("/tmp/mypass.cfg")
	cfg.InitConfiguration()
	if _, err := os.Stat("/tmp/mypass.cfg"); os.IsNotExist(err) {
		t.Fail()
	}
}

func TestSetGetConfiguration(t *testing.T) {
	err := os.Setenv("CONFIG_PATH","/tmp")
	if err != nil {
		log.Fatal(err)
	}
	cfg := MyPassConfiguration{}
	// delete file if exists
	_ = os.Remove("/tmp/mypass.cfg")
	cfg.InitConfiguration()
	key := cfg.Key
	cfg2 := MyPassConfiguration{}
	err = cfg2.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}
	if key != cfg2.Key {
		t.Fail()
	}
}
