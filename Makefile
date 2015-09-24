include .version

all: linux osx windows
bundles-dir:
	mkdir -p ./bundles
linux: bundles-dir
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 godep go build -a -installsuffix cgo -o ./bundles/test-runner_linux-amd64 .
osx: bundles-dir
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 godep go build -a -installsuffix cgo -o ./bundles/test-runner_darwin-amd64 .
windows: bundles-dir
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 godep go build -a -installsuffix cgo -o ./bundles/test-runner_windows-amd64 .
install:
	cp ./bundles/test-runner_linux-amd64 $(GOPATH)/bin/test-runner
bintray:
	curl -T main -uflower-pot:$(BINTRAY_API_KEY) https://api.bintray.com/content/emerald-ci/test-runner/binary/$(TEST_RUNNER_VERSION)/test-runner_linux-amd64
docker:
	docker build -t emeraldci/test-runner .
