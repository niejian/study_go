#!/bin/bash
if [ ! -n "$1" ]; then
  echo "请输入进程名, 如 watchDog"
  exit

fi

#dirName = "project_path=$(cd `dirname $0`; pwd)"


PID=$(ps -ef|grep $1|grep -v "grep"|grep -v "$0"|awk '{printf $2}')
echo "PID: $PID"


if [ $? -eq 0 ]; then
    echo "process id:$PID"
else
    echo "process $1 not exit"
    exit
fi

#for p_id in ${PID}
#do
#   echo "p_id: ${p_id}"
#done
#exit

if test -z "$PID"
then
  echo "$1 未启动, 开始启动"
else
  kill -9 ${PID}

  if [ $? -eq 0 ];then
      echo "kill $1 success"
  else
      echo "kill $1 fail"
  fi
fi

DirName=$(cd $(dirname $0); pwd)
cd $DirName

#CHMOD="chmod +x ./$1"
#echo -e "$CHMOD"
#$CHMOD

#COMMOND="nohup $DirName/$1 > watchdog.log 2>&1 &"
#COMMOND="nohup $DirName/$1 > watchdog.log 2>&1 &"
COMMOND="nohup $DirName/$1 2>&1 &"
echo "请自行运行启动命令："
echo -e "$COMMOND"
#$COMMOND
#sh $COMMOND
#
#echo "启动成功"

#/bin/sh $COMMOND

