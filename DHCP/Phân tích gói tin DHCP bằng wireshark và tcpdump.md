# Phân tích gói tin DHCP bằng wireshark và tcpdump

# Mục lục

## [1. Phân tích gói 4 bản tin DHCP bằng wirshark]()

## [2. Phân tích gói tin DHCP bằng TCP dump]

## 1.  Phân tích gói 4 bản tin DHCP bằng wirshark


<img src="https://imgur.com/Cna5vtR.png">

- DHCP Discovery
    <img src="https://imgur.com/rQP9kRa.png">

    - option (53): Kiểu tin nhắn.
    
    - option (55): Danh sách tham số yêu cầu.

    - User Datagram protocol: thông tin port nguồn , port đích DHCP sử dụng.

- DHCP offer

    <img src="https://imgur.com/x914ykv.png">

    - Internet Protocol v4: chứa địa chỉ IP nguồn, IP đích.

    - option (53): Kiểu tin nhắn.

    - option (54): DHCP server.

    - option (51): thời gian cấp địa chỉ IP.

    - option (1) : subnet mask.

    - option (28): địa chỉ broacast.

    - option (15): Domain Name.

    - option (6) : DNS.

    - option (3) : Deafault gateway.

- DHCP request

    <img src="https://imgur.com/i5sGA0K.png">

    - option (50): địa chỉ IP yêu cầu.

- DHCP ACK

    <img src="https://imgur.com/9RhP3qk.png">
