#!/bin/bash

cat go.mod | sed "s@_template_@${NAME}@g" > temp && rm -rf go.mod && mv temp go.mod
cp config/sample.yaml config/local.yaml
