# Cấu hình đăng nhập SSH Key cho user trên CentOS 7

# Mục lục

### [1. Trên máy CHients](https://github.com/phancong0897/Congphan/blob/master/SSH/Config-SSH-Keypair.md#1-tr%C3%AAn-m%C3%A1y-clients)

### [2. trên máy chủ Centos](https://github.com/phancong0897/Congphan/blob/master/SSH/Config-SSH-Keypair.md#2-tr%C3%AAn-m%C3%A1y-ch%E1%BB%A7-centos-1)

### [3. Kiểm tra](https://github.com/phancong0897/Congphan/blob/master/SSH/Config-SSH-Keypair.md#3-ki%E1%BB%83m-tra-1)

## 1. Trên máy Clients

- Trên Máy Clients dùng hệ điều hành window, Chúng ta tải phần mềm puttygen.exe để tạo key.

- Dowload xong , click vào phần mềm puttygen, chọn Generate để tạo key

<img src="https://imgur.com/Y67Ej8e.png">

- Click chọn Key passphrase để tạo mật khẩu, sau đó chọn save private key để lưu key, file này có đuôi .ppk

<img src="https://imgur.com/idSiajJ.png">

- Chúng ta copy key ở phần Public key for pasting into OpneSSH authorized_keys file và lưu vào notepadd có tên publickey. Đây là public key chúng ta dùng để share cho các máy chủ.

## 2. Trên máy chủ Centos

- Tạo thư mục .ssh bằng lệnh:

    - ` mkdir -p /home/congphan/.ssh `

- Phân quyền cho thư mục vừa tạo:

    - ` chmod 700 /home/congphan/.ssh`

- Tạo tệp authorized_keys để chưa key public :

    - ` touch /home/congphan/.ssh/authorized_keys `

- Phân quyền cho thư mục vừa tạo:

    - ` chmod 600 /home/congphan/.ssh/authorized_keys `

- Mở tệp /home/congphan/.ssh/authorized_keys và copy key public chúng ta vừa mới tạo ở clients có tên publickeys.

- Vào file config /etc/ssh/sshd_config để tắt chế độ xác thực bằng mặt khẩu, tìm đến dòng  PasswordAuthentication chuyển đổi yes thành no đẻ tắt tính năng xác thực bằng mật khẩu

<img src="https://imgur.com/z6pw6gI.png">

- Tiến hành restart lại dịch vụ sshd

    - ` systemctl restart sshd `

## 3. Kiểm tra

- Tiến hành kiểm tra bằng Mobaxterm

- Click chọn MobaXterm, chọn session --> chọn SSH

    - Nhập các tùy chọn : địa chỉ máy chủ, port, click phần advanced ssh setting--> chọn use private key và trỏ đường dẫn đến private key chúng ta lưu ở phần tạo key. sau đó chọn ok

    <img src="https://imgur.com/oEDkR4g.png">

- Nhập username, nhập pasword passphare tạo lúc đầu để đăng nhập ssh bằng key.

<img src="https://imgur.com/W3RKzBQ.png">

- Nếu hiển thị như hình thì bạn đã đăng nhập thành công SSH sử dụng ssh_keypair

<img src="https://imgur.com/JU7RGt7.png">
