# otelbuild-demo

执行命令查看效果
```
# 使用otelbuild代替go build, 编译期插桩代码
.\otelbuild.exe --keepbuilddir --debug --verbose --rule=+otel_hook/rules.json main.go`
# 执行
.\main.exe
```
