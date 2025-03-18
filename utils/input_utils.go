package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ParseTextFromStandardIn() ([]string, error) {
	if !IsInputPiped() {
		return nil, fmt.Errorf("input not piped")
	}

	slice := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		slice = append(slice, t)
	}

	return slice, nil
}

func IsInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
