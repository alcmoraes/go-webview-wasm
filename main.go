package main

import (
	"go-webview-wasm/desktop"
)

func main() {
	config := &desktop.Config{
		WasmBin: WASM_BIN,
		Width:   480,
		Height:  620 + 20,
		Title:   "GoWebview",
	}
	desktop.CreateDesktopApp(config)
}
