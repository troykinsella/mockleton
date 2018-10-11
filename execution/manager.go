package execution

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

func Marshal(spec *Spec) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	err := enc.Encode(spec)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Unmarshal(data []byte) (*Spec, error) {
	var s Spec
	err := json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func Capture() (*Spec, error) {
	var stdin *Stdin

	if hasStdin() {
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		stdin = NewStdin(string(in), "utf-8")
	}

	args := NewArguments(os.Args)

	env := NewEnvironment(os.Environ())

	execSpec := NewSpec(time.Now(), stdin, args, env)

	return execSpec, nil
}

func hasStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
