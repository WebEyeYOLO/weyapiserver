#!/bin/bash
if [ "$1" = "sh" ];then
    exec /bin/sh
elif [ "$1" = "bash" ];then
    exec /bin/bash
elif [ "$1" = "version" ];then
    echo $RELEASE_DESC
else
    exec /run/weyapiserver
fi