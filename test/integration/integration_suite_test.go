package integration_test

import (
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/troykinsella/mockleton/report"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var mockletonPath string

// For determining package name
type Noop struct{}

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func getMockletonPathPackage() string {
	thisPkg := reflect.TypeOf(Noop{}).PkgPath()
	parts := strings.Split(thisPkg, "/")
	return strings.Join(parts[0:len(parts)-2], "/") + "/cmd/mockleton"
}

var _ = SynchronizedBeforeSuite(func() []byte {
	binPath, err := gexec.Build(getMockletonPathPackage())
	Expect(err).NotTo(HaveOccurred())
	return []byte(binPath)
}, func(data []byte) {
	mockletonPath = string(data)
})

func Run(cmd *exec.Cmd, expectedExitCode int) *gexec.Session {
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session).Should(gexec.Exit(expectedExitCode))
	return session
}

func FixturePath(fixture string) string {
	_, basedir, _, ok := runtime.Caller(0)
	if !ok {
		// Don't assert here because it can be called outside of an It()
		panic(fmt.Errorf("Fixture not found: %s", fixture))
	}

	f := filepath.Join(basedir, "../fixtures/", fixture)
	return f
}

func LoadReport(path string) map[string]interface{} {
	if path == "" {
		path = report.DefaultReportFile
	}

	data, err := ioutil.ReadFile(path)
	Expect(err).ToNot(HaveOccurred())

	var r map[string]interface{}
	err = json.Unmarshal(data, &r)
	Expect(err).ToNot(HaveOccurred())

	return r
}
