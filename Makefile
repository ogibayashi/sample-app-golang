# ビルドするバイナリの名前
BINARY_NAME=sample-app-golang
MAIN=./main.go
TMPDIR=./.tmp
GOLANGCI_LINT_VERSION=v1.55.2

# デフォルトのターゲット
all: lint test build

# バイナリのビルド
build:
	CGO_ENABLED=0 go build -o bin/${BINARY_NAME} ${MAIN}

# バイナリの実行
run:
	CGO_ENABLED=0 go build -o bin/${BINARY_NAME} ${MAIN}
	./bin/${BINARY_NAME}

# テストの実行
test:
	go test -v ./...

# 依存関係のダウンロード
deps:
	go mod download

lint: golangci-lint
	${TMPDIR}/golangci-lint run

# ビルドしたバイナリとその他の生成ファイルの削除
clean:
	go clean
	rm ${BINARY_NAME}

golangci-lint:
ifeq ($(wildcard ${TMPDIR}/golangci-lint), )
	mkdir -p ${TMPDIR}
	curl -sSfL  https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh  | sh -s -- -b ${TMPDIR} ${GOLANGCI_LINT_VERSION}
endif
