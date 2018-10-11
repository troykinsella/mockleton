package report

import (
	"io/ioutil"
	"os"
)

const (
	DefaultReportFile = "mockleton.out"
)

type Manager struct {
	reportFile string
	appVersion string
}

func NewManager(reportFile string, appVersion string) *Manager {
	if reportFile == "" {
		reportFile = DefaultReportFile
	}

	return &Manager{
		reportFile: reportFile,
		appVersion: appVersion,
	}
}

func (m *Manager) LoadReport() (*R, error) {

	reportBytes, err := ioutil.ReadFile(m.reportFile)
	if err != nil {
		if os.IsNotExist(err) {
			return New(m.appVersion), nil
		}
		return nil, err
	}

	r, err := Unmarshal(reportBytes)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m *Manager) WriteReport(report *R) error {

	reportBytes, err := Marshal(report)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(m.reportFile, reportBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
