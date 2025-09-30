package analyzer

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func readByLine(filename string, fn func(string) error) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		trimmed := strings.TrimSpace(line)
		if fnErr := fn(trimmed); fnErr != nil {
			return fnErr
		}

		if err == io.EOF {
			break
		}
	}

	return nil
}
