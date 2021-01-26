load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "server_lib",
    srcs = ["main.go"],
    importpath = "github.com/nobu-k/bazel-go-test/echo/server",
    visibility = ["//visibility:private"],
    deps = [
        "//goproto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)
