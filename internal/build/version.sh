#!/bin/bash
COMMIT_HASH=$(git log -1 --pretty=%h)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
VERSION=$(echo ${BRANCH} | sed -e 's/.*\/v//g' | sed -e 's/\/.*//g')
echo -n "${VERSION}-${COMMIT_HASH}" > version.txt
