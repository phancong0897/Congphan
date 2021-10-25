# Triển khai VCenter

### Chuẩn bị:

- Server window ở đây ta sử dụng winserver 2016

- File iso của Vcenter dowload ở trang chủ vmware.

### Tiến hành cài đặt :

- Khởi động bộ cài đặt vCenter trên máy lên để deploy vCenter lên host ESXI. Chọn Install để tiến hành cài đặt.

<h3 align="center"><img src="../Images/Vcenter/1.png"></h3>

- Quá trình cài đặt sẽ gồm 2 stage:

    - Stage 1: Deploy vCenter

    - Stage 2: Setup vCenter

#### Stage 1: Deploy vCenter

- Nhấn next để tiến hành Stage 1 Deploy vCenter lên host ESXI.

<h3 align="center"><img src="../Images/Vcenter/2.png"></h3>

- Click chọn accept lincense

<h3 align="center"><img src="../Images/Vcenter/3.png"></h3>

- Nhập IP của Host ESXI muốn deploy lên. Khai báo username/password của Host ESXI đó, Port https để mặc định.

<h3 align="center"><img src="../Images/Vcenter/4.png"></h3>

- Nhập thông tin tên VM khi deploy và password để đăng nhập vào hệ thống vCenter.

<h3 align="center"><img src="../Images/Vcenter/5.png"></h3>

- Tiếp theo, chọn cấu hình muốn deploy cho vCenter.

<h3 align="center"><img src="../Images/Vcenter/6.png"></h3>

- Ở mục tiếp theo sẽ chọn datastore để cài đặt. Đối với bản 7.x trở đi hỗ trợ việc triển khai vSAN trực tiếp trên bộ cài đặt vCenter nên ở đây sẽ có 2 cách tạo phân vùng Datastore.

    - Cách 1 (Truyền thống): sẽ lấy 1 ổ đĩa tạo datastore và deploy vCenter lên đó. Sau khi setup xong vCenter sẽ tạo vSAN và migrate vCenter lên vSAN đó

    - Cách 2: sẽ tạo vSAN cluster trên host deploy vCenter và deploy vCenter lên đó. Sau đó sẽ setup network vSAN cho các host khác join vào cluster vSAN.

    - Đối với cách 1 sẽ chọn như hình. Lưu ý nhớ bật thin disk mode ( Thin disk mode sẽ cho phép VM sử dụng bộ nhớ tương ứng với bộ nhớ thật của VM. Không như thick disk mode sẽ lấy toàn bộ dung lượng ổ đĩa được phân sẵn)

    <h3 align="center"><img src="../Images/Vcenter/7.png"></h3>

    - Còn cách 2 sẽ chọn mục Install on a new vSAN cluster containning the target host. Ở đây sẽ nhóm các HDD và SSD lại thành vSAN Datastore. (HDD sẽ làm capacity tier và SSD sẽ làm cache tier).

    <h3 align="center"><img src="../Images/Vcenter/18.png"></h3>

- Xong mục setup datastore sẽ đến mục setup network cho VM vCenter. Ở đây lưu ý nên cấu hình FQDN nếu bạn public vCenter ra internet.

<h3 align="center"><img src="../Images/Vcenter/8.png"></h3>

- Kiểm tra lại cấu hình vCener và nhấn Next để tiến hành quá trình deploy.

<h3 align="center"><img src="../Images/Vcenter/9.png"></h3>

<h3 align="center"><img src="../Images/Vcenter/10.png"></h3>

<h3 align="center"><img src="../Images/Vcenter/11.png"></h3>

Sau khi deploy hoàn thành sẽ tiến hành Stage 2: Setup vCenter. Nếu khi deploy xong bạn lỡ tắt trình cài đặt thì có thể truy cập vào giao diện cài đặt stage 2 bằng giao diện web với đường dẫn  https://<IP_vCenter>:5480 để tiếp tục.

#### Stage 2: Setup vCenter

<h3 align="center"><img src="../Images/Vcenter/12.png"></h3>

- Bắt đầu config vCenter

<h3 align="center"><img src="../Images/Vcenter/13.png"></h3>

- Tiếp đến sẽ cáu hình password đăng nhập web GUI

<h3 align="center"><img src="../Images/Vcenter/14.png"></h3>

- Configure CEIP bỏ tick và chọn next

<h3 align="center"><img src="../Images/Vcenter/15.png"></h3>

- Bước cuối sẽ kiểm tra các thông tin của vCenter. Nếu có sai sót gì phải sửa ngay nếu không việc thay đổi về sau rất tốn thời gian.

<h3 align="center"><img src="../Images/Vcenter/16.png"></h3>

Sau khi Stage 2 hoàn tất, chúng ta có thể truy cập vào web GUI của vCenter để add các Host ESXI.
