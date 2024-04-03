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

    if parser.match(FUN) {
        return parser.function("function")
    }

    if parser.match(VAR) {
        return parser.varDeclaration()
    }

	return parser.statement()
}

func (parser *Parser) statement() AbstractStatement {

    if parser.match(FOR) {
        return parser.forStatement()
    }
    if parser.match(IF) {
        return parser.ifStatement()
    }
    if parser.match(PRINT) {
		return parser.printStatement()
	}
    if parser.match(RETURN) {
        return parser.returnStatement()
    }
    if parser.match(WHILE) {
        return parser.whileStatement()
    }
    if parser.match(LEFT_BRACE) {
        return Block{statements: parser.block()}
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

func (parser *Parser) function(kind string) Function {

    name := parser.consume(IDENTIFIER, "Expect " + kind + " name.")
    parser.consume(LEFT_PAREN, "Expect '(' after " + kind + " name.")

    parameters := []Token{}
    if !parser.check(RIGHT_PAREN) {

        //Keep matching params between ,
        for params := true; params; params = parser.match(COMMA) {

            parameters = append(parameters, parser.consume(IDENTIFIER, "Expect parameter name"))
        }
    }

    parser.consume(RIGHT_PAREN, "Expect ')' after parameters.")
    parser.consume(LEFT_BRACE, "Expect '{' before " + kind + " body.")
    body := parser.block()
    return Function{name:name, params:parameters, body:body}
}

func (parser *Parser) varDeclaration() AbstractStatement{
    name := parser.consume(IDENTIFIER, "Expect variable name.")

    var initializer AbstractExpression = nil
    if parser.match(EQUAL){
        initializer = parser.expression()
    }
    parser.consume(SEMICOLON, "Expected ';' after variable declaration")

    fmt.Println(fmt.Sprintf("%s %s", name.tokenType, initializer))
    return Var{name:name, initializer: initializer}
}

func (parser *Parser) block() []AbstractStatement {

    statements := []AbstractStatement{}

    for !parser.check(RIGHT_BRACE) && !parser.isAtEnd(){
        statements = append(statements, parser.declaration())
    }
    parser.consume(RIGHT_BRACE, "Expect '}' after block.")
    return statements
}

func (parser *Parser) ifStatement() AbstractStatement {

    parser.consume(LEFT_PAREN, "Expect '(' after 'if'")
    condition := parser.expression()
    parser.consume(RIGHT_PAREN, "Expect ')' after if condition")

    thenBranch := parser.statement()
    var elseBranch AbstractStatement = nil

    if parser.match(ELSE) {
        elseBranch = parser.statement()
    }

    return If{condition: condition, thenBranch: thenBranch, elseBranch: elseBranch}
}

func (parser *Parser) forStatement() AbstractStatement {

    parser.consume(LEFT_PAREN, "Expect '(' after 'for'")

    //Initializer i = 0
    var initializer AbstractStatement

    if parser.match(SEMICOLON) {
        initializer = nil
    } else if parser.match(VAR) {
        initializer = parser.varDeclaration()
    } else {
        initializer = parser.expressionStatement()
    }
    //Condition i < x
    var condition AbstractExpression = nil

    if !parser.check(SEMICOLON) {
        fmt.Println("Checking for condition")
        condition = parser.expression()
        fmt.Println(condition)
    }
    parser.consume(SEMICOLON, "Expect ';' after loop condition")

    //Increment i++
    var increment AbstractExpression = nil
    if !parser.check(RIGHT_PAREN) {
        increment = parser.expression()
    }
    parser.consume(RIGHT_PAREN, "Expect ')' after for clauses.")

    //body
    body := parser.statement()

    //Increment and evaluate
    if increment != nil {
        body_statements := []AbstractStatement{}
        body_statements = append(body_statements, body)
        body_statements = append(body_statements, Expression{expression: increment})
        body = Block{statements: body_statements}
    }

    if condition == nil {
        condition = Literal{value: true}
    }

    body = While{condition: condition, body: body}

    if initializer != nil {
        body_statements := []AbstractStatement{}
        body_statements = append(body_statements, initializer)
        body_statements = append(body_statements, body)

        body = Block{statements: body_statements}
    }

    return body
}

func (parser *Parser) whileStatement() AbstractStatement {

    parser.consume(LEFT_PAREN, "Expect ')' after 'while' ")
    condition := parser.expression()
    parser.consume(RIGHT_PAREN, "Expected ')' after condition")

    body := parser.statement()

    return While{condition: condition, body: body}
}

func (parser *Parser) returnStatement() AbstractStatement {

    keyword := parser.previous()
    var value AbstractExpression = nil

    if !parser.check(SEMICOLON) {
        value = parser.expression()
    }

    parser.consume(SEMICOLON, "Expect ';' after return value.")
    return Return{keyword: keyword, value: value}
}


/*
 * Starts the expression tree
 */
func (parser *Parser) expression() AbstractExpression {

	return parser.assignment()

}

func (parser *Parser) assignment() AbstractExpression {

	expr := parser.or()

	if parser.match(EQUAL) {
        equals := parser.previous()
		value := parser.assignment()

		variable, ok := expr.(Variable)
		if ok {

			return Assign{name: variable.name, value: value}
		} else {
			panic(equals.tokenType + "Invalid assignment target")
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

	for parser.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {

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
	return parser.call()
}

func (parser *Parser) call() AbstractExpression {

    expr := parser.primary()

    for parser.match(LEFT_PAREN) {
        expr = parser.finishCall(expr)
    }

    return expr
}

func (parser *Parser) finishCall(callee AbstractExpression) AbstractExpression {

    arguments := []AbstractExpression{}

    if !parser.check(RIGHT_PAREN) {

        for match_comma := true; match_comma; match_comma = parser.match(COMMA) {

            arguments = append(arguments, parser.expression())
        }

    }
    paren := parser.consume(RIGHT_PAREN, "Expect ')' after arguments")

    return Call{callee: callee, paren: paren, arguments: arguments }
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
