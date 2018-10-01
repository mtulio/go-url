package main

import (
	"sync"
)

// GlobalConfig is a main list of tests to go
type GlobalConfig struct {
	URLs        []URLTest `json:"urls"`
	ChanResp    chan URLTestResult
	WG          sync.WaitGroup
	OptForceDNS bool
	OptConfFile string
}

// URLTest is a test definition
type URLTest struct {
	URL     string `json:"url"`
	Proto   string `json:"proto"`
	Server  string `json:"server"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Path    string `json:"path"`
	Timeout int    `json:"timeout"`
}

type URLTestGroup struct {
	URLs []URLTest
}

// URLTestResult is a result struct for the tests
type URLTestResult struct {
	Message string
	Status  string
	Body    string
}
