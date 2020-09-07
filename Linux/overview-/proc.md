# Các tập lệnh cơ bản của thư mục /proc trong linux.

### 1. Hệ thống file /proc ( proc FS) 

- Có bao giờ các bạn tự hỏi là những chương trình mình hay sử dụng như top, ps, free hay KDE System Guard làm thế nào lấy được thông tin về hệ thống như các process đang chạy, lượng bộ nhớ còn trống, lượng CPU đang được sử dụng … hay không? Các bạn đang học hệ điều hành, đang tiếp xúc với lập trình hệ thống? Và các bạn muốn tự mình viết những chương trình như vậy? Vậy thì kiến thức về hệ thống file /proc là điều thiết yếu các bạn cần nắm. Bài viết này nhằm giới thiệu cho các bạn các thông tin cơ bản về /proc file system để dựa vào đó các bạn có thể tự bắt tay vào “vọc” hệ thống.

- Proc là hệ thống file ảo (pseudo file system), một hệ thống file thời gian thực (real time) và thường trú trong bộ nhớ (memory resident) để theo dõi các process đang chạy cùng với trạng thái của hệ thống.

- Proc là hệ thống file ảo bởi vì trên thực tế nó không tồn tại trong bất kì phương tiện lưu trữ nào. Nó tồn tại dựa trên bộ nhớ ảo và dữ liệu luôn thay đổi động cùng với trạgn thái của hệ thống. Hầu hết các dữ liệu trong proc FS được cập nhật liên tục để phù hợp với trạng thái hiện tại của hệ điều hành. Nội dung của proc FS có thể được đọc bởi user có quyền thích hợp, trong đó một số phần chỉ dó thể đọc bởi owner của process và root.

### 2. Một số tập tin thường gặp ở Hệ thống file /proc ( proc FS) 

- /proc/cpuinfo

Thông tin về bộ vi xử lý, chẳng hạn như loại vi xử lý, make, model, hay  hiệu năng.

- /proc/devices

Danh sách các driver thiết bị được cấu hình vào kernel hiện đang chạy

- /proc/dma

Hiển thị các kênh DMA ( Direct Memory Access ) đang được sử dụng tại thời điểm này.

- /proc/filesystems

Các Filesystems cấu hình tại kernel.

- /proc/interrupts

Hiển thị interrupts(xử lý ngắt) nào đang hoạt động.

Interrupts – ngắt là sự kiện dừng công việc hiện tại của CPU, buộc CPU thực hiện một việc nào đó rồi mới quay trở lại thực hiện tiếp công việc cũ. 

- /proc/ioports

Thông tin các  I/O ports đang được sử dụng tại thời điểm hiện tại.

- /proc/kmsg

Tin nhắn được tạo bởi kernel.Tin nhắn này cũng được chuyển vào syslog.

- /proc/loadavg

Thông tin load trung bình của hệ thống, ngoài 3 con số thông thường tương ứng với tình trạng load của hệ thống trong 1, 5 và 15 phút. Ngoài ra còn có 2 con số.

Số thread đang hoặc đợi chạy / tổng số thread hệ thống.

PID process mới nhất của hệ thống.

- /proc/meminfo

Thông tin về memeory usage, bao gồm cả thông tin về ram và swap (ram ảo).

- /proc/modules

Các module kernel được load tại thời điểm này.

- /proc/net

Thông tin trạng thái về các giao thức mạng.

- /proc/self

Một symbolic link đến thư mục của chương trình đang xem / Proc. Khi hai quá trình  sử dụng / Proc, chúng nhận được các liên kết khác nhau.
- /proc/stat

Thống kê khác nhau về hệ thống, chẳng hạn như số lỗi trang kể từ khi hệ thống được khởi động.

- /proc/uptime

Thời gian hệ thống hoạt động.

- /proc/version

Phiên bản của kernel .

- Các thư mục được đánh số như 1, 4567, 2385, 112, 40, 41 … chính là các Process ID (PID) của những process đang chạy trong hệ thống. Mỗi thư mục sẽ chứa thông tin về process đó. Bạn có thể dùng lệnh “ps –ef” để liệt kê các process đang chạy và so sánh với tên các thư mục trên để biết thư mục nào chứa thông tin của process nào. Nếu đang dùng giao diện KDE bạn có thể thử phím tắt Ctrl + ESC). Hãy cd vào một thư mục nào đó, số 1 chẳng hạn, và ls ra để xem có gì trong đó.