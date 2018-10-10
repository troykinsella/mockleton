package mockleton

import (
	"errors"
	"github.com/troykinsella/mockleton/config"
	"github.com/troykinsella/mockleton/execution"
	"github.com/troykinsella/mockleton/report"
	"io/ioutil"
	"os"
)

const (
	DefaultOutFile = "mockleton.out"
)

var (
	AppVersion = "0.0.0-dev.0" // Injected
)

type Mockleton struct {
	execSpec *execution.Spec
	mf       *config.Mockletonfile
}

func New(execSpec *execution.Spec, mf *config.Mockletonfile) *Mockleton {
	if execSpec == nil {
		panic(errors.New("execSpec required"))
	}
	if mf == nil {
		mf = &config.Mockletonfile{}
	}

	return &Mockleton{
		execSpec: execSpec,
		mf:       mf,
	}
}

func (m *Mockleton) Run() error {

	err := m.writeReport()
	if err != nil {
		return err
	}

	return nil
}

func (m *Mockleton) loadReport() (*report.R, error) {
	reportFile := DefaultOutFile
	reportBytes, err := ioutil.ReadFile(reportFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	report, err := report.Unmarshal(reportBytes)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (m *Mockleton) writeReport() error {

	reportFile := DefaultOutFile

	r, err := m.loadReport()
	if err != nil {
		return err
	}
	if r == nil {
		r = report.New(AppVersion)
	}

	r.Add(report.NewExecution(m.execSpec))

	reportBytes, err := report.Marshal(r)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(reportFile, reportBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
