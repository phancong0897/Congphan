# Tìm hiểu về CSF và Fail2ban

# Mục Lục

## [1. Tổng quan về CSF](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#1-t%E1%BB%95ng-quan-v%E1%BB%81-csf-1)

## [1.1 Khái niệm](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#11-kh%C3%A1i-ni%E1%BB%87m-1)

## [1.2 Cài đặt CSF](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#12-c%C3%A0i-%C4%91%E1%BA%B7t-csf-1)

## [1.3 Những file cấu hình CSF](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#13-nh%E1%BB%AFng-file-c%E1%BA%A5u-h%C3%ACnh-csf-1)

## [1.4 Một số lệnh thường dùng](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#14-m%E1%BB%99t-s%E1%BB%91-l%E1%BB%87nh-th%C6%B0%E1%BB%9Dng-d%C3%B9ng-1)

## [1.5 Xóa CSF](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#15-x%C3%B3a-csf-1)

## [2. Tổng quan về Fail2ban](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#2-t%E1%BB%95ng-quan-v%E1%BB%81-fail2ban)

## [2.1 Khái niệm](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#21-kh%C3%A1i-ni%E1%BB%87m-1)

## [2.2 Cài đặt Fai2ban](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#22-c%C3%A0i-%C4%91%E1%BA%B7t-fail2ban)

## [2.3 Cấu hình Fail2ban cho Ssh](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#23-c%E1%BA%A5u-h%C3%ACnh-fail2ban-cho-ssh-1)

## [Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/Linux/T%C3%ACm%20hi%E1%BB%83u%20s%E1%BB%AD%20d%E1%BB%A5ng%20CSF%20v%C3%A0%20Fail2ban.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o-1)

## 1. Tổng quan về CSF

## 1.1 Khái niệm

- CSF (ConfigServer & Firewall) là 1 gói ứng dụng hoạt động trên Linux như 1 Firewall được phát hành miễn phí để tăng tính bảo mật cho server (VPS và Dedicated). CSF hoạt động dựa trên iptables và tiến trình ldf để quyét các file log để phát hiện các dấu hiệu tấn công bất thường.

- CSF sẽ giúp server của bạn:

    - Chống DoS các loại.

    - Chống Scan Port.

    - Đưa ra các lời khuyên về việc cấu hình server (VD: Nên nâng cấp MySQL lên bản mới hơn).

    - Chống BruteForce Attack vào ftp server, web server, mail server,directadmin,cPanel…

    - Chống Syn Flood.

    - Chống Ping Flood.

    - Cho phép ngăn chặn truy cập từ 1 quốc gia nào đó bằng cách chỉ định Country Code chuẫn ISO.

    - Hỗ trợ IPv6 và IPv4.

    - Cho phép khóa IP tạm thời và vĩnh viễn ở tầng mạng (An toàn hơn ở tầng ứng dụng ) nên webserver ko phải mệt nhọc xử lý yêu cầu từ các IP bị cấm nữa.

    - Cho phép bạn chuyến hướng yêu cầu từ các IP bị khóa sang 1 file html để thông báo cho người dùng biết IP của họ bị khóa.

## 1.2 Cài đặt CSF

- Cài đặt module Perl cho CSF script
    
    - ` yum install perl-libwww-perl `

- Tải CSF
   
    - ` cd /tmp `

    - ` wget https://download.configserver.com/csf.tgz `

- Tiến hành giải nén và cài đặt CSF

    - ` tar -xzf csf.tgz `

    - ` cd csf `

    - ` sh install.sh `

- Mở file cấu hình CSF

    - ` nano /etc/csf/csf.conf `

<img src="https://imgur.com/Wl8xXvv.png">

- Khi đã cấu hình xong, tắt chế độ Testing bằng cách chuyển TESTING = “1” thành TESTING = “0” ở trong file /etc/csf/csf.conf

    - ` TESTING = "0" `

<img src="https://imgur.com/s5L5Gc0.png">

## 1.3 Những file cấu hình CSF

- Toàn bộ thông tin cấu hình và quản lý CSF được lưu ở các file trong folder /etc/csf. Nếu bạn chỉnh sửa các file này thì cần khởi động lại CSF để thay đổi có hiệu lực.

    - ` csf.conf  `: File cấu hình chính để quản lý CSF.

    - ` csf.allow `: Danh sách địa chỉ IP cho phép qua firewall.

    - ` csf.deny ` : Danh sách địa chỉ IP từ chối qua firewall.

    -  ` csf.ignore ` : Danh sách địa chỉ IP cho phép qua firewall và không bị block nếu có vấn đề.

    - ` csf.*ignore ` : Danh sách user, IP được ignore.

## 1.4 Một số lệnh thường dùng

- Một số câu lệnh sử dụng để add (-a) hoặc deny (-d) một địa chỉ IP.

    - ` csf -d IPADDRESS ` //Block địa chỉ IP

    - ` csf -dr IPADDRESS ` //Xóa địa chỉ IP đã bị block
    
    - ` csf -a IPADDRESS ` //Allow địa chỉ IP

    - ` csf -ar IPADDRESS ` //Xóa địa chỉ IP đã được allow

    - ` csf -g IPADDRESS ` //Kiểm tra địa chỉ IP có bị block không

    - ` csf -r ` //Khởi động lại CSF

    - ` csf -x ` //Tắt CSF

    - ` csf -e ` //Mở CSF

