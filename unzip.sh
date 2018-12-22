#!/bin/sh
sh close.sh
tar -xvf manan.tar
rm manan.tar
cd self-wechat &&\
sh start.sh
"关闭之前的web服务器，解压缩tar文件、进入到项目目录，然后执行启动脚本"
“unzip.sh和close.sh  manan.tar(项目的压缩文件),close.sh都是在同一个目录下”