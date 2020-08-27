# Tổng quan về Firewalld

## Mục lục

## 1. Giới thiệu về FirewallD

- FirewallD là giải pháp tường lửa mạnh mẽ tương tự Firewall CSF, được cài đặt mặc định trên RHEL 7 và CentOS 7, nhằm thay thế Iptables với những khác biệt cơ bản:

- FirewallD sử dụng “zones” và “services” thay vì “chain” và “rules” trong Iptables.

- FirewallD quản lý các quy tắc được thiết lập tự động, có tác dụng ngay lập tức mà không làm mất đi các kết nối và session hiện có.

### 1.1 Zone

- Trong FirewallD, zone là một nhóm các quy tắc nhằm chỉ ra những luồng dữ liệu được cho phép, dựa trên mức độ tin tưởng của điểm xuất phát luồng dữ liệu đó trong hệ thống mạng.

- Các zone được xác định trước theo mức độ tin cậy, theo thứ tự từ “ít-tin-cậy-nhất” đến “đáng-tin-cậy-nhất”:

    - drop: ít tin cậy nhất – toàn bộ các kết nối đến sẽ bị từ chối mà không phản hồi, chỉ cho phép duy nhất kết nối đi ra.

    - block: tương tự như drop nhưng các kết nối đến bị từ chối và phản hồi bằng tin nhắn từ icmp-host-prohibited (hoặc icmp6-adm-prohibited).

    - public: đại diện cho mạng công cộng, không đáng tin cậy. Các máy tính/services khác không được tin tưởng trong hệ thống nhưng vẫn cho phép các kết nối đến trên cơ sở chọn từng trường hợp cụ thể.

    - external: hệ thống mạng bên ngoài trong trường hợp bạn sử dụng tường lửa làm gateway, được cấu hình giả lập NAT để giữ bảo mật mạng nội bộ mà vẫn có thể truy cập.

    - internal: đối lập với external zone, sử dụng cho phần nội bộ của gateway. Các máy tính/services thuộc zone này thì khá đáng tin cậy.

    - dmz: sử dụng cho các máy tính/service trong khu vực DMZ(Demilitarized) – cách ly không cho phép truy cập vào phần còn lại của hệ thống mạng, chỉ cho phép một số kết nối đến nhất định.

    - work: sử dụng trong công việc, tin tưởng hầu hết các máy tính và một vài services được cho phép hoạt động.

    - home: môi trường gia đình – tin tưởng hầu hết các máy tính khác và thêm một vài services được cho phép hoạt động.

    - trusted: đáng tin cậy nhất – tin tưởng toàn bộ thiết bị trong hệ thống.

## 1.2. Quy tắc Runtime/Permanent

- Trong FirewallD, các quy tắc được cấu hình thời gian hiệu lực Runtime hoặc Permanent.

    - Runtime(mặc định): có tác dụng ngay lập tức, mất hiệu lực khi reboot hệ thống.

    - Permanent: không áp dụng cho hệ thống đang chạy, cần reload mới có hiệu lực, tác dụng vĩnh viễn cả khi reboot hệ thống.

- Ví dụ, thêm quy tắc cho cả thiết lập Runtime và Permanent:

```
# firewall-cmd --zone=public --add-service=http  (runtime)

# firewall-cmd --zone=public --add-service=http --permanent (permanent)

# firewall-cmd --reload

```

- Việc Restart/Reload sẽ hủy bộ các thiết lập Runtime đồng thời áp dụng thiết lập Permanent mà không hề phá vỡ các kết nối và session hiện tại. Điều này giúp kiểm tra hoạt động của các quy tắc trên tường lửa và dễ dàng khởi động lại nếu có vấn đề xảy ra.

- Các file quan trọng

    - Thông tin về service qua thông tin lưu tại /usr/lib/firewalld/services/

    - Thông tin các zone lưu ở /etc/firewalld/zones

    - Thông tin các service mình tự tạo ra lưu ở /etc/firewalld/services

## 2. Cấu hình FirewallD

### 2.1 Thiết lập các Zone

- Liệt kê tất cả các zone trong hệ thống

    - ` firewall-cmd --get-zones `

- Kiểm tra zone mặc định

    - ` firewall-cmd --get-default-zone `

- Kiểm tra zone active (được sử dụng bởi giao diện mạng)

    - Vì FirewallD chưa được thiết lập bất kỳ quy tắc nào nên zone mặc định cũng đồng thời là zone duy nhất được kích hoạt, điều khiển mọi luồng dữ liệu.

    - ` firewall-cmd --get-active-zones `

### 2.2 Thiết lập các quy tắc

- Liệt kê toàn bộ các quy tắc của các zones:

    - ` firewall-cmd --list-all-zones `

- Liệt kê toàn bộ các quy tắc trong zone mặc định và zone active

    - ` firewall-cmd --list-all `

- Thiết lập cho phép services trên FirewallD, sử dụng --add-service:

```

firewall-cmd --zone=public --add-service=http

firewall-cmd --zone=public --add-service=http --permanen

```

- Vô hiệu hóa services trên FirewallD, sử dụng --remove-service:

```

firewall-cmd --zone=public --remove-service=http
firewall-cmd --zone=public --remove-service=http --permanent

```

- Mở Port với tham số --add-port:

``` 

firewall-cmd --zone=public --add-port=9999/tcp
firewall-cmd --zone=public --add-port=9999/tcp --permanent

```

- Mở 1 dải port

```

firewall-cmd --zone=public --add-port=4990-5000/tcp

firewall-cmd --zone=public --add-port=4990-5000/tcp --permanent

```

- Đóng Port với tham số --remove-port:

```

firewall-cmd --zone=public --remove-port=9999/tcp
firewall-cmd --zone=public --remove-port=9999/tcp --permanent

```

- Lưu lại và khởi động lại FirewallD
firewall-cmd --reload