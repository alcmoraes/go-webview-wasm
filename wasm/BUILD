load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
    ],
    importpath = "go-webview-wasm/wasm",
    visibility = ["//visibility:private"],
    deps = [
        "//js/canvas:go_default_library",
    ],
)

go_binary(
    name = "wasm",
    embed = [":go_default_library"],
    goarch = "wasm",
    goos = "js",
    visibility = ["//visibility:public"],
)
