## 将SRGAN作为工具部署到Web平台

使用的SRGAN：[https://github.com/brade31919/SRGAN-tensorflow](https://github.com/brade31919/SRGAN-tensorflow)

---
### 该源码使用两种后端实现方式：
1. Python实现(`Tornado` + `Tensorflow`)
2. Golang实现(`Go` + `Tensorflow-serving`)
### 前端采用Vue实现(基于Vue-Cli3)
---
## 1.Python实现

### Denpendency
* Python 3.5 + (需要支持 `asyncio`)
    * `Tensorflow` (1.13.1测试通过)
    * `Tornado` (6.0.2测试通过)
### Getting Started
* 本仓库已经包含所有运行时所需的模型，安装依赖后可直接运行
```bash
# clone本仓库
git clone https://github.com/yphub/SRGAN_webtool.git

# 进入目录
cd SRGAN_webtool

# 请确认安装了tensorflow以及tornado
# pip install tensorflow tornado
# python3下
# pip3 install tensorflow tornado

# 阻塞式运行
python main.py
# 同时安装了python2与3的不要运行上面的语句，运行下面这一个
# python3 main.py

# 后台运行
nohup python main.py &
# 同时安装python2与3的：
# nohup python3 main.py &
```
开启后，可打开浏览器通过8082端口进行访问。例如http://localhost:8082

* 本仓库中的模型是由SRGAN_pre-trained的模型进一步抽离得到(仓库已经自带了)，如需重新生成，步骤如下

在[https://github.com/brade31919/SRGAN-tensorflow](https://github.com/brade31919/SRGAN-tensorflow)中作者的[googledrive](https://drive.google.com/a/gapp.nthu.edu.tw/uc?id=0BxRIhBA0x8lHNDJFVjJEQnZtcmc&export=download)下载得到SRGAN_pre-trained预训练模型

```bash
# Download the pre-trained model from the google-drive
# Go to https://drive.google.com/a/gapp.nthu.edu.tw/uc?id=0BxRIhBA0x8lHNDJFVjJEQnZtcmc&export=download
# Download the pre-trained model to SRGAN-tensorflow/
tar xvf SRGAN_pre-trained.tar

# 解压完毕后，执行 makeSRGAN_web.py生成SRGAN_web模型文件
# 重新生成之前，请将SRGAN_web文件夹删除
# rm -r SRGAN_web
python makeSRGAN_web.py

```

生成完毕之后可被本仓库的源码使用。如需要测试，请执行`makeSRGAN_web.py`中的testInference函数自行测试。

## 2.Grpc TF-Serving部署
Go的实现请进入`serving-go`阅读`readme.md`，本节仅讲解导出serving所需的模型步骤

* 仓库已经自带模型文件`SRGAN_grpc`，如需重新生成，步骤如下

以下步骤依赖SRGAN_web模型，请确认已经生成SRGAN_web模型(上一节的模型生成步骤`makeSRGAN_web.py`)

```bash
# 重新生成之前，请将SRGAN_grpc文件夹删除
# rm -r SRGAN_grpc
python makeSRGAN_grpc.py
```

* 部署tf-serving需要docker环境，请确认已经安装docker环境

[Serving a TensorFlow Model](https://www.tensorflow.org/tfx/serving/serving_basic)
```bash
# 查看docker环境
docker version

# 拉取tensorflow-serving
docker pull tensorflow/serving

# 开启serving和模型文件映射
# 8500为grpc端口
docker run -p 8500:8500 -v SRGAN_grpc:/models/SRGAN_grpc/000001 -e MODEL_NAME=SRGAN_grpc tensorflow/serving &
```

## 3.前端代码构建
前端代码构建请移步`webfront/README.md`