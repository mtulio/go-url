# go-url

Go-url is an tool to make HTTP requests on command line.

The big differencial are:

- make the request for each IP of the DNS endpoint
- read config from a file, it could schedulle the tests on an docker, for example

## Install

### from source code

* clone this repo
* create the config file, see [sample config](./hack/config-sample.json)
* run the code without build
`go run *.go -config hack/config-sample.json`

### install from latest version

* clone this repo
* `make build`
* see binary on `bin/` dir

## Usage

### `-config`

Read URLs tests from config:

`go run *.go -config hack/config-sample.json`

***[sample stdout](./samples-stdout.md#Option---config)***

### `-dns`

Force to resolve DNS and test on each IP address endpoint

`go run *.go -config hack/config-sample.json -dns`

***[sample stdout](./samples-stdout.md#Option--dns)***

### `-url`

Read url from option `-url`

`go run *.go -url https://www.google.com`

***[sample stdout](./samples-stdout.md#Option--url)***

### `-watch-*`

Add a option to watch requests (repeat requests):

```bash
go-url -dns -watch-period 20 -watch-interval 2 https://www.google.com
```

### Argument

Read url from argument (`argv[1]`)

`./bin/go-url https://www.google.com`

***[sample stdout](./samples-stdout.md#Argument)***

### Docker

Run with multiple options:

* `-dns` and `-config`

```bash
docker run \
    -v $PWD/hack/config-sample.json:/config.json \
    -i mtulio/go-url:docker \
    -dns -config /config.json
```

* argument

```bash
docker run \
    -v $PWD/hack/config-sample.json:/config.json \
    -i mtulio/go-url:docker https://g1.globo.com
```

***[sample stdout](./samples-stdout.md#Docker)***

### Metrics

* create local metrics stack (Prometheus + Pushgateway)

`make test-run-metrics-stack`

* send metrics to pushgateway using opt `-metric`

```bash
HOSTNAME=MyNode go run *.go \
    -dns \
    -url=http://www.google.com \
    -metric=http://localhost:9091
```

***[sample stdout](./samples-stdout.md#metrics)***

* look at the metric on the Pushgateway

![screenshot from 2018-10-19 02-27-23](https://user-images.githubusercontent.com/3216894/47199154-91acea00-d346-11e8-9ac1-eb7576ea1016.png)

## Contributing

1. Fork it
1. Create your feature branch (git checkout -b my-new-feature)
1. Commit your changes (git commit -am 'Added some feature')
1. Push to the branch (git push origin my-new-feature)
1. Create new Pull Request

Open an Issue or PR. =]
