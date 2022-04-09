## sqlx执行源码的过程(FindOne举例)
```go
func (c cacheNode) doTake(ctx context.Context, v interface{}, key string,
	query func(v interface{}) error, cacheVal func(v interface{}) error) error {
	logger := logx.WithContext(ctx)
	val, fresh, err := c.barrier.DoEx(key, func() (interface{}, error) {

        // 1 先从缓存中按key拿数据
		if err := c.doGetCache(ctx, key, v); err != nil {
            // 获取的是占位符*,说明1分钟内有人查过,但数据库里没有记录,
            // 在cache中设置占位符,防止高并发的时候,都去查数据库,造成数据库压力
            
			if err == errPlaceholder {
				return nil, c.errNotFound
            // cache应该是挂了,快速失败,不要去查数据库,人工介入
			} else if err != c.errNotFound {
				// why we just return the error instead of query from db,
				// because we don't allow the disaster pass to the dbs.
				// fail fast, in case we bring down the dbs.
				return nil, err
			}

            // cache中没有,查数据库
			if err = query(v); err == c.errNotFound {
                // 数据库中查不到记录,在cache中设置占位符
				if err = c.setCacheWithNotFound(ctx, key); err != nil {
					logger.Error(err)
				}

				return nil, c.errNotFound
            // 查询数据库报错
			} else if err != nil {
				c.stat.IncrementDbFails()
				return nil, err
			}
            // 查询数据库正常,拿到数据,设置cache
			if err = cacheVal(v); err != nil {
				logger.Error(err)
			}
		}

		return jsonx.Marshal(v)
	})
	if err != nil {
		return err
	}
	if fresh {
		return nil
	}

	// got the result from previous ongoing query
	c.stat.IncrementTotal()
	c.stat.IncrementHit()

	return jsonx.Unmarshal(val.([]byte), v)
}
```