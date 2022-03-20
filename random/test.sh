#!/bin/sh

# create necessary variables
DOCKERFILE="Dockerfile"
MODULES=( numpy scipy pandas )
FILENAME='./main.py'

# generate file content
CONTENT="FROM python:3\nRUN pip install ${MODULES[@]}\nCMD ['python', $FILENAME ]"

# pipe to the file
echo -e "$CONTENT" > $DOCKERFILE

# print sha sum
echo "$(sha256sum Dockerfile)"