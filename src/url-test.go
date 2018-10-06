package main

import (
	"fmt"
	"net"
	"net/http"
	netURL "net/url"
	"strconv"
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

	if u.Method == "" {
		u.Method = "GET"
	}

	/* Set timeout */
	if u.Timeout <= 0 {
		u.Timeout = int(defaultTimeout)
	}

	/* Set Host and Port fields */
	parsedURL, err := netURL.Parse(u.URL)
	if err != nil {
		panic(err)
	}

	host, port, err := net.SplitHostPort(parsedURL.Hostname())
	if err != nil {
		host = parsedURL.Hostname()
	}
	u.Server = host

	// Try to discovery the port
	u.Port, _ = strconv.Atoi(port)

	if u.Port == 0 {
		u.Port, _ = strconv.Atoi(parsedURL.Port())
	}

	if u.Proto == "" {
		u.Proto = parsedURL.Scheme
	}

	if u.Port == 0 {
		if parsedURL.Scheme == "https" {
			u.Port = 443
		} else {
			u.Port = 80
		}
	}
}

/* Setup URLs from config and call the start func */
func urlTestLauncher() {

	timeTotalStart := time.Now()

	config.ChanResp = make(chan URLTestResult)
	lenUrls := len(config.URLs)
	config.WG.Add(lenUrls)

	fmt.Printf("#> Found [%d] URLs to test, starting...\n", lenUrls)
	for k := range config.URLs {

		go urlTestSetup(&config.URLs[k])

	}

	/* Show all answers from the channel */
	for x := 0; x < lenUrls; x++ {
		r := <-config.ChanResp
		fmt.Println(r.Message)
	}

	/* Wait for all goroutines in a workgroup */
	config.WG.Wait()

	timeTotalTakenMs := int64(time.Since(timeTotalStart) / time.Millisecond)
	fmt.Printf("Total time taken: %vms\n", timeTotalTakenMs)
}

/* GOroutines to execute the URL check */
func urlTestSetup(u *URLTest) {

	defer config.WG.Done()
	var testResp URLTestResult
	urlValidate(u)

	if config.OptForceDNS {
		var testGroup URLTestGroup
		var testGroupResp URLTestResult

		ips, err := net.LookupIP(u.Server)
		if err != nil {
			testResp.Message += fmt.Sprintf("URL=[%50s]: [%4s] : DNS Err: %s", u.URL, "FAIL", err)
		}
		for _, ip := range ips {
			var uIP URLTest
			uIP = *u
			uIP.Host = uIP.Server
			uIP.Server = ip.String()
			uIP.URL = fmt.Sprintf("%s://%s:%d%s",
				u.Proto, uIP.Server,
				u.Port, u.Path)
			testGroup.URLs = append(testGroup.URLs, uIP)
		}
		// testGroupResp.Message = fmt.Sprintf(" Len servers=%d", len(testGroup.URLs))

		for _, uIP := range testGroup.URLs {
			urlTestStart(&uIP, &testGroupResp)
			respStr := testGroupResp.Message
			testResp.Message += fmt.Sprintf("\n%s", respStr)
		}

	} else {

		urlTestStart(u, &testResp)

	}

	config.ChanResp <- testResp
}

func urlTestStart(u *URLTest, r *URLTestResult) {
	// Check all fields was filled

	timeout := time.Duration((time.Duration)(u.Timeout) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(u.Method, u.URL, nil)

	timeStart := time.Now()
	resp, err := client.Do(req)
	timeTaken := time.Since(timeStart)

	if err != nil {

		r.Body = fmt.Sprintf("\n\t %s", err.Error())
		r.Status = "FAIL"
		// url_alert(u, &resp_url)

	} else {

		r.Status = "OK"
		r.Body = fmt.Sprintf("[%s] [%v ms]", resp.Status, int64(timeTaken/time.Millisecond))
		// url_alert_close(u, resp_url)

	}
	r.Message = fmt.Sprintf("URL=[%50s] [%4s] : %s", u.URL, r.Status, r.Body)
}
