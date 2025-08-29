#!/bin/bash

# 生成 proto 客户端代码的脚本

set -e

OUTPUT_DIR="generated"
REPO_NAME="github.com/gp16-hub/gp-test-proto-code"

echo "🚀 开始生成 Proto 客户端代码..."

# 清理并创建输出目录
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# 查找所有 proto 文件
proto_files=$(find . -name "*.proto" -type f)

if [ -z "$proto_files" ]; then
    echo "❌ 没有找到任何 proto 文件"
    exit 1
fi

echo "📁 找到以下 proto 文件:"
echo "$proto_files"
echo ""

# 为每个 proto 文件生成代码
for proto_file in $proto_files; do
    echo "🔄 处理: $proto_file"
    
    # 获取相对路径的目录
    proto_dir=$(dirname "$proto_file")
    proto_name=$(basename "$proto_file" .proto)
    
    # 创建输出目录
    output_dir="$OUTPUT_DIR/${proto_dir#./}"
    mkdir -p "$output_dir"
    
    # 使用 goctl 生成 RPC 客户端代码
    if goctl rpc protoc "$proto_file" \
        --go_out="$output_dir" \
        --go-grpc_out="$output_dir" \
        --zrpc_out="$output_dir" \
        --style=goZero; then
        echo "✅ 成功生成: $proto_file"
    else
        echo "⚠️  警告: 生成失败 $proto_file"
    fi
    echo ""
done

# 生成 go.mod 文件
echo "📝 生成 go.mod 文件..."
cat > "$OUTPUT_DIR/go.mod" << EOF
module $REPO_NAME

go 1.21

require (
    github.com/zeromicro/go-zero v1.6.0
    google.golang.org/grpc v1.60.0
    google.golang.org/protobuf v1.32.0
)
EOF

# 生成 README
echo "📝 生成 README.md 文件..."
cat > "$OUTPUT_DIR/README.md" << EOF
# Proto Client Code

这个仓库包含从 proto 文件自动生成的 Go 客户端代码。

## 生成来源
- 仓库: gp16-hub/gp-test
- 分支: $(git rev-parse --abbrev-ref HEAD)
- 提交: $(git rev-parse HEAD)

## 最后更新
$(date)

## 目录结构
\`\`\`
.
├── box/           # Box 服务客户端
├── external/      # External 服务客户端  
├── promotion/     # Promotion 服务客户端
├── tax/          # Tax 服务客户端
├── user/         # User 服务客户端
└── go.mod        # Go 模块文件
\`\`\`

## 使用方法

\`\`\`go
import (
    "$REPO_NAME/user/core"
    "$REPO_NAME/box/core"
    // ... 其他导入
)
\`\`\`

## 自动生成

此代码通过 GitHub Actions 自动生成和更新。
当源仓库中的 proto 文件发生变化时，会自动触发重新生成。
EOF

echo "📊 生成统计:"
echo "- Proto 文件数量: $(echo "$proto_files" | wc -l)"
echo "- 生成的 Go 文件数量: $(find "$OUTPUT_DIR" -name "*.go" | wc -l)"
echo ""
echo "🎉 代码生成完成！输出目录: $OUTPUT_DIR"
