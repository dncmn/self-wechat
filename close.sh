#!/bin/sh
let port=$(pidof "main")
echo $port
kill $port
"豪杰服务器关闭服务器的脚本"