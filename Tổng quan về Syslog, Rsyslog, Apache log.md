# Tổng quan về Syslog, Rsyslog, Apache log

# Mục lục

## [1. Log là gì?]()

## [2. Tổng quan về Syslog]()

## [3. Rsyslog]()

## [4. Apache log]()

## 1. Log là gì?

- Log ghi lại liên tục các thông báo về hoạt động của cả hệ thống hoặc của các dịch vụ được triển khai trên hệ thống và file tương ứng.

- Các file log có thể nói cho bạn bất cứ thứ gì bạn cần biết, để giải quyết các rắc rối mà bạn gặp phải miễn là bạn biết ứng dụng nào. Mỗi ứng dụng được cài đặt trên hệ thống có cơ chế tạo log file riêng của mình để bất cứ khi nào bạn cần thông tin cụ thể thì các log file là nơi tốt nhất để tìm.

- Các tập tin log được đặt trong thư mục /var/log. Bất kỳ ứng dụng khác mà sau này bạn có thể cài đặt trên hệ thống của bạn có thể sẽ tạo tập tin log của chúng tại /var/log. Dùng lệnh ls -l /var/log để xem nội dung của thư mục này.

- Ý nghĩa một số file log thông dụng có mặc định trên /var/log

    - /var/log/messages – Chứa dữ liệu log của hầu hết các thông báo hệ thống nói chung, bao gồm cả các thông báo trong quá trình khởi động hệ thống.

    - /var/log/cron – Chứa dữ liệu log của cron deamon. Bắt đầu và dừng cron cũng như cronjob thất bại.

    - /var/log/maillog hoặc /var/log/mail.log – Thông tin log từ các máy chủ mail chạy trên máy chủ.

    - /var/log/wtmp – Chứa tất cả các đăng nhập và đăng xuất lịch sử.

    - /var/log/btmp – Thông tin đăng nhập không thành công

    - /var/run/utmp – Thông tin log trạng thái đăng nhập hiện tại của mỗi người dùng.

    - /var/log/dmesg – Thư mục này có chứa thông điệp rất quan trọng về kernel ring buffer. Lệnh dmesg có thể được sử dụng để xem các tin nhắn của tập tin này.

    - /var/log/secure – Thông điệp an ninh liên quan sẽ được lưu trữ ở đây. Điều này bao gồm thông điệp từ SSH daemon, mật khẩu thất bại, người dùng không tồn tại, vv

## 2. Tổng quan Syslog

- Syslog là một giao thức client/server là giao thức dùng để chuyển log và thông điệp đến máy nhận log. Máy nhận log thường được gọi là syslogd, syslog daemon hoặc syslog server. Syslog có thể gửi qua UDP hoặc TCP. Các dữ liệu được gửi dạng cleartext. Syslog dùng port 514.

- Trong chuẩn syslog, mỗi thông báo đều được dán nhãn và được gán các mức độ nghiêm trọng khác nhau. Các loại phần mềm sau có thể sinh ra thông báo: auth, authPriv, daemon, cron, ftp, dhcp, kern, mail, syslog, user,… Với các mức độ nghiêm trọng từ cao nhất trở xuống Emergency, Alert, Critical, Error, Warning, Notice, Info, and Debug.

### Mục đích của Syslog

- Syslog được sử dụng như một tiêu chuẩn, chuyển tiếp và thu thập log được sử dụng trên một phiên bản Linux. Syslog xác định mức độ nghiêm trọng (severity levels) cũng như mức độ cơ sở (facility levels) giúp người dùng hiểu rõ hơn về nhật ký được sinh ra trên máy tính của họ. Log (nhật ký) có thể được phân tích và hiển thị trên các máy chủ được gọi là máy chủ Syslog.

### Định dạng tin nhắn Syslog?

<img src="https://imgur.com/UWD0oNw.png">

- Định dạng nhật ký hệ thống được chia thành ba phần, độ dài một thông báo không được vượt quá 1024 bytes:

    - PRI : chi tiết các mức độ ưu tiên của tin nhắn (từ tin nhắn gỡ lỗi (debug) đến trường hợp khẩn cấp) cũng như các mức độ cơ sở (mail, auth, kernel).

    - Header: bao gồm hai trường là TIMESTAMP và HOSTNAME, tên máy chủ là tên máy gửi nhật ký.

    - MSG: phần này chứa thông tin thực tế về sự kiện đã xảy ra. Nó cũng được chia thành trường TAG và trường CONTENT.

