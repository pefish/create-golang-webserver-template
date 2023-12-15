#!/bin/bash

set -euxo pipefail

projectName="_username_-_app-name_"
port="8000"
configPath="${HOME}/data/_username_/_app-name_-test"

git reset --hard && git pull && git checkout test && git pull

imageName="${projectName}:`git rev-parse --short HEAD`"

if [[ "`sudo docker images -q ${imageName} 2> /dev/null`" == "" ]]; then
  sudo docker build -t ${imageName} .
fi

sudo docker stop ${projectName}-test && sudo docker rm ${projectName}-test

sudo docker run --name ${projectName}-test -d -v ${configPath}:/app/config -p ${port}:${port} ${imageName}


# CONFIG_PATH=** ./script/ci-test.sh
