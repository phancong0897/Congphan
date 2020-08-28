# Tài liệu reset password user trên windows bằng hirenboot

## Mục lục

### [1. Hiren Boot]()

### [1.1 Hiren Boot là gì?]()

### [1.2 Các tính năng trên Hiren Boot]()

### [2. Reset password bằng Hiren Boot]()

### [2.1 Chuẩn bị]()

### [2.2 Các bước thực hiện]()

### [ Nguồn tham khảo]()

## 1. Hiren Boot

### 1.1 Hiren Boot là gì?

-  Boot là ứng dụng tạo file iso boot trên đĩa CD, USB một công cụ hữu ích giúp bạn có thể giải quyết nhiều vấn đề liên quan đến máy tính của mình thông qua các chương trình cài đặt bên trong Hiren Boot như: theo dõi, chỉnh sửa phân vùng, sao lưu, phục hồi dữ liệu khi mất, kiểm tra sức khỏe ổ ứng,...

### 1.2 Các tính năng chính trên Hiren Boot

- Một số tính năng có trên Hiren Boot được nhiều người sử dụng như:

    - Disk Partition Tool là công cụ giúp tạo phân vùng trên một ổ đĩa, bạn có thể tạo bốn phân vùng Primary, hay ba phân vùng Primary và một phân vùng Extended. Trong phân vùng Extended, bạn có thể tạo nhiều phân vùng con khác nhau.

    - Nhân bản phân vùng: sử dụng Hiren's boot bạn có thể tạo ra một phân vùng mới và sao lưu dữ liệu của phân vùng "gốc". Tương tự như một công cụ sao lưu và phục hồi khi cần thiết, hoặc di chuyển nhanh dữ liệu từ ổ đĩa này sang ổ đĩa khác.

    - Gộp và chia tách phân vùng: Bạn có thể gộp hai phân vùng có định dạng giống nhau như FAT, FAT32 mà không làm mất dữ liệu hoặc bạn cũng có thể sáp nhập phân vùng logic vào primary; ngược lại bạn cũng có thể chia tách phân vùng ra, dấu và phục hồi phân vùng đã xóa.

    - Antivirus Tools đây là công cụ diệt virus cho DOS với F-Prot Antivirus và McAfee Antivirus là 2 phần mềm hỗ trợ được nhiều người lựa chọn nhất. Khi sử dụng bạn hãy nhớ cập nhật phiên bản mới nhất để được hỗ trợ tối đa. Phần mềm này tích hợp công nghệ quét mạnh mẽ giúp xác định các dịch vụ ẩn, registry ẩn, tập tin ẩn đồng thời phát hiện và xóa bỏ rootkit, kể cả những phần tử nguy hiểm mới xuất hiện.

    - Testing Tool bao gồm nhiều công cụ giúp chuẩn đoán bệnh cho máy tính của bạn như: kiểm tra RAM, kiểm tra tốc độ toàn hệ thống, kiểm tra ổ cứng, CPU, card màn hình,...

    - Password & Registry Tools đây là công cụ cần thiết cho bạn giúp lấy lại mật khẩu windows khi bạn quên password.

    - System Info Tool khi cần kiểm tra và xem những thông tin cấu hình về hệ thống máy tính của mình bạn có thể sử dụng những công cụ có trong System Info Tool.

    - MultiMedia Tools đây là công cụ hỗ trợ bạn xem hình, nghe nhạc trong môi trường DOS mà không cần windows vẫn đảm bảo chất lượng hình ảnh và âm thanh chân thực.

## 2. Reset passord bằng hiren boot

### 2.1 Chuẩn bị

- Ta cần chuẩn bị một USB cài Boot Windows Server, USB từ 8GB trở lên để đáp ứng được dung lượng cả file ISO Windows Server.

- File ISO là bản cài Windows Server tương ứng với Windows Server hiện tại (2012, 2016,...)

- Phần mềm Rufus để cài USB boot, tham khảo [tại đây](https://kb.pavietnam.vn/reset-password-windows-server-2016.html)

### 2.2 Các bước thực hiện

- Boot máy chủ bằng USB boot vừa tạo, chọn CDROM Drive (1.0)

<img src="https://imgur.com/qB9rlwj.png">

- Chọn Next

<img src="https://imgur.com/yL8dKbL.png">

- Chọn Repair your computer

<img src="https://imgur.com/5p1jgO1.png">

- Chọn Trouble Shoot

<img src="https://imgur.com/cvCqidV.png">

- Chọn Command Prompt

<img src="https://imgur.com/bgzkwE0.png">

-  Xác định ổ chứa OS Windows Server

    - Ta phải xác định ổ chứa OS của Windows Server mà ta đã cài đặt, ví dụ trong bài này sẽ là ổ :/C

    - Các lệnh thực hiện:

    ```
    x:\source>c: (C – là partition OS chứa các file system)

    C:\>cd Windows\system32 (dùng lệnh ”cd ” truy cập vào đường dẫn Windows\system32)

    C:\Windows\system32>ren Utilman.exe Utilman.exe.old  (dùng lệnh “ren” để rename lại file Utilman.exe thành file Ultilman.exe.old)

    C:\Windows\system32>copy cmd.exe Utilman.exe ( dùng lệnh “copy” copy file “cmd.exe” đè lên “Utilman.exe” . Mục đích copy như vậy để khi login windows bấm phím Windows+U để mở cửa sổ Command Line)

    ```
- Sau khi hoàn thành ta đóng cửa sổ Command Prompt và chọn "Continue"

<img src="https://imgur.com/vFp1Q09.png">

- Để cho Windows Server khởi động bình thường, đến phần Login cho nhấn Tổ hợp phím Windows + U để mở cửa sổ Command Line

<img src="https://imgur.com/e6OTqhD.png">

- Sử dụng lệnh để thay đổi mật khẩu

    ` net user administrator (password cần đặt) `

    Thông báo thay đổi mật khẩu thành công:

    The command completed successfully

- Sau khi thay đổi mật khẩu thành công ta login lại với mật khẩu mới

Chúc các bạn thành công !

## Nguồn tham khảo

https://kb.pavietnam.vn/reset-password-windows-server-2016.html
