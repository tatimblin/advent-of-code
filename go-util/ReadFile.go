package util

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strconv"
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

func ReadMatrixInt(relativePath string, seperator string) ([][]int, error) {
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

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), seperator)
		var numbers []int
		for _, char := range chars {
			number, err := strconv.Atoi(char)
			if err == nil {
				numbers = append(numbers, number)
			} else {
				numbers = append(numbers, 0)
			}
		}
		lines = append(lines, numbers)
	}
	return lines, scanner.Err()
}
