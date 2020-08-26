# Tìm hiểu về LVM và KVM

# Mục lục

[1. Giới thiệu về LVM.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#1-gi%E1%BB%9Bi-thi%E1%BB%87u-v%E1%BB%81-lvm)

[1.1 Khái niệm.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#11-kh%C3%A1i-ni%E1%BB%87m)

[1.2 Ưu điểm và nhược điểm của KVM.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#12-%C6%B0u-%C4%91i%E1%BB%83m-v%C3%A0-nh%C6%B0%E1%BB%A3c-%C4%91i%E1%BB%83m-c%E1%BB%A7a-lvm)

[1.3 Các thành phần trong KVM.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#13-nh%E1%BB%AFng-th%C3%A0nh-ph%E1%BA%A7n-trong-lvm)

[2. Tổng quan về KVM.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#2-t%E1%BB%95ng-quan-v%E1%BB%81-kvm)

[2.1 Khái niệm và vai trò.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#21-kh%C3%A1i-ni%E1%BB%87m-v%C3%A0-vai-tr%C3%B2) 

[2.2 Cấu trúc của KVM.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#22-c%E1%BA%A5u-tr%C3%BAc-c%E1%BB%A7a-kvm)

[2.3 Mô hình vận hành.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#23-m%C3%B4-h%C3%ACnh-v%E1%BA%ADn-h%C3%A0nh)

[2.4 Cơ chế hoạt động.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#24-c%C6%A1-ch%E1%BA%BF-ho%E1%BA%A1t-%C4%91%E1%BB%99ng)

[2.5 Tính năng.](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#25-t%C3%ADnh-n%C4%83ng)

[Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/Linux/Tong-quan-LVM-KVM.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)



## 1. Giới thiệu về LVM

### 1.1 Khái niệm

- LVM (Logical Volume Management) là một công nghệ giúp quản lý các thiết bị lưu trữ dữ liệu trên các hệ điều hành Linux. Công nghệ này cho phép người dùng gom nhóm các ổ cứng vật lý lại và phân tách chúng thành những phân vùng nhỏ hơn, dễ dàng mở rộng các phân vùng này khi cần thiết.

- Một số ứng dụng của LVM:

  - Quản lý một lượng lớn ổ đĩa một cách dễ dàng.

  - Điều chỉnh phân vùng ổ cứng một cách linh động.

  - Backup hệ thống bằng cách snapshot các phân vùng ổ cứng (real-time).

  - Migrate dữ liệu dễ dàng.

- Một số khái niệm liên quan:

  - Physical Volume – PV: Ổ cứng vật lý từ hệ thống (đĩa cứng, partition, iSCSI LUN, SSD…) là đơn vị cơ bản để LVM dùng để khởi tạo các Volume Group. Trên mỗi một PV sẽ chứa khoảng 1 MB header ghi dữ liệu về cách phân bố của Volume Group chứa nó. Header này sẽ hỗ trợ rất nhiều trong việc phục hồi dữ liệu khi có sự cố xảy ra.

  - Volume Group – VG: là tập hợp các ổ cứng vật lý (PV) thành một kho lưu trữ chung với tổng dung lượng của các ổ đĩa con. Mỗi khi ta thêm một PV vào VG, LVM sẽ tự động chia dung lượng trên PV thành nhiều Physical Extent với kích cỡ bằng nhau. Và từ VG, ta có thể tạo ra nhiều Logical Volume và dễ dàng chỉnh sửa dung lượng của chúng.

  - Logical Volume – LV: là các phân vùng luận lý được tạo ra từ VG. Logical Volume tương tự như các partition trên ổ cứng bình thường nhưng linh hoạt hơn vì kích thước của LV có thể được dễ dàng thay đổi theo thời gian thực mà không lo làm gián đoạn hệ thống. Sở dĩ ta có thể dễ dàng thay đổi được kích thước của LV vì LV được chia thành nhiều Logical Extent, mỗi Logical Extent này sẽ được mapping tương ứng với 1 Physical Extent trên các ổ đĩa.

  - Extent: extent là đơn vị nhỏ nhất của VG. Mỗi một volume được tạo ra từ VG chứa nhiều extent nhỏ với kích thuớc cố định bằng nhau. Các extent trên LV không nhất thiết phải nằm liên tục với nhau trên ổ cứng vật lý bên dưới mà có thể nằm rải rác trên nhiều ổ cứng khác nhau. Extent chính là nền tảng cho công nghệ LVM, các LV có thể được mở rộng hay thu nhỏ lại bằng cách add thêm các extent hoặc lấy bớt các extent từ volume này.

  Tóm lại, với công nghệ LVM ta có thể gộp nhiều ổ cứng vật lý Physical Volume lại thành Volume Group để tổng hợp toàn bộ tài nguyên lưu trữ cần thiết. Sau đó, người quản trị có thể chia nhỏ Volume Group ra thành các Logical Volume một cách tùy ý và linh hoạt. Mỗi một Logical Volume gồm nhiều extent, khi cần mở rộng Logical Volume thì ta thêm vào một số extent, khi cần thu nhỏ thì ta lấy lại một số extent.

### 1.2 Ưu điểm và Nhược điểm của LVM

- Ưu điểm :

  - Không để hệ thống bị gián đoạn hoạt động.

  - Không làm hỏng dịch vụ.

  - Có thể kế hợp swap.

  - Có thể tạo ra các vùng dung lượng lớn nhỏ tuỳ ý.

- Nhược điểm:

  - Các bước thiết lập phức tạp và khó khăn hơn.

  - Càng gắn nhiều đĩa cứng và thiết lập càng nhiều LVM thì hệ thống khởi động càng lâu.

  - Khả năng mất dữ liệu cao khi một trong số các đĩa cứng bị hỏng.

  - Windows không thể nhận ra vùng dữ liệu của LVM. Nếu Dual-boot ,Windows sẽ không thể truy cập dữ liệu trong LVM.

### 1.3 Những thành phần trong LVM

- HDD : là một thiết bị lưu trữ máy tính. Nó là loại bộ nhớ không thay đổi và không bị mất dữ liệu khi ta ngừng cung cấp nguồn điện cho chúng.

- Partition: là các phân vùng của ổ cứng. Mỗi một ổ cứng có 4 partition. Trong đó bao gồm 2 loại là primary partition và extended partition.

  - Primary partition: còn được gọi là phân vùng chính, có thể khởi động và mỗi ổ cứng chỉ có tối đa 4 phân vùng này.

  - Extended partition: Hay còn được gọi là phân vùng mở rộng của ổ cứng.

- Cách thức hoạt động các tầng của LVM:

  - Tầng đầu tiên : hard drives là tầng các disk ban đầu khi chưa chia phân vùng.

  - Partitions: Sau đó ta chia các disk ra thành các phân vùng nhỏ hơn.

  - Physical volume : từ một partitions ta sẽ tạo ra được một physical.

  - Group volume : Ta sẽ ghép nhiều physical volume thành một group volume.

## 2. Tổng quan về KVM
### 2.1 Khái niệm và vai trò
- KVM (Kernel-based virtual machine) là giải pháp ảo hóa cho hệ thống linux trên nền tảng phần cứng x86 có các module mở rộng hỗ trợ ảo hóa (Intel VT-x hoặc AMD-V). Về bản chất, KVM không thực sự là một hypervisor có chức năng giải lập phần cứng để chạy các máy ảo. Chính xác KVM chỉ là một module của kernel linux hỗ trợ cơ chế mapping các chỉ dẫn trên CPU ảo (của guest VM) sang chỉ dẫn trên CPU vật lý (của máy chủ chứa VM). Hoặc có thể hình dung KVM giống như một driver cho hypervisor để sử dụng được tính năng ảo hóa của các vi xử lý như Intel VT-x hay AMD-V, mục tiêu là tăng hiệu suất cho guest VM.

- KVM ban đầu được phát triển bởi Qumranet – một công ty nhỏ, sau đó được Redhat mua lại vào tháng 9 năm 2008. Ta có thể thấy KVM là thế hệ tiếp theo của công nghệ ảo hóa. KVM được sử dụng mặc định từ bản RHEL (Redhat Enterprise Linux) từ phiên bản 5.4 và phiên bản Redhat Enterprise Virtualization dành cho Server.

- Qumranet phát hành mã của KVM cho cộng đồng mã nguồn mở. Hiện nay, các công ty nổi tiếng như IBM, Intel và ADM cũng đã cộng tác với dự án. Từ phiên bản 2.6.20, KVM trở thành một phần của hạt nhân Linux.

### 2.2 Cấu trúc của KVM
- Trong kiến trúc KVM, máy ảo là một tiến trình Linux, được lập lịch bởi chuẩn Linux scheduler. Trong thực tế mỗi CPU ảo xuất hiện như là một tiến trình Linux. Điều này cho phép KVM sử dụng tất cả các tính năng của Linux kernel.

- Linux có tất cả các cơ chế của một VMM cần thiết để vận hành (chạy) các máy ảo. Chính vì vậy các nhà phát triển không xây dựng lại mà chỉ thêm vào đó một vài thành phần để hỗ trợ ảo hóa. KVM được triển khai như một module hạt nhân có thể được nạp vào để mở rộng Linux bởi những khả năng này.
 
- Trong một môi trường linux thông thường mỗi process chạy hoặc sử dụng user-mode hoặc kernel-mode. KVM đưa ra một chế độ thứ 3, đó là guest-mode. Nó dựa trên CPU có khả năng ảo hóa với kiến trúc Intel VT hoặc AMD SVM, một process trong guest-mode bao gồm cả kernel-mode và user-mode.

- Trong cấu trúc tổng quan của KVM bao gồm 3 thành phần chính:

  - KVM kernel module:

    - Là một phần trong dòng chính của Linux kernel.

    - Cung cấp giao diện chung cho Intel VMX và AMD SVM (thành phần hỗ trợ ảo hóa phần cứng).

    - Chứa những mô phỏng cho các instructions và CPU modes không được support bởi Intel VMX và AMD SVM.

  - Quemu – kvm: là chương trình dòng lệnh để tạo các máy ảo, thường được vận chuyển dưới dạng các package “kvm” hoặc “qemu-kvm”. Có 3 chức năng chính:

    - Thiết lập VM và các thiết bị ra vào (input/output).

    - Thực thi mã khách thông qua KVM kernel module.

    - Mô phỏng các thiết bị ra vào (I/O) và di chuyển các guest từ host này sang host khác.

- libvirt management stack:

  - Cung cấp API để các tool như virsh có thể giao tiếp và quản lí các VM.

  - Cung cấp chế độ quản lí từ xa an toàn.

### 2.3 Mô hình vận hành

<img src="https://imgur.com/blC8wZ9.png">

Như ta thấy:

- User-mode: Các modul KVM gọi đến sử dụng ioclt() để thực thi mã khách cho đến khi hoạt động I/O khởi xướng bởi guest hoặc một sự kiện nào đó bên ngoài xảy ra. Sự kiện này có thể là sự xuất hiện của một gói tin mạng, cũng có thể là trả lời của một gói tin mạng được gửi bởi các may chủ trước đó. Những sự kiện như vậy được biểu diễn như là tín hiệu dẫn đến sự gián đoạn của thực thi mã khách.

- Kernel-mode: Kernel làm phần cứng thực thi các mã khách tự nhiên. Nếu bộ xử lý thoát khỏi guest do cấp phát bộ nhớ hay I/O hoạt động, kernel thực hiện các nhiệm vụ cần thiết và tiếp tục luồng thực hiện, Nếu các sự kiện bên ngoài như tín hiệu hoặc I/O hoạt động khởi xướng bởi các guest tồn tại, nó thoát tới user-mode.

- Guest-mode: Đây là cấp độ phần cứng, nơi mà các lệnh mở rộng thiết lập của một CPU có khả năng ảo hóa được sử dụng để thực thi mã nguồn gốc, cho đến khi một lệnh được gọi như vậy cần sự hỗ trợ của KVM, một lỗi hoặc một gián đoạn từ bên ngoài.

Khi một máy ảo chạy, có rất nhiều chuyển đổi giữa các chế độ. Từ kernel-mode tới guest-mode và ngược lại rất nhanh, bởi vì chỉ có mã nguồn gốc được thực hiện trên phần cứng cơ bản. Khi I/O hoạt động diễn ra các luồng thực thi tới user-mode, rất nhiều thiết bị ảo I/O được tạo ra, do vậy rất nhiều I/O thoát ra và chuyển sang chế độ user-mode chờ. Hãy tưởng tượng mô phỏng một đĩa cứng và 1 guest đang đọc các block từ nó. Sau đó QEMU mô phỏng các hoạt động bằng cách giả lập các hoạt động bằng các mô phỏng hành vi của các ổ đĩa cứng và bộ điều khiển nó được kết nối. Để thực hiện các hoạt động đọc, nó đọc các khối tương ứng từ một tập tin lớp và trả về dữ liệu cho guest. Vì vậy, user-mode giả lập I/O có xu hướng xuất hiện một nút cổ chai làm chậm việc thực hiện máy ảo.

### 2.4 Cơ chế hoạt động

- Để các máy ảo giao tiếp được với nhau, KVM sử dụng Linux Bridge và OpenVSwitch. Đây là 2 phần mềm cung cấp các giải pháp ảo hóa network. Trong bài này, chúng ta sẽ tìm hiểu về Linux Bridge.

- Linux bridge là một phần mềm đươc tích hợp vào trong nhân Linux để giải quyết vấn đề ảo hóa phần network trong các máy vật lý. Về mặt logic, Linux bridge sẽ tạo ra một con switch ảo để cho các VM kết nối được vào và có thể nói chuyện được với nhau cũng như sử dụng để ra mạng ngoài.
- Chú thích:

  - Bridge: tương đương với switch layer 2.

  - Port: tương đương với port của switch thật.

  - Tap (tap interface): có thể hiểu là giao diện mạng để các VM kết nối với bridge cho linux bridge tạo ra.

  - fd (forward data): chuyển tiếp dữ liệu từ máy ảo tới bridge.

- Các tính năng chính:

  - STP: Spanning Tree Protocol – giao thức chống lặp gói tin trong mạng.

  - VLAN: chia switch (do linux bridge tạo ra) thành các mạng LAN ảo, cô lập traffic giữa các VM trên các VLAN khác nhau của cùng một switch.

  - FDB (forwarding database): chuyển tiếp các gói tin theo database để nâng cao hiệu năng switch. Database lưu các địa chỉ MAC mà nó học được. Khi gói tin Ethernet đến, bridge sẽ tìm kiếm trong database có chứa MAC address không. Nếu không, nó sẽ gửi gói tin đến tất cả các cổng.

### 2.5 Tính năng 

- Security

  - Trong kiến trúc KVM, máy ảo được xem như các tiến trình Linux thông thường, nhờ đó nó tận dụng được mô hình bảo mật của hệ thống Linux như SELinux, cung cấp khả năng cô lập và kiểm soát tài nguyên. Bên cạnh đó còn có SVirt project – dự án cung cấp giải pháp bảo mật MAC (Mandatory Access Control – Kiểm soát truy cập bắt buộc) tích hợp với hệ thống ảo hóa sử dụng SELinux để cung cấp một cơ sở hạ tầng cho phép người quản trị định nghĩa nên các chính sách để cô lập các máy ảo. Nghĩa là SVirt sẽ đảm bảo rằng các tài nguyên của máy ảo không thể bị truy cập bởi bất kì các tiến trình nào khác; việc này cũng có thể thay đổi bởi người quản trị hệ thống để đặt ra quyền hạn đặc biệt, nhóm các máy ảo với nhau chia sẻ chung tài nguyên.

- Memory Management

  - KVM thừa kế tính năng quản lý bộ nhớ mạnh mẽ của Linux. Vùng nhớ của máy ảo được lưu trữ trên cùng một vùng nhớ dành cho các tiến trình Linux khác và có thể swap. KVM hỗ trợ NUMA (Non-Uniform Memory Access – bộ nhớ thiết kế cho hệ thống đa xử lý) cho phép tận dụng hiệu quả vùng nhớ kích thước lớn. KVM hỗ trợ các tính năng ảo của mới nhất từ các nhà cung cấp CPU như EPT (Extended Page Table) của Microsoft, Rapid Virtualization Indexing (RVI) của AMD để giảm thiểu mức độ sử dụng CPU và cho thông lượng cao hơn. KVM cũng hỗ trợ tính năng Memory page sharing bằng cách sử dụng tính năng của kernel là Kernel Same-page Merging (KSM).

- Storage

  - KVM có khả năng sử dụng bất kỳ giải pháp lưu trữ nào hỗ trợ bởi Linux để lưu trữ các Images của các máy ảo, bao gồm các ổ cục bộ như IDE, SCSI và SATA, Network Attached Storage (NAS) bao gồm NFS và SAMBA/CIFS, hoặc SAN thông qua các giao thức iSCSI và Fibre Channel. KVM tận dụng được các hệ thống lưu trữ tin cậy từ các nhà cung cấp hàng đầu trong lĩnh vực Storage. KVM cũng hỗ trợ các images của các máy ảo trên hệ thống tệp tin chia sẻ như GFS2 cho phép các images có thể được chia sẻ giữa nhiều host hoặc chia sẻ chung giữa các ổ logic.

- Live migration

  - KVM hỗ trợ live migration cung cấp khả năng di chuyển ác máy ảo đang chạy giữa các host vật lý mà không làm gián đoạn dịch vụ. Khả năng live migration là trong suốt với người dùng, các máy ảo vẫn duy trì trạng thái bật, kết nối mạng vẫn đảm bảo và các ứng dụng của người dùng vẫn tiếp tục duy trì trong khi máy ảo được đưa sang một host vật lý mới. KVM cũng cho phép lưu lại trạng thái hiện tại của máy ảo để cho phép lưu trữ và khôi phục trạng thái đó vào lần sử dụng tiếp theo.

- Performance and scalability

  - KVM kế thừa hiệu năng và khả năng mở rộng của Linux, hỗ trợ máy ảo với 16 CPUs ảo, 256GB RAM và hệ thống máy host lên tới 256 cores và trên 1TB RAM.

## Nguồn tham khảo:

https://blog.cloud365.vn/linux%20tutorial/tong-quan-lvm/

https://ductam.info/tong-quan-ve-kvm/#:~:text=KVM%20(Kernel%2Dbased%20virtual%20machine,%C4%91%E1%BB%83%20ch%E1%BA%A1y%20c%C3%A1c%20m%C3%A1y%20%E1%BA%A3o.




