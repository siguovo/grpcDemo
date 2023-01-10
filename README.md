### Project Introduction

    this demo is learn grpc with golang, from 0-1

### Note

    upgrade go in Linux
```bash
#!/bin/bash

if [ -z "$1" ]; then
    echo "./installgo.sh go-package.tar.gz"
fi

if [ -d "/usr/local/go" ]; then
    echo "Uninstall old version go"
    rm -rf /usr/local/go
fi

echo "install new go package......."

tar -C /usr/local -zxvf $1 

echo "install success!!!" 
```