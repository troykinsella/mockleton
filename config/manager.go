package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LocateMockletonfile() (string, error) {

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if path == "/" {
			break
		}

		mf, err := existsIn(path)
		if err != nil {
			return "", err
		}

		if mf != "" {
			return mf, nil
		}

		path = filepath.Dir(path)
	}

	return "", nil
}

func existsIn(dir string) (string, error) {
	path := filepath.Join(dir, FileName)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return path, nil
}

func LoadMockletonfileAt(path string) (*Mockletonfile, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result Mockletonfile
	err = yaml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	return fillNils(&result), nil
}

func LoadMockletonfile() (*Mockletonfile, error) {
	path, err := LocateMockletonfile()
	if err != nil {
		return nil, err
	}
	if path == "" {
		return NewMockletonfile(), nil
	}

	mf, err := LoadMockletonfileAt(path)
	if err != nil {
		return nil, err
	}

	return mf, nil
}
