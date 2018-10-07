
# go-url Makefile

include Makefile-ext.mk

# #####
# Build
build:
	$(call deps_dirs)
	go build -o bin/go-url *.go

# ##########
# Goreleaser
# https://goreleaser.com/introduction/

gorelease-install:
	echo $(GOPATH)
	go get -d github.com/goreleaser/goreleaser;
	cd $(GOPATH)/src/github.com/goreleaser/goreleaser && \
		dep ensure -vendor-only && \
		make setup build
	cd $(CPWD)

gorelease-install-rpm:
	yum install -y $(GORELEASE_URL_RPM)

gorelease-init:
	goreleaser init

# #######
# Release
# https://goreleaser.com/introduction/

tag:
	$(call deps_tag,$@)
	git tag -a $(version) -m "$(message)"
	git push origin $(version)

release:
	gorelease

