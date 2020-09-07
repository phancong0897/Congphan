# Tài liệu hướng dẫn chuyển hướng trang web từ non www sang www dùng file .htaccess

- Mặc định trong Apache sẽ có một file tên là .htaccess. File .htaccess này có tác dụng thay đổi truy cập mà không làm ảnh hưởng đến cấu hình.

- Nếu như trong Apache không có sẵn .htaccess ta cũng có thể tạo một file

- Trước khi bắt đầu, ta phải cho phép Apache đọc file .htaccess. Truy cập vào file:

    - vi /etc/httpd/conf/httpd.conf

    - Tìm đến đoạn cấu hình cho thư mục /var/www/html và thay đổi cấu hình AllowOverride None thành AllowOverride All.

### Cấu hình từ non-www sang www

- Sử dụng câu lệnh vi để vào file .htaccess
    
    - vi /var/www/html/.htaccess

    - Thêm nội dung sau vào file

```

RewriteEngine On
RewriteCond %{HTTP_HOST} ^congphan.abc [NC]
RewriteRule ^(.*)$ http://www.congphan.abc/$1 [L,R=301]

```

- Sau đó lưu lại, sử dụng lệnh curl để kiểm tra:

    ```

    [root@localhost ~]# curl -I http://congphan.abc
    HTTP/1.1 301 Moved Permanently
    Date: Mon, 07 Sep 2020 03:00:31 GMT
    Server: Apache/2.4.6 (CentOS) PHP/7.2.33
    Location: http://www.congphan.abc/
    Content-Type: text/html; charset=iso-8859-1

    ```

- Ta thấy Location: http://www.congphan.abc/ là kết quả mong muốn

### Cấu hình từ www sang non-www

- Vào file .htaccess

    - vi /var/www/html/.htacces

    - Thêm dội dung sau vào file

    ```
    
    RewriteEngine On
    RewriteCond % ^www.congphan.abc
    RewriteRule (.*) http://congphan.abc/$1 [R=301,L]

    ```

- Sử dụng curl để kiểm tra trang web

    ```
    
    [root@localhost ~]# curl -I http://www.congphan.abc
    HTTP/1.1 301 Moved Permanently
    Date: Mon, 07 Sep 2020 03:04:25 GMT
    Server: Apache/2.4.6 (CentOS) PHP/7.2.33
    Location: http://congphan.abc/
    Content-Type: text/html; charset=iso-8859-1

    ```

- Ta thấy Location: http://congphan.abc/ là kết quả mong muốn

### direct tên miền www tới non-www giao thức https bằng file .htaccess

- Sử dụng vi để chỉnh sửa file .htaccess

    - vi /var/www/html/.htaccess

    - Thêm nội dung sau vào file

    ```
    
    RewriteEngine on
    RewriteCond %{HTTPS} off
    RewriteRule (.*) https://%{HTTP_HOST}:443%{REQUEST_URI} [END]

    RewriteEngine On
    RewriteCond %{HTTP_HOST} ^www.congphan.abc [NC]
    RewriteRule ^(.*)$ https://congphan.abc/$1 [L,R=301]

    ```

- Sử dụng curl để kiểm tra trang web

    ```
    
    [root@localhost ~]# curl -I http://www.congphan.abc
    HTTP/1.1 302 Found
    Date: Mon, 07 Sep 2020 03:08:11 GMT
    Server: Apache/2.4.6 (CentOS) PHP/7.2.33
    Location: https://www.congphan.abc:443/
    Content-Type: text/html; charset=iso-8859-1
    
    ```

- Ta thấy Location: https://www.congphan.abc/ là kết quả mong muốn