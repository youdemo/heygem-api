# HeyGem 服务端部署教程

### 环境要求
- Docker 和 Docker Compose
- NVIDIA GPU (用于AI模型推理)
- CUDA 环境
- Linux环境需要安装 `nvidia-container-toolkit`
  - 参考安装文档：`https://mirrors.ustc.edu.cn/help/libnvidia-container.html#_3`

### 部署步骤

1. 克隆代码仓库
```shell
git clone https://github.com/it00021hot/heygem-api.git
cd heygem-api
```

2. 进入部署目录
```shell
cd manifest/deploy
```

3. 配置数据目录
在运行之前，请确保以下目录存在（如果不存在请创建）：
```shell
# Linux/Mac环境
mkdir -p /data/models/heygem/data/face2face
mkdir -p /data/models/heygem/data/mysql
```

4. 启动服务
```shell
docker-compose up -d
```

### 服务说明
部署后将启动以下服务：
- heygem-f2f: 数字人处理服务
- f5-tts: 语音合成服务
- mysql: 数据库服务
- heygem-api: API服务（默认端口8000）

### 验证部署
服务启动后，可以通过以下命令查看服务状态：
```shell
docker-compose ps
```

可以通过以下命令查看服务日志：
```shell
docker-compose logs -f
```
### 打开web页面: [http://ip:8000](http://localhost:8000)


### 注意事项
1. 确保服务器有足够的磁盘空间和GPU内存
2. 如果是Windows环境部署，需要修改docker-compose.yml中的volumes路径格式
3. 请确保所有需要的端口（如8000）未被其他服务占用

### 常见问题排查
如果服务无法正常启动，请检查：
1. Docker服务是否正常运行
2. NVIDIA驱动和CUDA是否正确安装
3. nvidia-container-toolkit是否正确配置
4. 数据目录权限是否正确

如需技术支持，请提交Issue到项目仓库。
