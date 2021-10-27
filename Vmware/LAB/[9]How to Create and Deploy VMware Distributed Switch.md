# Tạo và triển khai VMware Distributed Switch

Về VMware Distributed Switch là gì, chúng ta có thể tìm đọc để hiểu thêm [tại đây]()

Ở bài viết này chúng ta sẽ thực hành:

- Tạo VMware Distributed Switch

- Rename uplink

- Tạo Distributed Port Group 

- Assging uplink vào Distributed Port Group 

- Add host Esxi vào Distributed Switch

### Tạo VMware Distributed Switch

- Truy Cập vCenter --> Network --> Distributed switch --> New Distributed switch

<h3 align="center"><img src="../Images/Distributed switch/1.png"></h3>

- Đặt tên cho Distributed switch

<h3 align="center"><img src="../Images/Distributed switch/2.png"></h3>

- Chọn version Distributed switch

<h3 align="center"><img src="../Images/Distributed switch/3.png"></h3>

- Chọn số uplink

<h3 align="center"><img src="../Images/Distributed switch/4.png"></h3>

- Kiểm tra lại thông tin và finish

<h3 align="center"><img src="../Images/Distributed switch/5.png"></h3>

### Rename Up link

- Truy cập configure và chọn edit 

<h3 align="center"><img src="../Images/Distributed switch/6.png"></h3>

- Chọn edit uplink name

<h3 align="center"><img src="../Images/Distributed switch/7.png"></h3>

- Thay đổi tên

<h3 align="center"><img src="../Images/Distributed switch/8.png"></h3>

### Tạo Distributed Port Group

- CHọn new distributed port group

<h3 align="center"><img src="../Images/Distributed switch/9.png"></h3>

- Đặt tên cho port group

<h3 align="center"><img src="../Images/Distributed switch/10.png"></h3>

- Cấu hình port và next

<h3 align="center"><img src="../Images/Distributed switch/11.png"></h3>

- Kiểm tra lại thông tin và finish

<h3 align="center"><img src="../Images/Distributed switch/12.png"></h3>

Làm tương tự với các port group khác

### Assging uplink vào Distributed Port Group 

- Chọn Distributed Port Group vừa tạo và edit  --> Teaming and Failover

Chọn đúng uplink theo port group

<h3 align="center"><img src="../Images/Distributed switch/13.png"></h3>

Làm tương tự với các port group khác.

- Sau khi xong kiểm tra topology thì mỗi port group sẽ tương ưng với 1 uplink

<h3 align="center"><img src="../Images/Distributed switch/14.png"></h3>

<h3 align="center"><img src="../Images/Distributed switch/15.png"></h3>

<h3 align="center"><img src="../Images/Distributed switch/16.png"></h3>

### Add host Esxi vào Distributed Switch

- Chọn Add ang Manage Hosts

<h3 align="center"><img src="../Images/Distributed switch/17.png"></h3>

- Chọn Add hosts

<h3 align="center"><img src="../Images/Distributed switch/18.png"></h3>

- Select host cần add vào distributed switch

<h3 align="center"><img src="../Images/Distributed switch/19.png"></h3>

<h3 align="center"><img src="../Images/Distributed switch/20.png"></h3>

- Assign đúng card mạng với port group

<h3 align="center"><img src="../Images/Distributed switch/21.png"></h3>

<h3 align="center"><img src="../Images/Distributed switch/22.png"></h3>

- Add Manage VMkernel adapter

<h3 align="center"><img src="../Images/Distributed switch/23.png"></h3>

- Nếu Host đã có vm thì tích chọn Migarte để migrate network vm sang distributed switch

<h3 align="center"><img src="../Images/Distributed switch/24.png"></h3>

- Kiểm tra lại thông tin và finish

<h3 align="center"><img src="../Images/Distributed switch/25.png"></h3>

Vậy là chúng ta đã Tạo và triển khai VMware Distributed Switch.

Khi tạo VM ở phần New network sẽ có các port group để add vm vào

<h3 align="center"><img src="../Images/Distributed switch/26.png"></h3>

<h3 align="center"><img src="../Images/Distributed switch/27.png"></h3>



