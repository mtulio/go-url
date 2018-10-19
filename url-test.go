package main

import (
	"crypto/tls"
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

	if u.Method == "" {
		u.Method = "GET"
	}

	if u.Gzip == "" {
		u.Gzip = "no"
	}

	// Transport parameter
	if u.TrSSLSkip == "" {
		u.TrSSLSkip = "no"
	}
	if u.TrIpv6 == "" {
		u.TrIpv6 = "no"
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

	// Fill the path
	if u.Path == "" {
		if parsedURL.Path == "" {
			u.Path = "/"
		} else {
			u.Path = parsedURL.Path
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
		if config.OptMetric {
			sendMetrics(r.Metrics)
		}
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
		timeStart := time.Now()
		ips, err := net.LookupIP(u.Server)
		timeTaken := time.Since(timeStart)
		dnsTimeTaken := int64(timeTaken / time.Millisecond)
		if err != nil {
			testResp.Message += fmt.Sprintf(
				"URL=[%50s]: [%4s] - DNS [%v ms] Err: %s",
				u.URL, "FAIL",
				int64(timeTaken/time.Millisecond),
				err)
		}
		for _, ip := range ips {
			var uIP URLTest
			uIP = *u
			uIP.Host = uIP.Server
			uIP.Server = ip.String()
			uIP.TrSSLSkip = "yes"
			if ip.To4() == nil {
				uIP.TrIpv6 = "yes"
				uIP.Server = fmt.Sprintf("[%s]", uIP.Server)
			}
			uIP.URL = fmt.Sprintf("%s://%s:%d%s",
				u.Proto, uIP.Server,
				u.Port, u.Path)

			testGroup.URLs = append(testGroup.URLs, uIP)
		}

		// paralel test URLs for each IP
		uipResp := make(chan URLTestResult)
		var wg sync.WaitGroup
		wg.Add(len(testGroup.URLs))

		for i := range testGroup.URLs {
			go func(uIP URLTest) {
				defer wg.Done()
				var testGroupResp URLTestResult
				testGroupResp.DNSTimeTaken = dnsTimeTaken

				urlTestStart(&uIP, &testGroupResp)
				uipResp <- testGroupResp

			}(testGroup.URLs[i])
		}

		for i := 0; i < len(testGroup.URLs); i++ {
			resp := <-uipResp
			testResp.Message += fmt.Sprintf("\n%s [DNS %v ms]",
				resp.Message, resp.DNSTimeTaken)
			for i := range resp.Metrics {
				testResp.Metrics = append(testResp.Metrics, resp.Metrics[i])
			}
		}
		wg.Wait()

	} else {

		urlTestStart(u, &testResp)

	}

	config.ChanResp <- testResp
}

func urlTestStart(u *URLTest, r *URLTestResult) {

	// Setup transport Layer
	httpTr := &http.Transport{}
	timeout := time.Duration((time.Duration)(u.Timeout) * time.Second)
	if u.TrSSLSkip == "yes" {
		httpTr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	if u.TrIpv6 == "yes" {
		httpTr.Dial = (&net.Dialer{DualStack: true}).Dial
	}
	client := http.Client{
		Timeout:   timeout,
		Transport: httpTr,
	}

	// Setup HTTP request
	req, err := http.NewRequest(u.Method, u.URL, nil)
	if u.Host != "" {
		req.Host = u.Host
	}
	if u.Gzip == "yes" {
		req.Header.Add("Content-type", "application/gzip")
	}

	var metric Metric
	timeStart := time.Now()
	resp, err := client.Do(req)
	timeTaken := time.Since(timeStart)

	if err != nil {

		r.Body = fmt.Sprintf("\n\t %s", err.Error())
		r.Status = "FAIL"
		// url_alert(u, &resp_url)

	} else {

		r.Status = "OK"
		metric.HTTPCode = resp.Status
		metric.HTTPTimeTaken = int64(timeTaken / time.Millisecond)
		r.Body = fmt.Sprintf("[%s] [%v ms]",
			metric.HTTPCode, metric.HTTPTimeTaken)
		// url_alert_close(u, resp_url)

	}
	metric.HTTPHost = u.Host
	metric.HTTPServer = u.Server
	if config.OptForceDNS {
		testName := fmt.Sprintf("(%.30s) %29s", u.Host, u.URL)
		r.Message = fmt.Sprintf("URL=[%60s] [%4s] : %s",
			testName, r.Status, r.Body)
	} else {
		r.Message = fmt.Sprintf("URL=[%50s] [%4s] : %s", u.URL, r.Status, r.Body)
	}
	metric.DNSTimeTaken = r.DNSTimeTaken
	r.Metrics = append(r.Metrics, metric)
}
