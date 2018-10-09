
# go-url Makefile

include Makefile-defines.mk

# #####
# Build
build:
	$(call deps_dirs)
	go build -o bin/go-url *.go

run:
	go run *.go -url=http://www.google.com

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

.PHONY: docker-build
docker-build:
	docker build -t "$(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

.PHONY: docker-push
docker-push:
	docker push "$(DOCKER_REPO)/$(DOCKER_IMAGE_NAME)"

.PHONY: docker-tag-latest
docker-tag-latest:
	docker tag "$(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" "$(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):latest"
