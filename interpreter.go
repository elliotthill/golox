package main

import (
	"fmt"
	"reflect"
	"strconv"
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
func (interp *Interpreter) visitPrintStatement(stmt Print) interface{} {

	value := interp.evaluate(stmt.expression)
	fmt.Println(interp.stringify(value))
	return nil
}

func (interp *Interpreter) visitExpressionStatement(stmt Expression) interface{} {
	interp.evaluate(stmt.expression)
	return nil
}

type ReturnValue struct {
	value interface{}
}

func (interp *Interpreter) visitReturnStatement(stmt Return) interface{} {
	var value interface{} = nil

	if stmt.value != nil {
		value = interp.evaluate(stmt.value)
	}

	//throw new Return(value)
	panic(ReturnValue{value: value})
}

func (interp *Interpreter) visitWhileStatement(stmt While) interface{} {

	for interp.isTruthy(interp.evaluate(stmt.condition)) {
		interp.execute(stmt.body)
	}
	return nil
}

func (interp *Interpreter) visitBlockStatement(stmt Block) interface{} {

	env := NewEnvironment(interp.environment)
	interp.executeBlock(stmt.statements, env)
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

func (interp *Interpreter) visitVarStatement(stmt Var) interface{} {

	var value interface{} = nil
	if stmt.initializer != nil {
		value = interp.evaluate(stmt.initializer)
	}
	interp.environment.Define(stmt.name.lexeme, value)
	//interp.env[stmt.name.lexeme] = value
	return nil

}

func (interp *Interpreter) visitFunctionStatement(stmt Function) interface{} {

	function := RuntimeFunction{}
	function.declaration = stmt
    function.closure = interp.environment

	interp.environment.Define(stmt.name.lexeme, function)
	//interp.env[stmt.name.lexeme] = function
	return nil
}

func (interp *Interpreter) visitIfStatement(stmt If) interface{} {

	if interp.isTruthy(interp.evaluate(stmt.condition)) {
		interp.execute(stmt.thenBranch)
	} else if stmt.elseBranch != nil {
		interp.execute(stmt.elseBranch)
	}
	return nil
}

/*
* Expression
 */
func (interp *Interpreter) visitCallExpression(expr Call) interface{} {

	callee := interp.evaluate(expr.callee)
	var arguments []interface{}

	for _, arg := range expr.arguments {
		arguments = append(arguments, interp.evaluate(arg))
	}

	fn, ok := (callee).(callable)
	if !ok {
		panic("Can only call fucntions and classes")
	}

	if len(arguments) != fn.arity() {
		panic(fmt.Sprintf("Expected %d arguments but got %d.", fn.arity(), len(arguments)))
	}

	return fn.call(interp, arguments)
}


func (interp *Interpreter) visitFunctionExpression(expr FunctionExpression) interface{} {

	function := RuntimeFunction{}

    //We replace the expression with the function statement here
    functionStmt := Function{params: expr.params, body: expr.body}
	function.declaration = functionStmt;

    function.closure = interp.environment

	//interp.environment.Define(expr.name.lexeme, function)
	//interp.env[stmt.name.lexeme] = function
	return function
}


func (interp *Interpreter) evaluate(expr AbstractExpression) interface{} {
	return expr.Accept(interp)
}

func (interp *Interpreter) visitLiteralExpression(expr Literal) interface{} {
	return expr.value
}

func (interp *Interpreter) visitUnaryExpression(expr Unary) interface{} {
	right := interp.evaluate(expr.right)

	switch expr.operator.tokenType {
	case BANG:
		return !interp.isTruthy(right)
	case MINUS:
		return -interp.tryGetNumber(right)
	}

	//Unreachable
	return nil
}

func (interp *Interpreter) visitGroupingExpression(expr Grouping) interface{} {
	return interp.evaluate(expr.expression)
}

func (interp *Interpreter) visitBinaryExpression(expr Binary) interface{} {

	left := interp.evaluate(expr.left)
	right := interp.evaluate(expr.right)

	left_double, _ := left.(float64)
	right_double, _ := right.(float64)

	switch operator_type := expr.operator.tokenType; operator_type {
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

func (interp *Interpreter) visitVariableExpression(expr Variable) interface{} {

	val := interp.lookupVariable(expr.name.lexeme)

	if val == nil {
		panic("Undefined variable " + expr.name.lexeme + " ")
	}
	return val
}

func (interp *Interpreter) visitAssignExpression(expr Assign) interface{} {

	value := interp.evaluate(expr.value)
	//interp.env[expr.name.lexeme] = value
	interp.environment.Assign(expr.name.lexeme, value)

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
