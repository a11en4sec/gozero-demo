## 1 api文件定义

### 1.1 文件目录
```
api
 - order.api
 - user.api
```
### 1.2 注意事项
1. api一个模块一个文件，文件名不必相同
2. 每个api的定义文件中，service名字必须相同
3. 每个api的定义文件中，type中的req和resp不能相同，否则会在type中被覆盖
4. 定义的多个api文件中，需要有一个主api，并import其他api (import "order.api")
```
goctl api go -api user.api -dir ../ --style=gozero
```


