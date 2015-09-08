package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

const (
	inputFile       = "input.txt"
	matchString     = "src=\"([^\"]+.jpg)\""
	outputSeparator = "\n"
)

func main() {
	dat, err := ioutil.ReadFile(inputFile)
	check(err)

	r, err := regexp.Compile(matchString)
	check(err)

	matches := r.FindAllStringSubmatch(string(dat), -1)

	for _, match := range matches {
		fmt.Printf("%s%s", match[1], outputSeparator)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
