## Go 语言测试规范

- Go 测试与其测试的代码在同一目录
- 测试的文件的名称后缀_test
- 如果要测试的文件名为 strings.go ，则测试它的文件将名为 strings_test.go ，并位于文件 strings.go 所在的目录中

```
└── testCase
    ├── cal.go
    └── cal_test.go
```

-  测试为名称以单词 Test 打头的函数

```shell
func TestAddUpper(t *testing.T) {
  
}
```

## 基准测试

```shell
go test --bench=.
```

## 测试覆盖率

```shell
go test --cover xxx.go
```