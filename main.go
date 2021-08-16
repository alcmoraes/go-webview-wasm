package main

import (
	"go-webview-wasm/desktop"
)

func main() {
	config := &desktop.Config{
		WasmBin: WASM_BIN,
		Width:   1280,
		Height:  720 + 20, // add 20 for the top bar
		Title:   "GoWebview",
	}
	desktop.CreateDesktopApp(config)
}
