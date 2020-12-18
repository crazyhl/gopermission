# gopermission
基于 自己遇到场景抽象的 角色权限 判定组件，支持绑定模型，进行更进一步判定

## 说明
目前是仅仅做到满足自己的需求

数据库仅支持 `gorm` 和 `mysql`
数据验证用的的 `validator`
路由匹配的方案采用 `fiber` 的

这个库只考虑权限判定相关的东西，用户跟角色权限的绑定，交给使用方来自己决定

目前计划支持多个条件有一个条件允许就可以访问

条件解析用 `antlr` 实现，具体请看 `Condition.g4` 文件
条件的设定条件 `model` 指代关联模型，`user` 代表登录用户
条件可以写成 `model.category.id in user.categories || model.uid == user.id`


## 图示
![图示](./flow_chart.png)

