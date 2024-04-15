package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
    "github.com/elliotthill/golox/interpreter"
    "github.com/elliotthill/golox/parser"
    "github.com/elliotthill/golox/lexer"
)

func main() {

	var file string
	var debug bool

    flag.StringVar(&file, "f", "", "Input File")
	flag.BoolVar(&debug, "d", false, "Debug Mode")
	flag.Parse()

	if len(file) > 0 {
		sourceCode, error := loadFromFile(file)

		if error != nil {
			fmt.Println(error)
			return
		}

		Run(sourceCode, interpreter.NewInterpreter(), debug)

	} else {

		REPL(debug)
	}

}

func Run(source string, interpreter *interpreter.Interpreter, debug bool) {

	scanner := lexer.NewScanner(source)
	tokens := scanner.Scan()

	if debug {

		fmt.Println("== Tokens ==")

		for _, element := range tokens {
			fmt.Println(element)
		}
	}

	parser := parser.NewParser(tokens)
	statements := parser.Parse()

	if debug {
		fmt.Println("== Parse Tree ==")
		for _, element := range statements {

			fmt.Println(element)
		}

		fmt.Println("== Interp ==")
	}

    interpreter.SetStatements(statements);
	interpreter.Interpret()

}

func REPL(debug bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	interp := interpreter.NewInterpreter()

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			fmt.Println(err)
			os.Exit(64)
		}

		if string(line) == "exit" || string(line) == "exit()" {
			os.Exit(1)
		}

		Run(string(line), interp, debug)
		fmt.Print("> ")

	}

}

func loadFromFile(filename string) (string, error) {

	b, error := os.ReadFile(filename)
	if error != nil {
		fmt.Print(error)
		return "", errors.New("Could not read filename " + filename)
	}

	contents := string(b)
	fmt.Print(contents)
	return contents, nil
}
