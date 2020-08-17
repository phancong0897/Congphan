# Tài liệu cài đặt KVM trên máy ảo centos 7

# Mục lục

## [1. Chuẩn bị](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#1-chu%E1%BA%A9n-b%E1%BB%8B-1)

## [2. Cài đặt KVM](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#2-c%C3%A0i-%C4%91%E1%BA%B7t-kvm-1)

## [2.1 kiểm tra hỗ trợ ảo hóa](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#21-ki%E1%BB%83m-tra-h%E1%BB%97-tr%E1%BB%A3-%E1%BA%A3o-h%C3%B3a-1)

## [2.2 Cài đặt các gói cần thiết](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#22-c%C3%A0i-%C4%91%E1%BA%B7t-c%C3%A1c-g%C3%B3i-c%E1%BA%A7n-thi%E1%BA%BFt-1)

## [3. Sử dụng công cụ Virt-manager để cài đặt VM](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#3-s%E1%BB%AD-d%E1%BB%A5ng-c%C3%B4ng-c%E1%BB%A5-virt-manager-%C4%91%E1%BB%83-c%C3%A0i-vm)

## [3.1 Dowload file iso](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#31-dowload-file-iso-1)

## [3.2 Truy cập Virt-manager để cấu hình VM](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#32-truy-c%E1%BA%ADp-virt-manager-%C4%91%E1%BB%83-c%E1%BA%A5u-h%C3%ACnh-vm-1)

## [Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/T%C3%A0i%20li%E1%BB%87u%20c%C3%A0i%20%C4%91%E1%BA%B7t%20KVM.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o-1)


## 1. Chuẩn bị
  
- Một máy ảo centos 7 có hỗ trợ công nghệ ảo hóa: 2 CPU, 2GB RAM, Disk 30GB.

 - File iso của hệ điều hành cài lên máy ảo KVM.

 ## 2. Cài đặt KVM

### 2.1 Kiểm tra hỗ trợ ảo hóa

- Để kiểm tra máy có hỗ trợ ảo hóa hay không

   - egrep -c "svm|vmx" /proc/cpuinfo

<img src="https://imgur.com/1BjTLr5.png">

Nếu kết quả trả về 0 thì máy không hỗ trợ ảo hóa. Còn khác 0 tức là máy có hỗ trợ ảo hóa.

Nếu trên VMware, ta bật hỗ trợ ảo hóa trong Virtual Machine Settings của máy ảo. Đánh dấu vào 2 ô như hình dưới đây:

<img src="https://imgur.com/admGkD1.png">

### 2.2 Cài đặt các gói cần thiết

- Ban đầu chúng ta sẽ cài đặt gói qemu-kvm và qemu-img . Các gói này cung cấp KVM cấp người dùng và trình quản lý hình ảnh đĩa.

    - yum install qemu-kvm qemu-img

- Bây giờ, bạn có yêu cầu tối thiểu để triển khai nền tảng ảo trên máy chủ của mình, nhưng chúng tôi vẫn có các công cụ hữu ích để quản trị nền tảng của mình như:

    - Virt-manager cung cấp một công cụ GUI để quản lý các máy ảo của bạn.

    - libvirt-client cung cấp một công cụ CL để quản lý môi trường ảo của bạn, công cụ này có tên là virsh.
    
    - Virt-install cung cấp lệnh “Virt-install” để tạo các máy ảo của bạn từ CLI.

    - libvirt cung cấp các thư viện phía máy chủ và máy chủ để tương tác với người giám sát và hệ thống máy chủ.

- Hãy cài đặt các công cụ trên bằng lệnh sau:

    - yum install virt-manager libvirt libvirt-python libvirt-client

    - yum groupinstall virtualization-client virtualization-platform virtualization-tools

- Khởi động lại libvirtd và kiểm tra trạng thái của libvirtd

    - systemctl restart libvirtd

    - systemctl status libvirtd  


## 3. Sử dụng công cụ Virt-manager để cài VM

### 3.1 Dowload file iso

- Download và lưu file ISO bản Minimal vào thư mục /var/lib/libvirt/file-iso/

     - cd /var/lib/libvirt
     
     - mkdir file-iso
     
     - cd file-iso

- Nếu bạn có sắn file Iso, hay dũng winscp để tải lên thư mục vừa tạo, hoặc bạn có thể dowload trên internet bằng câu lệnh:

     - wget http://repos-va.psychz.net/centos/7.6.1810/isos/x86_64/CentOS-7-x86_64-Minimal-1810.iso

- Đối với bản Minimal thì để sử dụng công cụ đồ họa Virt-manager, ta cần cài gói X-window

     - yum install "@X Window System" xorg-x11-xauth xorg-x11-fonts-* xorg-x11-utils -y

### 3.2 Truy cập Virt-manager để cấu hình VM

- Truy cập Virt-manager ta chạy lệnh sau :

     - virt-manager

- Giao diện Virt-manager sẽ hiện ra :

 <img src="https://imgur.com/LywMfVD.png">

 - Để tạo máy ảo ta chọn New Virtua Machine

 <img src="https://imgur.com/N21w1nb.png">

 - Chọn kiểu cài đặt hệ điều hành :

 <img src="https://imgur.com/9JzEUSq.png">

 - Chọn đường dẫn file iso đã tải lên:

 <img src="https://imgur.com/zsi1CRq.png">

 - Cài đặt thông số cơ bản cho máy ảo:
   
   - Thiết lập RAM và CPU

<img src="https://imgur.com/hsXtD2q.png">

   - Thiết lập thông số Disk

<img src="https://imgur.com/8jGtAuZ.png">

   - Thông tin máy netwwork

<img src="https://imgur.com/tdQTqTI.png">

- Kiểm tra các thiết lập thông số máy ảo và chọn begin installation

<img src="https://imgur.com/Xqt2epk.png">

- Sau đó ta thực hiện cài dặt OS như bình thường

<img src="https://imgur.com/wpAyvDO.png">


## Nguồn tham khảo

https://news.cloud365.vn/kvm-huong-dan-cai-dat-kvm-tren-centos-7/

https://www.tecmint.com/install-and-configure-kvm-in-linux/



