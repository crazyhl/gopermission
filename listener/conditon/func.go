package conditon

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crazyhl/gopermission/v1/parser"
	"strings"
)

// 获取条件的结果
func GetConditionResult(conditionStr string, modelData map[string]interface{}, userData map[string]interface{}) bool {
	conditionStr = strings.Trim(conditionStr, " ")
	if conditionStr == "" {
		// 如果条件为空，返回 true
		return true
	}
	// 初始化 listener
	listener := &ConditionListener{
		ModelData: modelData,
		UserData:  userData,
	}
	is := antlr.NewInputStream(conditionStr)
	// Create the Lexer
	lexer := parser.NewConditionLexer(is)

	steam := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewConditionParser(steam)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())

	return listener.GetResult()
}
