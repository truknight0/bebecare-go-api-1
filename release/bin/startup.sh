#!/bin/sh

export BEBECARE_GO_API_1_HOME=/var/www/bebecare/bebecare-go-api-1
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
   exec `go build -o bebecare-go-api-1 server.go`
   echo "$DATE : $SERVICE_NAME startup"
   exec `./bebecare-go-api-1`
fi

echo "###################################################"
