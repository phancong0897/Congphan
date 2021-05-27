# Thổng quan về linux bridge

## [1. Tổng quan vê Linux Bridge](https://github.com/phancong0897/Congphan/blob/master/KVM/Linux-Bridge.md#1-t%E1%BB%95ng-quan-v%C3%AA-linux-bridge)

### [2. Tìm hiểu Linux Bridge](https://github.com/phancong0897/Congphan/blob/master/KVM/Linux-Bridge.md#2-t%C3%ACm-hi%E1%BB%83u-linux-bridge)

### [3. Ghi chép về một số khái niệm liên quan](https://github.com/phancong0897/Congphan/blob/master/KVM/Linux-Bridge.md#3-ghi-ch%C3%A9p-v%E1%BB%81-m%E1%BB%99t-s%E1%BB%91-kh%C3%A1i-ni%E1%BB%87m-li%C3%AAn-quan)

### [4. Cài đặt và quản lí linux bridge](https://github.com/phancong0897/Congphan/blob/master/KVM/Linux-Bridge.md#4-c%C3%A0i-%C4%91%E1%BA%B7t-v%C3%A0-qu%E1%BA%A3n-l%C3%AD-linux-bridge)


### 1. Tổng quan vê Linux Bridge

### 1.1. Giới thiệu

- Linux bridge là một soft switch – 1 trong 3 công nghệ cung cấp switch ảo trong hệ thống Linux (bên cạnh macvlan và OpenvSwitch), giải quyết vấn đề ảo hóa network bên trong các máy vật lý.

- Bản chất, linux bridge sẽ tạo ra các switch layer 2 kết nối các máy ảo (VM) để các VM đó giao tiếp được với nhau và có thể kết nối được ra mạng ngoài. Linux bridge thường sử dụng kết hợp với hệ thống ảo hóa KVM-QEMU.

- Linux Bridge thật ra chính là một switch ảo và được sử dụng với ảo hóa KVM/QEMU. Nó là 1 module trong nhân kernel. Sử dụng câu lệnh brctl để quản lý .

