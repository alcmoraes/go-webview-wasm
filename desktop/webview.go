// +build darwin,cgo

package desktop

import (
	"fmt"
	"net/url"

	"go-webview-wasm/desktop/cocoa"

	"github.com/webview/webview"
)

func CreateDesktopApp(config *Config) {
	indexHTML := STATIC_STRINGS["desktop/index.html"]
	wasmEXEC := STATIC_STRINGS["desktop/wasm_exec.js"]
	initJS := STATIC_STRINGS["desktop/init.js"]
	reactAPP := STATIC_STRINGS["desktop/index.js"]

	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetSize(config.Width, config.Height, webview.HintNone)

	w.Navigate(fmt.Sprintf("data:text/html,%s", url.PathEscape(indexHTML)))

	window := cocoa.NSWindow{w.Window()}
	window.Center()
	window.SetTitle(config.Title)

	w.Bind("getWASM", func() []byte { return config.WasmBin })
	w.Init(wasmEXEC)
	w.Init(initJS)
	w.Init(reactAPP)

	w.Run()
}
