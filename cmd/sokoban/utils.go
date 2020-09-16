// MIT License
//
// Copyright (c) 2020 Pedro Rodrigues
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"bufio"
	"strings"

	"github.com/markbates/pkger"
)

func readLinesFromFile(path string) ([]string, error) {
	file, err := pkger.Open(path)
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

func loadLevels(path string) [][]string {
	data, _ := readLinesFromFile(path)

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
