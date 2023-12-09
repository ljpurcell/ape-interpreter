package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ljpurcell/ape-interpreter/lexer"
	"github.com/ljpurcell/ape-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT)
        if !scanner.Scan() {
            return
        }

        line := scanner.Text()
        l := lexer.NewLexer(line)

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}
