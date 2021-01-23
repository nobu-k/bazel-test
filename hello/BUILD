load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_push")

go_binary(
    name = "hello",
    embed = [":hello_lib"],
    pure = "on",
    visibility = ["//visibility:public"],
    x_defs = {"Version": "{STABLE_HELLO_VERSION}"},
)

go_image(
    name = "image",
    binary = ":hello",
    visibility = ["//visibility:public"],
)

container_push(
    name = "push",
    image = ":image",
    format = "Docker",
    registry = "index.docker.io",
    repository = "nobu/bazel-test",
    tag = "0.1.0",
)

go_library(
    name = "hello_lib",
    srcs = ["main.go"],
    importpath = "github.com/nobu-k/bazel-go-test/hello",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_google_uuid//:uuid",
        "@com_github_makenowjust_heredoc//:heredoc",
    ],
)
