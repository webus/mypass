package base

import (
	"os"
	"log"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
)

const DEFAULT_EDITOR = "vi"
const DEFAULT_DATABASE_NAME = "mypass.db"

type MyPassConfiguration struct {
	Key string `json:"key"`
	DatabaseLocation string `json:"location"`
	DatabaseName string `json:"name"`
	Editor string `json:"editor"`
}

func getCfgPath() string {
	configPath := os.Getenv("HOME")
	if os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	}
	return filepath.Join(configPath, "mypass.cfg")
}

func (c *MyPassConfiguration) InitConfiguration() {
	err := c.GetConfiguration()
	if err != nil {
		c.Key = GetRandomKey()
		c.DatabaseLocation = os.Getenv("HOME")
		c.DatabaseName = DEFAULT_DATABASE_NAME
		c.Editor = DEFAULT_EDITOR
		err = c.SetConfiguration()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *MyPassConfiguration) GetConfiguration() error {
	cfgPath := getCfgPath()
	file, err := os.Open(cfgPath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		return err
	}
	return nil
}

func (c *MyPassConfiguration) SetConfiguration() error {
	json_data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	cfgPath := getCfgPath()
	if err = ioutil.WriteFile(cfgPath, json_data, 0644); err != nil {
		return err
	}
	return nil
}
