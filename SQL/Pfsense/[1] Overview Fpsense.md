# Tổng quan về Pfsense

### 1. Giới thiệu về pfSense

![Imgur](https://imgur.com/k0VaHKh.png)

PfSense là một ứng dụng có chức năng định tuyến, tường lửa và miễn phí, ứng dụng này sẽ cho phép bạn mở rộng mạng của mình mà không bị thỏa hiệp về sự bảo mật. Bắt đầu vào năm 2004, khi m0n0wall mới bắt đầu chập chững – đây là một dự án bảo mật tập trung vào các hệ thống nhúng – pfSense đã có hơn 1 triệu lượt download và được sử dụng để bảo vệ các mạng có tất cả kích cỡ, từ mạng gia đình đến các mạng lớn của các công ty/doanh nghiệp. Ứng dụng này có một cộng đồng phát triển rất tích cực và nhiều tính năng đang được bổ sung trong mỗi lần phát hành nhằm cải thiện hơn nữa tính bảo mật, sự ổn định và khả năng linh hoạt của nó. Và là một trong số ít những firewall có tính năng trạng thái, chỉ thường xuất hiện ở những firewall thương mại lớn như Cisco ASA, Checkpoint, Juniper …

    - PfSense bao gồm nhiều tính năng đặc biệt là firewall trạng thái mà bạn vẫn thấy trên các thiết bị tường lửa hoặc router thương mại lớn, chẳng hạn như giao diện người dùng (GUI) trên nền Web tạo sự quản lý một cách dễ dàng. Trong khi đó phần mềm miễn phí này còn có nhiều tính năng ấn tượng đối với firewall/router miễn phí, tuy nhiên cũng có một số hạn chế.

    - PfSense hỗ trợ lọc bởi địa chỉ nguồn và địa chỉ đích, cổng nguồn hoặc cổng đích hay địa chỉ IP. Nó cũng hỗ trợ chính sách định tuyến và cơ chế hoạt động trong chế độ brigde hoặc transparent, cho phép bạn chỉ cần đặt pfSense ở giữa các thiết bị mạng mà không cần đòi hỏi việc cấu hình bổ sung. PfSense cung cấp cơ chế NAT và tính năng chuyển tiếp cổng, tuy nhiên ứng dụng này vẫn còn một số hạn chế với Point-to-Point Tunneling Protocol (PPTP), Generic Routing Encapsulation (GRE) và Session Initiation Protocol (SIP) khi sử dụng NAT.
    
    - PfSense được dựa trên FreeBSD và giao thức Common Address Redundancy Protocol (CARP) của FreeBSD, cung cấp khả năng dự phòng bằng cách cho phép các quản trị viên nhóm hai hoặc nhiều tường lửa vào một nhóm tự động chuyển

Vì nó hỗ trợ nhiều kết nối mạng diện rộng (WAN) nên có thể thực hiện việc cân bằng tải. Tuy nhiên có một hạn chế với nó ở chỗ chỉ có thể thực hiện cân bằng lưu lượng phân phối giữa hai kết nối WAN và không thể chỉ định được lưu lượng cho qua một kết nối.

### 2. Một số lợi ích mà pfSense đem tới

- Hoàn toàn miễn phí, giá cả là ưu thế vượt trội của tường lửa pfSense. Tuy nhiên, rẻ không có nghĩa là kém chất lượng, tường lửa pfSense hoạt động cực kỳ ổn định với hiệu năng cao, đã tối ưu hóa mã nguồn và cả hệ điều hành. Cũng chính vì thê, pfSense không cần nền tảng phần cứng mạnh.

- Nếu doanh nghiệp không có đường truyền tốc độ cao, tường lửa pfSense chỉ cần cài đặt lên một máy tính cá nhân là có thể bắt đầu hoạt động. Điều đó càng góp phần làm giảm chi phí triển khai, đồng thời tạo nên sự linh hoạt, tính mở rộng/sẵn sàng chưa từng có, khi doanh nghiệp muốn có nhiều hơn một tường lửa.

- Không chỉ là một tường lửa, pfSense hoạt động như một thiết bị mạng tổng hợp với đầy đủ mọi tính năng toàn diện sẵn sàng bất cứ lúc nào. Khi có một vấn đề về hệ thống mạng phát sinh, thay vì phải loay hoay tìm thiết bị và mất thời gian đặt hàng, doanh nghiệp có thể kết hợp các tính năng đa dạng trên pfSense để tạo thành giải pháp hợp lý, khắc phục sự cố ngay lập tức.

- Không kém phần quan trọng đó là khả năng quản lý. Tường lửa pfSense được quản trị một cách dễ dàng, trong sáng qua giao diện web.



