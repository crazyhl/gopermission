package conditon

import "strings"

// 获取条件的结果
func GetConditionResult(conditionStr string, modelData map[interface{}]interface{}, userData map[interface{}]interface{}) bool {
	conditionStr = strings.Trim(conditionStr, " ")
	if conditionStr == "" {
		// 如果条件为空，返回 true
		return true
	}
	return false
}
