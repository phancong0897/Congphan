# Cấu hình mạng Multi Pathing để Truy cập storage VMWare ESXi

Như ở bài trước [[5]Configure Windows iSCSI Target Server for VMware vCenter or ESXi](https://github.com/phancong0897/Congphan/blob/master/Vmware/LAB/%5B5%5DConfigure%20Windows%20iSCSI%20Target%20Server%20for%20VMware%20vCenter%20or%20ESXi.md) chúng ta đã cấu hình máy chủ target Window ISCSI cho VMWare vCenter hoặc ESXI

Tiếp theo ta sẽ cấu hình kết nối VMware vCenter hoặc Esxi kết nối đến máy chủ targer window ISCSI với 2 NIC để tăng khả năng chịu lỗi khi Esxi bị mất 1 NIC.

### Cấu hình VCenter

- Kiểm tra host Esxi1 đã đủ 2 card mạng

<h3 align="center"><img src="../Images/ISCSI/26.png"></h3>

- Click virtual switches --> add networking

<h3 align="center"><img src="../Images/ISCSI/27.png"></h3>

- Chọn VMkernel Network adapter --> next

<h3 align="center"><img src="../Images/ISCSI/28.png"></h3>

- Chọn New standard switch để tạo mới

<h3 align="center"><img src="../Images/ISCSI/29.png"></h3>

- Add card mạng vào switch ảo 

<h3 align="center"><img src="../Images/ISCSI/30.png"></h3>

<h3 align="center"><img src="../Images/ISCSI/31.png"></h3>

- Đặt tên cho card mạng ảo:

<h3 align="center"><img src="../Images/ISCSI/32.png"></h3>

- Đặt IP cho mạng ảo

<h3 align="center"><img src="../Images/ISCSI/33.png"></h3>

- View lại và finish

<h3 align="center"><img src="../Images/ISCSI/34.png"></h3>

- Vậy là ta đã tạo thành công cho card mạng NIC1 có tên là ISCSI-01

<h3 align="center"><img src="../Images/ISCSI/35.png"></h3>

- Làm tương tự với NIC2 với tên là ISCSI-02

Do đã tạo swirch ảo ở NIC1 rồi nên ta chọn select an existing network và chọn vswitch1 vừa tạo

<h3 align="center"><img src="../Images/ISCSI/36.png"></h3>

<h3 align="center"><img src="../Images/ISCSI/37.png"></h3>

<h3 align="center"><img src="../Images/ISCSI/38.png"></h3>

- Tiến hành gán NIC cho từng card mạng ảo vừa tạo

    - NIC1

    <h3 align="center"><img src="../Images/ISCSI/39.png"></h3>

    <h3 align="center"><img src="../Images/ISCSI/40.png"></h3>

    - NIC2

    <h3 align="center"><img src="../Images/ISCSI/41.png"></h3>

- Kiểm tra lại VMkernel adapter các card mạng đã nhận IP

<h3 align="center"><img src="../Images/ISCSI/42.png"></h3>

### Kết nối host với ISCSI

- Chọn storage adapter --> add software adapter

<h3 align="center"><img src="../Images/ISCSI/43.png"></h3>

- Chọn add sofware ISCSI adapter

<h3 align="center"><img src="../Images/ISCSI/44.png"></h3>

- Như vâỵ ta đã tạo xong model ISCSI

<h3 align="center"><img src="../Images/ISCSI/45.png"></h3>

- Click vào model ISCSI vừa tạo và add card mạng ISCSI của Esxi

<h3 align="center"><img src="../Images/ISCSI/46.png"></h3>

<h3 align="center"><img src="../Images/ISCSI/47.png"></h3>

- Chọn dynamic discovery để kết nối đến ISCSI window server 2016

<h3 align="center"><img src="../Images/ISCSI/48.png"></h3>

<h3 align="center"><img src="../Images/ISCSI/49.png"></h3>

- Kiểm t ra lại devices đã kết nối đến ISCSI cua window server 2016

<h3 align="center"><img src="../Images/ISCSI/50.png"></h3>