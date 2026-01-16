#!/bin/bash


function DumbLog {
    if [ "$#" -ne 2 ]; then
        echo "Bad Programming, DumbLog expect two arguments (prefix, message)"
        exit 1
    fi
    RESET="\033[0m"
    BOLD="\033[1m"
    echo -e "$1 $RESET $BOLD $2 $RESET"
}

function SuccessLog {
    if [ "$#" -ne 1 ]; then
        echo "Bad Programming, SuccessLog expect one argument"
        exit 1
    fi
    GREEN="\033[32m"
    SuccessPrefix="$GREEN::SUCCESS::"
    DumbLog "$SuccessPrefix" "$1"
}

function InfoLog {
    if [ "$#" -ne 1 ]; then
        echo "Bad Programming, InfoLog expect one argument"
        exit 1
    fi
    BLUE="\033[34m"
    InfoPrefix="$BLUE::INFO::"
    DumbLog "$InfoPrefix" "$1"
}

# i'm still not writing a main func tbh

set -e
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <File/or/Folder/To/Compile>"
    echo "Note: include ./ if it's a folder"
    exit 1
fi

app_name="balafetch"
PLATFORMS=(
    "windows/amd64"
    "windows/arm64"
    "linux/amd64"
    "linux/arm64"
    "linux/arm"
    "darwin/amd64"
    "darwin/arm64"
    "freebsd/arm64"
    "freebsd/amd64"
    "freebsd/arm"
)
TARGET_FOLDER="dist"
COMPILE_PATH_TARGET="$1"

InfoLog "Creating $TARGET_FOLDER Folder"
mkdir -p $TARGET_FOLDER


declare -i current=1
for platform in ${PLATFORMS[@]}; do

    OS="${platform%/*}"
    ARCH="${platform#*/}"
    
    export GOOS=$OS
    export GOARCH=$ARCH
    
    output_file="$app_name-$OS-$ARCH"
    if [ $OS = "windows" ]; then
        output_file+='.exe'
    fi

    InfoLog "Compiling $output_file from $COMPILE_PATH_TARGET"
    go build -o $TARGET_FOLDER/$output_file $COMPILE_PATH_TARGET
    SuccessLog "Compiling Done ($current/${#PLATFORMS[@]})"
    current+=1
done

SuccessLog "Compilation to all platforms successful"
