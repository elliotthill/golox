package language

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
    Value interface{}
}
func (literal *Literal) String() string {

   return fmt.Sprintf("%v", literal.Value)
}
func (literal Literal) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitLiteralExpression(literal)
}


//Assign
type Assign struct{
    AbstractExpression
    Name Token
    Value AbstractExpression
}

func (assign Assign) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitAssignExpression(assign)
}

//Unary
type Unary struct{
    AbstractExpression
    Operator Token
    Right AbstractExpression
}

func (unary Unary) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitUnaryExpression(unary)
}

//Binary
type Binary struct{
    AbstractExpression
    Left AbstractExpression
    Operator Token
    Right AbstractExpression
}

func (binary Binary) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitBinaryExpression(binary)
}

//Ternary
type Ternary struct{
    AbstractExpression
    Left AbstractExpression
    LeftOperator Token
    Middle AbstractExpression
    RightOperator Token
    Right AbstractExpression
}

func (ternary Ternary) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitTernaryExpression(ternary)
}

//Grouping
type Grouping struct{
    AbstractExpression
    Expression AbstractExpression
}

func (grouping Grouping) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitGroupingExpression(grouping)
}

//Variable
type Variable struct{
    AbstractExpression
    Name Token
}

func (variable Variable) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitVariableExpression(variable)
}

//Logical
type Logical struct{
    AbstractExpression
    Left AbstractExpression
    Operator Token
    Right AbstractExpression
}

func (logical Logical) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitLogicalExpression(logical)
}

type Call struct{
    AbstractExpression
    Callee AbstractExpression
    Paren Token
    Arguments []AbstractExpression
}

func (call Call) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitCallExpression(call)
}


//Anonymous functions
type FunctionExpression struct{
    AbstractExpression
    Params []Token
    Body []AbstractStatement
}

func (funcExpr FunctionExpression) Accept(visitor ExpressionVisitor) interface{} {
    return visitor.VisitFunctionExpression(funcExpr);
}

type ExpressionVisitor interface {
    VisitAssignExpression(expression Assign) interface{}
    VisitBinaryExpression(expression Binary) interface{}
    VisitGroupingExpression(expression Grouping) interface{}
    VisitLiteralExpression(expression Literal) interface{}
    VisitLogicalExpression(expression Logical) interface{}
    VisitTernaryExpression(expression Ternary) interface{}
    VisitUnaryExpression(expression Unary) interface{}
    VisitVariableExpression(expression Variable) interface{}
    VisitCallExpression(expression Call) interface{}
    VisitFunctionExpression(expression FunctionExpression) interface{}
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
    Statements []AbstractStatement
}

func (block Block) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitBlockStatement(block)
}

//Expression
type Expression struct{
    AbstractStatement
    Expression AbstractExpression
}

func (expression Expression) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitExpressionStatement(expression)
}

//Print
type Print struct{
    AbstractStatement
    Expression AbstractExpression
}

func (print Print) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitPrintStatement(print)
}

//Var
type Var struct{
    AbstractStatement
    Name Token
    Initializer AbstractExpression
}

func (_var Var) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitVarStatement(_var)
}

//If
type If struct{
    AbstractStatement
    Condition AbstractExpression
    ThenBranch AbstractStatement
    ElseBranch AbstractStatement
}

func (_if If) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitIfStatement(_if)
}

//While
type While struct{
    AbstractStatement
    Condition AbstractExpression
    Body AbstractStatement
}

func (while While) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitWhileStatement(while)
}

//Return
type Return struct{
    AbstractStatement
    Keyword Token
    Value AbstractExpression
}

func (_return Return) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitReturnStatement(_return)
}


//Function
type Function struct{
    AbstractStatement
    Name Token
    Params []Token
    Body []AbstractStatement
}

func (_function Function) Accept(visitor StatementVisitor) interface{} {
    return visitor.VisitFunctionStatement(_function)
}

type StatementVisitor interface{
    VisitBlockStatement(statement Block) interface{}
    //visitClassStatement(statement Class)
    VisitExpressionStatement(statement Expression) interface{}
    VisitPrintStatement(statement Print) interface{}
    VisitVarStatement(statement Var) interface{}
    VisitIfStatement(statement If) interface{}
    VisitWhileStatement(statement While) interface{}
    VisitReturnStatement(statement Return) interface{}
    VisitFunctionStatement(statement Function) interface{}
}



