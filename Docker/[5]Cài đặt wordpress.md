# Thực hiện cài đặt wordpress trên container

- Mô hình như sau:

Ta sẽ tạo 3 container chạy 3 dịch vụ php, httpd, mysqld. 3 container này sẽ kết nối đến với nhau để cài đặt và sử dụng code wordpress.

### Tạo một mạng ảo để các container có thể kết nối được với nhau

` docker network create --driver bridge WP `

![Imgur](https://imgur.com/E1ZPDhz.png)

### Tạo container chạy PHP-FPM

- Dowload images php 7.3-fpm

    ` docker pull php:7.3-fpm `

    ![Imgur](https://imgur.com/h05evtq.png)

- Chaỵ lệnh tạo container với images php

    ` docker run -d --name c-php -h php -v /home/wp:/home/phpcode php:7.3-fpm --network WP `

- Các tham số tạo, chạy container như đã biết trong phần trước, ở đây cụ thể là:

    - -d : container sau khi tạo chạy luôn ở chế độ nền.

    - --name c-php : container tạo ra và đặt luôn cho nó tên là c-php (Tương tác với container dễ đọc hơn khi sử dụng tên thay vì phải sử dụng mã hash id, nếu không đặt tên thì docker sinh ra tên ngẫu nhiên).

    - -h php đặt tên HOSTNAME của container là php
    
    - -v /home/wp:/home/phpcode thư mục trên máy host /home/wp/ được gắn vào container ở đường dẫn /home/phpcode.
    
    - php:7.3-fpm là image khởi tạo ra container, nếu image này chưa có nó tự động tải về.

    - Muốn tương tác sử dụng câu lệnh trong container PHP dùng lệnh:

    `docker exect -it c-php bash`

    ![Imgur](https://imgur.com/hpdGGfp.png)

    - Thêm file info.php vào thư mục /home/phpcode  và chạy lệnh ` php info.php ` để kiểm tra

    ![Imgur](https://imgur.com/ETk9Pil.png)

    Như vậy là đã cài đặt xong container php 7.3-fpm
    
Như hình trên, một container có tên c-php được tạo ra, chạy với lệnh script mặc định docker-php-entrypoint, dịch vụ PHP-FPM chạy và lắng nghe ở cổng 9000.

 ` Chú ý: trong một số image, Entrypoint chạy script kết thúc ngay dẫn đến container kết thúc ngay, kể cả khi khởi chạy lại container bằng docker start thì nó cũng kết thúc luôn. Những loại image như vậy, cần tiến trình hoạt động dài để giữ cho container sống. Để khắc phục, thêm vào nó tham số -i ở lệnh docker run `


- Có sẵn một script có tên là docker-php-ext-configure và docker-php-ext-install để bạn cấu hình, cài các Extension cho PHP, ví dụ nếu muốn có thêm OpCache thì gõ:

    ` docker-php-ext-configure opcache --enable-opcache     && docker-php-ext-install opcache `

- Cài thêm mysqli - extension để kết nối PHP đến MySQL

    ` docker-php-ext-configure opcache --enable-mysqli      && docker-php-ext-install mysqli `
    
- Cài thêm pdo_mysql - extension để kết nối PHP đến MySQL với thư viện PDO

    ` docker-php-ext-configure opcache --enable-pdo_mysql   && docker-php-ext-install pdo_mysql `

- File php.ini mà PHP nạp nằm ở /usr/local/etc/php, để có file ini này gõ lệnh:

    ` cp /usr/local/etc/php/php.ini-production /usr/local/etc/php/php.ini `

### Tạo container chạy webserver apache và kết nối đến container c-php ta tạo trước đó

- Trước tiên pull images apache về

    ` docker pull httpd `

- Khi tạo container httpd, file cấu hình httpd.conf sẽ nằm ở thư mục    /usr/local/apache2/conf/httpd.conf , chính vì thế ta cần tạo trước container để copy file cấu hình ra máy host

    ` docker -rm -v /home/code/:/home/httpd httpd cp /usr/local/apache2/htdocs/httpd.conf /home/httpd `

    Tham số -rm ở đây, khi tạo container xong, nếu ta thoát container thì container sẽ tự động xóa.

- Sau khi copy file httpd.conf xong ta tiến hành chỉnh sửa file httpd.conf để cấu hình cũng như kết nối đến container c-php như sau:

    - bỏ comment ở hai dòng sau trong file httpd.conf

    ``` 
    
    LoadModule proxy_module modules/mod_proxy.so
    LoadModule proxy_fcgi_module modules/mod_proxy_fcgi.so

    ```

     - Thêm vào cuối file httpd.conf sau để kết nối đến container c-php

    ` AddHandler "proxy:fcgi://c-php:9000" .php `

     Ở đây ta có thể thay thế c-php bằng IP của container.

- Sau khi chỉnh sửa xong ta tiến hành tạo container c-httpd như sau:

    ` docker run -di -p 8080:80 -p 443:443 --name c-httpd -h httpd -v /home/wp:/home/phpcode -v /home/wp/httpd.conf:/usr/local/apache2/conf/httpd.conf httpd `

- Với:

    - -p 8080:80 truy cập port 80 của container bằng port 8080 của host

    - -v /home/wp:/home/phpcode thư mục trên máy host /home/wp/ được gắn vào container ở đường dẫn /home/phpcode.

    - -v /home/wp/httpd.conf:/usr/local/apache2/conf/httpd.conf file httpd.conf trên máy host ở thư mục /home/wp được gắn vào container ở đường dẫn /usr/local/apache2/conf/httpd.conf httpd .

- Sau khi tạo xong, tiến hành connect network của 2 container này nằm trên cùng 1 dải mạng:

```

docker network connect WP c-httpd
docker network connect WP c-php

WP là mạng bridge ta đã tạo trước đó

```

- Vậy ta đã kết nối xong container webserver với container php, ta kiểm tra bằng cách, dùng địa chỉ IP:8080/info.php

    ![Imgur](https://imgur.com/HTxRV8b.png)

Kết quả trả về file info.php ở trong thư mục /home/phpcode/www/info.php

### Tạo Mysql server

- Pull images Mysql

    ` docker pull mysql `

- Cũng giống như httpd, ta cần copy và chỉnh sửa file cấu hình của mysql nằm ở đường dẫn /etc/mysql/my.cnf, ta chạy lệnh sau:

    ` docker run --rm -v /home/wp:/home/phpcode mysql cp /etc/mysql/my.cnf /home/phpcode `

- Thêm dòng sau vào file my.cnf

    ```

    [mysqld]
    default-authentication-plugin=mysql_native_password 
    
    ```

- Tiến hành tạo container mysql:

    ` docker run -it --name c-mysql -h mysql -v /home/wp/data:/var/lib/mysql -v /home/wp/my.cnf:/etc/mysql/my.cnf -e MYSQL_ROOT_PASSWORD=abc123 mysql `

    Trong đó:

    - -e MYSQL_ROOT_PASSWORD=abc123 đặt password cho user root, quản trị mysql là abc123

    - -v /home/wp/data:/var/lib/mysql nơi lưu trữ các database là ở thư mục /home/wp/data của máy Host

    - -v /home/wp/my.cnf:/etc/mysql/my.cnf ánh xạ file my.cnf ở máy host lên container

- Tạo xong truy cập container c-mysql để tạo database:

    ```

    docker exec -it c-mysql bash                       #nhảy vào container

    mysql -uroot -pabc123                              #kết nối vào MySQL Server

    #Từ dấu nhắc MySQL, Tạo một user tên testuser với password là testpass

    CREATE USER 'wp'@'%' IDENTIFIED BY 'wp';

    #Tạo db có tên db_testdb
    create database wp;

    #Cấp quyền cho user testuser trên db - db_testdb
    GRANT ALL PRIVILEGES ON wp.* TO 'wp'@'%';
    flush privileges;

    show database;            #Xem các database đang có, kết quả có db bạn vừa tạo
    exit;                     #Ra khỏi MySQL Server

    ```

- Truy cập container c-httpd để thêm vhost vào file httpd.conf, thêm vào cuối file:

    ```

    <VirtualHost *:80>
        ServerName phancong.com
        ServerAdmin anhcong97ht@gmail.com
        DocumentRoot /home/phpcode/www/
        CustomLog /dev/null combined
        #LogLevel Debug
        ErrorLog /home/phpcode/www/error.log
        <Directory /home/phpcode/www/>
            Options -Indexes -ExecCGI +FollowSymLinks -SymLinksIfOwnerMatch
            DirectoryIndex index.php
            Require all granted
            AllowOverride None
        </Directory>
    </VirtualHost>

    ```

- Tiến hành restart lại container c-httpd

    `docker restart c-httpd `

- Truy cập phancong.com:8080 để tiến hành cài đặt wordpress

    Hiển thị như hình ảnh là các bạn đã cấu hình thành công.

    ![Imgur](https://imgur.com/yIfEonV.png)


