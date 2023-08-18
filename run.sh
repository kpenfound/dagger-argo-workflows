#!/bin/bash
set -euo pipefail

apt-get update -y
apt-get install -y wget

cd /tmp
wget https://github.com/dagger/dagger/releases/download/v0.8.4/dagger_v0.8.4_linux_arm64.tar.gz -O dagger.tar.gz
tar -xf dagger.tar.gz
ls
mv dagger /usr/local/bin/dagger
cd /work

go mod download
echo "starting dagger"
dagger run go run main.go
