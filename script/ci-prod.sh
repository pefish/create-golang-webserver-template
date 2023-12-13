#!/bin/bash

set -euxo pipefail

projectName="app-name"
port="8000"
srcPath=`pwd`
configPath=$CONFIG_PATH

cd ${srcPath}

git reset --hard && git pull && git checkout test && git pull

imageName="${projectName}:`git rev-parse --short HEAD`"

if [[ "`sudo docker images -q ${imageName} 2> /dev/null`" == "" ]]; then
  sudo docker build -t ${imageName} .
fi

sudo docker stop ${projectName}-prod && sudo docker rm ${projectName}-prod

sudo docker run --name ${projectName}-prod -d -v ${configPath}:/app/config -p ${port}:${port} ${imageName}
