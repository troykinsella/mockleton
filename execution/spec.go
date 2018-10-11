package execution

import (
	"github.com/troykinsella/mockleton/util"
	"time"
)

// Spec

type Spec struct {
	Timestamp time.Time   `json:"timestamp"`
	Stdin     *Stdin      `json:"stdin,omitempty"`
	Args      Arguments   `json:"args"`
	Env       Environment `json:"env"`
}

func NewSpec(ts time.Time, stdin *Stdin, args Arguments, env Environment) *Spec {
	return &Spec{
		Timestamp: ts,
		Stdin:     stdin,
		Args:      args,
		Env:       env,
	}
}

// Stdin

type Stdin struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

func NewStdin(content string, encoding string) *Stdin {
	return &Stdin{
		Content:  content,
		Encoding: encoding,
	}
}

// Arguments

type Arguments []string

func NewArguments(args []string) Arguments {
	return Arguments(args)
}

func (args Arguments) Count() int {
	return len(args) - 1
}

// Environment

type Environment map[string]string

func NewEnvironment(envSlice []string) Environment {
	m := util.KeyValuesToMap(envSlice)
	return Environment(m)
}
