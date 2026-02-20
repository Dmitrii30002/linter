package analyzer

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `Log linter checks log messages for common issues.

Rules:
  - Log messages must start with a lowercase letter
  - Log messages must be in English only
  - Log messages must not contain special characters or emoji
  - Log messages must not contain sensitive data (passwords, tokens, etc.)
`

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "linter",
		Doc:      Doc,
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)

		if logCall := extractLogCall(call); logCall != nil {
			checkMessage(pass, logCall)

			checkSensitiveData(pass, call, logCall)
		}
	})

	return nil, nil
}

func extractLogCall(call *ast.CallExpr) *LoggerCall {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil
	}

	pkgIdent, ok := sel.X.(*ast.Ident)
	if !ok {
		return nil
	}

	pkgName := pkgIdent.Name
	funcName := sel.Sel.Name

	if !isLogger(pkgName, funcName) {
		return nil
	}

	return &LoggerCall{
		Package: pkgName,
		Level:   funcName,
		Pos:     call.Pos(),
		Args:    call.Args,
		Message: extractStringFromExpr(call.Args[0]),
	}
}

func extractStringFromExpr(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			return strings.Trim(e.Value, "\"`")
		}
	case *ast.BinaryExpr:
		str := extractStringFromExpr(e.X)
		if str != "" {
			return str
		}
	}
	return ""
}

type LoggerCall struct {
	Package string
	Level   string
	Pos     token.Pos
	Args    []ast.Expr
	Message string
}

func isLogger(pkg, fn string) bool {
	switch pkg {
	case "log":
		switch fn {
		case "Print", "Printf", "Println",
			"Fatal", "Fatalf", "Fatalln",
			"Panic", "Panicf", "Panicln":
			return true
		}
	case "slog":
		switch fn {
		case "Info", "Infof", "Infoln",
			"Debug", "Debugf", "Debugln",
			"Warn", "Warnf", "Warnln",
			"Error", "Errorf", "Errorln":
			return true
		}
	case "zap":
		switch fn {
		case "Info", "Debug", "Warn", "Error":
			return true
		}
	}
	return false
}
