# Hướng dẫn tạo Swap trên Ubuntu 18.04

### 1. Kiểm tra Swap

Việc đầu tiên cần làm là kiểm tra xem hệ thống của các bạn đã có Swap hay chưa bằng cách chạy lệnh sau:
swapon --show

Nếu bạn không thấy kết quả nào được trả về, điều đó có nghĩa là hệ thống của bạn chưa có swap. Ngược lại nếu bạn thấy kết quả như ảnh dưới đây nghĩa là hệ thống của bạn đã có Swap và bạn không cần làm gì thêm

![Imgur](https://imgur.com/Oo2CLKV.png)

### 2. Tạo Swap

- Bước 1: Tạo swapfile

Chạy lệnh sau để tạo swapfile trên hệ thống Ubuntu 18.

sudo dd if=/dev/zero of=/swapfile bs=1024 count=2048k

Bạn có thể thay thế 2048k bằng dung lượng swapfile muốn tạo, ví dụ: 1024K sẽ là 1GB

- Bước 2: Phân quyền cho swapfile

Để đảm bảo rằng chỉ người dùng root mới có thể đọc và ghi vào swap các bạn chạy lần lượt hai lệnh dưới đây

    ```
    sudo chown root:root /swapfile

    sudo chmod 600 /swapfile

    ```

- Kích hoạt swap

Tiếp theo, các bạn chạy lệnh sau để tạo phân vùng swap:

    ` sudo mkswap /swapfile `

Chạy lệnh sau để kích hoạt swap:

    ` sudo swapon /swapfile `

- Thiết lập tự kích hoạt swapfile mỗi khi khởi động lại hệ thống

    - Để phân vùng swap không bị mất mỗi khi khởi động lại hệ thống các bạn cần chỉnh sửa file /etc/fstab bằng cách chạy lệnh sau:

` sudo echo '/swapfile   none    swap    sw    0   0' | sudo tee -a /etc/fstab `

    - Để kiểm tra xem swap đã được kích hoạt hay chưa hãy dùng lệnh sau:

    ` free -h ` 
    

