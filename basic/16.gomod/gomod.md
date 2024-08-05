# GoMod

## 什么是GoMod

Go modules 是 Go 语言的依赖解决方案，发布于 Go1.11，成长于 Go1.12，丰富于 Go1.13，正式于 Go1.14 推荐在生产上使用。

## GoMod解决的问题

1. Go 语言长久以来的依赖管理问题。
2. “淘汰”现有的 GOPATH 的使用模式。
3. 统一社区中的其它的依赖管理工具（提供迁移功能）。

## GoMod常用指令

| 命令              | 作用                   |
|-----------------|----------------------|
| go mod init     | 生成 go.mod 文件         |
| go mod download | 下载 go.mod 文件中指明的所有依赖 |
| go mod tidy     | 整理现有的依赖              |
| go mod graph    | 查看现有的依赖结构            |
| go mod edit     | 编辑 go.mod 文件         |
| go mod vendor   | 导出项目所有的依赖到vendor目录   |
| go mod verify   | 校验一个模块是否被篡改过         |
| go mod why      | 查看为什么需要依赖某模块         |

## GoMod有关环境变量

```markdown
go env
GO111MODULE="auto"
GOPROXY="https://proxy.golang.org,direct"
GONOPROXY=""
GOSUMDB="sum.golang.org"
GONOSUMDB=""
GOPRIVATE=""
...
```

1. **GO111MODULE**是是否开启Go modules，相关的参数有auto、on、off，推荐使用on，因为将会是后续版本的默认值
2. **GOPROXY**是Go modules的镜像代理站点，可以设置为[七牛云](https://goproxy.cn,direct)
   和[阿里云](https://mirrors.aliyun.com/goproxy/)的

> 而在刚刚设置的值中，我们可以发现值列表中有 “direct” 标识，它又有什么作用呢？\
> 实际上 “direct” 是一个特殊指示符，用于指示 Go 回源到模块版本的源地址去抓取（比如 GitHub 等），场景如下：当值列表中上一个 Go
> 模块代理返回 404 或 410 错误时，Go 自动尝试列表中的下一个，遇见 “direct” 时回源，也就是回到源地址去抓取，而遇见 EOF
> 时终止并抛出类似
> “invalid version: unknown revision...” 的错误。

3. **GOSUMDB**用于校验模块数据是否被更改。它的值是一个 Go checksum database，用于在拉取模块版本时（无论是从源站拉取还是通过
   Go module proxy 拉取）保证拉取到的模块版本数据未经过篡改，若发现不一致，也就是可能存在篡改，将会立即中止。
4. **GONOPROXY/GONOSUMDB/GOPRIVATE**，这三个环境变量都是用在当前项目依赖了私有模块，例如像是你公司的私有 git 仓库，又或是 github 中的私有库，都是属于私有模块，都是要进行设置的，否则会拉取失败。