- Cấp độ cơ sở Syslog (Syslog facility levels)

    - Một mức độ cơ sở được sử dụng để xác định chương trình hoặc một phần của hệ thống tạo ra các bản ghi.

    - Theo mặc định, một số phần trong hệ thống của bạn được cung cấp các mức facility như kernel sử dụng kern facility hoặc hệ thống mail của bạn bằng cách sử dụng mail facility.

    - Nếu một bên thứ ba muốn phát hành log, có thể đó sẽ là một tập hợp các cấp độ facility được bảo lưu từ 16 đến 23 được gọi là “local use” facility levels.

    - Ngoài ra, họ có thể sử dụng tiện ích của người dùng cấp độ người dùng (“user-level” facility), nghĩa là họ sẽ đưa ra các log liên quan đến người dùng đã ban hành các lệnh.

- Mức độ cảnh báo của Syslog?

    - Mức độ cảnh báo của Syslog được sử dụng để mức độ nghiêm trọng của log event và chúng bao gồm từ gỡ lỗi (debug), thông báo thông tin (informational messages) đến mức khẩn cấp (emergency levels).

    - Tương tự như cấp độ cơ sở Syslog, mức độ cảnh báo được chia thành các loại số từ 0 đến 7, 0 là cấp độ khẩn cấp quan trọng nhất.

## 3. Rsyslog

- Rsyslog là một sự phát triển của syslog, cung cấp các khả năng như các mô đun có thể cấu hình, được liên kết với nhiều mục tiêu khác nhau (ví dụ chuyển tiếp nhật ký Apache đến một máy chủ từ xa).

- Rsyslog cũng cung cấp tính năng lọc riêng cũng như tạo khuôn mẫu để định dạng dữ liệu sang định dạng tùy chỉnh.

- Rsyslog có thiết kế kiểu mô-đun. Điều này cho phép chức năng được tải động từ các mô-đun, cũng có thể được viết bởi bất kỳ bên thứ ba nào. Bản thân Rsyslog cung cấp tất cả các chức năng không cốt lõi như các mô-đun. Do đó, ngày càng có nhiều mô-đun. Có 6 modules cơ bản:

    - Output Modules

    - Input Modules

    - Parser Modules

    - Message Modification Modules

    - String Generator Modules

    - Library Modules

## 4. Apache log

- Với webserver bạn cần quan tâm đến 2 dạng log. Log truy cập (access log) ghi lại các thông tin người dùng truy cập vào website. Log lỗi (error log) ghi lại các cảnh báo các lỗi xảy ra với dịch vụ liên quan web server.

- Theo mặc định Apache server lưu file log tại : /var/log/hhtpd

- ACCESS_LOG:

    - Log này rất cần thiết đặc biệt khi website của bạn bị tấn công DDOS, nhờ nó bạn có thể xác định được nguồn tần công, tần xuất và quy mô để có action phù hợp như chặn theo User Agent hay IP nguồn.

    - Cấu trúc của acceslog

    <img src="https://imgur.com/05kgBLR.png">

    - Trong đó:

        - 66.0.0.254 là IP nguồn.

        - -- website không sử dụng xác thực.

        - [17/Sep...]: time local

        - "Get/... " : request từ phía khách hàng

        - 403: là mã lỗi

        - 4897: kicks thước body của website

        - "Mozila....": thông tin http của user agent.

- ERROR_LOG:

    - Log này để ghi lại thông tin các lỗi cài đặt cấu hình hay đơn giản chỉ là những cảnh báo giữa Web Server Nginx và các dịch vụ của nó.

    - Cấu trúc của error_log:

    <img src="https://imgur.com/7P88suR.png">

    - Trong đó:

        - [Tue Sep...]: time local

        - [core: notice]: cấp độ cảnh báo và mức cảnh báo.

        - [pid 11438]: Mã id

        [Selinux....]: mô tả chi tiết log.