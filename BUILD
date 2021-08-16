load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_embed_data")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

# gazelle:prefix github.com/alcmoraes/go-webview-wasm
gazelle(name = "gazelle")

buildifier(name = "buildifier")

# Asteroids binaries

go_embed_data(
    name = "wasm_embed",
    src = "//wasm",
    package = "main",
    string = False,
    var = "WASM_BIN",
)

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        ":wasm_embed",
    ],
    importpath = "github.com/alcmoraes/go-webview-wasm",
    visibility = ["//visibility:private"],
    deps = ["//desktop:go_default_library"],
)

# use lorca
go_binary(
    name = "webview-wasm-linux",
    embed = [":go_default_library"],
    goarch = "amd64",
    goos = "linux",
    visibility = ["//visibility:public"],
)

# use lorca
# go_binary(
#     name = "webview-wasm",
#     embed = [":go_default_library"],
#     goarch = "amd64",
#     goos = "darwin",
#     visibility = ["//visibility:public"],
# )

# # use webview
# go_binary(
#     name = "asteroids_darwin",
#     cgo = True,
#     embed = [":go_default_library"],
#     goarch = "amd64",
#     goos = "darwin",
#     pure = "off",
#     visibility = ["//visibility:public"],
# )

# # use lorca
# go_binary(
#     name = "asteroids_windows",
#     embed = [":go_default_library"],
#     goarch = "amd64",
#     goos = "windows",
#     visibility = ["//visibility:public"],
# )