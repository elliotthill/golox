package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Interpreter struct {
	ExpressionVisitor
	StatementVisitor
	statements []AbstractStatement
}

func NewInterpreter(statements []AbstractStatement) {
	interp := new(Interpreter)
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

func (interp *Interpreter) visitPrintStatement(stmt Print) interface{} {

	value := interp.evaluate(stmt.expression)
	fmt.Println(interp.stringify(value))
	return nil
}

func (interp *Interpreter) evaluate(expr AbstractExpression) interface{} {
	return expr.Accept(interp)
}

func (interp *Interpreter) visitLiteralExpression(expr Literal) interface{} {
	return expr.value
}

func (interp *Interpreter) visitBinaryExpression(expr Binary) interface{} {

	left := interp.evaluate(expr.left)
	right := interp.evaluate(expr.right)

	left_double, _ := strconv.ParseFloat(fmt.Sprintf("%v", left), 64)
	right_double, _ := strconv.ParseFloat(fmt.Sprintf("%v", right), 64)

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

func (interp *Interpreter) isEqual(a interface{}, b interface{}) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return reflect.DeepEqual(a, b)
}

func (interp *Interpreter) stringify(thing interface{}) string {

	if thing == nil {
		return "nil"
	}

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
