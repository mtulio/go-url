
# Makefile extensions assets

# ###########
# Global Vars

APP_NAME ?= go-url

CONTAINER_REPO			?= mtulio
CONTAINER_IMAGE_NAME	?= $(APP_NAME)
CONTAINER_IMAGE_TAG		?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

CONTAINER_ENGINE ?= podman

CPWD := $(PWD)

GORELEASE_VERSION	:= v0.86.1
GORELEASE_BASE_URL	:= https://github.com/goreleaser/goreleaser/releases/download/$(GORELEASE_VERSION)/goreleaser
GORELEASE_URL_RPM 	:= $(GORELEASE_BASE_URL)_amd64.rpm

# ##################
# Makefile functions

define show_usage
	echo -e " Usage: ";\
	echo -e "\t make tag version=v0.1.0 message='first commit'\n";
endef

define deps_tag
	@if [[ $(version)x == "x" ]]; then \
		echo -e "\n Error: the version was not specified."; \
		$(call show_usage) \
		exit 1; \
	fi
	@if [[ "$(message)"x == "x" ]]; then \
		echo -e "\n Error: the commit message was not provided."; \
		$(call show_usage) \
		exit 1; \
	fi
endef

