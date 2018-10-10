package report

import (
	"bytes"
	"encoding/json"
)

type R struct {
	Ver string      `json:"mockleton-version"`
	Seq []Execution `json:"sequence"`
}

func New(ver string) *R {
	return &R{
		Ver: ver,
		Seq: make([]Execution, 0),
	}
}

func Marshal(report *R) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	err := enc.Encode(report)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Unmarshal(data []byte) (*R, error) {
	var r R
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *R) Add(e *Execution) {
	r.Seq = append(r.Seq, *e)
}
