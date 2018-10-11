package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/troykinsella/mockleton/report"
	"os"
	"os/exec"
	"path/filepath"
)

var _ = Describe("sequence", func() {

	AfterEach(func() {
		os.Remove(report.DefaultReportFile)
		os.Remove(filepath.Join(FixturePath("sequence"), report.DefaultReportFile))
	})

	It("should produce configured outputs", func() {
		cmd := exec.Command(mockletonPath)
		cmd.Dir = FixturePath("sequence")
		session := Run(cmd, 11)

		Expect(string(session.Out.Contents())).To(Equal("foo\n"))
		Expect(len(session.Err.Contents())).To(Equal(0))

		cmd = exec.Command(mockletonPath)
		cmd.Dir = FixturePath("sequence")
		session = Run(cmd, 22)

		Expect(len(session.Out.Contents())).To(Equal(0))
		Expect(string(session.Err.Contents())).To(Equal("bar\n"))

		cmd = exec.Command(mockletonPath)
		cmd.Dir = FixturePath("sequence")
		session = Run(cmd, 33)

		Expect(string(session.Out.Contents())).To(Equal("baz\n"))
		Expect(len(session.Err.Contents())).To(Equal(0))

		// Back to default behaviour when sequence is exhausted
		cmd = exec.Command(mockletonPath)
		cmd.Dir = FixturePath("sequence")
		session = Run(cmd, 0)

		Expect(len(session.Out.Contents())).To(Equal(0))
		Expect(len(session.Err.Contents())).To(Equal(0))
	})

})
