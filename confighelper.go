package goocord

import (
	"io/ioutil"
	"os"
)

// ConfigHelper loads token from txt file
type ConfigHelper struct {
	File string // File token will be read from
}

// NewConfigHelper creates a new ConfigHelper
func NewConfigHelper(file string) (*ConfigHelper, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, err
	}

	return &ConfigHelper{file}, nil
}

// Token reads token from a file
func (c *ConfigHelper) Token() (string, error) {
	file, err := os.Open(c.File)
	if err != nil {
		return "", err
	}

	token, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return "", err
	}

	return string(token), nil
}
