package main

// URLConfig is a main list of tests to go
type URLConfig struct {
	URLs []URLTest `json:"urls"`
}

// URLTest is a test definition
type URLTest struct {
	URL     string `json:"url"`
	Proto   string `json:"proto"`
	Server  string `json:"server"`
	Port    int    `json:"port"`
	Path    string `json:"path"`
	Timeout int    `json:"timeout"`
}

// URLTestResult is a result struct for the tests
type URLTestResult struct {
	Message string
	Status  string
	Body    string
	AlertID string
}
