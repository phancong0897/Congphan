# Tạo user có quyền sudo

# Mục lục

## [1. Log in vào your server với tư cách root user]()

## [2. Dùng lệnh adduser để thêm user mới vào hệ thống]()

## [3. Dùng lệnh usermod để thêm user vừa tạo vào nhóm wheel]()

## [4. Kiếm tra truy cập sudo trên user mới]()

## 1. Log in vào your server với tư cách root user.

- ` ssh root@server_ip_address `

## 2. Dùng lệnh adduser để thêm một user mới vào hệ thống

- ` adduser username ` 

- Dùng lệnh passwd để câp nhật mật khẩu cho user vừa tạo.

    - ` passwd username `

<img src="https://imgur.com/0cDqPyj.png">

## 3. Dùng lệnh usermod để thêm user vừa tạo vào nhóm wheel .

- ` usermod -aG wheel username ` 

- Trên Centos 7, mặc định các user trong wheel được cấp quyền sudo.

## 4. Kiểm tra truy cập sudo trên user mới:

- Dùng lệnh su để chuyển tới user mới được tạo : 

    - ` su - username `

- Gõ sudo  phía trước lệnh bạn muốn chạy với quyền superuser:

    - ` sudo command_to_run `

- Ví dụ, lệnh hiển thị các thành phần của thư mục /root, vốn chỉ có thể thực hiện với root user:

    - ` sudo ls -la /root `

Lần đầu tiên sử dụng lệnh sudo , bạn sẽ phải nhập vào mật khẩu của user:

<img src="https://imgur.com/FNXog9x.png">


