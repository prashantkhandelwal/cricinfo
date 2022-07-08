BINARY_NAME=cricinfo
CACHE_DIR=cache
RELEASE_DIR=release

build:
	go build -o release/${BINARY_NAME} main.go urls.go t20.go config.go

build-win_amd64:
	GOARCH=amd64 GOOS=windows go build -o release/win-amd64/${BINARY_NAME}-win_amd64.exe main.go urls.go t20.go config.go

build-linux_amd64:
	GOARCH=amd64 GOOS=linux go build -o release/linux-amd64/${BINARY_NAME}-linux_amd64 main.go urls.go t20.go config.go

build-all: build-win_amd64 build-linux_amd64

clean:
	rm -rf ${CACHE_DIR}
	rm -rf ${RELEASE_DIR}