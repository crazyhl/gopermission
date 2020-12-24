# gopermission
[![Go Report Card](https://goreportcard.com/badge/github.com/crazyhl/gopermission)](https://goreportcard.com/report/github.com/crazyhl/gopermission)

角色权限判定模块，支持绑定模型，进行更进一步判定


## 致谢
* fiber
* gorm
* antlr



## 安装
```shell
go get -u github.com/crazyhl/gopermission
```

## 使用说明
引入后，在自己的角色 `model` 中关联 `rule` 或者 `permission`，可以根据自己的需求决定关联两个或者其中一个。


## 特殊说明
目前是仅仅做到满足自己的需求

数据库仅支持 `gorm` 和 `mysql`
数据验证用的的 `validator`
路由匹配的方案采用 `fiber` 的

这个库只考虑权限判定相关的东西，用户跟角色权限的绑定，交给使用方来自己决定

目前计划支持多个条件情况下有一个条件允许就可以访问

条件解析用 `antlr` 实现，具体请看 `Condition.g4` 文件
条件的设定条件 `model` 指代关联模型，`user` 代表登录用户
条件可以写成 `model.category.id in user.categories || model.uid == user.id`

条件支持：
* ==
* !=
* &gt;
* &gt;=
* &lt;
* &lt;=
* in
* ||
* &&
* ()

## 注意
1. 权限相关的绑定肯定是在开发期间做的，所以一定要做好充分测试。
2. 这边只负责传入的数据是否正确，所以，如果获取数据时候获取不到，返回结果肯定是 `false` 的
3. 书写条件的时候一定要善用括号，否则，大家看不懂，程序判断错误跟你预期不一致也是谁都不想看到的
4. == != 比较用的是转换为字符串
5. &gt; &gt;= &lt; &lt;= 比较会转换为数字
6. in 比较传入前面是值后面是对象数据，比较的字段会是最后一段的字段名称比如 `model.category.id in user.categories` 会比较model 的 category 的 id 的值是否在传入列表里面的 id 字段


## 图示
![图示](./flow_chart.png)

