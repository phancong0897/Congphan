# Tài liệu hướng dẫn migarate wordpress từ máy này qua máy khác

### Mục lục

### [1. Chuẩn bị](https://github.com/phancong0897/Congphan/blob/master/Wordpress/Migrate-wordpress.md#1-chu%E1%BA%A9n-b%E1%BB%8B-1)

### [2. Tiến hành migrate](https://github.com/phancong0897/Congphan/blob/master/Wordpress/Migrate-wordpress.md#2-ti%E1%BA%BFn-h%C3%A0nh-migrate-1)

### [2.1 Trên máy Cen1](https://github.com/phancong0897/Congphan/blob/master/Wordpress/Migrate-wordpress.md#21-tr%C3%AAn-m%C3%A1y-cen1-1)

### [2.2 Trên máy Cen2](https://github.com/phancong0897/Congphan/blob/master/Wordpress/Migrate-wordpress.md#22-tr%C3%AAn-m%C3%A1y-ceb2)

### 1. Chuẩn bị

- Ta chuẩn bị 2 máy Cen1 và Cen2, trong đó:

    - Cen1 có địa chỉ IP: 172.16.2.230 đã cài các dịch vụ:

        + Webserver apache.

        + MariaDB

        + PHP 7.2.

        + Wordpress đã cấu hình.

    - Cen2 có địa chỉ IP: 172.16.2.235 đã cài các dịch vụ:

        + Webserver apache.

        + MariaDB.

        + PHP 7.2.

### 2. Tiến hành migrate

### 2.1 Trên máy Cen1

### Backup Source code WordPress

- Tạo thư mục lưu trữ bản backup bằng câu lệnh:

    - mkdir -p /opt/backup/backup-code-wp

- Tạo file Script

    - vi backup-code.sh

    - Thêm nội dung dưới vào file vừa tạo:

    ```

    SRCDIR="/var/www/html/*"
    DESTDIR="/opt/backup/backup-code-wp/"
    FILENAME=wpcode-$(date +%-Y%-m%-d)-$(date +%-T).tgz
    tar -P --create --gzip --file=$DESTDIR$FILENAME $SRCDIR

    ````

- Gán quyền thực thi cho tệp Script vừa tạo

    - chmod +x /root/backup-code.sh

- Chạy Script vừa tạo

    - sh /root/backup-code.sh

- Sau khi tiến hành backup, file back up sẽ có dạng wpsourcecode-202039-08:09:40.tgz

### Backup Database

- Tạo thư mục backs up sql bằng câu lệnh:

    - mkdir -p /opt/backup/backup-sql

- Tạo file Script cho Cơ sở dữ liệu

    - vi backup-sql.sh

    - Ta thêm nội dung sau:


        ```
        NOW=$(date +"%Y-%m-%d-%H:%M")
        BACKUP_DIR="/opt/backup/backup-sql"

        DB_USER="root"
        DB_PASS="Quanganh1234"
        DB_NAME="wordpress"
        DB_FILE="wordpress.$NOW.sql"

        mysqldump -u$DB_USER -p$DB_PASS $DB_NAME > $BACKUP_DIR/$DB_FILE

        ```
- Gán quyền thực thi cho tệp Script vừa tạo

    - chmod +x /root/backup-sql

- Chạy Script vừa tạo

    - sh /root/backup-sql.sh

- Tệp được lưu lại sau khi chạy Script sẽ có dạng wordpress.2020-09-03-11:32.sql

### 2.2 Trên máy Ceb2

- Tạo cơ sở dữ liệu :

    - mysql -u root -p

    - CREATE DATABASE wordpress;

    - CREATE USER user@localhost IDENTIFIED BY 'pass';

    - GRANT ALL PRIVILEGES ON wordpress.* TO user@localhost IDENTIFIED BY 'pass';

    - FLUSH PRIVILEGES;

    - exit.

- Tiến hành chuyển cơ sở dữ liệu từ Cen1 sang Cen2  

    - mysql -u root -p wordpress < /backup/wordpress.sql

    - Trong đó /backup/wordpress.sql là đường dẫn của thư mục chứa file Backup Cơ sở dữ liệu.

- Sau khi Restore thành công ta khởi động lại MySQL hoặc MariaDB.

    - systemctl restart httpd

- Tiến hành copy soucode của wordpress ở Cen1 sang Cen2 bằng winscp. Lưu ý copy suorce từ Cen1 sang thư mục var/www/html của Cen2.

- Cấu hình WordPress trên hệ thống mới

    - Với việc đặt tên DB mới là wordpress thì ta không cần phải thay đổi lại DB_HOST trong cấu hình wordpress. Chúng ta cần phải thay đổi lại đường dẫn vì hiện tại cấu hình WordPress vẫn đang nhận cấu hình của hệ thống cũ.

    - Ta thay đổi cấu hình tại file wp-config.php. Ta thêm 2 dòng sau:

        ```
        define('WP_HOME','http://172.16.2.235');
        define('WP_SITEURL','http://172.16.2.235');

        ```
    - Trong đó 172.16.2.235 là địa chỉ IP của máy Cen2

- Sau đó lưu lại và khởi động lại httpd:

    - systemctl restart httpd

- Kiểm tra lại hoạt động của WordPress thành công:

<img src="https://imgur.com/SWQ0xbZ.png">


