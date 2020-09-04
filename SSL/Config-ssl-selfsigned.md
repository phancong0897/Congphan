# Tài liệu cáu hình ssl self-signed

- Bước 1: Cài đặt Mod SSL

    - Với Apache, mod_ssl là một Module để hỗ trợ việc cài đặt SSL, ta cần cài đặt Module này:

    ` yum install mod_ssl `

- Bước 2: Tạo chứng chỉ mới

    - Tạo một thư mục để lưu trữ Private key:

    ` mkdir /etc/ssl/private `

    - Gán quyền chỉ có user root mới có quyền chỉnh sửa thư mục này.
    
    ` chmod 700 /etc/ssl/private `

    - Sử dụng openssl để tạo chứng chỉ:

    ` openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/ssl/private/apache-selfsigned.key -out /etc/ssl/certs/apache-selfsigned.crt `

    - Tại đây sẽ ra thông báo ta thêm các thông tin như sau:

    ```

    Country Name (2 letter code) [XX]:
    State or Province Name (full name) []:HN
    Locality Name (eg, city) [Default City]:HN 
    Organization Name (eg, company) [Default Company Ltd]:NH
    Organizational Unit Name (eg, section) []:SP
    Common Name (eg, your name or your server's hostname) []:congphan.xyz
    Email Address []:congpv@nhanhoa.com.vn

    ```

    - Khi sử dụng SSL, ta nên tạo một nhóm Diffie-Hellman, được sử dụng để giao tiếp với Perfect Forward Secrecy . Sử dụng lệnh sau sẽ đợi một lúc:

    ` openssl dhparam -out /etc/ssl/certs/dhparam.pem 2048 `

    - Ta sao chép khóa dhparam.key vừa tạo ở trên vào chứng chỉ self-signed:
    
    ` cat /etc/ssl/certs/dhparam.pem | sudo tee -a /etc/ssl/certs/apache-selfsigned.crt `

- Bước 3: Cài đặt Virtual Host sử dụng SSL

    - vi /etc/httpd/conf.d/congphan.ssl.conf

    ```

    <VirtualHost *:443>
        ServerAdmin congpv@nhanhoa.com.vn
        ServerName congphan.xyz
        ServerAlias www.congphan.xyz

        DocumentRoot /var/www/html

        SSLEngine on
        SSLCertificateFile /etc/ssl/certs/apache-selfsigned.crt
        SSLCertificateKeyFile /etc/ssl/private/apache-selfsigned.key
    </VirtualHost>

    ```

    - Ta trỏ các đường dẫn thư mục của Chứng chỉ một cách chính xác.

    - Sau đó khởi động lại Apache:

    ` systemctl restart httpd `

    - Kiểm tra lại chứng chỉ vừa tạo:

    <img src="https://imgur.com/iLZc1Tf.png">

- Bước 4: Cấu hình Redirect HTTP sang HTTPS

    - Truy cập vào tệp cấu hình non-ssl:

    ` vi /etc/httpd/conf.d/congphan.conf `

    Ta thêm dòng sau:

    ```

    Redirect "/" "https://congphan.xyz/"

    ```

- Dòng này sẽ có tác dụng chuyển hướng HTTP sang HTTPS. Sau đó khởi động lại Apache:

    - systemctl restart httpd