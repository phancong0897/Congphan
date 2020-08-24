# Tạo user có quyền sudo

# Mục lục

## [1. Log in vào your server với tư cách root user](https://github.com/phancong0897/Congphan/blob/master/Linux/T%E1%BA%A1o%20user%20c%C3%B3%20quy%E1%BB%81n%20%20sudo.md#1-log-in-v%C3%A0o-your-server-v%E1%BB%9Bi-t%C6%B0-c%C3%A1ch-root-user-1)

## [2. Dùng lệnh adduser để thêm user mới vào hệ thống](https://github.com/phancong0897/Congphan/blob/master/Linux/T%E1%BA%A1o%20user%20c%C3%B3%20quy%E1%BB%81n%20%20sudo.md#2-d%C3%B9ng-l%E1%BB%87nh-adduser-%C4%91%E1%BB%83-th%C3%AAm-m%E1%BB%99t-user-m%E1%BB%9Bi-v%C3%A0o-h%E1%BB%87-th%E1%BB%91ng)

## [3. Dùng lệnh usermod để thêm user vừa tạo vào nhóm wheel](https://github.com/phancong0897/Congphan/blob/master/Linux/T%E1%BA%A1o%20user%20c%C3%B3%20quy%E1%BB%81n%20%20sudo.md#3-d%C3%B9ng-l%E1%BB%87nh-usermod-%C4%91%E1%BB%83-th%C3%AAm-user-v%E1%BB%ABa-t%E1%BA%A1o-v%C3%A0o-nh%C3%B3m-wheel-)

## [4. Kiếm tra truy cập sudo trên user mới](https://github.com/phancong0897/Congphan/blob/master/Linux/T%E1%BA%A1o%20user%20c%C3%B3%20quy%E1%BB%81n%20%20sudo.md#4-ki%E1%BB%83m-tra-truy-c%E1%BA%ADp-sudo-tr%C3%AAn-user-m%E1%BB%9Bi)

## [ Nguồn tham khảo]()

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


## Nguồn tham khảo

https://vicloud.vn/community/huong-dan-nhanh-cach-tao-sudo-user-tren-centos-7-367.html

