package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var (
	sensitiveKeywords = []string{
		"password", "passwd", "pwd",
		"secret", "key", "api", "key",
		"token", "jwt", "cvv", "security",
		"auth", "authorization",
		"certificate", "cert",
	}

	allowedСhars = map[rune]struct{}{
		' ': struct{}{},
	}
)

func checkMessage(pass *analysis.Pass, logCall *LoggerCall) {
	msg := logCall.Message

	//1. log messages must start with a lowercase letter
	checkFirstLetter(pass, logCall, msg)

	//2. log message should be in English only
	checkEnglishOnly(pass, logCall, msg)

	//3. log messages should not contain special characters or emojis
	checkSpecialChars(pass, logCall, msg)
}

func checkFirstLetter(pass *analysis.Pass, logCall *LoggerCall, msg string) {
	if msg == "" {
		return
	}

	first := rune(msg[0])
	if !unicode.IsLower(first) {
		pass.Reportf(logCall.Pos,
			"%s.%s: log messages must start with a lowercase letter (message: %s)",
			logCall.Package, logCall.Level, msg)
	}
}

func checkEnglishOnly(pass *analysis.Pass, logCall *LoggerCall, msg string) {
	for _, r := range msg {
		if unicode.IsLetter(r) && !(unicode.Is(unicode.Latin, r)) {
			pass.Reportf(logCall.Pos,
				"%s.%s: log message should be in English only (message: %s)",
				logCall.Package, logCall.Level, string(r))
			break
		}
	}
}

func checkSpecialChars(pass *analysis.Pass, logCall *LoggerCall, msg string) {
	for _, r := range msg {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			if _, ok := allowedСhars[r]; !ok {
				pass.Reportf(logCall.Pos,
					"%s.%s: log messages should not contain special characters or emojis (message: %s)",
					logCall.Package, logCall.Level, msg)
				break
			}
		}
	}
}

// 4. log message should not contain sensitive data
func checkSensitiveData(pass *analysis.Pass, call *ast.CallExpr, logCall *LoggerCall) {
	msg := strings.ToLower(logCall.Message)
	hasKeyword := false

	for _, kw := range sensitiveKeywords {
		if strings.Contains(msg, kw) {
			hasKeyword = true
			break
		}
	}

	if !hasKeyword {
		return
	}

	if len(call.Args) > 1 {
		pass.Reportf(logCall.Pos,
			"%s.%s: log message should not contain sensitive data (message: %s)",
			logCall.Package, logCall.Level, msg)
		return
	}

	if binary, ok := call.Args[0].(*ast.BinaryExpr); ok && binary.Op == token.ADD {
		pass.Reportf(logCall.Pos,
			"%s.%s: log message should not contain sensitive data (message: %s)",
			logCall.Package, logCall.Level, msg)
		return
	}
}
