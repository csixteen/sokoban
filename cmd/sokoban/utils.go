package main

import (
	"bufio"
	"os"
	"strings"
)

func ReadLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var res []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return res, nil
}

func LoadLevels(path string) [][]string {
	data, _ := ReadLinesFromFile(path)

	var allLevels = make([][]string, 0)

	var levelData = make([]string, 0)
	for _, line := range data {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			if len(levelData) > 0 {
				allLevels = append(allLevels, levelData)
				levelData = make([]string, 0)
			}
		} else {
			levelData = append(levelData, line)
		}
	}

	if len(levelData) > 0 {
		allLevels = append(allLevels, levelData)
	}

	return allLevels
}
