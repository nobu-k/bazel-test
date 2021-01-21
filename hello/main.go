package main

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/google/uuid"
)

var Version string

func main() {
	fmt.Println(Version, ": Hello ", uuid.New())
	fmt.Println(heredoc.Doc(`This was generated
	by heredoc.
	`))
}
