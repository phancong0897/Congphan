# Cách đổi port SSH

- Vì một số lí do bảo mật, quý khách có nhu cầu thay đổi port SSH sang một port khác để tránh bị hacker khai thác hoặc bruteforce. Quý khách có thể tham khảo bài viết này để thay đổi.

- Bước 1: SSH vào VM với quyền root.

- Bước 2: Thay đổi số port tại file config của dịch vụ này (Mặc định là #Port 22) và Save lại.

    - ` vi /etc/ssh/sshd_config `

<img src="https://imgur.com/Exj2bUD.png">

- Bước 3: Restart lại dịch vụ SSH.

    - ` /etc/init.d/sshd restart `

- Bước 4: Kiểm tra lại việc SSH và khai báo số port tương ứng đã thay đổi để kết nối vào Server, dùng lệnh sau để kiểm tra để biết chính xác.

    - ` netstat -lntp |grep ssh `

<img src="https://imgur.com/ToW7lnD.png">

- Lưu ý: Quý khách phải xem các port đang running trên hệ thống và tránh không sử dụng các port này. Quý khách xem với lệnh sau:

    - ` netstat -lntp `

<img src="https://imgur.com/ToW7lnD.png">