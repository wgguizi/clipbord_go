# /usr/bin/bash
dir=$(cd $(dirname $0);pwd);
binFile=$dir/main;
#echo $binFile;

chmod +x $binFile

startCmd="nohup $dir/main > /dev/null 2>&1 &";
stopCmd="ps -aux|grep $dir/main|grep -v grep|awk '{print \$2}'|xargs kill -9";
netCmd="netstat -tlnp";

if [ "start" == "$1" ]; then
#  echo $startCmd;
  eval $startCmd;
  eval $netCmd;
elif [ "stop" == "$1" ]; then
#  echo $stopCmd;
  eval $stopCmd;
  eval $netCmd;
elif [ "restart" == "$1" ]; then
  eval $stopCmd;
  eval $startCmd;
  eval $netCmd;
else
  echo -e "\nUsage: <bashFile> start|stop|restart\n"
  eval $netCmd;
fi
