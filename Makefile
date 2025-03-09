SHELL = /bin/bash
TEST_FLAGS :=
TEST_PKG := ./...
GO := go

build: synonym

synonym:
	${GO} build -o synonym .

run:
	${GO} run .

test:
	${GO} test ${TEST_FLAGS} ${TEST_PKG}

cover:
	${GO} test -coverprofile .coverageprofile ${TEST_FLAGS} ${TEST_PKG}
	${GO} tool cover -func=.coverageprofile

lint:
	@golangci-lint run

tidy:
	${GO} mod tidy

fmt:
	@${GO} fmt ./...

clean:
	${GO} clean -testcache
	rm -f synonym

.PHONY: run test tidy fmt lint clean cover vendor
