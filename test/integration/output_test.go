package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/troykinsella/mockleton"
	"os"
	"os/exec"
)

var _ = Describe("output", func() {

	AfterEach(func() {
		os.Remove(mockleton.DefaultOutFile)
	})

	It("should be absent by default", func() {
		cmd := exec.Command(mockletonPath)

		session := Run(cmd, 0)
		Expect(len(session.Out.Contents())).To(Equal(0))
		Expect(len(session.Err.Contents())).To(Equal(0))
	})
})
