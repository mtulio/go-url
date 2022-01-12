
# go-url Makefile

include Makefile-defines.mk

# #####
# Build
build:
	go build -o bin/go-url cmd/go-url/*.go

run-sample:
	go run cmd/go-url/*.go -dns -url='https://www.google.com/search?source=hp&q=google'

clean:
	$(call deps_clean)

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
tag:
	$(call deps_tag,$@)
	git tag -a $(version) -m "$(message)"
	git push origin $(version)

# Release tool
# https://goreleaser.com/introduction/
release:
	goreleaser --rm-dist

# ######
# Docker

.PHONY: container-run-sample
container-run-sample:
	$(CONTAINER_ENGINE) run -i \
		-v $(PWD)/hack/config-sample.json:/config-sample.json:Z \
		"$(CONTAINER_REPO)/$(CONTAINER_IMAGE_NAME):$(CONTAINER_IMAGE_TAG)" \
		-dns -config /config-sample.json

.PHONY: container-build
container-build:
	$(CONTAINER_ENGINE) build -t "$(CONTAINER_REPO)/$(CONTAINER_IMAGE_NAME):$(CONTAINER_IMAGE_TAG)" .

.PHONY: container-push
container-push:
	$(CONTAINER_ENGINE) push "$(CONTAINER_REPO)/$(CONTAINER_IMAGE_NAME)"

.PHONY: container-tag-latest
container-tag-latest:
	$(CONTAINER_ENGINE) tag \
		"$(CONTAINER_REPO)/$(CONTAINER_IMAGE_NAME):$(CONTAINER_IMAGE_TAG)"
		"$(CONTAINER_REPO)/$(CONTAINER_IMAGE_NAME):latest"

# ####
# Test

.PHONY: test-run-metrics-stack
test-run-metrics-stack:
	cd hack && docker-compose up -d
