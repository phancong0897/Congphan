# Tổng quan về TCP wrapper

### Mục lục

### [1. TCP wrapper là gì?](https://github.com/phancong0897/Congphan/blob/master/Firewall/TCP-wrapper.md#1-tcp-wrapper-l%C3%A0-g%C3%AC-1)

### [2. Cài đặt TCP wrapper](https://github.com/phancong0897/Congphan/blob/master/Firewall/TCP-wrapper.md#2-c%C3%A0i-%C4%91%E1%BA%B7t-tcp-wrappers)

## 1. TCP wrapper là gì?

- Các TCP Wrapper là danh sách điều khiển truy cập (ACL – Access Control List) dựa trên máy, và sử dụng để lọc truy cập mạng với các dịch vụ địa phương.

- Chúng xuất hiện vào những năm 1990 để bảo vệ các máy trạm Unix khỏi các cuộc tấn công mạng.

- Ưu điểm : TCP / Wrappers có một lợi thế lớn so với tường lửa thông thường: chúng hoạt động trong lớp 7 (Ứng dụng), do đó, chúng có thể, trong số những thứ khác, có thể lọc các truy vấn ngay cả khi sử dụng mã hóa.

- Nhược điểm : Tất cả các ứng dụng dịch vụ Unix phải được biên dịch để hoạt động với thư viện libwrap. Wrappers không hoạt động với các dịch vụ gọi thủ tục từ xa (RPC) qua TCP).

## 2. Cài đặt TCP Wrappers

- TCP Wrappers có sẵn trong kho chính thức của hầu hết các hệ điều hành Linux.

- Cài đặt tcp_wrapper trên centos7

    - ` yum -y install tcp_wrappers `

- Cấu hình hạn chế quyền truy cập vào máy chủ Linux bằng cách sử dụng TCP Wrappers.

- TCP Wrappers thực hiện kiểm soát truy cập với sự trợ giúp của hai tệp cấu hình: /etc/hosts.allow và /etc/hosts.deny .Hai tệp danh sách kiểm soát truy cập này quyết định liệu các máy khách cụ thể có được phép truy cập máy chủ Linux của bạn hay không.

    - /etc/hosts.allow : Tệp này chứa tên của các máy chủ được phép sử dụng các dịch vụ mạng.

    - /etc/hosts.deny : Tệp này chứa tên của máy chủ không thể sử dụng dịch vụ mạng.

- Nếu cùng một máy khách, người dùng hoặc ip được liệt kê trong cả hai tệp, hosts.allow có mức độ ưu tiên hơn hosts.deny, hãy cẩn thận với điều này.

- Cú pháp của các tệp này như sau:

    - ` list_of_service : list_of _ client [ : lệnh _ shell ] `

- Để bảo mật máy chủ Linux là chặn tất cả các kết nối đến và chỉ cho phép một vài máy chủ hoặc mạng cụ thể. Để làm như vậy, chỉnh sửa tệp /etc/hosts.deny :

    - ` vi /etc/hosts.deny `

    - Thêm dòng sau. Dòng này từ chối kết nối với tất cả các dịch vụ và tất cả các mạng.

        ` ALL: ALL ` 

<img src="https://imgur.com/DdVfqRV.png">

- Sau đó, chỉnh sửa tệp /etc/hosts.allow :

    - ` vi /etc/hosts.allow `

    - Thêm địa chỉ mạng cụ thể muốn kết nối:

        ` sshd : 172.16.2.141 `

Theo quy tắc trên, tất cả các kết nối đến sẽ bị từ chối cho tất cả các máy chủ ngoại trừ máy có địa chỉ IP: 172.16.2.141.

<img src="https://imgur.com/lPaQfnI.png">

- Bạn có thể xác minh điều này từ các tệp nhật ký của máy chủ Linux của bạn như hiển thị bên dưới.

    - ` cat /var/log/secure `

- Bạn có thể cho phép các kết nối đến từ tất cả các máy , nhưng không phải từ một máy cụ thể. Ví dụ, để cho phép các kết nối đến từ tất cả các máy trong mạng con 172.16.2. , nhưng không phải từ máy 172.16.2.100 , hãy thêm dòng sau vào tệp /etc/hosts.allow .

    - ` ALL: 172.16.2. EXCEPT 172.16.2.100 `
