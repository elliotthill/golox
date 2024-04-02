package main

import (
	"fmt"
)

/*
* Expressions
 */
type AbstractExpression interface {
    Accept(visitor ExpressionVisitor) interface{}
}

type Literal struct{
    AbstractExpression
    value interface{}
}
func (literal *Literal) String() string {

   return fmt.Sprintf("%v", literal.value)
}
func (literal Literal) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitLiteralExpression(literal)
}


//Assign
type Assign struct{
    AbstractExpression
    name Token
    value AbstractExpression
}

func (assign Assign) Accept(visitor ExpressionVisitor) interface{} {
    fmt.Println("Assign Expression")
    return visitor.visitAssignExpression(assign)
}

//Unary
type Unary struct{
    AbstractExpression
    operator Token
    right AbstractExpression
}

func (unary Unary) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitUnaryExpression(unary)
}

//Binary
type Binary struct{
    AbstractExpression
    left AbstractExpression
    operator Token
    right AbstractExpression
}

func (binary Binary) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitBinaryExpression(binary)
}

//Ternary
type Ternary struct{
    AbstractExpression
    left AbstractExpression
    leftOperator Token
    middle AbstractExpression
    rightOperator Token
    right AbstractExpression
}

func (ternary Ternary) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitTernaryExpression(ternary)
}

//Grouping
type Grouping struct{
    AbstractExpression
    expression AbstractExpression
}

func (grouping Grouping) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitGroupingExpression(grouping)
}

//Variable
type Variable struct{
    AbstractExpression
    name Token
}

func (variable Variable) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitVariableExpression(variable)
}

//Logical
type Logical struct{
    AbstractExpression
    left AbstractExpression
    operator Token
    right AbstractExpression
}

func (logical Logical) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitLogicalExpression(logical)
}


type ExpressionVisitor interface {
    visitAssignExpression(expression Assign) interface{}
    visitBinaryExpression(expression Binary) interface{}
    visitGroupingExpression(expression Grouping) interface{}
    visitLiteralExpression(expression Literal) interface{}
    visitLogicalExpression(expression Logical) interface{}
    visitTernaryExpression(expression Ternary) interface{}
    visitUnaryExpression(expression Unary) interface{}
    visitVariableExpression(expression Variable) interface{}
}


/*
* Statements
*/

type AbstractStatement interface{
    Accept(visitor StatementVisitor) interface{}
}
//Block
type Block struct {
    AbstractStatement
    statements []AbstractStatement
}

func (block Block) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitBlockStatement(block)
}

//Expression
type Expression struct{
    AbstractStatement
    expression AbstractExpression
}

func (expression Expression) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitExpressionStatement(expression)
}

//Print
type Print struct{
    AbstractStatement
    expression AbstractExpression
}

func (print Print) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitPrintStatement(print)
}

//Var
type Var struct{
    name Token
    initializer AbstractExpression
}

func (_var Var) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitVarStatement(_var)
}

type StatementVisitor interface{
    visitBlockStatement(statement Block) interface{}
    //visitClassStatement(statement Class)
    visitExpressionStatement(statement Expression) interface{}
    visitPrintStatement(statement Print) interface{}
    visitVarStatement(statement Var) interface{}
}



