build:
	go build cmd/chat/main.go

build_wasm:
	GOARCH=wasm GOOS=js go build -o web/static/lib.wasm internal/wasm/wasm.go