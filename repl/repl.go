package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tuantran1810/go-interpreter/evaluator"
	"github.com/tuantran1810/go-interpreter/lexer"
	"github.com/tuantran1810/go-interpreter/parser"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		obj := evaluator.Eval(program)

		io.WriteString(out, obj.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
