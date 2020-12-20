package base_struct

type CustomModelCheck func(modelName string, paramValue string, condition string, user map[string]interface{}) bool
