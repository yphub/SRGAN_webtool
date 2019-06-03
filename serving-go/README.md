## Go调用SRGAN后端

* 运行和构建之前请确保已经构建完`SRGAN_grpc`(serving模型导出)、`webfront/dist`(前端构建)
* 构建之前确保已经安装go环境
```bash
go version
```

* 从submodule安装依赖(该过程在母目录`SRGAN_webtool`下进行)
```bash
# 返回母目录
cd ..

# 从submodule安装依赖
git submodule init
git submodule update

# 返回该目录
cd serving-go

# 进行构建
sh build.sh

```

* 构建完毕后，在`build/server`中有构建结果。可执行文件必须与config.json同级目录，默认配置前端静态文件为仓库的`dist`文件夹，如果要将整个打包，可将`dist`复制到`build/server`中，并在配置修改`"StaticFile":"dist"`

```bash
# 复制dist文件夹
cp -r ../webfront/dist build/server/
# 修改配置
vi build/server/config.json

# 运行
cd build/server
# Linux下
./server
# Windows下
# ./server.exe

```