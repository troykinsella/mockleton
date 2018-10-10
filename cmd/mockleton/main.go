package main

import (
	"github.com/troykinsella/mockleton"
	"github.com/troykinsella/mockleton/config"
	"github.com/troykinsella/mockleton/execution"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {

	execSpec, err := makeExecSpec()
	if err != nil {
		panic(err)
	}

	mf, err := loadMockletonfile()
	if err != nil {
		panic(err)
	}

	m := mockleton.New(execSpec, mf)

	err = m.Run()
	if err != nil {
		panic(err)
	}
}

func makeExecSpec() (*execution.Spec, error) {

	var stdin *execution.Stdin

	if hasStdin() {
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		stdin = execution.NewStdin(string(in), "utf-8")
	}

	args := execution.NewArguments(os.Args)

	env := execution.NewEnvironment(os.Environ())

	execSpec := execution.NewSpec(time.Now(), stdin, args, env)

	return execSpec, nil
}

func hasStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}

func locateMockletonfile() (string, error) {

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if path == "/" {
			break
		}

		mf, err := locateMockletonfileIn(path)
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

func locateMockletonfileIn(dir string) (string, error) {
	path := filepath.Join(dir, config.FileName)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return path, nil
}

func loadMockletonfile() (*config.Mockletonfile, error) {
	path, err := locateMockletonfile()
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, nil
	}

	mf, err := config.LoadMockletonfile(path)
	if err != nil {
		return nil, err
	}

	return mf, nil
}
