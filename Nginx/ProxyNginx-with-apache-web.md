# Tài liệu cài đặt NginX làm proxy cho Apache Webserver

## Mục lục

### [1. Mô hình](https://github.com/phancong0897/Congphan/blob/master/Nginx/ProxyNginx-with-apache-web.md#1-m%C3%B4-h%C3%ACnh-1)

### [2. Triển khai](https://github.com/phancong0897/Congphan/blob/master/Nginx/ProxyNginx-with-apache-web.md#2-tri%E1%BB%83n-khai-1)

### [2.1 Cài đặt Apache](https://github.com/phancong0897/Congphan/blob/master/Nginx/ProxyNginx-with-apache-web.md#21-c%C3%A0i-%C4%91%E1%BA%B7t-apache-1)

### [2.2 Cài đặt NginX](https://github.com/phancong0897/Congphan/blob/master/Nginx/ProxyNginx-with-apache-web.md#22-c%C3%A0i-%C4%91%E1%BA%B7t-nginx-1)

### [2.3 Cấu hình nginx làm proxy](https://github.com/phancong0897/Congphan/blob/master/Nginx/ProxyNginx-with-apache-web.md#23-c%E1%BA%A5u-h%C3%ACnh-nginx-l%C3%A0m-proxy-1)

### [2.4 Cấu hình caching cho nginx](https://github.com/phancong0897/Congphan/blob/master/Nginx/ProxyNginx-with-apache-web.md#24-c%E1%BA%A5u-h%C3%ACnh-caching-cho-nginx-1)

### [ Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/Nginx/ProxyNginx-with-apache-web.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Mô hình 

<img src="https://imgur.com/nfyAGid.png">

- Trong mô hình này tôi sẽ thực hiện các cấu hình sau:

    - Cài đặt máy webs server sử dụng Apache, sau đó up các file ảnh hoặc tạo ra một site html đơn giản (có ảnh hoặc file js)

    - Cài đặt máy Nginx làm chức năng reverse và caching.

    - Cấu hình file host ở các máy của người dùng với domain websitetest01.com trỏ về IP của máy chủ nginx.

    - Người dùng mở trình duyệt hoặc dùng lệnh curl với tùy chọn -I để kiểm tra xem proxy có hoạt động hay không.

    - Người dùng truy cập nhiều lần vào web hoặc dùng curl để kiểm tra xem đã cache được hay chưa?

## 2. Triển khai

### 2.1 Cài đặt Apache

- Cài đặt các gói cơ bản

    - ` yum update -y `

    - ` yum install -y epel-release `

    - ` yum install -y wget byobu `

- Cài đặt Apache

    - ` yum install -y httpd `

- Khởi động httpd và kích hoạt http khi reboot hệ điều hành

    - ` systemctl start httpd `

    - ` systemctl enable httpd `

- Kiểm tra hoạt động của httpd

    - ` systemctl status httpd `

   <img src="https://imgur.com/plG98iz.png">

- Tạo một trang web có chứa ảnh hoặc các file tĩnh

    - Dùng lệnh dưới để tải file index.html có chứa nội dung ảnh về.

    ```
    wget -O /var/www/html/index.html https://gist.githubusercontent.com/congto/359e04f735162a987daf58d3f8d44fb6/raw/51ccab89265bff5717084af1212640dae6bbfa92/indext.html

    ```

- Truy cập website với địa chỉ IP của máy webserver

Mở trình duyệt và truy cập vào địa chỉ ở dưới, nội dung web sẽ hiển thị trên màn hình là ok.

http://172.16.2.230

<img src="https://imgur.com/uxbXKjX.png">

### 2.2 Cài đặt Nginx

- Đăng nhập với quyền root và thực hiện cài đặt các gói bổ trợ

    - ` yum update -y `

    - ` yum install -y epel-release `

    - ` yum install -y wget byobu `

- Cài đặt Nginx, ở đây tôi lựa chọn cài Nginx từ package

    - ` yum install -y nginx `

- Khởi động nginx và kích hoạt chế độ khởi động cùng OS.

    - ` systemctl start nginx `

    - ` systemctl enable nginx `

- Kiểm tra hoạt động của nginx

    - ` systemctl status nginx `

    <img src="https://imgur.com/K35vcfl.png">

### 2.3 Cấu hình nginx làm proxy

