# Keepalived là gì? Tìm hiểu dịch vụ Keepalived High Availability

### 1. Keepalived là gì ?

- Keepalived là một chương trình dịch vụ trên Linux cung cấp khả năng tạo độ sẵn sàng cao (High Availability) cho hệ thống dịch vụ và khả năng cân bằng tải (Load Balancing) đơn giản. Với sự gọn nhẹ, tối ưu trong dịch vụ HA của Keepalived mang đến cho quản trị viên một giải pháp Active-Backup dịch vụ rất tốt. Thế nhưng tính năng Load Balancing của Keepalived thì khá là chán, không mạnh mẽ, tuỳ biến linh hoạt như các giải pháp khác là Nginx hay HAProxy được.

- Keepalived cung cấp các bộ thư viện (framework) cho 2 chức năng chính là: cân bằng tải (load balancing) cùng cơ chế health checking và độ sẵn sàng cao cho hệ thống (high availability) với VRRP .

- Tính năng cân bằng tải sử dụng Linux Virtual Server (IPVS) module kernel trên Linux.

- Tính năng kiểm tra tình trạng sức khoẻ của các máy chủ backend cũng khá linh động giúp duy trì được pool server dịch vụ nào còn sống để cân bằng tải tốt.

- Tính sẵn sàng cao (HA) sẽ được Keepalived sử dụng kĩ thuật giao thức khá nổi tiếng VRRP (Virtual Redundancy Routing Protocol). VRRP được ứng dụng nhiều trong mô hình router failover, bạn có thể coi thêm tài liệu VRRP của Cisco.

<h3 align="center"><img src="../Images/17.png"></h3>

- Tóm gọn lại thì Keepalived cho bạn 2 quyền năng gồm:
    
    - Cân bằng tải: với chức năng health checking (kiểm tra tình trạng sức khoẻ) của các máy chủ trong mô hình HA và các phương thức cân bằng tải xuống server backend.
    
    - Tạo độ sẵn sàng cao (High Avaiability) : chức năng VRRP đảm nhận quản lý khả năng chịu lỗi của cụm server (Failover) với Virtual IP.

### 2. Keepalived Failover IP hoạt động như thế nào ?

- Keepalived sẽ gom nhóm các máy chủ dịch vụ nào tham gia cụm HA, khởi tạo một Virtual Server đại diện cho một nhóm thiết bị đó với một Virtual IP (VIP) và một địa chỉ MAC vật lý của máy chủ dịch vụ đang giữ Virtual IP đó. Vào mỗi thời điểm nhất định, chỉ có một server dịch vụ dùng địa chỉ MAC này tương ứng Virtual IP. Khi có ARP request gởi tới virtual IP thì server dịch vụ đó sẽ trả về địa chỉ MAC này.

- Các máy chủ dịch vụ sử dụng chung VIP phải liên lạc với nhau bằng địa chỉ multicast 224.0.0.18 bằng giao thức VRRP. Các máy chủ sẽ có độ ưu tiên (priority) trong khoảng từ 1 – 254, và máy chủ nào có độ ưu tiên cao nhất sẽ thành Master, các máy chủ còn lại sẽ thành các Slave/Backup, hoạt động ở chế độ chờ.

<h3 align="center"><img src="../Images/18.png"></h3>

- Như đã nói ở trên, các server dịch vụ dùng chung Virtual IP sẽ có 2 trạng thái là MASTER/ACTIVE và BACKUP/SLAVE. Cơ chế failover được xử lý bởi giao thức VRRP, khi khởi động dịch vụ, toàn bộ các server cấu hình dùng chung VIP sẽ gia nhập vào một nhóm multicast. Nhóm multicast này dùng để gởi/nhận các gói tin quảng bá VRRP. Các server sẽ quảng bá độ ưu tiên (priority) của mình, server với độ ưu tiên cao nhất sẽ được chọn làm MASTER. Một khi nhóm đã có 1 MASTER thì MASTER này sẽ chịu trách nhiệm gởi các gói tin quảng bá VRRP định kỳ cho nhóm multicast.

<h3 align="center"><img src="../Images/19.png"></h3>

- Nếu vì một sự cố gì đó mà các server BACKUP không nhận được các gói tin quảng bá từ MASTER trong một khoảng thời gian nhất định thì cả nhóm sẽ bầu ra một MASTER mới. MASTER mới này sẽ tiếp quản địa chỉ VIP của nhóm và gởi các gói tin ARP báo là nó đang giữ địa chỉ VIP này. Khi MASTER cũ hoạt động bình thường trở lại thì server này có thể lại trở thành MASTER hoặc trở thành BACKUP tùy theo cấu hình độ ưu tiên của các router.

#### VRRP Stack

Đây là tính năng quan trọng nhất của dịch vụ Keepalived. VRRP (Virtual Router Redundancy Protocol) . Một số đặc điểm của VRRP, hiểu cho khái niệm server thay router:

VRRP tạo ra một gateway dự phòng từ một nhóm các server. Node active được gọI là master server, tất cả các server còn lạI đều trong trạng thái backup. Server master là server có độ ưu tiên cao nhất trong nhóm VRRP.Chỉ số nhóm của VRRP thay đổI từ 0 đến 255; độ ưu tiên của router thay đổI từ 1 cho đến 254 (254 là cao nhất, mặc định là 100).Các gói tin quảng bá của VRRP được gửI mỗI chu kỳ một giây. Các server backup có thể học các chu kỳ quảng bá từ server master.Nếu có server nào có độ ưu tiên cao hơn độ ưu tiên của server master thì server đó sẽ chiếm quyền.VRRP dùng địa chỉ multicast 224.0.0.18, dùng giao thức IP.
