package conditon

import (
	"fmt"
	"github.com/crazyhl/gopermission/v1/parser"
	"strings"
)

type ConditionListener struct {
	*parser.BaseConditionListener

	ModelData map[interface{}]interface{}
	UserData  map[interface{}]interface{}

	// 存放结果的栈
	stack []bool
}

// 返回结果
func (l *ConditionListener) GetResult() bool {
	if len(l.stack) < 1 {
		return false
	}

	return l.pop()
}

// 这个方法目前来看是不需要进行任何逻辑操作的
//func (l *ConditionListener) ExitParensExpression(c *parser.ParensExpressionContext) {
//	fmt.Println("-------------括号运算--------------")
//	fmt.Println(c.GetChildCount())
//	fmt.Println(c.GetText())
//	fmt.Println("-------------括号运算--------------")
//}

// 这个就是单纯的比较
func (l *ConditionListener) ExitCompare(c *parser.CompareContext) {
	fmt.Println("-------------比较运算--------------")
	fmt.Println(c.GetOp().GetText())
	fmt.Println(c.GetChildCount())
	fmt.Println(c.GetText())
	fmt.Println(c.GetLeft().GetText())
	fmt.Println(c.GetRight().GetText())
	leftValue := l.getValue(c.GetLeft().GetText())
	rightValue := l.getValue(c.GetRight().GetText())
	switch c.GetOp().GetTokenType() {
	case parser.ConditionLexerEqualOP:
		// ==
		fmt.Println("等于比较")
		leftStr := fmt.Sprint(leftValue)
		rightStr := fmt.Sprint(rightValue)
		l.push(leftStr == rightStr)
	case parser.ConditionLexerNotEqualOP:
		// !=
		fmt.Println("不等于比较")
		leftStr := fmt.Sprint(leftValue)
		rightStr := fmt.Sprint(rightValue)
		l.push(leftStr != rightStr)
	case parser.ConditionLexerLargerOp:
		// >
		fmt.Println("大于比较")
	case parser.ConditionLexerLargerEqualOp:
		// >=
		fmt.Println("大于等于比较")
	case parser.ConditionLexerLessOp:
		// <
		fmt.Println("小于比较")
	case parser.ConditionLexerLessEqualOp:
		// <=
		fmt.Println("小于等于比较")
	case parser.ConditionLexerInOP:
		fmt.Println("in比较")
	}
	fmt.Println("-------------比较运算--------------")
}

// 这个是两个条件的或运算
func (l *ConditionListener) ExitOrCompare(c *parser.OrCompareContext) {
	fmt.Println("-------------Or 比较运算--------------")
	fmt.Println(c.GetChildCount())
	fmt.Println(c.GetText())
	fmt.Println(c.GetLeft().GetText())
	fmt.Println(c.GetRight().GetText())
	fmt.Println("-------------Or 比较运算--------------")
}

func (l *ConditionListener) ExitAndCompare(c *parser.AndCompareContext) {
	fmt.Println("-------------And 比较运算--------------")
	fmt.Println(c.GetChildCount())
	fmt.Println(c.GetText())
	fmt.Println(c.GetLeft().GetText())
	fmt.Println(c.GetRight().GetText())
	fmt.Println("-------------And 比较运算--------------")
}

// ---------------------- private function ------------------------

func (l *ConditionListener) push(i bool) {
	l.stack = append(l.stack, i)
}

func (l *ConditionListener) pop() bool {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *ConditionListener) getValue(paramStr string) interface{} {
	paramsArr := strings.Split(paramStr, ".")
	paramsLen := len(paramsArr)
	if paramsLen <= 1 {
		return nil
	}
	modelType := paramsArr[0]
	valueData := make(map[interface{}]interface{})
	switch modelType {
	case "model":
		// 从模型获取
		valueData = l.ModelData
	case "user":
		// 从用户获取
		valueData = l.UserData
	default:
		return ""
	}

	for idx := 1; idx < paramsLen; idx++ {
		param := paramsArr[idx]
		if idx == paramsLen-1 {
			return valueData[param]
		}
		if valueData[param] == nil {
			return nil
		}
		valueData = valueData[param].(map[interface{}]interface{})
	}

	return nil
}
