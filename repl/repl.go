package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/andrearcaina/bamboo/lexer"
	"github.com/andrearcaina/bamboo/token"
)

// PROMPT is the prompt for the REPL (Read-Eval-Print Loop)
const PROMPT = ">> "

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("{Type: '%s', Literal: '%s'}\n", tok.Type, tok.Literal)
		}
	}
}
