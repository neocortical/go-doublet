package main

import (
	"os"
	"bufio"
	"regexp"
	"strings"
	"fmt"
	"time"
)

type Dictionary map[string]bool

func initDictionary(fName string) *Dictionary {
	startTime := time.Now()
	fmt.Printf("Loading dictionary...")
	
	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	dict := make(Dictionary)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// words must be lowercase and alpha-only
		word := scanner.Text()
		word = strings.ToLower(word)
		valid, _ := regexp.MatchString("^[a-z]+$", word)
		if valid {
			dict[word] = true
		}
	}
	
	fmt.Printf("Loaded dictionary in %d ms\n", time.Since(startTime) / 1000000)
	
	return &dict
}

func (dict Dictionary) Contains(word string) bool {
	return dict[strings.ToLower(word)]
}