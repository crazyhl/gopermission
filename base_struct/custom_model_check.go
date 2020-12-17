package base_struct

type CustomModelCheck func(modelName string, getModelFieldName string, paramValue string, condition string, user map[interface{}]interface{}) bool
