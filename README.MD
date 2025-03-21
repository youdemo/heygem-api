# heygem服务端/客户端异地部署


### 部署服务端
自行安装docker，cuda
linux环境需要安装`nvidia-container-toolkit`, 参考`https://mirrors.ustc.edu.cn/help/libnvidia-container.html#_3`

```shell
# 下载docker-compose.yml文件到服务器
wget https://raw.githubusercontent.com/it00021hot/HeyGem.ai/refs/heads/main/deploy/docker-compose.yml
# 运行(默认配置的路径为linux/mac路径格式，如果是windows部署请自行修改volumes)
docker-compose up -d
```
### 部署客户端
自行安装node环境
#### 拉取代码
```shell
git clone https://github.com/it00021hot/HeyGem.ai.git && cd HeyGem.ai
```
#### 修改 `.env` 配置
```
# 改成自己的服务端ip
FACE_API=http://10.8.5.210:8383/easy
TTS_API=http://10.8.5.210:18180
FILE_API=http://10.8.5.210:8000

# 本地语音视频存放地址，改成自己想存放的位置
MODEL_PATH=/Users/zhihu/heygem_data/face2face/temp
TTS_PRODUCT_PATH=/Users/zhihu/heygem_data/face2face/temp
TTS_ROOT_PATH=/Users/zhihu/heygem_data/voice/data
TTS_TRAIN_PATH=/Users/zhihu/heygem_data/voice/data/origin_audio
```
#### run
```shell
npm install
npm run dev
```
#### build
```shell
# windows
npm run build:win
# mac 
npm run build:mac
```
