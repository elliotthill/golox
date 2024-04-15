package interpreter

import (
	"fmt"
	"reflect"
	"strconv"
    ."github.com/elliotthill/golox/language"
)

type Interpreter struct {
	ExpressionVisitor
	StatementVisitor
	statements  []AbstractStatement
	environment *Environment
	globals     *Environment
}

func NewInterpreter() *Interpreter {
	interp := new(Interpreter)

	//interp.environment = NewEnv(nil)
	interp.globals = NewEnvironment(nil)
	interp.environment = interp.globals
	fmt.Println("Initialized interpreter")
	return interp
}

func (interp *Interpreter) SetStatements(statements []AbstractStatement) {
   interp.statements = statements
}

func (interp *Interpreter) Interpret() {

	for _, stmt := range interp.statements {

		interp.execute(stmt)
	}

}

func (interp *Interpreter) execute(stmt AbstractStatement) {

	stmt.Accept(interp)
}

/*
* Statements
 */
func (interp *Interpreter) VisitPrintStatement(stmt Print) interface{} {

	value := interp.evaluate(stmt.Expression)
	fmt.Println(interp.stringify(value))
	return nil
}

func (interp *Interpreter) VisitExpressionStatement(stmt Expression) interface{} {
	interp.evaluate(stmt.Expression)
	return nil
}

type ReturnValue struct {
	value interface{}
}

func (interp *Interpreter) VisitReturnStatement(stmt Return) interface{} {
	var value interface{} = nil

	if stmt.Value != nil {
		value = interp.evaluate(stmt.Value)
	}

	//throw new Return(value)
	panic(ReturnValue{value: value})
}

func (interp *Interpreter) VisitWhileStatement(stmt While) interface{} {

	for interp.isTruthy(interp.evaluate(stmt.Condition)) {
		interp.execute(stmt.Body)
	}
	return nil
}

func (interp *Interpreter) VisitBlockStatement(stmt Block) interface{} {

	env := NewEnvironment(interp.environment)
	interp.executeBlock(stmt.Statements, env)
	return nil
}

func (interp *Interpreter) executeBlock(statements []AbstractStatement, env *Environment) {

	previous := interp.environment
	defer func() {
		interp.environment = previous
	}()

	interp.environment = env
	for _, statement := range statements {

		interp.execute(statement)
	}
}

func (interp *Interpreter) VisitVarStatement(stmt Var) interface{} {

	var value interface{} = nil
	if stmt.Initializer != nil {
		value = interp.evaluate(stmt.Initializer)
	}
	interp.environment.Define(stmt.Name.Lexeme, value)
	//interp.env[stmt.name.lexeme] = value
	return nil

}

func (interp *Interpreter) VisitFunctionStatement(stmt Function) interface{} {

	function := RuntimeFunction{}
	function.declaration = stmt
    function.closure = interp.environment

	interp.environment.Define(stmt.Name.Lexeme, function)
	//interp.env[stmt.name.lexeme] = function
	return nil
}

func (interp *Interpreter) VisitIfStatement(stmt If) interface{} {

	if interp.isTruthy(interp.evaluate(stmt.Condition)) {
		interp.execute(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		interp.execute(stmt.ElseBranch)
	}
	return nil
}

/*
* Expression
 */
func (interp *Interpreter) VisitCallExpression(expr Call) interface{} {

	callee := interp.evaluate(expr.Callee)
	var arguments []interface{}

	for _, arg := range expr.Arguments {
		arguments = append(arguments, interp.evaluate(arg))
	}

	fn, ok := (callee).(callable)
	if !ok {
		panic("Can only call functions and classes")
	}

	if len(arguments) != fn.arity() {
		panic(fmt.Sprintf("Expected %d arguments but got %d.", fn.arity(), len(arguments)))
	}

	return fn.call(interp, arguments)
}


func (interp *Interpreter) VisitFunctionExpression(expr FunctionExpression) interface{} {

	function := RuntimeFunction{}

    //We replace the expression with the function statement here
    functionStmt := Function{Params: expr.Params, Body: expr.Body}
	function.declaration = functionStmt;

    function.closure = interp.environment

    //Anonymous functions aren't defined
	//interp.environment.Define(expr.name.lexeme, function)
	return function
}


func (interp *Interpreter) evaluate(expr AbstractExpression) interface{} {
	return expr.Accept(interp)
}

func (interp *Interpreter) VisitLiteralExpression(expr Literal) interface{} {
	return expr.Value
}

func (interp *Interpreter) VisitUnaryExpression(expr Unary) interface{} {
	right := interp.evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case BANG:
		return !interp.isTruthy(right)
	case MINUS:
		return -interp.tryGetNumber(right)
	}

	//Unreachable
	return nil
}

func (interp *Interpreter) VisitGroupingExpression(expr Grouping) interface{} {
	return interp.evaluate(expr.Expression)
}

func (interp *Interpreter) VisitBinaryExpression(expr Binary) interface{} {

	left := interp.evaluate(expr.Left)
	right := interp.evaluate(expr.Right)

	left_double, _ := left.(float64)
	right_double, _ := right.(float64)

	switch operator_type := expr.Operator.TokenType; operator_type {
	case GREATER:
		return left_double > right_double
	case GREATER_EQUAL:
		return left_double >= right_double
	case LESS:
		return left_double < right_double
	case LESS_EQUAL:
		return left_double <= right_double
	case MINUS:
		return left_double - right_double
	case SLASH:
		return left_double / right_double
	case STAR:
		return left_double * right_double
	case PLUS:
		//add string concatenation here
		return left_double + right_double
	case BANG_EQUAL:
		return !interp.isEqual(left, right)
	case EQUAL_EQUAL:
		return interp.isEqual(left, right)

	}
	//Unreachable...in theory
	fmt.Println("Error")
	return nil
}

func (interp *Interpreter) VisitVariableExpression(expr Variable) interface{} {

	val := interp.lookupVariable(expr.Name.Lexeme)

	if val == nil {
		panic("Undefined variable " + expr.Name.Lexeme + " ")
	}
	return val
}

func (interp *Interpreter) VisitAssignExpression(expr Assign) interface{} {

	value := interp.evaluate(expr.Value)
	//interp.env[expr.name.lexeme] = value
	interp.environment.Assign(expr.Name.Lexeme, value)

	return value
}

func (interp *Interpreter) lookupVariable(name string) interface{} {

	value := interp.environment.Get(name)

	if value == nil {
		value = interp.globals.Get(name)
	}

	return value
}

func (interp *Interpreter) isEqual(a interface{}, b interface{}) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return reflect.DeepEqual(a, b)
}

func (interp *Interpreter) isTruthy(obj interface{}) bool {
	if obj == nil {
		return false
	}

	switch v := obj.(type) {
	case bool:
		return v
	case string:
		return len(v) > 0
	default:
		return true
	}

}

func (interp *Interpreter) tryGetNumber(operand interface{}) float64 {

	switch v := operand.(type) {
	case string:
		try_float, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic("Operand must be a number")
		}
		return try_float

	default:
		panic("Operand must be a number")
	}
}

func (interp *Interpreter) stringify(thing interface{}) string {

	if thing == nil {
		return "nil"
	}
	return fmt.Sprint(thing)

	switch v := thing.(type) {
	case bool:
		return strconv.FormatBool(v)
	case int:
		return fmt.Sprintf("Integer: %v", v)
	case float64:
		return fmt.Sprintf("Float64: %v", v)
	case string:
		return v
	default:
		return "<glox object>"
	}

}
