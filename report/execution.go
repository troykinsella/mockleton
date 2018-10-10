package report

import "github.com/troykinsella/mockleton/execution"

type Execution struct {
	Spec    *execution.Spec `json:"exec-spec"`
	Results *Results        `json:"results"`
}

func NewExecution(execSpec *execution.Spec) *Execution {
	return &Execution{
		Spec: execSpec,
	}
}
