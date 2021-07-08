# Hướng dẫn cài đặt Checkmk 2.0 trên centos 7

## Các bước cài đặt

### 1 Cài đặt gói wget

` yum install wget -y `

### 2. Khai báo repo

` yum install epel-release -y `

### 3. Sử dụng wget để download check_mk

`  wget https://download.checkmk.com/checkmk/2.0.0p5/check-mk-free-2.0.0p5-el7-38.x86_64.rpm `

### 4. Cài đặt check_mk bằng link vừa download. Ở đây tôi lấy link của tôi đã download về. Hãy lấy link của các bạn nhé

`  yum install check-mk-free-2.0.0p5-el7-38.x86_64.rpm `

### 5. Mở port 80 để sử dụng dịch vụ httpd

```

 firewall-cmd --permanent --add-port=80/tcp
 firewall-cmd --reload 

```

### 6. Tắt selinux

` setenforce 0  `

### 7. Tạo và khởi động site

```
omd create monitoring
omd start monitoring

```

### 8. Đặt mật khẩu cho site

```
su - monitoring
htpasswd -m ~/etc/htpasswd omdadmin

```

### 9. Đăng nhập vào trang web bằng tài khoản admin

` http://10.10.10.73/monitoring/check_mk `

![Imgur](https://imgur.com/80ZpuBs.png)

