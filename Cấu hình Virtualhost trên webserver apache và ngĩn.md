# Tài liệu cấu hình Virtualhost trên Apache và NginX

# Mục lục

## [1. Cấu hình Virtualhost trên NginX]()

## [2. Cấu hình Virtualhost trên Apache]()

## 1. Cấu hình Virtualhost trên NginX

- Thêm kho lưu trữ CentOS 7 EPEL: 

    - ` sudo yum install epel-release `

- Cài đặt Nginx

    - ` yum install nginx `

- Khởi động Nginx:

    - ` systemctl start nginx `

- Nếu bạn đang chạy tường lửa, hãy chạy các lệnh sau để cho phép lưu lượng HTTP và HTTPS:

    - ` firewall-cmd --permanent --zone=public --add-service=http `

    - ` firewall-cmd --permanent --zone=public --add-service=https `

    - ` firewall-cmd --reload `

- Kiểm tra nginx

    - http://server_domain_name_or_IP/

- Trước khi tiếp tục, có thể bạn sẽ muốn kích hoạt Nginx khởi động khi hệ thống của bạn khởi động. Để làm như vậy, hãy nhập lệnh sau:

    - ` systemctl enable nginx `

- Thư mục gốc máy chủ mặc định là /usr/share/nginx/html. Các tệp được đặt trong đó sẽ được phục vụ trên máy chủ web của bạn. Vị trí này được chỉ định trong tệp cấu hình khối máy chủ mặc định đi kèm với Nginx, được đặt tại /etc/nginx/conf.d/default.conf.

- Mở file /etc/nginx/conf.d/default.conf để cấu hình virtual host:

<img src="https://imgur.com/piI5LTh.png">

- Trong đó cần lưu ý :

    - servername là tên máy chủ web;

    - root: thư mục chưa source của máy chủ web.

    - Nếu có nhiều trang web thì mình sẽ cấu hình nhiều virtualhost để có thể truy cập đến máy chủ web.

- Cấu hình xong , restart lại nginx và truy cấp máy chủ web trên trình duyệt để kiểm tra.

## 2. Cấu hình Virtualhost trên apache

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

-  Để tạo tệp cấu hình máy chủ ảo:

    - ` vi /etc/httpd/conf.d/vhost.conf `

    - Thêm vào tệp cấu hình đoạn sau:
```
    <VirtualHost *:80>
    ServerName default:80
    DocumentRoot /var/www/html
</VirtualHost>

<VirtualHost *:80>
    ServerName congphan.abc
    DocumentRoot /var/www/html/congphan
    <Directory "/var/www/html/congphan">
        Options Indexes FollowSymLinks
        AllowOverride All
        Order allow,deny
        Allow from all
    </Directory>
</VirtualHost>

<VirtualHost *:80>
    ServerName test.abc
    DocumentRoot /var/www/html/test
    <Directory "/var/www/html/test">
        Options Indexes FollowSymLinks
        AllowOverride All
        Order allow,deny
        Allow from all
    </Directory>
</VirtualHost>
```
<img src="https://imgur.com/XGySv38.png">

- Sau đó restart lại apache và kiểm tra bằng cách truy cập máy chủ web.


