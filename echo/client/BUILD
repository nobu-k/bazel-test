load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "client_lib",
    srcs = ["main.go"],
    importpath = "github.com/nobu-k/bazel-test/echo/client",
    visibility = ["//visibility:private"],
    deps = [
        "//goproto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "client",
    embed = [":client_lib"],
    visibility = ["//visibility:public"],
)
