# Proto Client Code

这个仓库包含从 proto 文件自动生成的 Go 客户端代码。

## 生成来源
- 仓库: gp16-hub/gp-test
- 分支: master
- 提交: 12bb136ce70fa6ac1f10cab027a536e3b8880afe

## 最后更新
Fri Aug 29 19:11:27 CST 2025

## 目录结构
```
.
├── box/           # Box 服务客户端
├── external/      # External 服务客户端  
├── promotion/     # Promotion 服务客户端
├── tax/          # Tax 服务客户端
├── user/         # User 服务客户端
└── go.mod        # Go 模块文件
```

## 使用方法

```go
import (
    "github.com/gp16-hub/gp-test-proto-code/user/core"
    "github.com/gp16-hub/gp-test-proto-code/box/core"
    // ... 其他导入
)
```

## 自动生成

此代码通过 GitHub Actions 自动生成和更新。
当源仓库中的 proto 文件发生变化时，会自动触发重新生成。
