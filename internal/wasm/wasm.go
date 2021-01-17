package main

import (
	"syscall/js"

	"github.com/HunterGooD/Chat-Go/internal/app"

	"github.com/HunterGooD/Chat-Go/pkg/crypto"
)

func base64decode() js.Func {
	return js.FuncOf(func(this js.Value, input []js.Value) interface{} {
		return crypto.Base64Decode(input[0].String())
		//return "test decode"
	})
}

func base64encode() js.Func {
	return js.FuncOf(func(this js.Value, input []js.Value) interface{} {
		return crypto.Base64Encode(input[0].String())
		//return "test encode"
	})
}

func getAesKey() js.Func {
	return js.FuncOf(func(this js.Value, input []js.Value) interface{} {
		return app.AESKEY
	})
}

func main() {
	js.Global().Set("base64decode", base64decode())
	js.Global().Set("base64encode", base64encode())
	js.Global().Set("getAesKey", getAesKey())
	<-make(chan bool)
}