![Imgur](https://imgur.com/s1m8QJ7.png)

### 1.2. Cấu trúc hệ thống sử dụng Linux bridge

![Imgur](https://imgur.com/USZ88W9.png)

Khái niệm về physical port và virtual port:

- Virtual Computing Device: Thường được biết đến như là máy ảo VM chạy trong host server

- Virtual NIC (vNIC): máy ảo VM có virtual network adapters(vNIC) mà đóng vai trò là NIC cho máy ảo.

- Physical swtich port: Là port sử dụng cho Ethernet switch, cổng vật lý xác định bởi các port RJ45. Một port RJ45 kết nối tới port trên NIC của máy host.

- Virtual swtich port: là port ảo tồn tại trên virtual switch. Cả virtual NIC (vNIC) và virtual port đều là phần mềm, nó liên kết với virtual cable kết nối vNIC

### 2. Tìm hiểu Linux Bridge

### 2.1. Cấu trúc và các thành phần

- Cấu trúc Linux Bridge:

![Imgur](https://imgur.com/HgkaoAo.png)

- Một số khái niệm liên quan tới linux bridge:

    - Port: tương đương với port của switch thật

    - Bridge: tương đương với switch layer 2

    - Tap: hay tap interface có thể hiểu là giao diện mạng để các VM kết nối với bridge cho linux bridge tạo ra (nó nằm trong nhân kernel, hoạt động ở lớp 2 của mô hình OSI)

    - fd: forward data - chuyển tiếp dữ liệu từ máy ảo tới bridge.

### 2.2. Các tính năng

- STP: Spanning Tree Protocol - giao thức chống loop gói tin trong mạng.

- VLAN: chia switch (do linux bridge tạo ra) thành các mạng LAN ảo, cô lập traffic giữa các VM trên các VLAN khác nhau của cùng một switch.

- FDB: chuyển tiếp các gói tin theo database để nâng cao hiệu năng switch.

### 3. Ghi chép về một số khái niệm liên quan

### 3.1. Port

- Trong networking, khái niệm port đại diện cho điểm vào ra của dữ liệu trên máy tính hoặc các thiết bị mạng. Port có thể là khái niệm phần mềm hoặc phần cứng. Software port là khái niệm tồn tại trong hệ điều hành. Chúng thường là các điểm vào ra cho các lưu lượng của ứng dụng. Tức là khái niệm port mức logic. Ví dụ: port 80 trên server liên kết với Web server và truyền các lưu lượng HTTP.

- Hardware port (port khái niệm phần cứng): là các điểm kết nối lưu lượng ở mức khái niệm vật lý trên các thiết bị mạng như switch, router, máy tính, … ví dụ: router với cổng kết nối RJ45 (L2/Ethernet) kết nối tới máy tính của bạn.

- Physical switch port: Thông thường chúng ta hay sử dụng các switch L2/ethernet với các cổng RJ45. Một đầu connector RJ45 kết nối port trên switch tới các port trên NIC của máy tính.

- Virtual switch port: giống như các physical switch port mà tổn tại như một phần mềm trên switch ảo. cả virtual NIC và virtual port đều duy trì bởi phần mềm, được kết nối với nhau thông qua virtual cable.

### 3.2. Uplink port

- Uplink port là khái niệm chỉ điểm vào ra của lưu lượng trong một switch ra các mạng bên ngoài. Nó sẽ là nơi tập trung tất cả các lưu lượng trên switch nếu muốn ra mạng ngoài.

![Imgur](https://imgur.com/uxZLGEG.png)

- Khái niệm virtual uplink switch port được hiểu có chức năng tương đương, là điểm để các lưu lượng trên các máy guest ảo đi ra ngoài máy host thật, hoặc ra mạng ngoài. Khi thêm một interface trên máy thật vào bridge (tạo mạng bridging với interface máy thật và đi ra ngoài), thì interface trên máy thật chính là virtual uplink port.

### 3.3 Tap interface

- Ethernet port trên máy ảo VM (mô phỏng pNIC) thường gọi là vNIC (Virtual NIC). Virtual port được mô phỏng với sự hỗ trợ của KVM/QEMU.

- Port trên máy ảo VM chỉ có thể xử lý các frame Ethernet. Trong môi trường thực tế (không ảo hóa) interface NIC vật lý sẽ nhận và xử lý các khung Ethernet. Nó sẽ bóc lớp header và chuyển tiếp payload (thường là gói tin IP) tới lên cho hệ điều hành. Tuy nhiên, với môi trường ảo hóa, nó sẽ không làm việc vì các virtual NIC sẽ mong đợi các khung Ethernet.

- Tap interface là một khái niệm về phần mềm được sử dụng để nói với Linux bridge là chuyến tiếp frame Ethernet vào nó. Hay nói cách khác, máy ảo kết nối tới tap interface sẽ có thể nhận được các khung frame Ethernet thô. Và do đó, máy ảo VM có thể tiếp tục được mô phỏng như là một máy vật lý ở trong mạng.

Nói chung, tap interface là một port trên switch dùng để kết nối với các máy ảo VM.

### 4. Cài đặt và quản lí linux bridge

### 4.1. Chuẩn bị môi trường KVM

Chuẩn bị Một Server có cài đặt KVM - Hệ điều hành CentOS 7

Trên Server có Cổng mạng Vật lý eth0, trước tiên ta cần cấu hình Cổng vật lý em1 nối với Bridge br0

- Cấu hình trên Cổng vật lý eth0:

    ` vi /etc/sysconfig/network-scripts/ifcfg-eth0 `

    ```
    
    DEVICE=eth0
    
    TYPE=Ethernet
    
    ONBOOT=yes
    
    BRIDGE=br0
    
    HWADDR=52:54:00:fb:d0:ab
    
    ```
- Tạo cấu hình cho br0:

    ` vi /etc/sysconfig/network-scripts/ifcfg-br0 `

    ```

    ONBOOT="YES"
    
    DEVICE="br0"
    
    BOOTPROTO="static"
    
    IPADDR="10.10.34.121"  # IP card vật lý
    
    PREFIX="24"
    
    GATEWAY="10.10.34.1"
    
    DNS1="8.8.8.8"

    ```
### 4.2. Cài đặt VM trên KVM nhận card Bridge

Chọn Network vào br0 để VM nhận card mạng bridge.

![Imgur](https://imgur.com/0kOi8Rh.png)

Sau khi hoàn thành cài đặt, kiểm tra thấy VM đã nhận IP với dải IP của Card vật lý

![Imgur](https://imgur.com/C6cHQjE.png)

