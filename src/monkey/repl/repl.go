package repl

import (
	"bufio"
	"fmt"
	"io"
	"interpreter/src/monkey/lexer"
	"interpreter/src/monkey/token"
	"strings"
)


const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if strings.TrimSpace(line) == "exit" {
			fmt.Fprintln(out, "Goodbye!")
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

	}
}