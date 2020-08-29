# Taì liệu hướng dẫn chuyển nameserver lên cloudflare

- Bước 1: Đăng ký tài khoản CloudFlare
Để đăng ký tài khoản CloudFlare, ta truy cập vào trang: https://www.cloudflare.com/a/sign-up

- Bước 2: Nhập Domain mà ta muốn đăng ký lên CloudFlare

<img src="https://imgur.com/LfnJS6i.png">

- Bước 3: Thêm các bản ghi DNS vào CloudFlare

Tại đây có thể cấu hình thiết lập các record DNS đã có sẵn sau khi quá trình scan hoàn tất hoặc thêm các record DNS khác cho domain cũng như cho phép các record DNS của domain tương ứng phân giải thông qua CloudFlare (fake IP) và chọn Continue.

<img src="https://imgur.com/GPDojVs.png">

- Một số Trạng thái cần chú ý:

<img src="https://imgur.com/2TQwkou.png">

- Buowcs4: Tiến hành thay đổi Nameserver sang Nameserver của CloudFlare

Thông báo thay đổi Nameserver mặc định sang Nameserver của CloudFlare

<img src="https://imgur.com/XqoixIs.png">

<img src="https://imgur.com/vMidOnR.png">

- Bước 5: Truy cập vào trang quản lý tên miền cũ để thay đổi Nameserver

Truy cập vào https://zonedns.vn/ và đăng nhập với tên miền muốn thay đổi

<img src="https://imgur.com/Jex7sSK.png">

- Bước 6: Xác nhận bên CloudFlare đã thay đổi Nameserver

Ta nên đợi một lúc để các thay đổi được hoàn thành

<img src="https://imgur.com/dW8BATA.png">