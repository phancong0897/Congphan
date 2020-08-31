# Tài liệu Cấu hình Static IP cho các hệ điều hành Linux

## 1. Hệ điều hành CentOS 7

### 1.1. Cấu hình bằng chỉnh sửa file

- Bước 1: Kiểm tra tên card mạng mà ta sẽ cấu hình:

`ip a`

- Bước 2: Để chỉnh sửa file cấu hình của card mạng đó, ta sử dụng lệnh:

`vi /etc/sysconfig/network-scripts/ifcfg-[tên card mạng]`

Ví dụ card mạng của bạn là `eth0`:

`vi /etc/sysconfig/network-scripts/ifcfg-eth0`

- Bước 3: Cấu hình cho card mạng

Tại đây sẽ hiện ra cấu hình trước đó của card mạng `etho0`:

```
DEVICE="eth0"
ONBOOT=yes
NETBOOT=yes
UUID="41171a6f-bce1-44de-8a6e-cf5e782f8bd6"
IPV6INIT=yes
BOOTPROTO=dhcp
HWADDR="00:08:a2:0a:ba:b8"
TYPE=Ethernet
NAME="eth0"
```

Trong đó `BOOTPROTO=dhcp` là đang chạy mặc định dịch vụ DHCP để cấp IP cho card mạng này, ta nên thay bằng `none` hoặc `static` để cấu hình IP tĩnh.

Tiến hành chỉnh sửa và thêm một số cấu hình như sau:

```
HWADDR=00:08:A2:0A:BA:B8
TYPE=Ethernet
BOOTPROTO=none
# Server IP #
IPADDR=172.16.2.230
# Subnet Mask #
PREFIX=20
# Set default gateway IP #
GATEWAY=172.16.10.1
# Set dns servers #
DNS1=172.16.10.1
DNS2=8.8.8.8
DNS3=8.8.4.4
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
# Disable ipv6 #
IPV6INIT=no
NAME=eth0
# This is system specific and can be created using 'uuidgen eth0' command #
UUID=41171a6f-bce1-44de-8a6e-cf5e782f8bd6
DEVICE=eth0
ONBOOT=yes
```

**Lưu ý**: Việc cấu hình địa chỉ IP, Subnetmask, DNS tương ứng với cấu hình mạng tại nơi ta thay đổi. Nếu không đúng thì sẽ không thể kết nối mạng.

- Bước 4: Lưu lại cấu hình và tiến hành khởi động lại network:

`systemctl restart network`

### 1.2. Cấu hình Static IP bằng lệnh

Việc cấu hình bằng lệnh cũng giống như việc thay đổi file cấu hình của card mạng. Chúng ta có thể lựa chọn 1 trong 2 cách và hiệu quả đều như nhau. Cách này ta sử dụng lệnh `nmcli`:

- Cấu hình địa chỉ IP và Subnet Mask:

`sudo nmcli connection modify eth0 IPv4.address 172.16.2.235`

- Cấu hình địa chỉ Gateway:

`sudo nmcli connection modify eth0 IPv4.gateway 172.16.10.1`

- Cấu hình địa chỉ DNS:

`sudo nmcli connection modify eth0 IPv4.dns 172.16.10.1`

- Khởi động lại card mạng:

`sudo nmcli connection down eth0 && sudo nmcli connection up eth0`

## 2. Hệ điều hành CentOS 8

### 2.1. Cấu hình bằng cách chỉnh sửa file

Nhìn chung file cấu hình card mạng của CentOS 8 cũng giống như CentOS 7. Các bước thực hiện cũng gần như tương tự CentOS 7.

- Bước 1: Kiểm tra tên card mạng ta muốn thay đổi cấu hình

`ip addr`

- Bước 2: Truy cập vào file để cấu hình card mạng tại `vi /etc/sysconfig/network-scripts/ifcfg-eth0`

```
TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="none"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="eth0"
UUID="d5f41bf4-de0a-43b3-b633-7e2ec6212e58"
DEVICE="eth0"
ONBOOT="yes"
IPADDR=172.16.2.235
PREFIX=20
GATEWAY=172.16.10.1
DNS1=172.16.10.1
```

- Bước 3: Khởi động lại card mạng

`sudo nmcli connection down eth0 && sudo nmcli connection up eth0`

### 2.2. Sử dụng lệnh cấu hình Static IP

Các bước thực hiện giống với CentOS 7

### 2.3. Sử dụng nmcli shell

Với phương pháp này, ta sẽ cấu hình trong shell của câu lệnh `nmcli`

- Kích hoạt cấu hình nmcli shell cho card mạng:

`sudo nmcli connection edit eth0`

- Cấu hình địa chỉ IP, Gateway, DNS:

```
nmcli> set IPv4.address 172.16.2.235
nmcli> set IPv4.gateway 172.16.10.1
nmcli> set IPv4.dns 172.16.10.1
nmcli> set IPv4.method manual
```

- Sau khi cấu hình xong ta lưu lại bằng lệnh:

```
nmcli> save
Connection 'eth0' (d5f41bf4-de0a-43b3-b633-7e2ec6212e58) successfully updated.
```

- Sau đó `nmcli> quit` và khởi động lại card mạng:

`sudo nmcli connection down eth0 && sudo nmcli connection up eth0`
