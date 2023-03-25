# 环境安装
https://zhuanlan.zhihu.com/p/512877083
生成grpc文件：protoc --go_out=plugins=grpc:. protocol/helloworld/helloworld.proto 

# gateway-grpc
教程：https://www.liwenzhou.com/posts/Go/grpc-gateway/
命令：protoc --go_out=plugins=grpc:./  --grpc-gateway_out=./  helloworld/helloworld.proto


# Git问题
## remote: Support for password authentication was removed on August 13, 2021
https://blog.csdn.net/qq_41646249/article/details/119777084