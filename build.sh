#!/bin/bash

GOOS=linux go build -o main;

platform='unknown'
if [ "$(uname)" == "Darwin" ]; then
    platform='darwin'     
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    platform='linux'
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
    platform='win'
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
    platform='win'
fi

if [[ $platform != "win" ]]; then
    zip main.zip main;
elif [[ $platform == "win" ]]; then
    build-lambda-zip -o main.zip main;
fi

rm main;