# Viết playbook cài đặt lamp wordpress trên centos7

```

- hosts: centos7
  tasks:
  - name: install pakage
    yum:
      name: '{{item}}'
      state: latest
    with_items:
    - httpd
    - mariadb
    - mariadb-server

  - name: enable service
    service:
      name: '{{item}}'
      state: started
      enabled: true
    with_items:
    - httpd
    - mariadb

  - name: open firewalld
    firewalld:
      service: '{{item}}'
      state: enabled
      permanent: true
      immediate: true
    with_items:
    - http
    - https

  - name: install wget
    yum:
      name: wget
      state: latest

  - name: dowload wp
    shell:  wget https://wordpress.org/latest.tar.gz

  - name: copy file
    copy:
      src: /root/latest.tar.gz
      dest: /var/www/html
      remote_src: yes

  - name: extract
    unarchive:
      src: /var/www/html/latest.tar.gz
      dest: /var/www/html
      remote_src: yes

  - name: copy wp html
    shell: mv /var/www/html/wordpress/* /var/www/html

  - name: set permission
    shell: chown -R apache:apache /var/www/html

  - name: install mysql-python
    yum:
      name: MySQL-python
      state: present

  - name: mv config
    shell: mv /var/www/html/wp-config-sample.php /var/www/html/wp-config.php

  - name: create db
    mysql_db:
      name: wordpress1
      state: present

  - name: create user
    mysql_user:
      name: wordpress
      host: localhost
      password: Vnpay@123
      priv: 'wordpress.*:ALL'
      state: present

  - name: config db
    replace:
      path: /var/www/html/wp-config.php
      regexp: 'database_name_here'
      replace: 'wordpress'

  - name: config user
    replace:
      path: /var/www/html/wp-config.php
      regexp: 'username'
      replace: 'wordpress'

  - name: config password
    replace:
      path: /var/www/html/wp-config.php
      regexp: 'password'
      replace: 'Vnpay@123'

  - name: update
    yum:
      name: epel-release
      state: present

  - name: rpm
    shell: rpm -ivh https://rpms.remirepo.net/enterprise/remi-release-7.rpm

  - name: yum pakage
    yum:
      name: yum-utils
      state: present

  - name: config php7.4
    shell: yum-config-manager --enable remi-php74
  
  - name: install php 7.4
    yum:
      name: '{{item}}'
      state: present
    with_items:
    - php
    - php-bcmath
    - php-cli
    - php-common
    - php-devel
    - php-gd
    - php-imap
    - php-intl
    - php-json
    - php-ldap
    - php-lz4
    - php-mbstring
    - php-mysqlnd
    - php-soap
    - php-intl
    - php-opcache
    - php-xml
    - php-pdo

  - name: restart httpd
    service:
      name: httpd
      state: restarted
```