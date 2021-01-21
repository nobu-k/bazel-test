load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "hello",
    embed = [":hello_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "hello_lib",
    srcs = ["main.go"],
    importpath = "github.com/nobu-k/bazel-go-test/hello",
    visibility = ["//visibility:private"],
    deps = ["@com_github_google_uuid//:go_default_library"],
)
