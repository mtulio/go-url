package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* JSON File */
// Sample: {"urls": [{"url": "http://www.google.com"}]}
func configParserFromFile(c *URLConfig, filename *string) {

	/* Read Config */
	filePath := string(*filename)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("> error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, c)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
