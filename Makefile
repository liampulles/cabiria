VERSION=$(shell cat ./pkg/meta/constants.go | grep "ProgramVersion =" | cut -f 3 -d " " | cut -c 2- | rev | cut -c 2- | rev)
CURRENT_VERSION=$(shell git tag | tail -n 1)
PREDICTOR_PATH=$(shell cat ./pkg/intertitle/constants.go | grep "PredictorPath =" | cut -f 3 -d " " | cut -c 2- | rev | cut -c 2- | rev)
PREDICTOR_FILENAME=$(shell cat ./pkg/intertitle/constants.go | grep "PredictorFilename =" | cut -f 3 -d " " | cut -c 2- | rev | cut -c 2- | rev)

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
	sudo mkdir -p $(PREDICTOR_PATH)
	sudo cp ./data/intertitle/$(PREDICTOR_FILENAME) $(PREDICTOR_PATH)/$(PREDICTOR_FILENAME)
inspect: build
	golint ./...
pre-commit: clean coverage.txt inspect
	go mod tidy
release: pre-commit
	@echo -n "Going to release $(VERSION) (current version is $(CURRENT_VERSION)). Proceed? [y/N] " && read ans && [ $${ans:-N} = y ]
	git add -A
	git commit -m "Release $(VERSION)"
	git tag $(VERSION)
	git push origin --tags
	git push
	go list -m github.com/liampulles/cabiria@$(VERSION)
clean:
	rm -f ${GOPATH}/bin/cabiria*
	rm -f coverage.txt