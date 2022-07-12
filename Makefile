BINARY_NAME=cricinfo
CACHE_DIR=cache
RELEASE_DIR=release

build:
	go build -o release/${BINARY_NAME} .

build-win_amd64:
	GOARCH=amd64 GOOS=windows go build -o release/win-amd64/${BINARY_NAME}-win_amd64.exe .

build-linux_amd64:
	GOARCH=amd64 GOOS=linux go build -o release/linux-amd64/${BINARY_NAME}-linux_amd64 .

build-all: build-win_amd64 build-linux_amd64

lint:
	golangci-lint

clean:
	rm -rf ${CACHE_DIR}
	rm -rf ${RELEASE_DIR}

test: 
	go test -v coverprofile coverage.txt
	go tool cover -html coverage.html