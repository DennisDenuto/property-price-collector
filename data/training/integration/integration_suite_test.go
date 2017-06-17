package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"fmt"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"io"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pachyderm Integration Suite")
}

var portForwardContextCancel context.CancelFunc

var _ = BeforeSuite(func() {
	startMinikube()
	waitForPods()
	pachPortForward()
})

var _ = AfterSuite(func() {
	if portForwardContextCancel != nil {
		portForwardContextCancel()
	}
})

func startMinikube() {
	cmd := exec.Command("../../../bin/start-local-pachyderm.sh")
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(session, 5*time.Minute).Should(gbytes.Say("done"))
}

func waitForPods() {
	getPodsFunc := func() *gexec.Session {
		cmd := exec.Command("kubectl", "get", "pods")
		sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		sess.Wait()
		return sess
	}

	Eventually(getPodsFunc, 5*time.Minute, 2*time.Second).Should(gbytes.Say("etcd.*1/1"))
	Eventually(getPodsFunc, 5*time.Minute, 2*time.Second).Should(gbytes.Say("pachd.*1/1"))
}

func pachPortForward() {
	var portForwardContext context.Context
	portForwardContext = context.Background()
	portForwardContext, portForwardContextCancel = context.WithCancel(portForwardContext)

	go func() {
		defer GinkgoRecover()
		sess, err := gexec.Start(exec.CommandContext(portForwardContext, "pachctl", "port-forward"), GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())

		io.Copy(GinkgoWriter, sess.Out)
		select {
		case <-portForwardContext.Done():
			sess.Kill()
		}
	}()
}

func ListRepos() []string {
	cmd := exec.Command("/bin/bash", "-c", "pachctl list-repo | tail -n +2 | awk '{print $1}'")
	sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())

	return strings.Split(string(sess.Wait(2*time.Second).Out.Contents()), "\n")
}

func ListCommits(repoName string) []string {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("pachctl list-commit %s | tail -n +2 | awk '{print $1}'", repoName))
	sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())

	return strings.Split(string(sess.Wait(2*time.Second).Out.Contents()), "\n")
}

func GetFile(repoName, commitId, path string) string {
	cmd := exec.Command("pachctl", "get-file", repoName, commitId, path)
	sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())

	return string(sess.Wait(2 * time.Second).Out.Contents())
}
