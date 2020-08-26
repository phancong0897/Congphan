# Tài liệu cài đặt Wordpress trên 2 server

## Mục lục

### [1. Mô hình](https://github.com/phancong0897/Congphan/blob/master/Wordpress/install-wordpress-2-server.md#1-m%C3%B4-h%C3%ACnh-1)

### [2. Cài đặt](https://github.com/phancong0897/Congphan/blob/master/Wordpress/install-wordpress-2-server.md#2-c%C3%A0i-%C4%91%E1%BA%B7t-1)

### [2.1 Máy chủ SQL](https://github.com/phancong0897/Congphan/blob/master/Wordpress/install-wordpress-2-server.md#21-m%C3%A1y-ch%E1%BB%A7-sql-1)

### [2.2 Máy chủ web](https://github.com/phancong0897/Congphan/blob/master/Wordpress/install-wordpress-2-server.md#22-tr%C3%AAn-m%C3%A1y-web)

### [Nguồn tham khảo] (https://github.com/phancong0897/Congphan/blob/master/Wordpress/install-wordpress-2-server.md#nu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Mô hình

- Chuẩn bị 2 máy chạy HĐH CentOS7.

- Trong đó một máy sẽ đóng vai trò MySQL Server và một máy sẽ là Web Server. Hai máy có cấu hình tối thiểu như sau:

- Cấu hình máy Web Server

    - 1 CPU

    - 512M RAM

    - 2 interface bao gồm:

        ens33 dùng để truy cập ra internet có địa chỉ : 172.16.2.230.

        ens37 dùng trong local có địa chỉ : 192.168.14.133

- Cấu hình máy MySQL

    - 1 CPU

    - 512M RAM

    - 2 interface
     
        ens33 dùng để truy cập internet có địa chỉ : 172.16.2.235

        ens37 dùng trong local có địa chỉ: 192.168.14.134

- Lưu ý: bạn thực hiện mở port 80 trên Web Server và mở kết nối tới DB còn trên máy MySQL cần mở port 3306.

## 2. Cài đặt

### 2.1 máy chủ SQL

- Cài đặt dịch vụ MySQL

```
    - wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm

    - rpm -ivh mysql-community-release-el7-5.noarch.rpm

    - yum install mysql-server
```

- Khởi động dịch vụ MySQL

    - ` systemctl start mysqld `

- Cho phép cải thiện tính bảo mật bằng lệnh sau :

    - ` mysql_secure_installation `

- Để bắt đầu, ta đăng nhập vào MySQL bằng tài khoản root bằng lệnh sau :

    - ` mysql -u root -p `

```

    - mysql> creat database congphan;

    - mysql> create user cong@192.168.14.133 identified by '1234';

    - mysql> GRANT ALL PRIVILEGES ON congphan.* to cong@192.168.14.133 identified by '1234';

    - mysql> flush privileges;

    - mysql> exit;

```

- Lưu Ý:

    - congphan: là tên database.

    - 192.168.14.133 : là địa chỉ của của máy Web truy cập MySQL

    - admin : là user để wordpress sử dụng để đăng nhập vào MySQL

    - 1234 : là mật khẩu của user1

### 2.2 Trên máy Web

### Cài đặt Apache

- Để cài đặt Apache ta sử dụng lệnh sau :

    - ` yum install httpd -y `

- Sau khi cài đặt, ta khởi động dịch vụ Apache và cho phép nó khởi động cùng hệ thống :

    - ` systemctl start httpd `

    - ` systemctl enable httpd `

- Để kiểm tra dịch vụ ta truy cập địa chỉ ip công cộng của máy Web server trong trình duyệt Web. Ở đây địa chỉ là 172.16.2.230.

    - ` http://172.16.2.230 `

- Ta sẽ thấy trang Web Apache mặc định của CentOS 7 như sau:

<img src="https://imgur.com/q2y8PDO.png">

### Cài đặt PHP

- Ở đây mình muốn cài đặt php phiên bản 5.6 nên trước tiên ta phải bật kho lưu trữ EPEL và Remi cho hệ thống CentOS 7 bằng các lệnh dưới đây:

```
    - yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm

    - yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm
```

- Tiếp theo cài đặt Yum-util để tăng cường các chức năng cũng như cung cấp cho nó các tùy chọn quản lý gói nâng cao và cũng giúp sử dụng dễ dàng hơn.

    - ` yum install yum-utils `

- Sau khi cài đặt xong, ta tiến hành cài đặt php 5.6 với những lệnh sau :

```
    - yum -config-manager --enable remi-php56

    - yum install php php-mysql php-gd php-pear –y

```

- Kiểm tra xem php đã được cài chưa bằng lệnh

    - ` php -v `

- Nhập lệnh sau để tạo 1 tệp php :

    - ` echo "<?php phpinfo(); ?>" > /var/www/html/info.php `

Bây giờ ta cần khởi đông lại dịch vụ Apache để nó cập nhật module mới :

    - ` systemctl restart httpd `

- Vào trình duyệt Web và truy cập địa chỉ http://địa_chỉ_ip/info.php, địa chỉ của mình là 172.16.2.230 nên sẽ là như sau:

    - ` http://172.16.2.230/info.php `

- Nếu hiển thị như này là đã thành công .

<img src="https://imgur.com/sJsVC5G.png">

### Cài đặt Wordpress

- Trước tiên, ta truy cập vào thư mục /var/ww/html/ sau đó tiến hành download WordPres từ internet vào thư mục này để tránh việc phải sao chép lại thư mục wordpress vào đây.

```
    - cd /var/www/html

    - wget https://wordpress.org/latest.tar.gz

```

- sau khi tải về nó là 1 tập tin nén, ta tiến hành giải nén tập tin.

    - ` tar xzvf latest.tar.gz `

Tiếp theo, trước khi định cấu hình cho WordPress ta cần đổi tên nó thành wp-config.php. Thực hiện điều đó như sau :

```

    - mv wordpress/* /var/www/html/

    - mv wp-config-sample.php wp-config.php

```

- Bây giờ ta sẽ cấu hình bằng cách thay đổi 1 số thông tin cập nhật theo cơ sở dữ liệu trong file wp-config.php .

    - ` vi /var/www/html/wp-config.php `

<img src="https://imgur.com/f4F33Q0.png">

- Tiếp theo ta tiến hành tắt cổng IFace ens33 trên máy sql để máy không thể đi ra ngoài mạng và chỉ sử dụng để kết nối trong local .

- Vào trình duyệt Web và điều hướng đến địa chỉ ip của Web server.

    - ` http://tên_miền_or_ip_address `

    - ` http://172.16.2.230 `

- Nếu hiển thì như hình thì ta đã cài đặt thành công và chúng ta tiến hành cài đặt wordpress như bình thường.

<img src="https://imgur.com/NC5qYg6.png">


## Nuồn tham khảo:

https://news.cloud365.vn/cai-dat-wordpress-tren-2-server-bang-centos-7/



