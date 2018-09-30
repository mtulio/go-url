package main

import (
	"fmt"
	"net"
	"net/http"
	netURL "net/url"
	"strconv"
	"sync"
	"time"
)

/* Validate URL structure */
func urlValidate(u *URLTest) {

	/* Set URL */
	if u.URL == "" {
		if u.Proto == "" {
			u.Proto = string(defaultSchema)
		}

		if u.Path == "" {
			u.Proto = string(defaultURLPath)
		}

		u.URL = fmt.Sprintf("%s://%s:%d%s",
			u.Proto, u.Server,
			u.Port, u.Path)
	}

	/* Set timeout */
	if u.Timeout <= 0 {
		u.Timeout = int(defaultTimeout)
	}

	/* Set Host and Port fields */
	newURL, err := netURL.Parse(u.URL)
	if err != nil {
		panic(err)
	}

	host, port, _ := net.SplitHostPort(newURL.Hostname())
	u.Server = host
	u.Port, _ = strconv.Atoi(port)

}

/* Setup URLs from config and call the start func */
func urlTestSetup(config *URLConfig) {

	var wg sync.WaitGroup
	timeTotalStart := time.Now()
	chanResp := make(chan URLTestResult)
	lenUrls := len(config.URLs)

	wg.Add(lenUrls)

	fmt.Printf("#> Found [%d] URLs to test, starting...\n", lenUrls)
	for k := range config.URLs {

		go urlTestStart(chanResp, &wg, config.URLs[k])

	}

	/* Show all answers from the channel */
	go func() {
		for resp := range chanResp {
			fmt.Println(resp.Message)
		}
	}()

	/* Wait for all goroutines in a workgroup */
	wg.Wait()

	timeTotalTakenMs := int64(time.Since(timeTotalStart) / time.Millisecond)
	fmt.Printf("Total time taken: %vms\n", timeTotalTakenMs)
}

/* GOroutines to execute the URL check */
func urlTestStart(chanResp chan<- URLTestResult, wg *sync.WaitGroup, u URLTest) {

	defer wg.Done()
	var testResp URLTestResult

	// Check all fields was filled
	urlValidate(&u)

	timeout := time.Duration((time.Duration)(u.Timeout) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	timeStart := time.Now()
	resp, err := client.Get(u.URL)
	timeTaken := time.Since(timeStart)

	if err != nil {

		testResp.Body = err.Error()
		testResp.Status = "FAIL"
		// url_alert(u, &resp_url)

	} else {

		testResp.Status = "OK"
		testResp.Body = fmt.Sprintf("[%s] [%v ms]", resp.Status, int64(timeTaken/time.Millisecond))
		// url_alert_close(u, resp_url)

	}
	testResp.Message = fmt.Sprintf("[%4s] URL=[%50s] : %s", testResp.Status, u.URL, testResp.Body)
	chanResp <- testResp
}
