# Tài liệu cài đặt wordpress với letsencypt + Apache trên centos7
# Mục lục

## [1. Cài đặt Apache](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20wordpress%20v%E1%BB%9Bi%20letsencrypt%20%2B%20apache%20tr%C3%AAn%20cent07.md#1-c%C3%A0i-%C4%91%E1%BA%B7t-apache-1)

## [2. Cài đặt Maria-db](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20wordpress%20v%E1%BB%9Bi%20letsencrypt%20%2B%20apache%20tr%C3%AAn%20cent07.md#2-c%C3%A0i-%C4%91%E1%BA%B7t-maria-db-1)

## [3. Cài đặt PHP](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20wordpress%20v%E1%BB%9Bi%20letsencrypt%20%2B%20apache%20tr%C3%AAn%20cent07.md#3-c%C3%A0i-%C4%91%E1%BA%B7t-php-1)

## [4. Cài đặt wordpress](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20wordpress%20v%E1%BB%9Bi%20letsencrypt%20%2B%20apache%20tr%C3%AAn%20cent07.md#4-c%C3%A0i-%C4%91%E1%BA%B7t-wordpress-1)

## [5. Cài đặt lensencrypt](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20wordpress%20v%E1%BB%9Bi%20letsencrypt%20%2B%20apache%20tr%C3%AAn%20cent07.md#5-c%C3%A0i-%C4%91%E1%BA%B7t-lensencrypt-1)

## [Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20wordpress%20v%E1%BB%9Bi%20letsencrypt%20%2B%20apache%20tr%C3%AAn%20cent07.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o-1)

## 1. Cài đặt Apache

- Cài đặt httpd

    - ` yum -y install httpd `

- Cài xong, tiến hành khởi động lại service:

    - ` systemctl start httpd `

    - ` systemctl enable httpd `

- Bạn có thể check lại trang thái hoạt động của service bằng cách gõ:

    - ` systemctl status httpd `

<img src="https://imgur.com/A5Gtj6E.png">

- Nếu bạn sử dụng hệ điều hành trên máy ảo, bạn có thể tắt firewall để có thể truy cập trên browser của máy thực:

    - ` systemctl stop firewalld `

## 2. Cài đặt Maria-db

- Trên cửa sổ terminal, tiến hành cài đặt mariadb:

    - ` yum -y install mariadb mariadb-server `

- Tiến hành khởi động mariadb service:

    - ` systemctl start mariadb `

- Cài lại mật khẩu mật khẩu cho quyền root của cơ sở dữ liệu:

    - ` sudo mysql_secure_installation `

- Sau khi thiết lập xong, kích hoạt mariadb để khởi động cùng hệ thống:

    - ` systemctl enable mariadb `

## 3. Cài đặt PHP

- Để cài đặt PHP 5.6 , bạn phải cài đặt và kích hoạt kho EPEL và Remi cho hệ thống CentOS 7 của mình bằng các lệnh bên dưới:

    - ` yum install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm `

    - ` yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm `

- Tiếp theo, cài đặt yum-utils:

    - ` yum install yum-utils `

- Cài đặt PHP 5.6 với tất cả các mô-đun cần thiết như sau:

    - ` yum-config-manager --enable remi-php56 `

    - ` yum install php php-mcrypt php-cli php-gd php-curl php-mysql php-ldap php-zip php-fileinfo `

- Kiểm tra version php:
    - ` php -v `

<img src="https://imgur.com/MACZ2ES.png">

## 4. Cài đặt wordpress

- Tạo cơ sở dữ liệu và tài khoản cho WordPress:

    - ` mysql -u root -p `

- Tiếp theo bạn sẽ tạo cơ sở dữ liệu cho wordpress.

    - ` CREATE DATABASE wordpress; `

- Tạo một tài khoản riêng để quản lí cơ sở dữ liệu cho WordPress. 

    - ` CREATE USER user@localhost IDENTIFIED BY 'pass'; `

- Tiến hành cấp quyền quản lí cơ sở dữ liệu wordpress cho user mới tạo.

    - ` GRANT ALL PRIVILEGES ON wordpress.* TO user@localhost IDENTIFIED BY 'pass'; `

