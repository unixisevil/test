#!/bin/bash

: ${HADOOP_PREFIX:=/usr/local/hadoop}

role=$1

$HADOOP_PREFIX/etc/hadoop/hadoop-env.sh

rm /tmp/*.pid
#host_ip=$(ifconfig eth0|awk -F '[: ]*' '/inet addr/{print $4}')

if [ -z $HADOOP_MASTER ];then
    echo "HADOOP_MASTER not set" && exit 1
fi

if [ -z $HADOOP_SLAVE ];then
    echo "HADOOP_SLAVE not set" && exit 1
fi

sed -i  s/master/$HADOOP_MASTER/  \
        $HADOOP_CONF_DIR/core-site.xml  \
        $HADOOP_CONF_DIR/yarn-site.xml

service sshd start

if  [ $role = "master" ]; then

    sed -i  s/master/$HADOOP_MASTER/   $HADOOP_CONF_DIR/mapred-site.xml
    echo > $HADOOP_CONF_DIR/slaves
    echo y | hdfs namenode -format
        start-dfs.sh
        start-yarn.sh
    while true;do
           sleep 5
           if [ -z $HADOOP_SLAVE ];then
               continue
           fi
           for  s  in  $(dig +short  $HADOOP_SLAVE); do
               echo  $s  >> $HADOOP_CONF_DIR/slaves
           done
           content=$(cat $HADOOP_CONF_DIR/slaves|sort -u)
           echo -e  "$HADOOP_MASTER\n$content"  >  $HADOOP_CONF_DIR/slaves
    done
fi

while true;do
    sleep  100
done
