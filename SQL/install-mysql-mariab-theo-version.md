# Tài liệu hướng dẫn cài đặt mysql, mariadb theo bản chỉ định

## Mysql

### Cài đặt mysql bản 8.0

- Kích hoạt MySQL 8.0 bằng lệnh sau:

    - ` yum localinstall https://dev.mysql.com/get/mysql80-community-release-el7-1.noarch.rpm `

- Cài đặt gói MySQL 8.0:
    
    - ` yum install mysql-community-server`

- Sau khi cài đặt xong sử dụng câu lệnh sau để kiểm tra version của mysql:

    - ` mysql -V `

### Cài đặt mysql bản 5.7 trên Centos 7

- Kích hoạt MySQL 5.7 bằng lệnh sau:

    - ` yum localinstall https://dev.mysql.com/get/mysql57-community-release-el7-11.noarch.rpm `

- Cài đặt gói MySQL 5.7 bằng lệnh sau:

    - ` yum install mysql-community-server `

- Sau khi cài đặt xong bạn sử dụng câu lệnh sau để kiểm tra version

    - ` mysql -V `

## MariaDB

- Tạo file repo

    - Theo mặc định repo của CentOS chỉ có sẵn MariaDB 5. Để cài đặt MariaDB 10 các bạn cần tạo repo riêng.

    - ` vi /etc/yum.repos.d/mariadb.repo `

    - Dán nội dung dưới đây vào sau đó bấm Ctrl + o và nhần Enter để lưu file, nhấn Ctrl + x để thoát khỏi nano.

    ```
        [mariadb]
        name = MariaDB
        baseurl = http://yum.mariadb.org/10.5.2/rhel7-amd64/
        gpgkey=https://yum.mariadb.org/RPM-GPG-KEY-MariaDB
        gpgcheck=1

     ```   
Lưu ý: Tại thời điểm viết bài phiên bản MariaDB 10 mới nhất là 10.5.2. Trước khi tạo repo các bạn có thể truy cập http://yum.mariadb.org để kiểm tra phiên bản, bạn muốn cài đặt phiên bản nào thì chỉ cần tìm đến phiên bản đó ở trên http://yum.mariadb.org và copy link dowload thay thế cho phần baseurl ở trên.

- Cài đặt MariaDB

    - Sau khi tạo file repo các bạn tiến hành cài đặt MariaDB bằng lệnh sau

        ` yum install MariaDB-server MariaDB-client -y `

- Sau khi cài đặt xong bạn sử dụng câu lệnh sau để kiểm tra version

    - ` mysql -V `

### PHP

- Phiên bản có sẵn trong repo của CentOS đang là 5.4. Phiên bản này khá cũ và sẽ khiến bạn gặp một số vấn đề xảy ra khi tiến hành cài đặt wordpress. Vì vậy bạn cần phải cài đặt phiên bản 7x để khắc phục. Bạn cần tiến hành thêm kho vào Remi CentOS:

    - ` rpm -Uvh http://rpms.remirepo.net/enterprise/remi-release-7.rpm `

- Cài yum-utils vì chúng ta cần tiện ích yum-config-manager để cài đặt:

    - ` yum -y install yum-utils `

- Tiến hành cài đặt php. Ở đây ta cần lưu ý về phiên bản cài đặt như sau:

```

Bản 7.0:

yum-config-manager --enable remi-php70
yum -y install php php-opcache php-mysql

Bản 7.1:

yum-config-manager --enable remi-php71
yum -y install php php-opcache php-mysql

Bản 7.2:

yum-config-manager --enable remi-php72
yum -y install php php-opcache php-mysql

Bản 7.3:
yum-config-manager --enable remi-php73
yum -y install php php-opcache php-mysql

```
- Trong bài này, mình cài phiên bản 7.2.

- Sau khi cài đặt xong, sử dụng câu lệnh sau để kiểm tra version

    - ` php -v `

