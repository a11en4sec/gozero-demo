# 1 主要配置项
```go
// A LogConf is a logging config.
type LogConf struct {
	ServiceName         string `json:",optional"`
	Mode                string `json:",default=console,options=[console,file,volume]"`
	Encoding            string `json:",default=json,options=[json,plain]"`
	TimeFormat          string `json:",optional"`
	Path                string `json:",default=logs"`
	Level               string `json:",default=info,options=[info,error,severe]"`
	Compress            bool   `json:",optional"`
	KeepDays            int    `json:",optional"`
	StackCooldownMillis int    `json:",default=100"`
}
```

# 2 配置文件中的一些设置
```yaml
Log:
  ServiceName: user-api
  Level: error
  Encoding: plain
  Mode: file
  Path: logs
```
- Encoding: 默认是json,编译kafka收集日志,但在查看堆栈信息的时候,不直观;可以改成plain
- Level: 
  - info级别 会打印详细的sql
  - info级别 会每隔1分钟,统计stat信息(cpu/内存/gc等)

# 3 打印日志的方法
```go
logx.Error()
logx.Info()
```

# 4 关闭stat的方法
在main函数中:
```go
logx.DisableStat()
```
