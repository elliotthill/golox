package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Scanner struct {
    source string;
    current int;
    start int;
    line int;
    tokens []Token
}

func NewScanner(source string) *Scanner {
    scanner := new(Scanner)
    scanner.source = source
    scanner.current = 0
    scanner.start = 0
    scanner.line = 0
    //scanner.tokens = make([]Token, 0)
    return scanner
}

func (scanner *Scanner) Scan() []Token {

    for scanner.notAtEnd() {

        scanner.start = scanner.current

        switch c := scanner.advance(); c {

        case "(":
            scanner.addToken(LEFT_PAREN)
        case ")":
            scanner.addToken(RIGHT_PAREN)
        case "{":
            scanner.addToken(LEFT_BRACE)
        case "}":
            scanner.addToken(RIGHT_BRACE)
        case ",":
            scanner.addToken(COMMA)
        case ".":
            scanner.addToken(DOT)
        case "-":
            scanner.addToken(MINUS)
        case "+":
            scanner.addToken(PLUS)
        case ";":
            scanner.addToken(SEMICOLON)
        case "*":
            scanner.addToken(STAR)
        case "!":
            if scanner.match("=") {
                scanner.addToken(BANG_EQUAL)
            } else {
                scanner.addToken(EQUAL)
            }
        case "=":
            if scanner.match("=") {
                scanner.addToken(EQUAL_EQUAL)
            } else {
                scanner.addToken(EQUAL)
            }
        case "<":
            if scanner.match("=") {
                scanner.addToken(LESS_EQUAL)
            } else {
                scanner.addToken(LESS)
            }
        case ">":
            if scanner.match("=") {
                scanner.addToken(GREATER_EQUAL)
            } else {
                scanner.addToken(GREATER)
            }
        case "/":
            if scanner.match("/") {
                for scanner.peek() != "\n" && scanner.notAtEnd() {
                    scanner.advance()
                }
            } else {
                scanner.addToken(SLASH)
            }
        case " ":
        case "\r":
        case "\t":
        case "\n":
            scanner.line++
        case "\"":
            scanner.string()
        case "'":
            scanner.string()
        default:
            if scanner.isDigit(c) {
                scanner.number()
                //scanner.number()
            } else if scanner.isAlpha(c) {
                scanner.identifier()
            } else {
                fmt.Print("ERROR")
            }
        }
    }
    scanner.addToken(EOF)
    return scanner.tokens
}

func (scanner *Scanner) advance() string {

    val := string(scanner.source[scanner.current])
    scanner.current++
    return val
}

func (scanner *Scanner) addTokenLiteral(tokenType TokenType, literal interface{}) {

    text := scanner.source[scanner.start:scanner.current]
    newToken :=Token{tokenType, text, literal, scanner.line}
    scanner.tokens = append(scanner.tokens, newToken)
}

func (scanner *Scanner) addToken(tokenType TokenType) {

    scanner.addTokenLiteral(tokenType, nil)
}

//Match next character and consume
func (scanner *Scanner) match(expected string) bool {

    if (!scanner.notAtEnd()) {
        return false
    }

    if (string(scanner.source[scanner.current]) != expected) {
        return false //Not the character we want, exit
    }

    //Consume the expected token
    scanner.current++
    return true
}

//Lookahead one character
func (scanner *Scanner) peek() string {

    if scanner.isAtEnd() {
        return ""
    }
    return string(scanner.source[scanner.current])
}

func (scanner *Scanner) peekNext() string {

    if scanner.current + 1 >= len(scanner.source) {
        return ""
    }

    return string(scanner.source[scanner.current+1])
}

func (scanner *Scanner) isAlphaNumeric(char string) bool {
    return scanner.isAlpha(char) || scanner.isDigit(char)
}

func (scanner *Scanner) isDigit(char string) bool {
    _, error := strconv.Atoi(char)

    return error == nil
}

func (scanner *Scanner) isAlpha(char string) bool {

    for _, c := range char {
        return unicode.IsLetter(c) || c == '_'
    }
    return false
}


//Match string until ' or "
func (scanner *Scanner) string() {

    for scanner.peek() != "'" && scanner.peek() != "\"" && scanner.notAtEnd() {
        if scanner.peek() == "\n" {
            scanner.line++
        }
        scanner.advance();
    }

    if (scanner.isAtEnd()) {
        fmt.Println("Unterminated string")
        return
    }

    scanner.advance()

    value := scanner.source[scanner.start+1:scanner.current-1]
    scanner.addTokenLiteral(STRING, value)
}

func (scanner *Scanner) identifier() {

    for scanner.isAlphaNumeric(scanner.peek()) {
        scanner.advance()
    }

    text := scanner.source[scanner.start:scanner.current]
    tokenType, ok := keywords[text]

    if (!ok) {
        tokenType = IDENTIFIER
    }

    scanner.addToken(tokenType)
}

func (scanner *Scanner) number() {

    for scanner.isDigit(scanner.peek()) {
        scanner.advance()
    }

    if scanner.peek() == "." && scanner.isDigit(scanner.peekNext()) {
        scanner.advance()

        for scanner.isDigit(scanner.peek()) {
            scanner.advance()
        }
    }

    float,error := strconv.ParseFloat(scanner.source[scanner.start:scanner.current],64)

    if (error != nil) {
        fmt.Print("ERROR cannot parse float " + scanner.source[scanner.start:scanner.current])
    }
    scanner.addTokenLiteral(NUMBER, float)

}

func (scanner *Scanner) notAtEnd() bool {

    return scanner.current < len(scanner.source)
}
func (scanner *Scanner) isAtEnd() bool {
    return !scanner.notAtEnd()
}
