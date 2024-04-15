package main
/*
import (
	"fmt"
)

type DebugPrint struct{
    ExpressionVisitor
    StatementVisitor
}

func (a DebugPrint) print(expr Expression) string {
	return expr.Accept(a).(string)
}

func (a DebugPrint) VisitAssignExpr(expr Assign) interface{} {
	return a.parenthesize("var "+expr.name.lexeme+ " = ")
}

func (a DebugPrint) VisitVariableExpr(expr Variable) interface{} {
	return expr.name.lexeme
}

func (a DebugPrint) VisitTernaryExpr(expr Ternary) interface{} {
	return a.parenthesize(expr.leftOperator.String() + expr.rightOperator.String(),
    expr.left, expr.middle, expr.right)
}

func (a DebugPrint) VisitBinaryExpr(expr Binary) interface{} {
	return a.parenthesize(expr.operator.String(), expr.left, expr.right)
}

func (a DebugPrint) VisitGroupingExpr(expr Grouping) interface{} {
	return a.parenthesize("group", expr.Expression)
}

func (a DebugPrint) VisitLiteralExpr(expr Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprint(expr.Value)
}

func (a DebugPrint) VisitUnaryExpr(expr Unary) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a DebugPrint) VisitLogicalExpr(expr Logical) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a DebugPrint) parenthesize(name string, exprs ...Expression) string {
	var str string

	str += "(" + name
	for _, expr := range exprs {
		str += " " + a.print(expr)
	}
	str += ")"

	return str
}
*/
