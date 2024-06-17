package typechecker

import (
	"fmt"
	"walrus/frontend/ast"
)

func Evaluate(astNode ast.Node, env *Environment) RuntimeValue {
	switch node := astNode.(type) {
	case ast.NumericLiteral:
		return IntegerValue{
			Type:  "int",
			Value: int(node.Value),
			Size:  64,
		}
	case ast.StringLiteral:
		return StringValue{
			Type:  "string",
			Value: node.Value,
		}
	case ast.BooleanLiteral:
		return BooleanValue{
			Type:  "bool",
			Value: node.Value,
		}
	case ast.NullLiteral:
		return NullValue{
			Type: "null",
		}
	case ast.ProgramStmt:
		return EvaluateProgramBlock(node, env)
	case ast.VariableDclStml:
		return EvaluateVariableDeclarationStmt(node, env)
	case ast.AssignmentExpr:
		return EvaluateAssignmentExpr(node, env)
	case ast.UnaryExpr:
		return EvaluateUnaryExpression(node, env)
	case ast.BinaryExpr:
		return EvaluateBinaryExpr(node, env)
	case ast.IdentifierExpr:
		return EvaluateIdenitifierExpr(node, env)
	case ast.IfStmt:
		return EvaluateControlFlowStmt(node, env)
	case ast.FunctionDeclStmt:
		return EvaluateFunctionDeclarationStmt(node, env)
	default:
		panic(fmt.Sprintf("This ast node is not implemented yet: %v", node))
	}
}