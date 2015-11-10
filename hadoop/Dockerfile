from sequenceiq/hadoop-docker:2.7.0
env  PATH=$PATH:/usr/local/hadoop/bin:/usr/local/hadoop/sbin
add  core-site.xml    $HADOOP_PREFIX/etc/hadoop/core-site.xml
add  mapred-site.xml  $HADOOP_PREFIX/etc/hadoop/mapred-site.xml
add  yarn-site.xml    $HADOOP_PREFIX/etc/hadoop/yarn-site.xml
add  hdfs-site.xml    $HADOOP_PREFIX/etc/hadoop/hdfs-site.xml
add  bootstrap.sh     /etc/bootstrap.sh
run  yum install -y   bind-utils && \
		yum clean all && \
		chown root:root  /etc/bootstrap.sh && \
		chmod 700 /etc/bootstrap.sh
cmd ["/etc/bootstrap.sh"]
