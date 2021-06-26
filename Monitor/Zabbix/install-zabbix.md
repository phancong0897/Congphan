# Cài đặt Zabbix 5 trên Centos 7 (sử dụng với mariadb)

Zabbix 5.0 yêu cầu phiên bản php 7.2 trở lên

### 1. Cài php

### Cài php 7.2:

```

sudo yum install -y epel-release yum-utils

sudo yum install -y http://rpms.remirepo.net/enterprise/remi-release-7.rpm

yum-config-manager --enable remi-php72

yum install -y php php-common php-opcache php-mcrypt php-cli php-gd php-curl php-mysqlnd

yum update -y

```

### Kiểm tra phiên bản php đã cài

` php -v `

### Tắt firewall và selinux

```

sudo systemctl disable firewalld

sudo systemctl stop firewalld

sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/sysconfig/selinux

sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config

setenforce 0

```

### 2. Cài MariaDB 10.5

### Add repository to CentOS 7:

```

sudo tee /etc/yum.repos.d/mariadb.repo<<EOF

[mariadb]

name = MariaDB

baseurl = http://yum.mariadb.org/10.5/centos7-amd64

gpgkey=https://yum.mariadb.org/RPM-GPG-KEY-MariaDB

gpgcheck=1

EOF

```

### Xác nhận kho lưu trữ đang hoạt động bằng cách cập nhật bộ nhớ cache.

` yum makecache `

### Liệt kê các kho lưu trữ có sẵn:

` yum repolist `

### Cài đặt MariaDB 10.5

` yum install MariaDB-server MariaDB-client `

### Start và enable dịch vụ mariadb

` sudo systemctl enable --now mariadb `

### Tạo database cho zabbix

` mysql -u root -p `

```

CREATE DATABASE zabbix CHARACTER SET UTF8 COLLATE UTF8_BIN;

CREATE USER 'zabbix'@'%' IDENTIFIED BY 'yourpassword';

GRANT ALL PRIVILEGES ON zabbix.* TO 'zabbix'@'%';

quit;

```

### 3. Cài đặt zabbix 5.0 repository

` sudo yum install -y https://repo.zabbix.com/zabbix/5.0/rhel/7/x86_64/zabbix-release-5.0-1.el7.noarch.rpm `

### Cài Zabbix 5.0 Server và frontend với MySQL support:

` sudo yum install -y zabbix-server-mysql zabbix-agent zabbix-get `

### Cài Zabbix Frontend:

```

sudo yum-config-manager --enable zabbix-frontend

sudo yum -y install centos-release-scl

sudo yum -y install zabbix-web-mysql-scl zabbix-apache-conf-scl

```

### Import database đối với mysql/mariadb

` zcat /usr/share/doc/zabbix-server-mysql*/create.sql.gz | mysql -uzabbix -p zabbix `

### Sửa nội dung file /etc/zabbix/zabbix_server.conf

```

DBHost=localhost

DBName=zabbix

DBUser=zabbix

DBPassword=yourpassword             #Nhập mật khẩu của bạn tại đây


```

### Cấu hình PHP cho zabbix frontend. Sửa file /etc/opt/rh/rh-php72/php-fpm.d/zabbix.conf

Sửa dòng

` php_value date.timezone Asia/Ho_Chi_Minh `

### Khởi động Zabbix server và các tiến trình agent và enable chúng khởi động khi hệ thống boot

```

sudo systemctl restart zabbix-server zabbix-agent httpd rh-php72-php-fpm

sudo systemctl enable zabbix-server zabbix-agent httpd rh-php72-php-fpm

```

### Khởi động lại httpd

` systemctl restart httpd `

Truy cập vào http://(Zabbix server’s hostname or IP address)/zabbix/ để cài đặt

Đăng nhập bằng tài khoản mặc định ID/pass: Admin/zabbix