package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func RunCli(command, arguments string) (string, error) {
	cmd := exec.Command(command, strings.Split(arguments, " ")...)

	output, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	errScanner := bufio.NewScanner(stderr)
	errScanner.Split(bufio.ScanWords)
	errMessage := ""
	for errScanner.Scan() {
		errMessage += errScanner.Text() + " "
	}

	outScanner := bufio.NewScanner(output)
	outMessage := ""
	for outScanner.Scan() {
		outMessage += outScanner.Text()
	}

	cmd.Wait()

	fmt.Println(command, arguments)

	if strings.Contains(errMessage, "Conversion failed") {
		fmt.Println("CMD ERROR:")
		fmt.Println(errMessage)
		return "", errors.New(errMessage)
	}

	return outMessage, nil
}
