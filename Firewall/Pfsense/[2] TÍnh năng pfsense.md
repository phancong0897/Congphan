# Các tính năng của pfsense

### 1. Aliases

- Aliases có thể giúp bạn tiết kiệm một khối lượng thời gian nếu bạn sử dụng chúng một cách chính xác.

- Một Aliases ngắn cho phép bạn sử dụng cho một host, cổng hoặc mạng có thể được sử dụng khi tạo các rule trong pfSense.

- Sự dụng Aliases sẽ giúp bạn cho phép bạn lưu trữ nhiều mục trong một nơi duy nhất, có nghĩa là bạn không cần phải tạo ra nhiều rule cho nhóm các máy hoặc cổng.

- Aliases là tính năng vô cùng hiệu quả của pfSense. Một Aliases có thể định nghĩa được rất nhiều port hoặc một host hoặc nhiều dãy IP.Kiểu như là một số điện thoại là 0935005235 được lưu là jetking.vn thì Aliases ở đây chính là jetking.vn.

- Một ví dụ đơn giản bạn muốn block IP Facebook – Facebook là một dãy IP rất nhiều lớp mạng, ta không thể làm từng rule trên Firewall để chặn Facebook, nó làm ta mất rất nhiều thời gian, làm chậm hệ thống, làm khó khăn cho việc quản lý.

    - Và chỉ cần ta tạo một Aliases có tên là IP_Facebook chứa một dãy IP facebook. Sau đó trên Firewall ta chỉ cần làm một rule là chặn Aliases IP_Facebook thì block được facebook.com.

    ![Imgur](https://imgur.com/OX4C88Z.png)

    - Các thành phần trong Aliases:

    Host: tạo nhóm các địa chỉ IP

    Network: tạo nhóm các mạng
    
    Port: Cho phép gom nhóm các port nhưng không cho phép tạo nhóm các protocol. Các protocol được sử dụng trong các rule

### 2. Rules (Luật)

- Nơi lưu các rules (Luật) của Firewall. Để vào Rules của pfSense, ta vào Firewall -> Rules.

![Imgur](https://imgur.com/oLHcqd0.png)

- Mặc định pfSense cho phép mọi traffic ra/vào hệ thống. Bạn phải tạo các rules để quản lý mạng bên trong Firewall.

- Một số lựa chọn trong Destination và Source.

    Any: Tất cả

    Single host or alias: Một địa chỉ ip hoặc là một bí danh.

    Lan subnet: Đường mạng Lan

    Network: địa chỉ mạng

    Lan address: Tất cả địa chỉ mạng nội bộ
    
    Wan address: Tất cả địa chỉ mạng bên ngoài
    
    PPTP clients: Các clients thực hiện kết nối VPN sử dụng giao thức PPTP
    
    PPPoE clients: Các clients thực hiện kết nối VPN sử dụng giao thức PPPoE

### 3. Firewall Schedules

Các Firewall rules có thể được sắp xếp để nó chỉ hoạt động vào các thời điểm nhất định trong ngày hoặc vào những ngày nhất định cụ thể hoặc các ngày trong tuần.

Đây là một cơ chế rất hay vì nó thực tế với những yêu cầu của các doanh nghiệp muốn quản lí nhân viên sử dụng internet trong giờ hành chính.

Đề tạo một Schedules mới, ta vào Firewall -> Schedules: Nhấn dấu +

Ví dụ: ở đây Tạo lịch tên ThoiGianLamViec của tháng Từ thứ hai đến thứ bảy và thời gian từ 7 giờ đến 16 giờ.

![Imgur](https://imgur.com/IYxSFBz.png)

### 4. NAT

Tương tự như mọi Firewall khác, pfSense cũng hỗ trợ tính năng NAT để quản trị mạng có thể Publish Web, Mail ra internet

Trong Firewall bạn cũng có thể cấu hình các thiết lập NAT nếu cần sử dụng cổng chuyển tiếp cho các dịch vụ hoặc cấu hình NAT tĩnh (1:1) cho các host cụ thể.

Thiết lập mặc định của NAT cho các kết nối outbound là Automatic outbound NAT …, tuy nhiên bạn có thể thay đổi kiểu Manual outbound NAT .. nếu cần.

Ví dụ ở đây ta NAT qua port 1723 (PPTP) cho cấu hình VPN với IP NAT là 192.168.2.100

![Imgur](https://imgur.com/WzgDy0V/png)

### 5. DHCP Server

Dịch vụ này cho phép pfSense cấp địa chỉ IP và các thông tin cấu hình cho client trong mạng LAN.

Tính năng này nằm trong Services => DHCP server

Bật tính năng cấp IP động cho các máy client.

![Imgur](https://imgur.com/araBzYF.png)

### 6. Load Balancer (cân bằng tải)

- Với chức năng này bạn có thể điều phối mạng hay còn gọi là cân bằng tải mạng

- Có 2 loại load balancing trên pfSense:

    - Gateway load balancing: được dùng khi có nhiều kết nối WAN. Client bên trong LAN khi muốn kết nối ra ngoài Internet thì pfSense lựa chọn card WAN để chuyển packet ra card đó giúp cho việc cân bằng tải cho đường truyền.

    - Server load balancing: cho phép cân bằng tải cho các server của mình. Được dùng phổ biến cho các web server, mail server và server không hoạt động nữa thì sẽ bị remove.

Chức năng cân bằng tải của pfSense có những đặc điểm:

- Ưu điểm

    - Miễn phí.

    - Có khả năng bổ sung thêm tính năng bằng gói dịch vụ cộng thêm.
    
    - Dễ cài đặt, cấu hình
    
- Hạn chế

    - Phải trang bị thêm modem nếu không có sẵn.

    - Không được hỗ trợ từ nhà sản xuất như các thiết bị cân bằng tải khác

    - Vẫn chưa có tính năng lọc URL như các thiết bị thương mại.
    
    - Đòi hỏi người dùng phải có kiến thức cơ bản về mạng để cấu hình.
    
    - Để cấu hình loadbalancung vào Services => Load Balancer

