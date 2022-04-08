> 环境: goctl<=1.3.3 
> 数据库: 没有任何唯一索引
> 生成model代码的时候,使用了-c cache参数


## bug复现的过程
1. 数据库里面没有记录
2. 使用findOne查询,返回没有记录,并且会同时在redis cache中,设置: *,过期时间为一分钟的缓存
3. 通过insert插入记录,但代码里面没有删除cache的操作
4. 通过findOne查询,仍然查不到记录

## 原因
insert的时候没有清缓存，其他如findOne、update用到了缓存

## 解决办法
1. 使用goctl template init 生成模板,并修改model下的insert逻辑
2. 升级goctl至新版本

