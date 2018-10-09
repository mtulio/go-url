#!/bin/bash

WORKDIR=$(dirname $0)

test -x ${WORKDIR}/go-url || chmod +x ${WORKDIR}/bin/go-url

${WORKDIR}/go-url $@