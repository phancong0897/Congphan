# Tổng quan về DHCP

# Mục lục

## [1. Khái niệm](https://blog.cloud365.vn/ccna/dhcp-tong-quan/)

## [2. Nguyên lý hoạt động của DHCP](https://github.com/phancong0897/Congphan/blob/master/DHCP/T%E1%BB%95ng%20quan%20v%E1%BB%81%20DHCP.md#2-nguy%C3%AAn-l%C3%BD-ho%E1%BA%A1t-%C4%91%E1%BB%99ng-c%E1%BB%A7a-dhcp-1)

## [ Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/DHCP/T%E1%BB%95ng%20quan%20v%E1%BB%81%20DHCP.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Khái niệm

- DHCP (Dynamic Host Configuration Protocol) là giao thức cấu hình host động. Nó cung cấp cho máy tính địa chỉ ip ; subnet mask; default gateway. Và nó thường được cấp phát bởi DHPC server được tích hợp sẵn trên router.

- DHCP giao tiếp bằng UDP và sử dụng port 67 và 68. DHCP server sử dụng port 67 để nghe thông tin từ các client và sử dụng port 68 để reply thông tin

- Ưu điểm khi sử dụng DHCP

    - Tập trung quản trị thông tin cấu hình host

    - Cấu hình động các máy

    - Cấu hình IP cho các máy một cách liền mạch.

    - Sự linh hoạt

    - Đơn giản hóa vài trò quản trị của việc cauas hình địa chỉ IP của client.

    - Sự linh hoạt

- DHCP làm việc theo mô hình client server và thành phần chính của DHCP là:

    - DHCP client: Là thiết bị dùng để kết nối vào mạng

    - DHCP server: Là thiết bị dùng để cấp phát địa chỉ cho client

<img src="https://imgur.com/Q93mSGj.png">

## 2. Nguyên lý hoạt động của DHCP

- Thành phần chính của DHCP bao gồm 4 bản tin:

    - DISCOVERY

    - OFFER

    - REQUEST

    - ACK

- Quá trình cấp phát địa chỉ IP trong giao thức DHCP bao gồm các bước sau : 
    
    <img src="https://imgur.com/1rexBMa.png">

    -  Đầu tiên máy client sẽ gửi đi 1 gói tin quảng bá tên là DHCP discover, nhằm yêu cầu cho việc lấy các thông tin cấu hình IP address, subnet mask, defaut getway, preferred DNS ….. lúc này, vì clinet chưa có địa chỉ ip cho nên nó sẽ dùng một địa chỉ source (nguồn) là 0.0.0.0, đồng thời cũng không biết một địa chỉ broadcast là 255.255.255.255 và sau đó gói tin DHCP Discover này sẽ quảng bá đi toàn mạng. gói tin này chứa một địa chỉ MAC  (là địa chỉ mà mỗi một card mạng được nhà sản xuất cấp cho là mã để phân biệt các card mạng với nhau). Ngoài ra nó còn chứa tên của máy client để server có thể biết được client nào gởi yêu cầu đến.

    - – Sau khi nhận được gói tin DHCP Discover của client, nếu có một DHCP server hợp lệ (nghĩa là nó có khả năng cung cấp địa chỉ IP cho client) thì nó sẽ trả lời lại bằng một goí tin DHCP offer. gói tin này chứa một địa chỉ ip đề nghị cho thuê trong một khoảng thời gian nhất định (mặc định là 8 ngày, sau một khoảng thời gian 50% tức 4 ngày, nó sẽ tự thu hồi IP address đã cấp nếu như client không sử dụng) kèm theo là địa chỉ MAC của client được cấp, một subnet mask và địa chỉ IP của DHCP server đã cấp phát. trong thời gian này server sẽ không cấp phát địa chỉ IP vừa đề nghị cho một client nào khác. 
    
    - Máy client sau khi nhận được những lời đề nghị là gói tin DHCP Offer trên mạng ( trường hợp trong mạng có nhiều hơn 1 DHCP server) sẽ tiến hành chọn lọc một gói tin phù hợp và sau đó phản hồi lại bằng một gói tin là DHCP Request (bao gồm thông tin về DHCP server cấp phát địa chỉ cho nó) để chấp nhận lời đề nghị đó. Điều này giúp cho việc các gói tin còn lại không được chấp nhận sẽ được các server rút lại và dùng để cấp phát cho client khác.

    -  Khi DHCP server nhận được DHCP request, nó sẽ trả lời lại DHCP client bằng gói tin là DHCP Ack nhằm mục đích thông báo là đã chấp nhận cho DHCP client đó thuê địa chỉ IP. Gói tin này bao gồm địa chỉ IP và các thông tin cấu hình khác (DNS server, Wins server….) cuối cùng client nhận được gói DHCP Ack thì cũng có nghĩa là kết thúc quá trình thuê và cấp phát địa chỉ IP và địa chỉ IP này chính thức được client sử dụng

## Nguồn tham khảo:

https://blog.cloud365.vn/ccna/dhcp-tong-quan/