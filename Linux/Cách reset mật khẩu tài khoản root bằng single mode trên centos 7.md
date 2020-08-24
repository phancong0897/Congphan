# Cách reset mật khẩu tài khoản root bằng single mode trên centos7

## Mục lục

### [1. khởi động lại hệ thống, chỉnh sửa ‘grub2’]( https://github.com/phancong0897/Congphan/blob/master/Linux/C%C3%A1ch%20reset%20m%E1%BA%ADt%20kh%E1%BA%A9u%20t%C3%A0i%20kho%E1%BA%A3n%20root%20b%E1%BA%B1ng%20single%20mode%20tr%C3%AAn%20centos%207.md#1-kh%E1%BB%9Fi-%C4%91%E1%BB%99ng-l%E1%BA%A1i-h%E1%BB%87-th%E1%BB%91ng-ch%E1%BB%89nh-s%E1%BB%ADa-grub2)

### [2. Chỉnh thông số entry cần thiết](https://github.com/phancong0897/Congphan/blob/master/Linux/C%C3%A1ch%20reset%20m%E1%BA%ADt%20kh%E1%BA%A9u%20t%C3%A0i%20kho%E1%BA%A3n%20root%20b%E1%BA%B1ng%20single%20mode%20tr%C3%AAn%20centos%207.md#2-ch%E1%BB%89nh-th%C3%B4ng-s%E1%BB%91-entry-c%E1%BA%A7n-thi%E1%BA%BFt-1)

### [3. Remount filesystem và chuyển chế độ chroot](https://github.com/phancong0897/Congphan/blob/master/Linux/C%C3%A1ch%20reset%20m%E1%BA%ADt%20kh%E1%BA%A9u%20t%C3%A0i%20kho%E1%BA%A3n%20root%20b%E1%BA%B1ng%20single%20mode%20tr%C3%AAn%20centos%207.md#3-remount-filesystem-v%C3%A0-chuy%E1%BB%83n-ch%E1%BA%BF-%C4%91%E1%BB%99-chroot-1)

### [Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/Linux/C%C3%A1ch%20reset%20m%E1%BA%ADt%20kh%E1%BA%A9u%20t%C3%A0i%20kho%E1%BA%A3n%20root%20b%E1%BA%B1ng%20single%20mode%20tr%C3%AAn%20centos%207.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o-1)

## 1. khởi động lại hệ thống, chỉnh sửa ‘grub2’

- Khởi động lại hệ thống và tinh chỉnh chế độ GRUB2 ở màn hình boot GRUB2.

- Bấm nút ESC để màn hình dừng lại, sau đó ấn nút ‘e’ để thực hiện chỉnh sửa.

<img src="https://imgur.com/0K0c9h2.png">

## 2. Chỉnh thông số entry cần thiết

- Xoá 2 thông số “rhgb quiet” để kích hoạt log message hệ thống khi thực hiện đổi mật khẩu root, sau đó thêm vào dòng init=/bin/bash/

<img src="https://imgur.com/PwJKoue.png">

- Sau đó Ctrl X để lưu và tự động boot vào môi trường initramfs.

## 3. Remount filesystem và chuyển chế độ chroot

- Hệ thống filesystem hiện tại đang ở chế độ “read only” được mount ở thư mục /sysroot/, để thực hiện khôi phục mật khẩu root thì ta cần thêm quyền ghi (write) trên filesystem. Ta sẽ tiến hành remount lại filesystem /sysroot/ với quyền đọc-ghi (read-write).

- Kiểm tra lại xem filesystem đã được remount quyền đọc-ghi hay chưa.

    - ` mount -o remount, rw /sysroot `

    - ` mount | grep -w “/sysroot“ `

- Lúc này filesystem đã được mount và ta sẽ chuyển đổi sang môi trường filesystem .

    - ` chroot /sysroot `
 
- Tiến hành reset password user root.

    - ` passwd root `

- Chạy lệnh sau để update lại các thông số cấu hình SELINUX, nếu bạn có sử dụng SELINUX. Nguyên nhân khi ta update file /etc/passwd chứa mật khẩu thì các thông số SELINUX security contex sẽ khác nên cần update lại.

    - ` touch /.autorelabel `

    - ` touch /.autorelabel `

- Remount filesystem “/“ ở chế độ read-only.

    - ` mount -o remount, ro / `

- Thoát môi trường chroot và Khởi động lại hệ thống.

    - ` exit `

    - ` exec /sbin/reboot `
 
- Bạn sẽ thấy hệ thống reboot và chậm hơn bình thường , do hệ thống đang tiến hành hoạt động SELINUX relabel. Sau khi boot vào hệ thống prompt console thành công thì bạn có thể đăng nhập bằng mật khẩu mới.

## Nguồn tham khảo

https://news.cloud365.vn/huong-dan-reset-password-bang-single-mode-tren-linux-1/

https://cuongquach.com/khoi-phuc-mat-khau-root-tren-centos7-rhel7.html

