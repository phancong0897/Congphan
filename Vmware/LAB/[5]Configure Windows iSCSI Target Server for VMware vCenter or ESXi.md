# Cấu hình máy chủ Window ISCSI cho Vcenter

Các bạn có thể tìm hiểu lý thuyể về ISCSI và các thành phần [ tại đây nhé ](https://github.com/phancong0897/Congphan/blob/master/Vmware/Docs/ISCSI.md).

### Chuẩn bị

- Chuẩn bị máy chủ window 2016 cài đặt ISCSI có 2 card mạng
- Chuẩn bị Esxi1 2 card mạng

Mô hình IP như sau:

<h3 align="center"><img src="../Images/ISCSI/24.png"></h3>

### Máy chủ window server 2016

#### Cài đặt iSCSI Target Server:

- Để triển khai dịch vụ, chúng ta tiến hành cài đặt Server roles iSCSI Target Server. Tại Server Roles mở rộng File and Storage Services -> File and iSCSI  Services và chọn iSCSI Target Server.

<h3 align="center"><img src="../Images/ISCSI/1.png"></h3>

<h3 align="center"><img src="../Images/ISCSI/2.png"></h3>

Sau khi cài đặt xong Server Roles, bạn chọn Close.

#### Tạo New iSCSI Virtual Disk…

- Trong Server Manager, chọn lên File and Storage Services, chọn iSCSI. Trong iSCSI Virtual Disks, chọn Tasks -> New iSCSI Virtaul Disk…

<h3 align="center"><img src="../Images/ISCSI/10.png"></h3>

- Tiếp theo, chọn nơi lưu trữ cho iSCSI Virtual Disk , ở đấy mình chọn ổ E

<h3 align="center"><img src="../Images/ISCSI/11.png"></h3>

- Đặt tên cho iSCSI Virtual Disk

<h3 align="center"><img src="../Images/ISCSI/12.png"></h3>

- Tiếp theo, chọn size cho iSCSI Virtual Disk. Ở đây, nếu chọn Dynamically expanding thì dung lượng của fiile ổ cứng iSCSI Virtual Disk sẽ tăng theo dung lượng mà chúng ta sử dụng và cao nhất là bằng với số dung lượng mà bạn đã chọn ở mục Size. Dynamically expanding cho phép tiết kiệm dung lượng của phân vùng chưa file iSCSI Virtual Disk khi file này chưa sử dụng đến dung lượng mà bạn khai báo ở Size.

<h3 align="center"><img src="../Images/ISCSI/13.png"></h3>

- Tiếp theo, tạo New iSCSI target

<h3 align="center"><img src="../Images/ISCSI/14.png"></h3>

- Đặt tên cho iSCSI target

<h3 align="center"><img src="../Images/ISCSI/15.png"></h3>

- Tiếp theo, mục Access Server chúng ta sẽ tiến hành khai báo iSCSI initiator, chọn Add.

<h3 align="center"><img src="../Images/ISCSI/16.png"></h3>

- Tiếp theo nhập IP Esxi dùng để cấu hình ISCSI là: 192.168.60.20, 192.168.60.120

<h3 align="center"><img src="../Images/ISCSI/25.png"></h3>

- Tiếp theo, bạn có thể chọn phương thức chứng thực cho kết nối từ iSCSI Initiator đến iSCSI Target. Ở đây, chúng ta bỏ qua tùy chọn này.

<h3 align="center"><img src="../Images/ISCSI/19.png"></h3>

- Màn hình Confirmation, cho phép bạn có cái nhìn tổng thể về những cấu hình mà bạn đã chọn. Bạn click Create để tiến hành tạo iSCSI Virtual Disk.

<h3 align="center"><img src="../Images/ISCSI/17.png"></h3>

- Sau khi tạo xong iSCSI Virtual Disk, bạn sẽ thấy  ổ đỉa ảo này trong thư iSCSI của Server Manager.

<h3 align="center"><img src="../Images/ISCSI/18.png"></h3>

Vậy là cấu hình ISCSI trên window đã hoàn tất, Tiếp theo ta sẽ cấu hình để [ ESXi có thể kết nối đến ISCSI ](https://github.com/phancong0897/Congphan/blob/master/Vmware/LAB/%5B6%5DConfiguring%20Multi%20Pathing%20Network%20for%20VMWare%20ESXi%20Storage%20Access.md)

