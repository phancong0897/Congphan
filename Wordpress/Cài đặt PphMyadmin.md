# Tài liệu cài đặt PhpMyadmin

## Mục lục

### [1. Giới thiệu](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20PHP.md#1-gi%E1%BB%9Bi-thi%C3%AAu)

### [2. Cài đặt](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20PHP.md#2-c%C3%A0i-%C4%91%E1%BA%B7t)

### [ Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/Wordpress/C%C3%A0i%20%C4%91%E1%BA%B7t%20PHP.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Giới thiêu

- PhpMyAdmin là công cụ miễn phí (free tool) quản trị MySQL bằng giao diện GUI. Cài đặt phpMyAdmin được thực hiện sau khi đã cài đặt Apache, MySQL và PHP. Việc quản lý MariaDB bằng dòng lệnh khó khăn đối với bạn, thì việc cài đặt phpMyAdmin sẽ giúp bạn giải quyết vấn đền này một cách trực quan, thuận tiện với giao diện web.

- LƯU Ý: Cài đặt phpMyAdmin được thực hiện sau khi đã cài đặt Apache, MySQL và PHP.

## 2. Cài đặt

- Cài đặt Apache
 
    - `  yum -y install httpd `

- Để cài đặt bạn phải sử Fedora Projects EPEL (Extra Packages for Enterprise Linux) repo.

    - ` rpm -iUvh http://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm`

- Sau khi cài đặt Fedora Projects EPEL repo chúng ta sử dụng lệnh sau để kiểm tra lại thông tin phiên bản “phpmyadmin”.

    - ` yum info phpMyAdmin `

- Tiến hành cài đặt phpmyadmin trên CentOS 7 Minimal bằng lệnh sau.

    - ` yum -y install phpmyadmin `

- Sau khi cài đặt xong phpmyadmin trên CentOS 7 bạn không thể truy cập vào địa chỉ “http://yourip/phpmyadmin” bởi vì phpMyAdmin giới hạn quyền truy cập từ các ip khác. Để thay đổi, bạn sửa File phpMyAdmin.conf

    - ` vi /etc/httpd/conf.d/phpMyAdmin.conf `

    - Bạn hãy thêm "# trước 2 dòng"
        ` # Require ip 127.0.0.1 `
        
        ` # Require ip ::1 `
    
    - Và thêm dòng ` Require all granted ` như hình :

<img src="https://imgur.com/I5WbVTH.png">

- Cài đặt PHP:
    
    - yum install -y php

- Sau đó các bạn lưu file “phpmyadmin.conf” mà bạn đã cấu hình ở trên lai và restart lại dịch vụ “httpd” trên CentOS 7 Minimal.

    - ` systemctl restart httpd `

- Các bạn nhớ mở rule trên firewall cho phép dịch vụ httpd có thể đi qua nhé.

    - ` firewall-cmd --add-service=http `

- Sau khi cài đặt xong, các bạn truy cập vào phpmyadmin bằng cách truy cập địa chỉ
` http://yourip/phpmyadmin `

<img src="https://imgur.com/JrtGiMr.png">

Nếu bạn đăng nhập địa chỉ và trình duyệt hiển thị như hình thì bạn đã cài đặt thành công PhpmMadmin

## Nguồn tham khảo

https://tuanvd.com/huong-dan-cai-dat-phpmyadmin-tren-centos-7/
