# samples stdout

Stdout for sample commands from [README.md](README.md)

## Option -config

```text
#> Reading config from JSON file: hack/config-sample.json
#> Found [4] URLs to test, starting...
URL=[                     http://grafana.internal:3000/] [FAIL] : 
	 Get http://grafana.internal:3000/: dial tcp: lookup grafana.internal on 8.8.8.8:53: no such host
URL=[                            http://www.google.com/] [  OK] : [200 OK] [198 ms]
URL=[                           https://www.google.com/] [  OK] : [200 OK] [410 ms]
URL=[                        http://ifconfig.co:80/json] [  OK] : [200 OK] [1470 ms]
Total time taken: 1471ms
```

## Option -dns

```text
#> Reading config from JSON file: hack/config-sample.json
#> Found [4] URLs to test, starting...
URL=[                     http://grafana.internal:3000/]: [FAIL] - DNS [44 ms] Err: lookup grafana.internal on 8.8.8.8:53: no such host

URL=[              (www.google.com)     http://216.58.202.228:80/] [  OK] : [200 OK] [153 ms] [DNS 42 ms]
URL=[       (www.google.com) http://[2800:3f0:4004:800::2004]:80/] [  OK] : [200 OK] [168 ms] [DNS 42 ms]

URL=[              (www.google.com)   https://216.58.202.228:443/] [  OK] : [200 OK] [288 ms] [DNS 42 ms]
URL=[     (www.google.com) https://[2800:3f0:4004:800::2004]:443/] [  OK] : [200 OK] [299 ms] [DNS 42 ms]

URL=[                 (ifconfig.co)   http://104.28.18.94:80/json] [  OK] : [200 OK] [629 ms] [DNS 43 ms]
URL=[                 (ifconfig.co)   http://104.28.19.94:80/json] [  OK] : [200 OK] [642 ms] [DNS 43 ms]
URL=[      (ifconfig.co) http://[2606:4700:30::681c:125e]:80/json] [  OK] : [200 OK] [1315 ms] [DNS 43 ms]
URL=[      (ifconfig.co) http://[2606:4700:30::681c:135e]:80/json] [  OK] : [200 OK] [1321 ms] [DNS 43 ms]
Total time taken: 1365ms
```

## Option -url

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

## Argument

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

## Docker


```bash
 $ docker run -v $PWD/hack/config-sample.json:/config.json -i mtulio/go-url:docker -dns -config /config.json 
#> Reading config from JSON file: /config.json
#> Found [4] URLs to test, starting...
URL=[                     http://grafana.internal:3000/]: [FAIL] - DNS [14 ms] Err: lookup grafana.internal on 181.213.132.2:53: no such host

URL=[              (www.google.com)     http://172.217.29.132:80/] [  OK] : [200 OK] [132 ms] [DNS 14 ms]
URL=[       (www.google.com) http://[2800:3f0:4001:805::2004]:80/] [FAIL] : 
	 Get http://[2800:3f0:4001:805::2004]:80/: dial tcp [2800:3f0:4001:805::2004]:80: connect: cannot assign requested address [DNS 14 ms]

URL=[              (www.google.com)   https://172.217.29.132:443/] [  OK] : [200 OK] [230 ms] [DNS 14 ms]
URL=[     (www.google.com) https://[2800:3f0:4001:805::2004]:443/] [FAIL] : 
	 Get https://[2800:3f0:4001:805::2004]:443/: dial tcp [2800:3f0:4001:805::2004]:443: connect: cannot assign requested address [DNS 14 ms]

URL=[                 (ifconfig.co)  http://104.27.143.41:80/json] [  OK] : [200 OK] [1115 ms] [DNS 308 ms]
URL=[                 (ifconfig.co)  http://104.27.142.41:80/json] [  OK] : [200 OK] [549 ms] [DNS 308 ms]
URL=[      (ifconfig.co) http://[2606:4700:30::681b:8e29]:80/json] [FAIL] : 
	 Get http://[2606:4700:30::681b:8e29]:80/json: dial tcp [2606:4700:30::681b:8e29]:80: connect: cannot assign requested address [DNS 308 ms]
URL=[      (ifconfig.co) http://[2606:4700:30::681b:8f29]:80/json] [FAIL] : 
	 Get http://[2606:4700:30::681b:8f29]:80/json: dial tcp [2606:4700:30::681b:8f29]:80: connect: cannot assign requested address [DNS 308 ms]
Total time taken: 1975ms
```

```bash
$ docker run -v $PWD/hack/config-sample.json:/config.json -i mtulio/go-url:docker https://g1.globo.com 
#> Reading config from Param
#> Found [1] URLs to test, starting...
URL=[                              https://g1.globo.com] [  OK] : [200 OK] [192 ms]
Total time taken: 193ms
```

## metrics

```bash
HOSTNAME=MyNode go run *.go -dns -url=http://www.google.com -metric=http://localhost:9091
#> Reading config from Param
#> Found [1] URLs to test, starting...

URL=[      (www.google.com) http://[2607:f8b0:4004:80b::2004]:80/] [  OK] : [200 OK] [353 ms] [DNS 160 ms]
URL=[              (www.google.com)      http://172.217.7.228:80/] [  OK] : [200 OK] [341 ms] [DNS 160 ms]
Total time taken: 858ms

```
