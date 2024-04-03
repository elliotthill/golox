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
		parser.statements = append(parser.statements, parser.declaration())
	}

	return parser.statements
}

func (parser *Parser) declaration() AbstractStatement {

	/*if parser.match(FUN) {
	      return parser.function("function")
	  }

	  if parser.match(VAR) {
	      return parser.varDeclaration()
	  }*/

	return parser.statement()
}

func (parser *Parser) statement() AbstractStatement {

	if parser.match(PRINT) {
		return parser.printStatement()
	}
	/*if parser.match(LEFT_BRACE) {
	    return Block{}
	}*/

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

/*
 * Starts the expression tree
 */
func (parser *Parser) expression() AbstractExpression {

	return parser.assignment()

	expr := parser.comparison()

	return expr
}

func (parser *Parser) assignment() AbstractExpression {

	expr := parser.or()

	if parser.match(EQUAL) {
		parser.previous()
		value := parser.assignment()

		variable, ok := expr.(Variable)
		if ok {
			return Assign{name: variable.name, value: value}
		} else {
			panic("Invalid assignment target")
		}
	}

	return expr
}

func (parser *Parser) or() AbstractExpression {

	expr := parser.and()

	for parser.match(OR) {
		operator := parser.previous()
		right := parser.and()
		expr = Logical{left: expr, operator: operator, right: right}
	}
	return expr
}

func (parser *Parser) and() AbstractExpression {

	expr := parser.equality()

	for parser.match(AND) {

		operator := parser.previous()
		right := parser.equality()
		expr := Logical{left: expr, operator: operator, right: right}
		return expr
	}

	return expr
}

func (parser *Parser) equality() AbstractExpression {

	expr := parser.comparison()

	for parser.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := parser.previous()
		right := parser.comparison()
		expr := Binary{left: expr, operator: operator, right: right}
		return expr
	}
	return expr
}

func (parser *Parser) comparison() AbstractExpression {

	expr := parser.term()

	for parser.match(GREATER) {

		operator := parser.previous()
		right := parser.primary()
		expr := Binary{left: expr, operator: operator, right: right}
		return expr
	}

	return expr
}

func (parser *Parser) term() AbstractExpression {

	expr := parser.factor()
	for parser.match(MINUS, PLUS) {
		operator := parser.previous()
		right := parser.factor()
		expr = Binary{left: expr, operator: operator, right: right}
		return expr
	}

	return expr
}

func (parser *Parser) factor() AbstractExpression {

	expr := parser.unary()

	for parser.match(SLASH, STAR) {
		operator := parser.previous()
		right := parser.unary()
		expr := Binary{left: expr, operator: operator, right: right}
		return expr
	}

	return expr
}

func (parser *Parser) unary() AbstractExpression {

	if parser.match(BANG, MINUS) {
		operator := parser.previous()
		right := parser.unary()
		expr := Unary{operator: operator, right: right}
		return expr
	}
	return parser.primary()
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

/*
* Control flow functions
 */
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
