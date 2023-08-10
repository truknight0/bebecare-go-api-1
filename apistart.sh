#!/bin/sh

export BEBECARE_GO_API_1_HOME=/var/www/bebecare/bebecare-go-api-1
SERVICE_NAME=bebecare-go-api-1

DATE=`date`

echo "####################################################"
echo "$SERVICE_NAME Start Ready..."

Cnt=`ps -ef | grep $SERVICE_NAME | grep -v grep | wc -l`
PROCESS=`ps -ef | grep $SERVICE_NAME | grep -v grep | awk '{print $2}'`

if [ $Cnt -ne 0 ]
then
   echo "$DATE : $SERVICE_NAME(PID : $PROCESS) is already running"
   echo "$DATE : $SERVICE_NAME(PID : $PROCESS) Shutdown Begin..."
   kill -9 $PROCESS
   echo "$DATE : $SERVICE_NAME(PID : $PROCESS) Shutdown Complete!"
   echo "$DATE : $SERVICE_NAME Service Restart!"
   exec `go build -o $SERVICE_NAME server.go`
   echo "$DATE : $SERVICE_NAME Startup!"
   exec `./$SERVICE_NAME`
else
   exec `go build -o $SERVICE_NAME server.go`
   echo "$DATE : $SERVICE_NAME Startup!"
   exec `./$SERVICE_NAME`
fi

echo "####################################################"
