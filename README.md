# 项目导航
- 原项目地址: [AMllChat-抹茶](https://github.com/MallChat-Go/MallChat-Go/blob/main/README.md)
- 使用Go语言重构的项目
- 此项目仅供学习交流使用，请勿用于商业用途。


# api 文件快速生成go开发环境
```go

    -- 生成api文件
     goctl api go  -api .\app\api\mallchatgo.api --dir .\app\
    
    -- 生成model文件
    goctl model mysql datasource -url="root:password@tcp(127.0.0.1)/mallchatgo" -table="*"  -dir="./app/internal/model"

    -- 生成api文档,需要工具(go install github.com/zeromicro/goctl-swagger@latest)
    goctl api plugin -p goctl-swagger="swagger -filename mallchatgo.json" --api .\app\api\mallchatgo.api --dir .

    -- 启动项目
    go run .\app\mallchatgo.go -f .\app\etc\mallchatgo.yaml
```