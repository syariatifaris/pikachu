#!/bin/bash

export NOW=$(shell date --rfc-3339=ns)
export PKGS=$(shell go list ./... | grep -v vendor/)

all: update build

build:
	@echo "${NOW} ===BUILDING CORE==="
	@go build -o pikachu .
	@echo "${NOW} ====Done==="

update:
	@echo "${NOW} ===UPDATING DEP==="
	@dep ensure -v
	@echo "${NOW} ====Done==="