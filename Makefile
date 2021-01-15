build: build_front \
	build_goland

build_wasm:
	GOARCH=wasm GOOS=js go build -o web/public/javascript/lib/utils.wasm internal/wasm/wasm.go

build_front:
	cd web/ && yarn build


build_goland:
	go build -o dist/chat/main cmd/chat/main.go

install:
	cd web
	npm i
	yarn build
	cd ../cmd/chat
	go build main.go -o ../../main

clean:
	rm -rf web/dist
	rm -rf dist

help:
	@echo "                                                             "
	@echo "                  Существующие команды                       "
	@echo "-------------------------------------------------------------"
	@echo "install       : Установка всех модулей и сборка прокта	    "
	@echo "clean         : Очистить проект                              "
	@echo "build_wasm    : Сборка WASM модулей               	        "
	@echo "-------------------------------------------------------------"
	@echo "                                                             "