- Sau đó xác thực lại những thay đổi về quyền:

    - ` FLUSH PRIVILEGES; `

- Sau khi hoàn tất, thoát khỏi mariadb:

    - ` exit `

### Tải và cài đặt WordPress

- Trước khi bắt đầu tiến hành cài gói hỗ trợ php-gd:

    - `yum -y install php-gd `

- Tiến hành tải xuống WordPress với phiên bản mới nhất.

    - ` wget http://wordpress.org/latest.tar.gz `

- Tiến hành giải nén file latest.tar.gz:

    - ` tar xvfz latest.tar.gz `

- Copy các file trong thư mục WordPress tới đường dẫn /var/www/html như sau:

    - ` cp -Rvf /root/wordpress/* /var/www/html `

### Cấu hình WordPress

- Ta di chuyển đường dẫn tới thư mục chứa các file cài đặt WordPress như sau:

    - ` cd /var/www/html `

- File cấu hình wordpress là wp-config.php. Tuy nhiên tại đây chỉ có file wp-config-sample.php. Tiến hành copy lại file cấu hình như sau:

    - ` cp wp-config-sample.php wp-config.php `

- Mở file config với vi để sửa:

    - ` vi wp-config.php ` 

<img src="https://imgur.com/CEqioTw.png">

- Hoàn tất phần cài đặt giao diện ,trên trình duyệt, gõ địa chỉ ip server trên thành url, trình duyệt sẽ xuất hiện như sau:

## 5. Cài đặt lensencrypt

- Cài đặt mod_ssl

    - ` yum -y install mod_ssl `

- Định cấu hình Apache

    - Tạo một thư mục gốc tài liệu cho trang web của bạn:

        ` mkdir /var/www/congphan `

    - Tạo tệp cấu hình máy chủ ảo cho trang web của bạn bằng cách mở tệp đó bằng nano và sau đó dán các nội dung sau vào bên trong:

        ` nano /etc/httpd/conf.d/congphan.xyz.conf `

    ```

<VirtualHost *:80>

    ServerAdmin admin@test.com

    DocumentRoot "/var/www/html"

    ServerName congphan.xyz

    ServerAlias congphan.xyz

    ErrorLog "/var/log/httpd/test.error_log"

    CustomLog "/var/log/httpd/test.access_log" common

</VirtualHost>

    ```

- Lưu tệp và thay đổi chủ sở hữu của tệp '/var/www/congphan' thành người dùng Apache để Apache có thể đọc tệp:

    - ` chown -R apache: apache /var/www/congphan `

- Để cài đặt certbot, trước tiên chúng ta cần đảm bảo rằng chúng ta đã bật kho lưu trữ EPEL, để thực hiện điều đó, hãy thực hiện lệnh sau:

    - ` yum -y cài đặt epel-release `

- Đảm bảo rằng yum-utils đã được cài đặt:

    - ` yum -y cài đặt yum-utils `

- Sau đó cài đặt certbot cho Apache:

    - ` yum -y install certbot-apache `

- Bây giờ chúng ta đã cài đặt certbot, hãy chạy certbot bằng lệnh sau:

    - ` certbot --apache -d congphan.xyz -d www.congphan.xyz `

- Do cài đặt wordpress ta để source ở thư mục /var/www/congphan, nhưng khi cài đặt lensencrypt ta để thư mục Documentroot là /var/www/html, nên cần chuyển các file từ /var/www/congphan sang /var/www/html bằng câu lệnh sau:

    - ` cd /var/www/ `

    - ` cp -ra congphan/* html/ `

- Cấu hình hoàn tất, kiểm tra lại bằng cách truy cấp đường linh congphan.xyz

## Nguồn tham khảo:

- https://www.rosehosting.com/blog/how-to-install-lets-encrypt-with-apache-on-centos-7/

- https://news.cloud365.vn/cai-dat-wordpress-tren-may-centos/

- https://www.tecmint.com/install-php-5-6-on-centos-7/
