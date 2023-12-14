#!/bin/bash

cat go.mod | sed "s@package-name@${PACKAGE_NAME}@g" > temp && rm -rf go.mod && mv temp go.mod

cat README.md | sed "s@package-name@${PACKAGE_NAME}@g" > temp && rm -rf README.md && mv temp README.md

cat README.md | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf README.md && mv temp README.md

cat README_zh-cn.md | sed "s@package-name@${PACKAGE_NAME}@g" > temp && rm -rf README_zh-cn.md && mv temp README_zh-cn.md

cat README_zh-cn.md | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf README_zh-cn.md && mv temp README_zh-cn.md

cat Dockerfile | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf Dockerfile && mv temp Dockerfile

cat version/version.go | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf version/version.go && mv temp version/version.go

cat cmd/app-name/main.go | sed "s@package-name@${PACKAGE_NAME}@g" > temp && rm -rf cmd/app-name/main.go && mv temp cmd/app-name/main.go

cat cmd/app-name/main.go | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf cmd/app-name/main.go && mv temp cmd/app-name/main.go

cat cmd/app-name/command/default.go | sed "s@package-name@${PACKAGE_NAME}@g" > temp && rm -rf cmd/app-name/command/default.go && mv temp cmd/app-name/command/default.go

cat pkg/controller/user.go | sed "s@package-name@${PACKAGE_NAME}@g" > temp && rm -rf pkg/controller/user.go && mv temp pkg/controller/user.go

cat pkg/route/user.go | sed "s@package-name@${PACKAGE_NAME}@g" > temp && rm -rf pkg/route/user.go && mv temp pkg/route/user.go

cat script/ci-prod.sh | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf script/ci-prod.sh && mv temp script/ci-prod.sh

cat script/ci-prod.sh | sed "s@username@${USERNAME}@g" > temp && rm -rf script/ci-prod.sh && mv temp script/ci-prod.sh

cat script/ci-test.sh | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf script/ci-test.sh && mv temp script/ci-test.sh

cat script/ci-test.sh | sed "s@username@${USERNAME}@g" > temp && rm -rf script/ci-test.sh && mv temp script/ci-test.sh

cat .github/workflows/deploy_main.yml | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf .github/workflows/deploy_main.yml && mv temp .github/workflows/deploy_main.yml

cat .github/workflows/deploy_main.yml | sed "s@username@${USERNAME}@g" > temp && rm -rf .github/workflows/deploy_main.yml && mv temp .github/workflows/deploy_main.yml

cat .github/workflows/deploy_test.yml | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf .github/workflows/deploy_test.yml && mv temp .github/workflows/deploy_test.yml

cat .github/workflows/deploy_test.yml | sed "s@username@${USERNAME}@g" > temp && rm -rf .github/workflows/deploy_test.yml && mv temp .github/workflows/deploy_test.yml

cat Makefile | sed "s@app-name@${APP_NAME}@g" > temp && rm -rf Makefile && mv temp Makefile

cp config/sample.yaml config/local.yaml

mv ./cmd/app-name ./cmd/"${APP_NAME}"