package conditon

import (
	"fmt"
	"github.com/crazyhl/gopermission/v1/parser"
	"reflect"
	"strconv"
	"strings"
)

type ConditionListener struct {
	*parser.BaseConditionListener

	ModelData map[string]interface{}
	UserData  map[string]interface{}

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
	// 如果两个选项有一个为空就返回空
	if leftValue == nil || rightValue == nil {
		l.push(false)
		return
	}
	switch c.GetOp().GetTokenType() {
	case parser.ConditionLexerEqualOP:
		// ==
		leftStr := fmt.Sprint(leftValue)
		rightStr := fmt.Sprint(rightValue)
		l.push(leftStr == rightStr)
	case parser.ConditionLexerNotEqualOP:
		// !=
		leftStr := fmt.Sprint(leftValue)
		rightStr := fmt.Sprint(rightValue)
		l.push(leftStr != rightStr)
	case parser.ConditionLexerLargerOp:
		// >
		leftNumber, leftErr := strconv.Atoi(fmt.Sprint(leftValue))
		rightNumber, rightErr := strconv.Atoi(fmt.Sprint(rightValue))
		if leftErr != nil || rightErr != nil {
			l.push(false)
		} else {
			l.push(leftNumber > rightNumber)
		}
	case parser.ConditionLexerLargerEqualOp:
		// >=
		leftNumber, leftErr := strconv.Atoi(fmt.Sprint(leftValue))
		rightNumber, rightErr := strconv.Atoi(fmt.Sprint(rightValue))
		if leftErr != nil || rightErr != nil {
			l.push(false)
		} else {
			l.push(leftNumber >= rightNumber)
		}
	case parser.ConditionLexerLessOp:
		// <
		leftNumber, leftErr := strconv.Atoi(fmt.Sprint(leftValue))
		rightNumber, rightErr := strconv.Atoi(fmt.Sprint(rightValue))
		if leftErr != nil || rightErr != nil {
			l.push(false)
		} else {
			l.push(leftNumber < rightNumber)
		}
	case parser.ConditionLexerLessEqualOp:
		// <=
		leftNumber, leftErr := strconv.Atoi(fmt.Sprint(leftValue))
		rightNumber, rightErr := strconv.Atoi(fmt.Sprint(rightValue))
		if leftErr != nil || rightErr != nil {
			l.push(false)
		} else {
			l.push(leftNumber <= rightNumber)
		}
	case parser.ConditionLexerInOP:
		// in
		fmt.Println("in比较")
		paramArr := l.getParamArr(c.GetLeft().GetText())
		paramsLen := len(paramArr)
		if paramsLen <= 1 {
			l.push(false)
		} else {
			leftStr := fmt.Sprint(leftValue)
			checkFieldName := paramArr[paramsLen-1]
			fmt.Println(leftStr)
			fmt.Println(checkFieldName)
			fmt.Println(reflect.TypeOf(rightValue))

			switch rightValue.(type) {
			case []interface{}:
				rightValueSlice := rightValue.([]interface{})
				for _, v := range rightValueSlice {
					valueMap := v.(map[string]interface{})
					if leftStr == fmt.Sprint(valueMap[checkFieldName]) {
						l.push(true)
						return
					} else {
						l.push(false)
					}
				}
			default:
				l.push(false)
			}
		}
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

func (l *ConditionListener) getParamArr(paramStr string) []string {
	return strings.Split(paramStr, ".")
}

func (l *ConditionListener) getValue(paramStr string) interface{} {
	paramsArr := l.getParamArr(paramStr)
	paramsLen := len(paramsArr)
	if paramsLen < 1 {
		return nil
	}
	if paramsLen == 1 {
		// 建判断一下是不是 int
		intValue, err := strconv.Atoi(paramStr)
		if err != nil {
			return nil
		}
		return intValue
	}
	modelType := paramsArr[0]
	valueData := make(map[string]interface{})
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
		valueData = valueData[param].(map[string]interface{})
	}

	return nil
}
