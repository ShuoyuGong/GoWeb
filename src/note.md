# Go WEB 开发笔记

## 1.模板

- 嵌套输出

  ```Go
  {{with .UserInfo}}
    {{.Name}};{{.Age}};{{.Sex}}
  {{end}}
  ```

- 数组循环输出

  ```Go
  {{range .Nums}}
    {{.}}
  {{end}}
  ```

- 模板变量的使用

  ```Go
  // Controller
    this.Data["TplVal"] = "hey,gsy."
  // View
    {{$tpl_var := .TplVal}}
    {{$tpl_var}}
  ```

- HTml 字符串直接显示，beego 内置模板函数

  ```Go
  // Controller
    this.Data["Html"] = "<div>hello,web!</div>"
  // View
    {{str2html .Html}}
  ```

- 模板 Html 的嵌套

  ```Go
  // View
  // 定义内容
  {{define "Test"}}
    <div>
      this is test template.
    </div>
  {{end}}

  //使用内容
  {{template "Test"}}
  ```

## 2. 数据库的建立于连接

- 注册数据库

  ```Go
  // 参数1        数据库的别名，用来在ORM中切换数据库使用
  // 参数2        driverName
  // 参数3        对应的链接字符串
  orm.RegisterDataBase("default", "mysql", "root:root@/db_Name?charset=utf8")

  ```

`
