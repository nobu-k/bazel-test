load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "proto",
    srcs = ["goproto.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "goproto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/nobu-k/bazel-go-test/goproto",
    proto = ":proto",
    visibility = ["//visibility:public"],
)
