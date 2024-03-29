# TDD pratice

![build-status](https://travis-ci.com/26huitailang/go_tdd.svg?branch=master)

## 参考

- [learn-go-with-tests](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/integers)

## 测试覆盖

    go test -cover

输出结果：

    go test -v -coverprofile cover.out [FILENAME(可选)]
    go tool cover -html=cover.out -o cover.html

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

mocking：

- 没有对代码中重要的区域进行 mock 将会导致难以测试。在我们的例子中，我们不能测试我们的代码在每个打印之间暂停，但是还有无数其他的例子。调用一个 可能 失败的服务？想要在一个特定的状态测试您的系统？在不使用 mocking 的情况下测试这些场景是非常困难的。
- 如果没有 mock，你可能需要设置数据库和其他第三方的东西来测试简单的业务规则。你可能会进行缓慢的测试，从而导致 缓慢的反馈循环。
- 当不得不启用一个数据库或者 webservice 去测试某个功能时，由于这种服务的不可靠性，你将会得到的是一个 脆弱的测试。

测试成都：

有时候很难知道到底要测试到 什么级别，但是这里有一些我试图遵循的思维过程和规则。

- 重构的定义是代码更改，但行为保持不变。 如果您已经决定在理论上进行一些重构，那么你应该能够在没有任何测试更改的情况下进行提交。所以，在写测试的时候问问自己。
    - 我是在测试我想要的行为还是实现细节？
    - 如果我要重构这段代码，我需要对测试做很多修改吗？
- 虽然 Go 允许你测试私有函数，但我将避免它作为私有函数与实现有关。
- 我觉得如果一个测试 超过 3 个模拟，那么它就是警告 —— 是时候重新考虑设计。
- 小心使用监视器。监视器让你看到你正在编写的算法的内部细节，这是非常有用的，但是这意味着你的测试代码和实现之间的耦合更紧密。如果你要监视这些细节，请确保你真的在乎这些细节。

## 不要过早优化

遵循三个步骤:
[使它运作，使它正确，使它快速](http://wiki.c2.com/?MakeItWorkMakeItRightMakeItFast)

[过早的优化是万恶之源 —— Donald Knuth](http://wiki.c2.com/?PrematureOptimization)

## interface

如果你想实现函数的多态性，请考虑是否可以围绕接口（不是 interface 类型，这里容易让人困惑）设计它，以便用户可以用多种类型来调用你的函数，这些类型实现了函数工作所需要的任何方法。

## 注意事项

- 需要注意的是 reflect.DeepEqual 不是「类型安全」的，所以有时候会发生比较怪异的行为。
- **不要过度测试，特别是在掌握了mocking之后**，根据系统的工作方式来安排测试，而不是它做了什么。
