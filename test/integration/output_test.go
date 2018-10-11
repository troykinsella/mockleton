package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/troykinsella/mockleton/report"
	"os"
	"os/exec"
	"path/filepath"
)

var _ = Describe("output", func() {

	AfterEach(func() {
		os.Remove(report.DefaultReportFile)
		os.Remove(filepath.Join(FixturePath("basic"), report.DefaultReportFile))
	})

	It("should be absent by default", func() {
		cmd := exec.Command(mockletonPath)

		session := Run(cmd, 0)
		Expect(len(session.Out.Contents())).To(Equal(0))
		Expect(len(session.Err.Contents())).To(Equal(0))
	})

	It("should exit with specified code", func() {
		cmd := exec.Command(mockletonPath)
		cmd.Dir = FixturePath("basic")

		Run(cmd, 123)
	})

	It("should output to stdout and stderr", func() {
		cmd := exec.Command(mockletonPath)
		cmd.Dir = FixturePath("basic")

		session := Run(cmd, 123)
		Expect(string(session.Out.Contents())).To(Equal("foo\nbaz\n"))
		Expect(string(session.Err.Contents())).To(Equal("bar\n"))
	})
})