- Trước khi cầu hình Nginx làm reverse proxy thì việc truy cập vào website của máy 172.16.2.235 sẽ thông qua địa chỉ IP hoặc domain được trỏ ở file host. Trong phần này tôi sẽ khai báo cấu hình cho Nginx làm chức năng reverse proxy cho máy web server. Tức là người dùng sẽ truy cập vào IP hoặc domain được trỏ về IP 172.16.2.235, lúc này Nginx sẽ làm nhiệm vụ điều phối truy cập vào máy 172.16.2.230 để lấy nội dung trang web trả về cho người dùng.

- Các bước chi tiết như sau

    - Sao lưu file cấu hình mặc định của nginx

        ` cp /etc/nginx/nginx.conf /etc/nginx/nginx.conf.bka `

    - Di chuyển vào thư mục /etc/nginx/conf.d/để khai báo file cấu hình làm nhiệm vụ reverse proxy.
        
        ` cd /etc/nginx/conf.d/ `

    - Tạo file với đuôi mở rộng là .conf

        ` vi /etc/nginx/conf.d/proxy.conf `
    
    Nội dung của file proxy.conf sẽ là:

    ```
    server {
    listen 80;
    server_name congphan.abc;
    access_log /var/log/nginx/access.log;
    error_log /var/log/nginx/error.log;
    
    location / {
        proxy_pass http://172.16.2.235:80/;
        # Input any other settings you may need that are not already contained in the default snippets.
        }
    }

    ```
    - Lưu ý trong dòng proxy_pass ta sẽ khai báo về địa chỉ của máy cài httpd.

    - Sau khi khai báo xong, kiểm tra lại xem khai báo này đã đúng chưa bằng lệnh

        ``` nginx -t ```

        <img src="https://imgur.com/GKfsVjn.png">

    - Tiến hành khởi động lại nginx

        ``` systemctl restart nginx ```

    - Lưu ý, nếu làm lab các bạn nhớ trỏ domain thông qua file host. Xin lưu ý rằng mặc dù website được đặt trên máy chủ 172.16.2.230 nhưng domain sẽ được trỏ về 172.16.2.235, đây chính là việc sử dụng tính năng reverse proxy của nginx để điều hướng các kết nối của người dùng và máy chủ gốc (origin).

    - Kiểm tra nội dung web và xem revert proxy hoạt động hay chưa bằng cách truy cập vào địa chỉ web1cloud365.vn. Trong ảnh này tôi sử dụng thêm mode của phím F12 để có thể show được các kết nối từ client tới webserver. Ta quan sát thấy nội dung của web được hiển thị và địa chỉ IP của Nginx reserver proxy.

    <img src="https://imgur.com/5nSmrVU.png">

### 2.4 Cấu hình caching cho nginx

- Mở file proxy.conf ở phần trước và thêm các dòng này vào phần đầu của file (nằm ngoài block server)

    ``` 
    proxy_cache_path /var/lib/nginx/cache levels=1:2 keys_zone=backcache:8m max_size=50m;

    proxy_cache_key "$scheme$request_method$host$request_uri$is_args$args";

    proxy_cache_valid 200 302 10m;

    proxy_cache_valid 404 1m;

    ```
- Tiếp tục thêm dòng dưới vào directive localtion

    ```

    proxy_cache backcache;
    
    add_header X-Proxy-Cache $upstream_cache_status;

    ```
- Chúng ta sẽ có file conf hoàn chỉnh sau đấy:

<img src="https://imgur.com/MMwotjU.png">

- Khởi động lại nginx:

    ``` systemctl restart nginx ```

-  Sử dụng trình duyệt để kiểm tra:

    - Lần 1:

    <img src="https://imgur.com/UglqGUL.png">

    - Lần 2:

    <img src="https://imgur.com/ZLdF6Rg.png">

- Truy cập vào trong máy chủ nginx và kiểm tra thư mục cache, ta sẽ thấy có các file được hash, đó chính là nội dung của web đã được cache, sử dụng lệnh: du -ah /var/lib/nginx/cache/

```
[root@localhost ~]# du -ah /var/lib/nginx/cache/

4.0K    /var/lib/nginx/cache/7/ef/235fe61149a0e9c2d9c6dd137cec5ef7

4.0K    /var/lib/nginx/cache/7/ef

4.0K    /var/lib/nginx/cache/7

4.0K    /var/lib/nginx/cache/

```
## Nguồn tham khảo:

https://gist.github.com/congto/bb0a2c37348ced72284659a86cd24052