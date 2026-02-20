PLUGIN_PATH := .

build_plugin:
	go build -buildmode=plugin -o ${PLUGIN_PATH}/linter.so cmd/linter.go

.PHONY: test build_plugin install clean

install: dependencies build_plugin clean

test:
	go test ./pkg/analyzer

dependencies:
	go work init
	go work use .
	git clone https://github.com/golangci/golangci-lint.git && go work use golangci-lint
	cd golangci-lint && git checkout tags/v1.61.0 && go build -o ../golangci ./cmd/golangci-lint && cd ../

clean:
	rm -rf golangci-lint
	rm -f go.work go.work.sum

build_windows:
	go build -o ./linter.exe ./cmd/linter.go
