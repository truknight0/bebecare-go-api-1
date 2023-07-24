#!/bin/sh

export BEBECARE_GO_API_1_HOME=/home/ubuntu/bebecare-go-api-1
SERVICE_NAME=bebecare-go-api-1

DATE=`date`

echo "####################################################"
echo "$SERVICE_NAME startup..."

Cnt=`ps -ef | grep $SERVICE_NAME | grep -v grep | wc -l`
PROCESS=`ps -ef | grep $SERVICE_NAME | grep -v grep | awk '{print $2}'`

if [ $Cnt -ne 0 ]
then
   echo "$DATE : $SERVICE_NAME(PID : $PROCESS) is already running"
else
   nohup $BEBECARE_GO_API_1_HOME/bin/bebecare-go-api-1-server -service=$SERVICE_NAME >> $BEBECARE_GO_API_1_HOME/logs/bebecare-go-api-1-server.log 2>&1 &
   echo "$DATE : $SERVICE_NAME startup"
fi

echo "###################################################"
