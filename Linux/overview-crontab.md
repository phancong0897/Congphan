# Tổng quan về Crontab

## Mục lục

### [1. Cron là gì?](https://github.com/phancong0897/Congphan/blob/master/Linux/overview-crontab.md#1-cron-l%C3%A0-g%C3%AC-1)

### [2. Cron làm việc như thế nào?](https://github.com/phancong0897/Congphan/blob/master/Linux/overview-crontab.md#2-cron-l%C3%A0m-vi%E1%BB%87c-th%E1%BA%BF-n%C3%A0o)

### [3. Cài đặt crontab](https://github.com/phancong0897/Congphan/blob/master/Linux/overview-crontab.md#3-c%C3%A0i-%C4%91%E1%BA%B7t-crontab-1)

### [4. Cấu trúc của crontab](https://github.com/phancong0897/Congphan/blob/master/Linux/overview-crontab.md#4-c%E1%BA%A5u-tr%C3%BAc-c%E1%BB%A7a-crontab-1)

### [5. Ví dụ cụ thể](https://github.com/phancong0897/Congphan/blob/master/Linux/overview-crontab.md#5-v%C3%AD-d%E1%BB%A5-c%E1%BB%A5-th%E1%BB%83-1)

### [ Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/Linux/overview-crontab.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

### 1. Cron là gì?

- Cron là một tiện ích cho phép thực hiện các tác vụ một cách tự động theo định kỳ, ở chế độ nền của hệ thống. Crontab (CRON TABle) là một file chứa đựng bảng biểu (schedule) của các entries được chạy.

### 2. Cron làm việc thế nào?

- Một cron schedule đơn giản là một text file. Mỗi người dùng có một cron schedule riêng. Crontab files không cho phép bạn tạo hoặc chỉnh sửa trực tiếp với bất kỳ trình text editor nào, trừ phi bạn dùng lệnh crontab.

- Một số lệnh thường dùng:

    - crontab -e: tạo hoặc chỉnh sửa file crontab 

    - crontab --l: hiển thị file crontab 

    - crontab -r: xóa file crontab
- Hầu hết tất cả VPS đều được cài đặt sẵn crontab, tuy nhiên vẫn có trường hợp VPS không có. 

- Nếu bạn sử dụng lệnh crontab -l mà thấy output trả lại -bash: crontab: command not found thì cần tự cài crontab thủ công.

### 3. Cài đặt crontab

- Sử dụng lệnh:

    - ` yum install cronie `

- Start crontab và tự động chạy mỗi khi reboot:

    - ` service crond start `

    - ` chkconfig crond on `

### 4. Cấu trúc của crontab

- Một crontab file có 5 trường xác định thời gian, cuối cùng là lệnh sẽ được chạy định kỳ, cấu trúc như sau:

```

    *     *     *     *     *     command to be executed
    -     -     -     -     -
    |     |     |     |     |
    |     |     |     |     +----- day of week (0 - 6) (Sunday=0)
    |     |     |     +------- month (1 - 12)
    |     |     +--------- day of month (1 - 31)
    |     +----------- hour (0 - 23)
    +------------- min (0 - 59)

```

- Nếu một cột được gán ký tự *, nó có nghĩa là tác vụ sau đó sẽ được chạy ở mọi giá trị cho cột đó.

- Ví dụ:

    - Chạy script 30 phút 1 lần:

        ` 0,30 * * * * command `

    - Chạy script 15 phút 1 lần:

        ` 0,15,30,45 * * * * command `

    - Chạy script vào 3 giờ sáng mỗi ngày:

        ` 0 3 * * * command `
### 5. Ví dụ cụ thể

- Giả sử mình viết một đoạn script sao lưu toàn bộ thư mục /home/domain.com/public_html/ và chuyển file nén .zip vào thư mục /root/ như sau:

```

#!/bin/bash

zip -r /root/backup_domain.com_$(date +"%Y-%m-%d").zip /home/domain.com/public_html/ -q

```

- Script này lưu lại ở đường dẫn /etc/backup.sh (gán quyền execute – chmod +x nếu là bash script).

- Sau đó mình cho script này chạy định kỳ vào 15h thứ Hai và thứ Năm hàng tuần bằng cách tạo một file crontab như sau:

    - ` crontab -e `

- Nhấn o (chữ o) để thêm dòng mới với nội dung:

    - ` 0 15 * * 1,4 sh /etc/backup.sh `

- Để lưu lại và thoát bạn nhấn ESC, rồi gõ vào :wq nhấn Enter.

- Cuối cùng, nhớ khởi động lại cron daemon:

    - ` systemctl restart crond `

### Nguồn tham khảo

- https://hocvps.com/tong-quat-ve-crontab/#:~:text=Cron%20l%C3%A0%20m%E1%BB%99t%20ti%E1%BB%87n%20%C3%ADch,quan%2C%20r%E1%BA%A5t%20d%E1%BB%85%20s%E1%BB%AD%20d%E1%BB%A5ng.

- https://news.cloud365.vn/tim-hieu-ve-cron/
