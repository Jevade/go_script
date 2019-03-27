#!/bin/bash
SERVER="58spy.linux"
INTERVAL=2
BASE_DIR=$(pwd)

ARGS="-c conf/config.yaml"

function select_os(){
    if [ "$(uname)" == "Darwin" ]; then
        SERVER="58spy.mac" 
	echo "mac start 58spy"
        # Do something under Mac OS X platform        
    elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
        SERVER="58spy.linux" 
	echo "linux start 58spy"
        # Do something under GNU/Linux platform
    elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
        SERVER="58spy.exe" 
	echo "MINGW32_NT start 58spy"
        # Do something under 32 bits Windows NT platform
    elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
        SERVER="58spy.exe" 
	echo "MINGW_64 start 58spy"
        # Do something under 64 bits Windows NT platform
    fi
 }

function start(){
    	
    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        echo "$SERVER already running"
        exit 1
    fi

    select_os
    nohup $BASE_DIR/$SERVER $ARGS  2>&1  &
    echo $?
    echo $BASE_DIR/$SERVER 
    echo "sleeping....." && sleep $INTERVAL
    if [ "`pgrep $SERVER -u $UID`" == "" ];then
        echo "$SERVER start failed "
        exit 1
    fi

}
function status(){
    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        echo "$SERVER is running"
    else
        echo "$SERVER is not running"
    fi
}
function stop(){
    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        kill -9 `pgrep $SERVER -u $UID`
    fi

    echo "sleeping....." && sleep $INTERVAL

    if [ "`pgrep $SERVER -u $UID`" != "" ];then
        echo "$SERVER stop field"
        exit 1
    fi
}

select_os
echo $SERVER
case "$1" in 
        'start')
        start
        ;;
        'stop')
        stop
        ;;
        'status')
        status
        ;;
        'restart')
        stop && start
        ;;
        *)
        echo "usage:$0 {start|stop|restart|status}"
        exit 1
        ;;
esac
