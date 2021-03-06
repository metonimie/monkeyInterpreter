package repl

import (
	"io"
	"bufio"
	"fmt"
	"github.com/metonimie/monkeyInterpreter/lexer"
	"github.com/metonimie/monkeyInterpreter/parser"
	"github.com/metonimie/monkeyInterpreter/evaluator"
	"github.com/metonimie/monkeyInterpreter/object"
)

const PROMPT = ">"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			goto exit
		}
		line := scanner.Text()

		if line == "exit" {
			goto exit
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		//io.WriteString(out, program.String())
		//io.WriteString(out, "\n")
		ev := evaluator.Eval(program, env)
		if ev != nil {
			io.WriteString(out, ev.Inspect())
			io.WriteString(out, "\n")
		}
	}

exit:
	return
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
