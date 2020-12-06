package dbgo

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PathToLines(pathname string) []string {
	f, err := os.Open(pathname)
	Check(err)
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Lines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}


func System(command string) {
	// run a shell command, ignore output
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	Check(err)
}

func System2(command string) (string, string, error) {
	// run shell command and get back stringified stdout and stderr
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}