#!/bin/bash

: ${HADOOP_PREFIX:=/usr/local/hadoop}

role=$1

$HADOOP_PREFIX/etc/hadoop/hadoop-env.sh

rm /tmp/*.pid
#host_ip=$(ifconfig eth0|awk -F '[: ]*' '/inet addr/{print $4}')

sed -i  s/master/$HADOOP_MASTER/  \
        $HADOOP_CONF_DIR/core-site.xml  \
        $HADOOP_CONF_DIR/mapred-site.xml \
        $HADOOP_CONF_DIR/yarn-site.xml

service sshd start

if  [ $role = "master" ]; then
    echo > $HADOOP_CONF_DIR/slaves
	hdfs namenode -format
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
           echo "$content"  >  $HADOOP_CONF_DIR/slaves
    done
fi

while true;do
    sleep  100
done
