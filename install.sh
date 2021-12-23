#!/usr/bin/sh

# Check go install
if ! command -v go &> /dev/null
then
    echo -e "[\033[5;38;5;160mERROR\033[0m]\033[38;5;160m Go isn't installed\033[0m"
    echo -e "[\033[5;38;5;114mHELP\033[0m] Refer to --> https://go.dev/doc/install"
    exit 1
fi

# Check go version
if [[ $(go version) == "go version go1.17."* ]]
then
    echo -e "[\033[5;38;5;75mINFO\033[0m] You have a go version >= 1.17"
else 
    echo -e "[\033[5;38;5;160mERROR\033[0m]\033[38;5;160m Go version isn't compatible\033[0m"
    echo -e "[\033[5;38;5;114mHELP\033[0m] Download new version or version >= 1.17. Refer to --> https://go.dev/doc/install"
    exit 1
fi

# Build gktool
cd $GOPATH
go build
if test $? -eq 0
then
    echo -e "[\033[5;38;5;34mSUCCESS\033[0m] BUILD bin"
fi

# Move build in bin path
sudo mv gktool /usr/bin
if [[ $? -eq 0 ]]
then
    echo -e "[\033[5;38;5;34mSUCCESS\033[0m] MV gktool in bin path"
fi
