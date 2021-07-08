# Hướng dẫn cài đặt agent của checkmk trên centos 7

## Các bước thực hiện với client centos 7

### 1. Tìm agent phù hợp

![Imgur](https://imgur.com/DZqCfGr.png)

Trên các website của check_mk server khi bạn cài đặt và đăng nhập vào nó sẽ hỗ trợ và hiển thị cho bạn các agent mà checkmk hỗ trợ. Việc của bạn là chọn agent phù hợp với hệ điều hành của mình. Ở đây tôi cài đặt agent trên centos 7 nên tôi sẽ chọn agent có đuôi là .rpm để tiến hành cài đặt

![Imgur](https://imgur.com/VKf5fVw.png)

### 2. Cài đặt gói wget

` yum install wget -y  `

### 3. Download agent đã chọn ở bước trên

Ở đây tôi dùng máy tính cá nhân dowload bản rpm về máy tính, sau đó scp lên server cần giám sát, agent có dung lượng 3mb nên việc dowload và scp không mất quá nhiều thời gian.

![Imgur](https://imgur.com/rtg84aN.png)

### 4. Cấp quyền thực thi cho file agent vừa scp

` chmod +x check-mk-agent-2.0.0p5-52a3af09bf17f961.noarch.rpm `

### 5. Cài đặt agent

` rpm -ivh check-mk-agent-2.0.0p5-52a3af09bf17f961.noarch.rpm `

### 6. Cài đặt xinetd

` yum install xinetd -y `

### 7. Khởi động xinetd

```

systemctl start xinetd
systemctl enable xinetd

```

### 8. Cài đặt gói net-tools để kiểm tra dễ dàng hơn

` yum install net-tools -y  `

### 9. Mở port trên client để có thể giao tiếp với check_mk server

` vi /etc/xinetd.conf `

![Imgur](https://imgur.com/W4m3vG8.png)

### 10. Kiểm tra port mặc định của check_mk sử dụng để giám sát được chưa

```

[root@congkt-lab ~]# netstat -npl | grep 6556

tcp        0      0 0.0.0.0:6556            0.0.0.0:*               LISTEN      1/systemd

```

### 11. Mở port trên firewall

```

firewall-cmd --add-port=6556/tcp --permanent
firewall-cmd --reload

```

### 12. Tắt selinux

` setenforce 0 `









