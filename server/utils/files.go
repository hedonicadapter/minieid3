package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(path string) (string, error) {
	var sb strings.Builder
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		sb.WriteString(line)
	}

	return sb.String(), nil
}
