https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/integers

## 测试覆盖

    go test -cover

## 基准测试

    go test -bench=.

## 测试

### 辅助函数

### 表格驱动测试

## 注意事项

- 需要注意的是 reflect.DeepEqual 不是「类型安全」的，所以有时候会发生比较怪异的行为。
