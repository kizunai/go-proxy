#!bin/bash
set -x

Usage="sh ./$Self <APP NAME>"

start(){
    ./bin/GoProxy &
    echo "start go-proxy"
    ps -ef | grep GoProxy
}

stop(){
    ps -ef | grep GoProxy | grep -v grep | awk '{print $2}' | xargs kill -9
    echo "stop go-proxy"
}

restart(){
    stop
    sleep 0.5
    start
}

case $1 in
start)
    start
    ;;
stop)
    stop
    ;;
restart)
    restart
    ;;
*)
    echo $Usage $0 {start|stop|restart}
    ;;
esac
