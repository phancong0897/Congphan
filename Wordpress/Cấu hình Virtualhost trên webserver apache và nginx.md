# Tài liệu cấu hình Virtualhost trên Apache và NginX

# Mục lục

## [1. Cấu hình Virtualhost trên NginX](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%E1%BA%A5u%20h%C3%ACnh%20Virtualhost%20tr%C3%AAn%20webserver%20apache%20v%C3%A0%20nginx.md#1-c%E1%BA%A5u-h%C3%ACnh-virtualhost-tr%C3%AAn-nginx-1)

## [2. Cấu hình Virtualhost trên Apache](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%E1%BA%A5u%20h%C3%ACnh%20Virtualhost%20tr%C3%AAn%20webserver%20apache%20v%C3%A0%20nginx.md#2-c%E1%BA%A5u-h%C3%ACnh-virtualhost-tr%C3%AAn-apache-1)

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

## Tạo các Virtual host trên file default.conf

- Thư mục gốc máy chủ mặc định là /usr/share/nginx/html. Các tệp được đặt trong đó sẽ được phục vụ trên máy chủ web của bạn. Vị trí này được chỉ định trong tệp cấu hình khối máy chủ mặc định đi kèm với Nginx, được đặt tại /etc/nginx/conf.d/default.conf.

- Mở file /etc/nginx/conf.d/default.conf để cấu hình virtual host:

<img src="https://imgur.com/piI5LTh.png">

- Trong đó cần lưu ý :

    - servername là tên máy chủ web;

    - root: thư mục chưa source của máy chủ web.

    - Nếu có nhiều trang web thì mình sẽ cấu hình nhiều virtualhost để có thể truy cập đến máy chủ web.

- Cấu hình xong , restart lại nginx và truy cấp máy chủ web trên trình duyệt để kiểm tra.

## Tạo trên nhiều file Virtualhost

- TẠO CÂY THƯ MỤC CHO WEBSITE MỚI

    - Có 2 website là congphan.xyz và qatest.xyz

    - Tạo cây thư mục cho congphan.xyz

        - ` mkdir -p /var/www/congphan/public_html `

        - ` mkdir -p /var/www/congphan/logs `

    - Tạo cây thư mục cho qatest.xyz

        - ` mkdir -p /var/www/qa/public_html `

        - ` mkdir -p /var/www/qa/logs `

- Phân quyền quản lý thư mục

    - ` chown -R nginx:nginx /var/www `

- Dùng Text editor VI tạo file config virtual host cho congphan.xyz  vi /etc/nginx/conf.d/congphan.conf có nội dung như sau

```

server {

listen  80;

server_name congphan.xyz;

access_log /var/www/congphan/logs/access.log;

error_log /var/www/congphan/logs/error.log;

root /var/www/congphan/public_html; 

index index.html index.htm index.php;

}

``` 
- - Dùng Text editor VI tạo file config virtual host cho qatest.xyz  vi /etc/nginx/conf.d/qatest.conf có nội dung như sau

```

server {

listen  80;

server_name qatest.xyz;

access_log /var/www/qatest/logs/access.log;

error_log /var/www/qatest/logs/error.log;

root /var/www/qatest/public_html; 

index index.html index.htm index.php;

}

``` 
- Trên đây là thông số cấu hình tối thiểu nhất để tạo virtual host trên nginx. Toàn bộ virtualhost được bao trong block server.

    - Listen: xác định IP và port lắng nghe request đến website, port 80 là giá trị mặc định.
 
    - server_name: gán hostname cho website, ở đây mình đặt hostname là demo1.thuysys.com

    - access_log: nếu đã đọc phần I cơ bản về nginx bạn sẽ thấy chỉ thị access_log đã được khai báo trước đó trong block http (block cha của block server). Việc khai báo trong block server sẽ vô hiệu hóa chỉ thị access_log trong http và quy định lại đường dẫn đến file ghi log truy cập mới. Theo mình mỗi virtual host bạn cấu hình log riêng để dễ quản lý, giờ bạn có thể xóa chỉ thị access_log ngoài block http đi cho đỡ rối.
    
    - error_log: quy định đường dẫn đến file ghi log lỗi, cũng tương tự access_log bạn xem lại file /etc/nginx/nginx.conf để cài đặt log cho phù hợp.

    - root: chỉ định đường dẫn đến source code của trang web.

    - index: các file nginx sẽ xử lý.
- Tiếp theo bạn lưu lại file cấu hình, tạo 2 file test để kiểm tra 

    - Tạo file test cho congphan.xyz

        - vi /var/www/congphan/public_html/index.html và thêm html vào

    ```
    <html>
    
    <head>
    
    <title>congphan.xyz</title>
    
     </head>
    
    <body>
    
    <p>xin chao day la website congphan.xyz</p>
    
     </body>
    
    </html>

    ```
    - Tạo file test cho qatest.xyz

        - vi /var/www/qatest/public_html/index.html và thêm html vào

            ```
    <html>
    
    <head>
    
    <title>congphan.xyz</title>
    
     </head>
    
    <body>
    
    <p>xin chao day la website congphan.xyz</p>
     
     </body>
    
    </html>

    ```
- Tiếp theo bạn restrat lại nginx và truy cập trang web để kiểm tra mình đã cấu hình đúng chưa, Nếu kết quả trả về giống trang index.html thì bạn đã cấu hình đúng.

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
- Sau đó restart lại apache và kiểm tra bằng cách truy cập máy chủ web.


