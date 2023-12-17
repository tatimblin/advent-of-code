package util

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadLines(relativePath string) ([]string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find relative path")
	}

	absolutePath := path.Join(path.Dir(filename), relativePath)

	file, err := os.Open(absolutePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadMatrix(relativePath string) ([][]string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find relative path")
	}

	absolutePath := path.Join(path.Dir(filename), relativePath)

	file, err := os.Open(absolutePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	return lines, scanner.Err()
}
