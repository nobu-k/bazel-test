package enums

// NOTE: go generate below didn't work because bazel run seems to change its CWD
//go:generate bazel run @org_golang_x_tools//cmd/stringer -- -output lib_string.go -type Enum
type Enum int

const (
	A Enum = iota
	B
	C
)
