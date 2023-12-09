package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ljpurcell/ape-interpreter/repl"
	"github.com/ljpurcell/ape-interpreter/utils"
)

var input = "input.txt"

func main() {
    user, err := user.Current()
    utils.Check(err, "Could not get OS user")

    fmt.Printf("Hello, %v. This is the Ape interpreter!\n", user.Username)
    repl.Start(os.Stdin, os.Stdout)
}
