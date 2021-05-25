# Cài đặt KVM trên CentOS


### 1. Chuẩn bị

- Một máy ảo centos 7 có hỗ trợ công nghệ ảo hóa: 3 CPU, 3GB RAM, Disk 50GB.

- File iso của hệ điều hành cài lên máy ảo KVM.

### 2. Cài đặt KVM

### 2.1 Kiểm tra

- Để kiểm tra máy có hỗ trợ ảo hóa hay không

    ` egrep -c "svm|vmx" /proc/cpuinfo `
    
    ![Imgur](https://imgur.com/3OD14lJ.png)

- Nếu kết quả trả về 0 thì máy không hỗ trợ ảo hóa. Còn khác 0 tức là máy có hỗ trợ ảo hóa.

### 2.2 Cài đặt các gói cần thiết

- Ban đầu chúng ta sẽ cài đặt gói qemu-kvm và qemu-img . Các gói này cung cấp KVM cấp người dùng và trình quản lý hình ảnh đĩa.

    - ` yum install qemu-kvm qemu-img `

- Bây giờ, bạn có yêu cầu tối thiểu để triển khai nền tảng ảo trên máy chủ của mình, nhưng chúng tôi vẫn có các công cụ hữu ích để quản trị nền tảng của mình như:

    - Virt-manager cung cấp một công cụ GUI để quản lý các máy ảo của bạn.

    - libvirt-client cung cấp một công cụ CL để quản lý môi trường ảo của bạn, công cụ này có tên là virsh.
    
    - Virt-install cung cấp lệnh “Virt-install” để tạo các máy ảo của bạn từ CLI.

    - libvirt cung cấp các thư viện phía máy chủ và máy chủ để tương tác với người giám sát và hệ thống máy chủ.

- Hãy cài đặt các công cụ trên bằng lệnh sau:

    - ` yum install virt-manager libvirt libvirt-python libvirt-client `

    - ` yum groupinstall virtualization-client virtualization-platform virtualization-tools `

- Khởi động lại libvirtd và kiểm tra trạng thái của libvirtd

    - ` systemctl restart libvirtd `

    - ` systemctl status libvirtd  ` 

### 3. Sử dụng công cụ Virt-manager để cài VM

- Download và lưu file ISO bản Minimal vào thư mục /var/lib/libvirt/file-iso/

     -  `cd /var/lib/libvirt `
     
     -  ` mkdir file-iso `
     
     -  ` cd file-iso `

- Nếu bạn có sắn file Iso, hay dũng winscp để tải lên thư mục vừa tạo, hoặc bạn có thể dowload trên internet bằng câu lệnh:

     -  ` wget http://repos-va.psychz.net/centos/7.6.1810/isos/x86_64/CentOS-7-x86_64-Minimal-1810.iso `

- Đối với bản Minimal thì để sử dụng công cụ đồ họa Virt-manager, ta cần cài gói X-window

     -  ` yum install "@X Window System" xorg-x11-xauth xorg-x11-fonts-* xorg-x11-utils -y `

### 3.2 Truy cập Virt-manager để cấu hình VM

- Truy cập Virt-manager ta chạy lệnh sau :

     - ` virt-manager `

- Giao diện Virt-manager sẽ hiện ra :

    ![Imgur](https://imgur.com/5A9fEC0.png)

 - Để tạo máy ảo ta chọn New Virtua Machine

    ![Imgur](https://imgur.com/HxkJu3j.png)

 - Chọn kiểu cài đặt hệ điều hành :

    ![Imgur](https://imgur.com/HxkJu3j.png)

 - Chọn đường dẫn file iso đã tải lên:

    ![Imgur](https://imgur.com/fpaFffe.png)

 - Cài đặt thông số cơ bản cho máy ảo:
   
   - Thiết lập RAM và CPU

    ![Imgur](https://imgur.com/q9awtmZ.png)

   - Thiết lập thông số Disk

    ![Imgur](https://imgur.com/RjgjZPf.png)

   - Thông tin máy netwwork

    ![Imgur](https://imgur.com/NyvqDdm.png)

- Kiểm tra các thiết lập thông số máy ảo và chọn begin installation

    ![Imgur](https://imgur.com/qmZOJcj.png)

- Sau đó ta thực hiện cài dặt OS như bình thường
    ![Imgur](https://imgur.com/D0vxDPw.png)



