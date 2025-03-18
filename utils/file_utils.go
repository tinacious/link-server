package utils

import (
	"bufio"
	"os"
)

func ParseLinksFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	slice := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		slice = append(slice, t)
	}

	return slice, nil
}
