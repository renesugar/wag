include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile
.PHONY: test
PKG := github.com/Clever/wag
PKGS := $(shell go list ./... | grep -v /vendor)
$(eval $(call golang-version-check,1.6))

test:
	rm -rf generated/*
	go run main.go genclients.go -file swagger.yml -package $(PKG)/generated
	cd impl && go build
	cd test && go test

vendor: golang-godep-vendor-deps
	$(call golang-godep-vendor,$(PKGS))
