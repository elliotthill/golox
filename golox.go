package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

    if len(os.Args) < 2 {
        fmt.Println("Error: please provide a file arg")
        return
    }

    sourceCode, error := loadFromFile(os.Args[1])

    if error != nil {
        fmt.Println(error);
        return;
    }

    scanner := Scanner{source: sourceCode}
    tokens := scanner.Scan()


    for _, element := range tokens{
        fmt.Println(element)
    }

    parser := Parser{tokens: tokens}
    statements := parser.Parse()


    fmt.Println("== Parsed Tree ==")
    for _, element := range statements {

        fmt.Println(element)
    }

    fmt.Println("== Interp ==")
    interpreter := Interpreter{statements: statements}
    interpreter.Interpret()

    return str

}



func loadFromFile (filename string) (string, error) {

    b, error := os.ReadFile(filename)
    if error != nil {
        fmt.Print(error)
        return "", errors.New("Could not read filename");
    }

    contents := string(b)
    fmt.Print(contents)
    return contents, nil
}