Trong trường hợp bạn quên những lệnh trên, hãy sử dụng csf sẽ liệt kê toàn bộ danh sách các option.

## 1.5 Xóa CSF

- Nếu bạn muốn xóa hoàn toàn CSF, chỉ cần sử dụng script sau:

    - ` /etc/csf/uninstall.sh `

- Việc này sẽ xóa toàn bộ CSF nên bạn cần cân nhắc khi dùng. Nếu muốn tạm thời tắt CSF thì có thể chuyển chế độ TESTING sang 1.

## 2. Tổng quan về Fail2ban

## 2.1 Khái niệm

- Fail2ban là ứng dụng theo dõi các file log, phát hiện và ngăn chặn kết nối từ những địa chỉ IP có những dấu hiệu độc hại như đăng nhập sai mật khẩu SSH nhiều lần, sử dụng iptables firewall để block địa chỉ IP trong một thời gian nhất định.

## 2.2 Cài đặt Fail2ban

- Để cài đặt Fail2Ban trên CentOS 7, trước tiên chúng ta sẽ phải cài đặt EPEL (Gói bổ sung cho Enterprise Linux). EPEL chứa các gói bổ sung cho tất cả các phiên bản CentOS, một trong những gói bổ sung này là Fail2Ban.

    - ` yum -y install epel-release `

    - ` yum -y install fail2ban fail2ban-systemd `

- Sau khi cài đặt, chúng ta sẽ phải cấu hình và tùy chỉnh Fail2Ban với tệp cấu hình jail.local. Tệp jail.local được sử dụng để cập nhật cấu hình tùy chỉnh cho Fail2Ban.

- Đầu tiên hãy tạo một bản sao của tệp jail.conf và lưu nó với tên jail.local:

    - ` cp -pf /etc/fail2ban/jail.conf /etc/fail2ban/jail.local `

- Mở tệp /etc/fail2ban/jail.local để chỉnh sửa. Để dễ dàng hơn trong việc chỉnh sửa tệp tin các bạn có thể login VPS với SFTP:

- Tiếp theo các bạn sẽ cần tìm và chỉnh sửa các thông số sau

    - Ignoreip: Được sử dụng để thiết lập danh sách IP sẽ không bị block. Danh sách các địa chỉ IP được phân cách bởi khoảng trắng (Phím cách). Tham số này được sử dụng để đặt địa chỉ IP cá nhân của bạn (nếu bạn truy cập máy chủ từ một IP tĩnh).

    - bantime: Khoảng thời gian block một IP

    - findtime: Khoảng thời gian một IP phải login thành công.

    - maxretry: Số lần đăng nhập sai tối đa

<img src="https://imgur.com/X03JZYk.png">

- Nếu bạn không muốn chỉnh sửa các thông số này thì các bạn có thể để nguyên thông số mặc định của Fail2ban và chuyển sang bước tiếp theo

## 2.3 Cấu hình Fail2ban cho SSH

- Để bảo vệ SSH khỏi Bruteforce các bạn tạo file /etc/fail2ban/jail.d/sshd.local với nội dung như sau 

<img src="https://imgur.com/vizYjxA.png">

- Trong đó:

    - enabled: kích hoạt bảo vệ

    - filter: giữ mặc định để sử dụng file cấu hình /etc/fail2ban/filter.d/sshd.conf

    - action: fail2ban sẽ ban địa chỉ IP nếu match filter /etc/fail2ban/action.d/iptables.conf

    - logpath: đường dẫn file log fail2ban
    
    - maxretry: số lần login false tối đa

    - bantime: thời gian ban IP 3600 giây = 1 giờ, bạn có thể chỉnh nếu muốn.

- Sau khi cấu hình hoàn tất các bạn tiến hành khởi động Fail2ban

    - ` systemctl enable fail2ban `

    - ` systemctl start fail2ban `

- Kiểm tra trạng thái của dịch vụ bằng câu lệnh sau:

    - ` systemctl status fail2ban `

- Theo dõi nhật ký của Fail2ban để biết bản ghi các hành động gần đây:

    - ` sudo tail -F /var/log/fail2ban.log `

- Bạn có thể sử dụng lệnh sau để biết được VPS/Server đã từng bị tấn công SSH chưa:

    - ` cat /var/log/secure | grep 'Failed password' | sort | uniq -c `

<img src="https://imgur.com/luFZwJU.png">

- Để xem danh sách các IP bị block các bạn chạy lệnh như sau

    - ` fail2ban-client status sshd `

<img src="https://imgur.com/UXlBBKT.png">

- Để mở khoá một IP các bạn có thể sử dụng lệnh như sau

    - ` fail2ban-client set sshd unbanip IP-Bị-Block `

<img src="https://imgur.com/hUIozLv.png">

## Nguồn tham khảo

https://news.cloud365.vn/cach-su-dung-fail2ban-de-bao-mat-may-chu-centos/

https://blog.hostvn.net/chia-se/cai-dat-fail2ban-tren-centos-7-chong-bruteforce-attack.html