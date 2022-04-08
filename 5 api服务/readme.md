# 1 基于api文件生成markdown文档
```
goctl api doc --dir ./
```

# 2 @server里面的group是便于文件分类的文件夹
是为了文件夹,不是路由的前缀

# 3 @server里面的prefix是路由的前缀
```
@server(
	// jwt:   Auth
	group: foo
	prefix: order/v1
```