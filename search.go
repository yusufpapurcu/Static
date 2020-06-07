package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	SEARCH_FILE    = 0
	SEARCH_KEYWORD = 1
)

//Search function
func Search(dir, keyword string, mode int) []string {
	var result []string
	switch mode {
	case SEARCH_FILE:
		result = SearchFile(dir, keyword)
		break
	case SEARCH_KEYWORD:
		result = SearchKeyword(dir, keyword)
		break
	}
	return result
}

//SearchFile func
func SearchFile(dir, keyword string) []string {
	result := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Name() == keyword {
			result = append(result, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return result
}

//SearchKeyword function
func SearchKeyword(dir, keyword string) []string {
	result := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				return nil
			}
			r := bufio.NewScanner(f)
			var line int
			for r.Scan() {
				line++
				if strings.Contains(r.Text(), keyword) {
					result = append(result, fmt.Sprintf("Found in line %v in %v : %v", line, path, r.Text()))
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return result
}
