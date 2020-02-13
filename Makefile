VERSION := $(shell cat ./pkg/meta/contants.go | grep "ProgramVersion =" | cut -f 3 -d " " | cut -c 2- | rev | cut -c 2- | rev)

# Keep test at the top so that it is default when `make` is called.
# This is used by Travis CI.
coverage.txt:
	mkdir -p /tmp/cabiria
	go test -race -coverprofile=/tmp/cabiria/pkg_coverage.txt -covermode=atomic -coverpkg=./pkg/... ./test/pkg/...
	cat /tmp/cabiria/*_coverage.txt > coverage.txt
view-cover: coverage.txt
	go tool cover -html=coverage.txt
test: build
	go test ./test/...
build:
	go build ./...
install: build
	go install ./...
inspect: build
	golint ./...
pre-commit: clean coverage.txt inspect
	go mod tidy
release: pre-commit
	@echo -n "Going to tag this branch as $(VERSION). Proceed? [y/N] " && read ans && [ $${ans:-N} = y ]
	git add -A
	git commit -m "Release $(VERSION)"
	git tag -a $(VERSION) -m "Tagging for release of $(VERSION)"
	git push
	go list -m
clean:
	rm -f coverage.txt