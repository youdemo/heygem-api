#!/bin/sh

# 定义路径
ASSETS_DIR="app/resource/assets"
FAVICON_FILE="app/resource/favicon.ico"
INDEX_FILE="app/resource/index.html"
PUBLIC_ASSETS_DIR="app/resource/public/assets"

# 检查 app/resource/assets 是否存在
if [ -d "$ASSETS_DIR" ]; then
    echo "检测到 $ASSETS_DIR 文件夹存在。"

    # 删除 app/resource/public/assets（如果存在）
    if [ -d "$PUBLIC_ASSETS_DIR" ]; then
        echo "删除旧的 $PUBLIC_ASSETS_DIR 文件夹..."
        rm -rf "$PUBLIC_ASSETS_DIR"
    fi
    # 移动 app/resource/assets 到 app/resource/public/
    echo "移动 $ASSETS_DIR 到 app/resource/public/ ..."
    mv "$ASSETS_DIR" "app/resource/public/"
    mv "$FAVICON_FILE" "app/resource/public/"
    mv "$INDEX_FILE" "app/resource/public/"
else
    echo "$ASSETS_DIR 文件夹不存在，无需操作。"
fi

# 启动 heygem-api 应用
echo "启动 heygem-api ..."
exec /app/heygem-api
