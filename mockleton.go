package mockleton

import (
	"fmt"
	"github.com/troykinsella/mockleton/config"
	"github.com/troykinsella/mockleton/execution"
	"github.com/troykinsella/mockleton/report"
	"os"
)

var (
	AppVersion = "0.0.0-dev.0" // Injected
)

type Mockleton struct {
	reportM *report.Manager

	execSpec *execution.Spec
	mf       *config.Mockletonfile
}

func New() *Mockleton {
	return &Mockleton{}
}

func (m *Mockleton) Run() (int, error) {
	var err error

	m.execSpec, err = execution.Capture()
	if err != nil {
		return 0, err
	}

	m.mf, err = config.LoadMockletonfile()
	if err != nil {
		return 0, err
	}

	m.reportM = report.NewManager(m.mf.Report.File, AppVersion)

	r, err := m.reportM.LoadReport()
	if err != nil {
		return 0, err
	}

	activeSeq := len(r.Seq)
	exitCode, err := m.runSequence(activeSeq)
	if err != nil {
		return 0, err
	}

	r.Add(report.NewExecution(m.execSpec))

	err = m.reportM.WriteReport(r)
	if err != nil {
		return 0, err
	}

	return exitCode, nil
}

func (m *Mockleton) runSequence(i int) (int, error) {
	if i < len(m.mf.Sequence) {
		step := m.mf.Sequence[i]

		if step.Output != nil {
			return m.processOutput(step.Output)
		}
	}

	return 0, nil
}

func (m *Mockleton) processOutput(output *config.Output) (int, error) {

	for _, p := range output.Print {
		if p["out"] != "" {
			fmt.Println(p["out"])
		} else if p["err"] != "" {
			fmt.Fprintln(os.Stderr, p["err"])
		} else {
			return 0, fmt.Errorf("invalid print operation: %v", p)
		}
	}

	return output.ExitCode, nil
}
