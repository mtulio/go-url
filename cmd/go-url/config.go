package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* JSON File */
// Sample: {"urls": [{"url": "http://www.google.com"}]}
func configParserFromFile() {

	/* Read Config */
	filePath := config.OptConfFile
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("> error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	assertConfig()
}

func configParserFromParam() {
	var u URLTest
	u.URL = config.OptURL
	u.Headers = config.OptHeader
	config.URLs = append(config.URLs, u)
	assertConfig()
}

func assertConfig() {
	if config.Location == "" {
		config.Location = os.Getenv("HOSTNAME")
	}
	if config.MetricPush == "" {
		config.MetricPush = "http://localhost:9091"
	}
}
