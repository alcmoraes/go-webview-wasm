# Go Webview WASM
A go example project using Webview, WASM and Bazel

## Requirements

- [Bazel](https://bazel.build/)

## Usage

```bash
    # For linux
    bazel run :webview-wasm
    # For mac
    bazel run :webview-wasm-osx
    # For mac (cocoa)
    bazel run :webview-wasm-osx-cocoa
    # For windows
    bazel run :webview-wasm-windows
```

Open inspect and check the console after clicking the button. A message should be trigger that comes from Go compiled WASM.
