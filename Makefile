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
clean:
	rm coverage.txt