# Thực hiện cài đặt wordpress trên container

- Mô hình như sau:

Ta sẽ tạo 3 container chạy 3 dịch vụ php, httpd, mysqld. 3 container này sẽ kết nối đến với nhau để chạy code wordpress.

### Tạo một mạng ảo để các container có thể kết nối được với nhau

` docker network create --driver bridge WP `

![Imgur](https://imgur.com/E1ZPDhz.png)

### Tạo container chạy PHP-FPM

- Dowload images php 7.3-fpm

    ` docker pull php:7.4-fpm `

    ![Imgur](https://imgur.com/MPcFo5C.png)

- Chaỵ lệnh tạo container với images php

    ` docker run -d -v /home/code/:/home/dulieu --network WP --name c-php -h php php:7.4-fpm `

- Các tham số tạo, chạy container như đã biết trong phần trước, ở đây cụ thể là:

    - -d : container sau khi tạo chạy luôn ở chế độ nền.

    - --network WP : tạo container ở dải mạng WP mà ta tạo trước đó.

    - --name c-php : container tạo ra và đặt luôn cho nó tên là c-php (Tương tác với container dễ đọc hơn khi sử dụng tên thay vì phải sử dụng mã hash id, nếu không đặt tên thì docker sinh ra tên ngẫu nhiên).
    - -h php đặt tên HOSTNAME của container là php
    
    - -v /home/code/:/home/dulieu thư mục trên máy host /home/code/ được gắn vào container ở đường dẫn /home/dulieu.
    
    - php:7.4-fpm là image khởi tạo ra container, nếu image này chưa có nó tự động tải về.

    - Muốn tương tác sử dụng câu lệnh trong container PHP dùng lệnh:

    `docker exect -it c-php bash`

    ![Imgur](https://imgur.com/hpdGGfp.png)

    - Thêm file info.php vào thư mục /home/dulieu  và chạy lệnh ` php info.php ` để kiểm tra

    ![Imgur](https://imgur.com/198pTZu.png)

    Như vậy là đã cài đặt xong container php 7.4-fpm

![Imgur](https://imgur.com/27cIC6j)
    
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

