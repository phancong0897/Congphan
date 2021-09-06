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

    - Để try cập container PHP dùng lệnh:

    `docker exect -it c-php bash`

    ![Imgur](https://imgur.com/hpdGGfp.png)

    - Thêm file info.php vào thư mục /home/dulieu  và chạy lệnh ` php info.php ` để kiểm tra

    ![Imgur](https://imgur.com/198pTZu.png)

    Như vậy là đã cài đặt xong container php 7.4-fpm

