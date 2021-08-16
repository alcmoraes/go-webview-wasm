load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "ns_window.go",
        "ns_window.h",
        "ns_window.m",
    ],
    cgo = True,
    clinkopts = ["-framework Cocoa"],
    copts = ["-x objective-c"],
    importpath = "go-webview-wasm/desktop/cocoa",
    visibility = ["//visibility:public"],
)
