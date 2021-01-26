package main

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/google/uuid"
	"github.com/nobu-k/bazel-test/enums"
)

var Version string

func main() {
	fmt.Println(Version, ": Hello ", uuid.New())
	fmt.Println(enums.A, enums.B, enums.C)
	fmt.Println(heredoc.Doc(`This was generated
	by heredoc.
	`))
}
