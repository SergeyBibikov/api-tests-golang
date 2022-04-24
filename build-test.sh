#! /bin/bash
# Place this file in the dir where the repo with tests 
# and the repo with the service were cloned to

killall app > /dev/null 2>&1

cd ./api-for-tests &&./start.sh && cd ..
cd ./api-tests-golang && go test

killall app > /dev/null 2>&1
docker stop $(docker ps -aq)