package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	servicesMap := make(map[string]string)
	servicesMap["gatewayService"] = "./gateway"
	cleanServices()
	for serviceName, servicePath := range servicesMap {
		fmt.Println("Building service " + serviceName + " with path : " + servicePath)
		buildService(serviceName, servicePath, []string{})
	}
}

func buildService(serviceName, servicePath string, args []string) {
	serviceBinaryPath := "./bin/" + serviceName
	buildArgs := []string{"build", "-ldflags", ldflags()}
	if len(args) > 0 {
		buildArgs = append(buildArgs, "-tags", strings.Join(args, ","))
	}
	buildArgs = append(buildArgs, "-o", serviceBinaryPath)
	buildArgs = append(buildArgs, servicePath)

	executeCommand("go", buildArgs...)
}

func cleanServices() {
	serviceBinaryPath := "./bin"
	deleteFiles([]string{serviceBinaryPath}...)
}

func ldflags() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(" -X main.gitCommitNumber=%s", getGitCommit()))
	return b.String()
}

func getGitCommit() string {
	commitID, err := runCommbinedOutput("git", "rev-parse", "HEAD")
	if err != nil {
		return "undefined"
	}
	return commitID
}

func executeCommand(command string, args ...string) {
	execCommand := exec.Command(command, args...)
	execCommand.Stdout = os.Stdout
	execCommand.Stderr = os.Stderr
	err := execCommand.Run()
	if err != nil {
		panic("Execution command failed")
	}
}

func runCommbinedOutput(command string, args ...string) (string, error) {
	execCommand := exec.Command(command, args...)
	responseByte, err := execCommand.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(responseByte), nil
}

func deleteFiles(paths ...string) {
	for _, path := range paths {
		os.RemoveAll(path)
	}
}
