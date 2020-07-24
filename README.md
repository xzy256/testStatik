在需要使用的go文件中加入以下两行，其中src具体设置成对应的目录
```
//go:generate statik -src=../assets
//go:generate go fmt statik/statik.go

```
-----------------
项目结构
```
.
├── assets
│   └── 1.sh
├── cmd
│   └── main.go
├── go.mod
├── Gopkg.lock
├── Gopkg.toml
├── go.sum
└── pkg
    ├── hello.go
    └── statik
        └── statik.go
```

-------------------
build

```
go get github.com/rakyll/statik
export PATH=$PATH:$GOPATH/bin
go generate
go build -o ./bin/testStatik ./cmd

```
参考
https://blog.fatedier.com/2016/08/01/compile-assets-into-binary-file-with-statik-in-golang/
