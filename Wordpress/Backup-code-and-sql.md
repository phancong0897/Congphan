# Tạo Scripts Backup Source Code và Database WordPress

### Backup Source code WordPress

- Tạo thư mục lưu trữ bản backup bằng câu lệnh:

    - ` mkdir -p /opt/backup/backup-code-wp `

-  Tạo file Script

    - ` vi backup-code.sh `

    - Thêm nội dung dưới vào file vừa tạo:

    ``` 
    SRCDIR="/var/www/html/*"
    DESTDIR="/opt/backup/backup-code-wp/"
    FILENAME=wpcode-$(date +%-Y%-m%-d)-$(date +%-T).tgz
    tar -P --create --gzip --file=$DESTDIR$FILENAME $SRCDIR

    ```

- Gán quyền thực thi cho tệp Script vừa tạo

    - ` chmod +x /root/backup-code.sh `

- Chạy Script vừa tạo

    - ` sh /root/backup-code.sh `

Sau khi tiến hành backup, file back up sẽ có dạng wpsourcecode-202039-08:09:40.tgz

### Backup Database

- Tạo thư mục backs up sql bằng câu lệnh:

    - ` mkdir -p /opt/backup/backup-sql `

- Tạo file Script cho Cơ sở dữ liệu

    - ` vi backup-sql.sh `

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

    - ` chmod +x /root/backup-sql `

- Chạy Script vừa tạo
    
    - ` sh /root/backup-sql.sh `

Tệp được lưu lại sau khi chạy Script sẽ có dạng wordpress.2020-09-03-11:32.sql

### Lưu vào Crontab để tự động chạy Scripts hàng ngày

- crontab -e

    ```

    0 8 * * * /root/backup-code
    0 8 * * * /root/backup-sql

    ```

Hai Scripts vừa tạo sẽ được chạy tự động lúc 8 giờ sáng mỗi ngày.

- Để kiểm tra các tiến trình chạy trong crontab thì ta dùng lệnh 

    - ` crontab -l `