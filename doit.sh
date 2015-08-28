#!/bin/sh
set -xe

#bunzip2 -c coreos_production_image.bin.bz2  > /dev/vda
#blockdev -v --rereadpt  /dev/vda
#mount -t ext4  /dev/vda9   /mnt
mkdir -p  /mnt/var/lib/coreos-install

res=$(echo "hello"|openssl passwd -stdin -1)

cat <<-EOF  > /mnt/var/lib/coreos-install/user_data
#cloud-config
coreos:
  units:
    - name: docker.service
      command: start
      drop-ins:
        - name: 50-insecure-registry.conf
          content: |
            [Service]
            Environment=DOCKER_OPTS='--insecure-registry="d.nicescale.com:5000"'
    - name: mongodb.service
      command: start
      enable: true
      content: |
        [Unit]
        Description=High-performance, schema-free document-oriented database
        After=network.target
        
        [Service]
        User=root
        ExecStart=/usr/bin/mongod --quiet --config /etc/mongodb.conf

        [Install]
        WantedBy=multi-user.target
    - name: prometheus.service
      command: start
      enable: true
      content: |
        [Unit]
        Description=An open-source service monitoring system and time series database
        After=network.target
        
        [Service]
        User=root
        ExecStart=/usr/bin/prometheus -config.file /etc/prometheus.yml -storage.local.path /data/monitor -alertmanager.url http://127.0.0.1 -storage.local.memory-chunks 80000 -storage.local.max-chunks-to-persist 80000 -storage.local.index-cache-size.fingerprint-to-metric 5242880 -storage.local.index-cache-size.fingerprint-to-timerange 2621440 -storage.local.index-cache-size.label-name-to-label-values 5242880 -storage.local.index-cache-size.label-pair-to-fingerprints 10485760

        [Install]
        WantedBy=multi-user.target
    - name: csphere.service
      command: start
      enable: true
      content: |

write_files:
  - path: /etc/prometheus.yml
    permissions: 0644
    owner: root
    content: |
      global:
        scrape_interval:     30s
        evaluation_interval: 45s
      rule_files:
        - "/data/alarm-rules/*.rule"
      scrape_configs:
        - job_name: 'csphere-exporter'
          basic_auth:
            username: 'csphere'
            password: 'helloworld'
          scrape_interval: 30s
          scrape_timeout: 10s
          metrics_path: '/api/metrics'
          target_groups:
            - targets: ['127.0.0.1']
  - path: /etc/mongodb.conf
    permissions: 0644
    owner: root
    content: |
      dbpath=/data/db
      logpath=/data/logs/mongodb.log
      logappend=true
      bind_ip = 0.0.0.0
      port = 27017
      journal=true
      smallfiles=true

users:
  - name: yj
    passwd: $res
    groups:
      - sudo
      - docker
ssh_authorized_keys:
  - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDUaII2dtAVwg5eqfQcA+xIkvumcjtNm0owQxsMqCXjb/GXrfh1MjTC9m4TmBhj9zKQZ3vudlE3ki1ne0cApKgI0TJAobxrQezwzjbfAtW2BYONUzX7dR1g8IbyleDL/sy5AtXcU7f8SFhzKSoo5C41lwc2Ac2G09i07VtfyHn8mqMy2nGMpbSVxe+kLhK4P0ZK3SuJdoKu5nPSKbWpt1q8+j05cUZe3mYwtEStCAg+6JsiwRDCwZTQtYz/NaVuxpPM6VqWiZ/0fO+AB9Q6/P88KO19EKvs3f+a3AcyLvsAfuC0/xpjrgcKdml4cCWcAu9MJO/E4sQzLKKIN8GiLPkX yj@yj-Inspiron-N5010
EOF

#umount /mnt
#reboot
