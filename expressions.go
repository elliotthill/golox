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

type Call struct{
    AbstractExpression
    callee AbstractExpression
    paren Token
    arguments []AbstractExpression
}

func (call Call) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.visitCallExpression(call)
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
    visitCallExpression(expression Call) interface{}
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
    AbstractStatement
    name Token
    initializer AbstractExpression
}

func (_var Var) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitVarStatement(_var)
}

//If
type If struct{
    AbstractStatement
    condition AbstractExpression
    thenBranch AbstractStatement
    elseBranch AbstractStatement
}

func (_if If) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitIfStatement(_if)
}

//While
type While struct{
    AbstractStatement
    condition AbstractExpression
    body AbstractStatement
}

func (while While) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitWhileStatement(while)
}

//Return
type Return struct{
    AbstractStatement
    keyword Token
    value AbstractExpression
}

func (_return Return) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitReturnStatement(_return)
}


//Function
type Function struct{
    AbstractStatement
    name Token
    params []Token
    body []AbstractStatement
}

func (_function Function) Accept(visitor StatementVisitor) interface{} {
    return visitor.visitFunctionStatement(_function)
}

type StatementVisitor interface{
    visitBlockStatement(statement Block) interface{}
    //visitClassStatement(statement Class)
    visitExpressionStatement(statement Expression) interface{}
    visitPrintStatement(statement Print) interface{}
    visitVarStatement(statement Var) interface{}
    visitIfStatement(statement If) interface{}
    visitWhileStatement(statement While) interface{}
    visitReturnStatement(statement Return) interface{}
    visitFunctionStatement(statement Function) interface{}
}



