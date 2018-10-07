
# Makefile extensions assets

# ###########
# Global Vars

CPWD := $(PWD)

BASE_DIRS := ./bin
BASE_DIRS += ./dist

GORELEASE_VERSION := v0.86.1
GORELEASE_BASE_URL := https://github.com/goreleaser/goreleaser/releases/download/$(GORELEASE_VERSION)/goreleaser
GORELEASE_URL_RPM := $(GORELEASE_BASE_URL)_amd64.rpm

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

# Ensure directories exists
define deps_dirs
	@mkdir -p $(BASE_DIRS)
endef

# Ensure directories are removed
define deps_clean
	@for d in $(BASE_DIRS); do \
		rm -rvf $$d; \
	done
endef