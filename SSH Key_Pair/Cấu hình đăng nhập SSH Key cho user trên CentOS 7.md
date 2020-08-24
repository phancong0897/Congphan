# Cấu hình đăng nhập SSH Key cho user trên CentOS 7

# Mục lục

## [1. Thực hiện cấu hình trên Server](https://github.com/phancong0897/Congphan/blob/master/SSH%20Key_Pair/C%E1%BA%A5u%20h%C3%ACnh%20%C4%91%C4%83ng%20nh%E1%BA%ADp%20SSH%20Key%20cho%20user%20tr%C3%AAn%20CentOS%207.md#1-th%E1%BB%B1c-hi%E1%BB%87n-c%E1%BA%A5u-h%C3%ACnh-tr%C3%AAn-server-1)

## [2. Thực hiện cài đặt phía Client trên window](https://github.com/phancong0897/Congphan/blob/master/SSH%20Key_Pair/C%E1%BA%A5u%20h%C3%ACnh%20%C4%91%C4%83ng%20nh%E1%BA%ADp%20SSH%20Key%20cho%20user%20tr%C3%AAn%20CentOS%207.md#2-th%E1%BB%B1c-hi%C3%AAn-c%C3%A0i-%C4%91%E1%BA%B7t-ph%C3%ADa-client-tr%C3%AAn-window)

## [ Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/SSH%20Key_Pair/C%E1%BA%A5u%20h%C3%ACnh%20%C4%91%C4%83ng%20nh%E1%BA%ADp%20SSH%20Key%20cho%20user%20tr%C3%AAn%20CentOS%207.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Thực hiện cấu hình trên Server

- SSH vào server bằng user root và thực hiện như sau:

- Tạo 1 cặp ssh key bằng câu lệnh sau:

    - ` ssh-keygen -t rsa `

    - Tại vị trí thứ 1 : Nhập vào nơi lưu key hoặc enter để tự động lưu vào file mặc định, bấm enter.
    
    - Tại vị trí thứ 2 : Nhập vào chuỗi mật khẩu để tăng tính bảo mật. Có thể sử dụng mật khẩu này để đăng nhập nếu không có mật khẩu của user.

    - Tại vị trí thứ 3 : Nhập lại chuỗi mật khẩu passphrase 1 lần nữa.

<img src="https://imgur.com/iPo90vk.png">

- Di chuyển Public Key đến vị trí mặc định của nó là /.ssh/authorized_keys

mv /root/.ssh/id_rsa.pub /root/.ssh/authorized_keys

- Phân quyền cho các file và thư mục chứa keys.

    - `chmod 600 /root/.ssh/authorized_keys `

    -  `chmod 700 .ssh `

- Chỉnh sửa file cấu hình

    Chạy các lệnh sau đối với file /etc/ssh/sshd_config để khai báo thư mục đặt key.

    - ` sudo sed -i 's|#LoginGraceTime 2m|LoginGraceTime 2m|' /etc/ssh/sshd_config `

    - ` sudo sed -i 's|#StrictModes yes|StrictModes yes|' /etc/ssh/sshd_config `

    - ` sudo sed -i 's|#PubkeyAuthentication yes|PubkeyAuthentication yes|' /etc/ssh/sshd_config `

- Vào file cấu hình /etc/ssh/sshd_config chỉnh sửa để tắt tính năng đăng nhập bằng mật khẩu :

    - PasswordAuthentication yes , đổi yes thành no.

<img src="https://imgur.com/GiRINCt.png">


- Khởi động lại dịch vụ ssh

    - ` systemctl restart sshd `

- Sao chép file chứa Key

    Để client có thể đăng nhập vào server thông qua xác thực key, ta cần copy file chứa key để xác thực phía client. File chứa key có vị trí tại :

    - ` /root/.ssh/id_rsa `

## 2. Thực hiên cài đặt phía Client trên window

- Sau khi copy key phía server, sử dụng trình soan thảo văn bản hoặc sử dụng công cụ editor để lưu lại. Key được lưu trong file có phần mở rộng là .ppk

- Tại đây mình sẽ lưu file với tên là key. Sau khi xác định đường dẫn file. Mình sẽ sử dụng MobaXterm để thực hiện load key.

### Load key với MobaXterm

- Trên MobaXterm, click vào tab Tools, chọn MoBaKeyGen (SSh key generator)

<img src="https://imgur.com/iDZnmSw.png">

- Tiếp theo, chọn Load để load private key.

<img src="https://imgur.com/0S3TogZ.png">

- Chọn đường dẫn đến file copy key từ server đã lưu ở phần trên. Sau đó sẽ có 1 của sổ hiện ra như bên dưới, nhập vào mật khẩu Passphrase sau đó chọn Ok để tiếp tục.

<img src="https://imgur.com/B9FGdKx.png">

- Như vậy ta đã load xong. Tiếp tục chọn Save private key để lưu lại private key, sau này dùng cho việc ssh.

<img src="https://imgur.com/81W484q.png">

### Kiểm tra sử dụng SSH Key

- Để kiểm tra, ta cần sử dụng key và ssh vào server. Hãy làm theo các bước sau :

    - Trên server chọn tab Session sau đó chọn New session, chọn SSH:

    <img src="https://imgur.com/5NYg3oG.png">

    - Nhập địa chỉ IP vào Remote host, và click chọn Advanced SSH setting

    <img src="https://imgur.com/LH0Wtfb.png">

    - Chọn Use private key, chọn đường dẫn lưu privatekey của mình sau đó chọn Ok

    <img src="https://imgur.com/Ks6nh0p.png">

    Nếu màn hình hiển thi như ảnh ở trên thì bạn đã cấu hình thành công SSHkey

## Nguồn tham khảo

https://news.cloud365.vn/ssh-phan-2-cau-hinh-dang-nhap-ssh-key-cho-user-tren-centos-7/

