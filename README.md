# go-url

Go-url is an tool to make HTTP requests on command line.

The big differencial are:
- make the request for each IP of the DNS endpoint
- read config from a file, it could schedulle the tests on an docker, for example

## Usage

### from source code

* clone this repo

* create the config file, see [sample config](./src/config-sample.json)

* run the code without build

`go run src/*.go -config src/config-sample.json`

### install from latest version

<TODO>


## Options

* Read URLs tests from config

`go run src/*.go -config src/config-sample.json`
```text
#> Reading config from JSON file: src/config-sample.json
#> Found [4] URLs to test, starting...
URL=[                        http://ifconfig.co:80/json] [  OK] : [200 OK] [628 ms]
URL=[                            http://www.google.com/] [  OK] : [200 OK] [640 ms]
URL=[                           https://www.google.com/] [  OK] : [200 OK] [1052 ms]
URL=[                     http://grafana.internal:3000/] [FAIL] : 
	 Get http://grafana.internal:3000/: dial tcp: lookup grafana.internal on 10.50.0.2:53: no such host
Total time taken: 1448ms

```

* Force to resolve DNS and test on each IP address endpoint

` go run src/*.go -config src/config-sample.json -dns`
```text
#> Reading config from JSON file: src/config-sample.json
#> Found [4] URLs to test, starting...

URL=[      (www.google.com) http://[2607:f8b0:4004:800::2004]:80/] [  OK] : [200 OK] [344 ms] [DNS 170 ms]
URL=[              (www.google.com)        http://172.217.8.4:80/] [  OK] : [200 OK] [349 ms] [DNS 170 ms]
URL=[                     http://grafana.internal:3000/]: [FAIL] - DNS [1213 ms] Err: lookup grafana.internal on 10.50.0.2:53: no such host

URL=[    (www.google.com) https://[2607:f8b0:4004:800::2004]:443/] [  OK] : [200 OK] [675 ms] [DNS 170 ms]
URL=[              (www.google.com)      https://172.217.8.4:443/] [  OK] : [200 OK] [679 ms] [DNS 170 ms]

URL=[      (ifconfig.co) http://[2606:4700:30::681b:8e29]:80/json] [  OK] : [200 OK] [555 ms] [DNS 176 ms]
URL=[      (ifconfig.co) http://[2606:4700:30::681b:8f29]:80/json] [  OK] : [200 OK] [407 ms] [DNS 176 ms]
URL=[                 (ifconfig.co)  http://104.27.142.41:80/json] [  OK] : [200 OK] [3549 ms] [DNS 176 ms]
URL=[                 (ifconfig.co)  http://104.27.143.41:80/json] [  OK] : [200 OK] [561 ms] [DNS 176 ms]
Total time taken: 5251ms

``` 

* read url from arg `-url`

```bash
$ ./bin/go-url -url https://www.google.com
#> Reading config from Param
#> Found [1] URLs to test, starting...
URL=[                            https://www.google.com] [  OK] : [200 OK] [952 ms]
Total time taken: 953ms

$ ./bin/go-url -url https://www.google.com -dns
#> Reading config from Param
#> Found [1] URLs to test, starting...

URL=[    (www.google.com) https://[2607:f8b0:4008:811::2004]:443/] [  OK] : [200 OK] [962 ms] [DNS 106 ms]
URL=[              (www.google.com)   https://172.217.29.132:443/] [  OK] : [200 OK] [377 ms] [DNS 106 ms]
Total time taken: 1446ms


```


* read url from arg `argv[1]`

```bash
$ ./bin/go-url https://www.google.com
#> Reading config from Param
#> Found [1] URLs to test, starting...
URL=[                            https://www.google.com] [  OK] : [200 OK] [942 ms]
Total time taken: 942ms

$ ./bin/go-url -dns https://www.google.com
#> Reading config from Param
#> Found [1] URLs to test, starting...

URL=[    (www.google.com) https://[2607:f8b0:4008:811::2004]:443/] [  OK] : [200 OK] [1870 ms] [DNS 113 ms]
URL=[              (www.google.com)    https://172.217.30.68:443/] [  OK] : [200 OK] [236 ms] [DNS 113 ms]

```

## Contributing

<TODO>

