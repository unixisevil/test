#!/bin/sh
set -ex

for  card in  $(lshw -class network|grep 'logical name'|awk -F ':' '{print $2}');do
     echo "now config : $card in promisc mode"
     ifconfig  $card  promisc 
     echo "now create one new bridge, and add ${card} to it"
     ip=$(ifconfig $card |grep 'inet addr' |awk -F ':[ \t]*|[ \t]+' '{print $4}')
     mask=$(ifconfig $card |grep 'inet addr' |awk -F ':[ \t]*|[ \t]+' '{print $8}')
     echo "got $card,  ip: $ip , mask: $mask"
     brctl  addbr  br${card}
     brctl  addif  br${card}  ${card}
     ifconfig  br${card}  $ip   netmask $mask
done
