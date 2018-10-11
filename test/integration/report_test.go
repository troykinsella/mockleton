package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/troykinsella/mockleton/report"
	"os"
	"os/exec"
)

var _ = Describe("report", func() {

	AfterEach(func() {
		os.Remove(report.DefaultReportFile)
	})

	It("should be created", func() {
		cmd := exec.Command(mockletonPath)

		Run(cmd, 0)
		Expect(report.DefaultReportFile).To(BeAnExistingFile())
	})

	It("should have an exec spec sequence", func() {
		cmd := exec.Command(mockletonPath)
		Run(cmd, 0)

		report := LoadReport("")
		Expect(report).To(HaveKey("sequence"))
		Expect(report["sequence"]).To(HaveLen(1))

		sequence := report["sequence"].([]interface{})
		execution := sequence[0].(map[string]interface{})

		Expect(execution).To(HaveKey("exec-spec"))
		Expect(execution).To(HaveKey("results"))

		execSpec := execution["exec-spec"].(map[string]interface{})

		Expect(execSpec).To(HaveKey("timestamp"))
		Expect(execSpec["timestamp"]).To(BeAssignableToTypeOf(""))
		Expect(execSpec).To(HaveKey("args"))
		Expect(execSpec["args"]).To(BeAssignableToTypeOf([]interface{}{}))
		Expect(execSpec).To(HaveKey("env"))
		Expect(execSpec).ToNot(HaveKey("stdin"))
	})

	It("should report zero arguments", func() {
		cmd := exec.Command(mockletonPath)
		Run(cmd, 0)

		report := LoadReport("")
		sequence := report["sequence"].([]interface{})
		execution := sequence[0].(map[string]interface{})
		execSpec := execution["exec-spec"].(map[string]interface{})

		Expect(execSpec["args"]).To(HaveLen(1))
		Expect(execSpec["args"]).To(ConsistOf(ContainSubstring("/mockleton")))
	})

	It("should report one argument", func() {
		cmd := exec.Command(mockletonPath, "hi")
		Run(cmd, 0)

		report := LoadReport("")
		sequence := report["sequence"].([]interface{})
		execution := sequence[0].(map[string]interface{})
		execSpec := execution["exec-spec"].(map[string]interface{})

		Expect(execSpec["args"]).To(HaveLen(2))
		Expect(execSpec["args"]).To(ConsistOf(
			ContainSubstring("/mockleton"),
			"hi",
		))
	})

	It("should report multiple arguments", func() {
		cmd := exec.Command(mockletonPath, "foo", "bar", "baz", "biz")
		Run(cmd, 0)

		report := LoadReport("")
		sequence := report["sequence"].([]interface{})
		execution := sequence[0].(map[string]interface{})
		execSpec := execution["exec-spec"].(map[string]interface{})

		Expect(execSpec["args"]).To(HaveLen(5))
		Expect(execSpec["args"]).To(ConsistOf(
			ContainSubstring("/mockleton"),
			"foo",
			"bar",
			"baz",
			"biz",
		))
	})

	It("should capture environment variables", func() {
		cmd := exec.Command(mockletonPath)
		cmd.Env = []string{
			"FOO=bar",
		}
		Run(cmd, 0)

		report := LoadReport("")
		sequence := report["sequence"].([]interface{})
		execution := sequence[0].(map[string]interface{})
		execSpec := execution["exec-spec"].(map[string]interface{})

		Expect(execSpec["env"]).To(Equal(map[string]interface{}{"FOO": "bar"}))
	})

	It("should append sequence elements", func() {
		cmd := exec.Command(mockletonPath)
		Run(cmd, 0)

		report := LoadReport("")
		Expect(report).To(HaveKey("sequence"))
		Expect(report["sequence"]).To(HaveLen(1))

		cmd = exec.Command(mockletonPath)
		Run(cmd, 0)

		report = LoadReport("")
		Expect(report).To(HaveKey("sequence"))
		Expect(report["sequence"]).To(HaveLen(2))
	})

})
