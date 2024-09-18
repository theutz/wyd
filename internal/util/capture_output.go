package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CaptureOutput(function func() error) (string, error) {
	read, write, err := os.Pipe()
	if err != nil {
		err = fmt.Errorf("error: setting up pipe: %w", err)

		return "", err
	}

	defer read.Close()
	defer write.Close()

	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()

	stderr := os.Stderr
	defer func() {
		os.Stderr = stderr
	}()

	os.Stdout = write
	os.Stderr = write

	err = function()

	write.Close()

	str := strings.Builder{}

	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		str.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
	}

	if scanner.Err() != nil {
		err = fmt.Errorf("error: while scanning output: %w", err)

		return "", err
	}

	out := str.String()

	return out, err
}
