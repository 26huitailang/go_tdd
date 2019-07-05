https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/integers

## 测试覆盖

    go test -cover

## 基准测试

    go test -bench=.

## 测试

### 辅助函数

### 表格驱动测试

### errcheck

有一种情况我们还没有测试过。要找到它，在一个终端中运行以下命令来安装 errcheck，这是许多可用的 linters（代码检测工具）之一。

    go get -u github.com/kisielk/errcheck

然后，在你的代码目录中运行 errcheck .。

### 依赖注入

dependency injection

- 你不需要一个框架
- 它不会过度复杂化你的设计
- 它易于测试
- 它能让你编写优秀和通用的函数

## 注意事项

- 需要注意的是 reflect.DeepEqual 不是「类型安全」的，所以有时候会发生比较怪异的行为。
