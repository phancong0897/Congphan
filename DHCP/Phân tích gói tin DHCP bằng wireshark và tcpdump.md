# Phân tích gói tin DHCP bằng wireshark và tcpdump

# Mục lục

## [1. Phân tích gói 4 bản tin DHCP bằng wirshark](https://github.com/phancong0897/Congphan/blob/master/DHCP/Ph%C3%A2n%20t%C3%ADch%20g%C3%B3i%20tin%20DHCP%20b%E1%BA%B1ng%20wireshark%20v%C3%A0%20tcpdump.md#1--ph%C3%A2n-t%C3%ADch-g%C3%B3i-4-b%E1%BA%A3n-tin-dhcp-b%E1%BA%B1ng-wirshark)

## [2. Phân tích gói tin DHCP bằng TCP dump]()

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
