package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

const (
	FileName = "Mockletonfile"
)

type Mockletonfile struct {
	Vars     map[string]interface{} `yaml:"vars"`
	Runfile  Runfile                `yaml:"runfile"`
	Sequence Sequence               `yaml:"sequence"`
	Report   Report                 `yaml:"output"`
}

type Runfile struct {
	File   string
	Format string
}

type Sequence struct {
	Steps []Step
}

type Step struct {
	Label string `yaml:"label"`
	//Expect Expect `yaml:"expect"`
	Output Output `yaml:"output"`
}

//type Expect yomega.Spec

type Output struct {
	ExitCode int `yaml:"exit-code"`
	Print    []OutputOp

	// or

	Exec string // path to executable
}

type OutputOp struct {
	Type    string // out, err
	Content string
}

type Report struct {
	File   string
	Format string
}

func LoadMockletonfile(path string) (*Mockletonfile, error) {
	var result *Mockletonfile
	var err error

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if isJsonFile(path) {
		result, err = UnmarshalJsonMockletonfile(content)
	} else if isYamlFile(path) {
		result, err = UnmarshalYamlMockletonfile(content)
	} else {
		err = fmt.Errorf("unsupported Mockelton file format: %s", path)
	}

	if err != nil {
		return nil, err
	}

	return result, err
}

func isJsonFile(path string) bool {
	path = strings.ToLower(path)
	return strings.HasSuffix(path, "json")
}

func isYamlFile(path string) bool {
	path = strings.ToLower(path)
	return strings.HasSuffix(path, "yaml") ||
		strings.HasSuffix(path, "yml")
}

func UnmarshalJsonMockletonfile(content []byte) (*Mockletonfile, error) {
	var result Mockletonfile
	err := json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func UnmarshalYamlMockletonfile(content []byte) (*Mockletonfile, error) {
	var result Mockletonfile
	err := yaml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
