package main

import (
	"sync"
)

// GlobalConfig is a main list of tests to go
type GlobalConfig struct {
	URLs          []URLTest `json:"urls"`
	Location      string    `json:"location"`
	MetricPush    string    `json:"metric_push"`
	ChanResp      chan URLTestResult
	WG            sync.WaitGroup
	OptForceDNS   bool
	OptConfFile   string
	OptURL        string
	OptEmpty      bool
	OptMetric     bool
	WatchPeriod   int
	WatchInterval int
}

// URLTest is a test definition
type URLTest struct {
	URL       string            `json:"url"`
	Proto     string            `json:"proto"`
	Server    string            `json:"server"`
	Host      string            `json:"host"`
	Port      int               `json:"port"`
	Path      string            `json:"path"`
	Timeout   int               `json:"timeout"`
	Method    string            `json:"method"`
	Gzip      string            `json:"gzip"`
	TrSSLSkip string            `json:"ssl_skip"`
	TrIpv6    string            `json:"ipv6"`
	Headers   map[string]string `json:"headers"`
}

// URLTestGroup is a group of URLTest used when
// -dns is enabled, to split one test case in groups.
type URLTestGroup struct {
	URLs []URLTest
}

// URLTestResult is a result struct for the tests
type URLTestResult struct {
	Message      string
	Status       string
	Body         string
	DNSTimeTaken int64
	Metrics      []Metric
}

// Metric is a results
type Metric struct {
	HTTPCode      string
	HTTPTimeTaken int64
	HTTPHost      string
	HTTPServer    string
	DNSTimeTaken  int64
}
