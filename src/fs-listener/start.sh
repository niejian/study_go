#!/bin/bash
# 启动
if [ ! -n "$1" ]; then
  echo "请输入进程名, 如 watchDog"
  exit

fi
DirName=$(cd $(dirname $0); pwd)
cd $DirName

chmod +x ./$1

#COMMOND="nohup $DirName/$1 > watchdog.log 2>&1 &"
#COMMOND="nohup $DirName/$1 > watchdog.log 2>&1 &"
COMMOND="$DirName/$1"
echo "启动命令：$COMMOND"
$COMMOND
echo "启动成功"