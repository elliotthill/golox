package main

import "fmt"

type Parser struct {
	tokens     []Token
	current    int
	statements []AbstractStatement
}

func NewParser(tokens []Token) {

	parser := new(Parser)
	parser.current = 0
	parser.tokens = tokens
}

func (parser *Parser) Parse() []AbstractStatement {

	for !parser.isAtEnd() {
		parser.statements = append(parser.statements, parser.statement())
	}

	return parser.statements
}

func (parser *Parser) statement() AbstractStatement {

	if parser.match(PRINT) {
		return parser.printStatement()
	}

	return parser.expressionStatement()
}

func (parser *Parser) printStatement() AbstractStatement {

	value := parser.expression()
	parser.consume(SEMICOLON, "Expected ; after value.")
	return Print{expression: value}
}

func (parser *Parser) expressionStatement() AbstractStatement {

	expr := parser.expression()
	parser.consume(SEMICOLON, "Expect ; after expression.")

	expr_statement := Expression{expression: expr}
	return expr_statement
}

func (parser *Parser) isAtEnd() bool {
	return parser.peek().tokenType == EOF
}

func (parser *Parser) peek() Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) previous() Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) advance() Token {
	if !parser.isAtEnd() {
		parser.current++
	}

	return parser.previous()
}

func (parser *Parser) match(tokenTypes ...TokenType) bool {

	for _, tokenType := range tokenTypes {
		if parser.check(tokenType) {
			parser.advance()
			return true
		}
	}

	return false
}

func (parser *Parser) check(tokenType TokenType) bool {

	if parser.isAtEnd() {
		return false
	}
	return parser.peek().tokenType == tokenType
}

func (parser *Parser) consume(tokenType TokenType, message string) Token {
	if parser.check(tokenType) {
		return parser.advance()
	}

	m := fmt.Sprintf("%s %s", parser.peek().tokenType, message)
	panic(m)
}

/*
 * Starts the expression tree
 */

func (parser *Parser) expression() AbstractExpression {

	expr := parser.comparison()

	return expr
}

func (parser *Parser) comparison() AbstractExpression {

	expr := parser.primary()

	for parser.match(GREATER) {

		operator := parser.previous()
		right := parser.primary()
		expr := Binary{left: expr, operator: operator, right: right}
		return expr
	}

	return expr
}

func (parser *Parser) primary() AbstractExpression {

	if parser.match(FALSE) {
		return Literal{value: false}
	}

	if parser.match(TRUE) {
		return Literal{value: true}
	}

	if parser.match(NIL) {
		return Literal{value: nil}
	}

	if parser.match(NUMBER, STRING) {
		expr := Literal{value: parser.previous().literal}
		return expr
	}

	if parser.match(IDENTIFIER) {
		return Variable{name: parser.previous()}
	}

	if parser.match(LEFT_PAREN) {
		expr := parser.expression()
		parser.consume(RIGHT_PAREN, "Expect ')' after expression. ")
		return Grouping{expression: expr}
	}

	panic("Expected expression")
}
