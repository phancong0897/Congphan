# Tổng quan về KVM

### 1. Khái niệm và vai trò

- KVM (Kernel-based virtual machine) là giải pháp ảo hóa cho hệ thống linux trên nền tảng phần cứng x86 có các module mở rộng hỗ trợ ảo hóa (Intel VT-x hoặc AMD-V). Về bản chất, KVM không thực sự là một hypervisor có chức năng giả lập phần cứng để chạy các máy ảo. Chính xác KVM chỉ là một module của kernel linux hỗ trợ cơ chế mapping các chỉ dẫn trên CPU ảo (của guest VM) sang chỉ dẫn trên CPU vật lý (của máy chủ chứa VM). Hoặc có thể hình dung KVM giống như một driver cho hypervisor để sử dụng được tính năng ảo hóa của các vi xử lý như Intel VT-x hay AMD-V, mục tiêu là tăng hiệu suất cho guest VM.

- KVM ban đầu được phát triển bởi Qumranet – một công ty nhỏ, sau đó được Redhat mua lại vào tháng 9 năm 2008. Ta có thể thấy KVM là thế hệ tiếp theo của công nghệ ảo hóa. KVM được sử dụng mặc định từ bản RHEL (Redhat Enterprise Linux) từ phiên bản 5.4 và phiên bản Redhat Enterprise Virtualization dành cho Server.

- Qumranet phát hành mã của KVM cho cộng đồng mã nguồn mở. Hiện nay, các công ty nổi tiếng như IBM, Intel và ADM cũng đã cộng tác với dự án. Từ phiên bản 2.6.20, KVM trở thành một phần của hạt nhân Linux.

### 2. Cấu trúc của KVM

- Trong kiến trúc KVM, máy ảo là một tiến trình Linux, được lập lịch bởi chuẩn Linux scheduler. Trong thực tế mỗi CPU ảo xuất hiện như là một tiến trình Linux. Điều này cho phép KVM sử dụng tất cả các tính năng của Linux kernel.

- Cấu trúc tổng quan:

            ![Imgur](https://imgur.com/JWNSsCd.png)

    - Linux có tất cả các cơ chế của một VMM cần thiết để vận hành (chạy) các máy ảo. Chính vì vậy các nhà phát triển không xây dựng lại mà chỉ thêm vào đó một vài thành phần để hỗ trợ ảo hóa. KVM được triển khai như một module hạt nhân có thể được nạp vào để mở rộng Linux bởi những khả năng này.

    - Trong một môi trường linux thông thường mỗi process chạy hoặc sử dụng user-mode hoặc kernel-mode. KVM đưa ra một chế độ thứ 3, đó là guest-mode. Nó dựa trên CPU có khả năng ảo hóa với kiến trúc Intel VT hoặc AMD SVM, một process trong guest-mode bao gồm cả kernel-mode và user-mode.

- Trong cấu trúc tổng quan của KVM bao gồm 3 thành phần chính:

    - KVM kernel module:

    Là một phần trong dòng chính của Linux kernel.

    Cung cấp giao diện chung cho Intel VMX và AMD SVM (thành phần hỗ trợ ảo hóa phần cứng)

    Chứa những mô phỏng cho các instructions và CPU modes không được support bởi Intel VMX và AMD SVM

    - quemu – kvm: là chương trình dòng lệnh để tạo các máy ảo, thường được vận chuyển dưới dạng các package “kvm” hoặc “qemu-kvm”. Có 3 chức năng chính:

    Thiết lập VM và các thiết bị ra vào (input/output)

    Thực thi mã khách thông qua KVM kernel module

    Mô phỏng các thiết bị ra vào (I/O) và di chuyển các guest từ host này sang host khác

    - libvirt management stack:

    Cung cấp API để các tool như virsh có thể giao tiếp và quản lí các VM

    Cung cấp chế độ quản lí từ xa an toàn

### 3. Mô hình vận hành

- Hình dưới đây mô tả mô hình vận hành của KVM. Đây là một vòng lặp của các hành động diễn ra để vận hành các máy ảo. Những hành động này được phân cách bằng 3 phương thức chúng ta đã đề cập trước đó: user-mode, kernel-mode, guest-mode.

            ![Imgur](https://imgur.com/nIhWWMs.png)

- Như ta thấy:

    - User-mode: Các modul KVM gọi đến sử dụng ioclt() để thực thi mã khách cho đến khi hoạt động I/O khởi xướng bởi guest hoặc một sự kiện nào đó bên ngoài xảy ra. Sự kiện này có thể là sự xuất hiện của một gói tin mạng, cũng có thể là trả lời của một gói tin mạng được gửi bởi các may chủ trước đó. Những sự kiện như vậy được biểu diễn như là tín hiệu dẫn đến sự gián đoạn của thực thi mã khách.

    - Kernel-mode: Kernel làm phần cứng thực thi các mã khách tự nhiên. Nếu bộ xử lý thoát khỏi guest do cấp phát bộ nhớ hay I/O hoạt động, kernel thực hiện các nhiệm vụ cần thiết và tiếp tục luồng thực hiện, Nếu các sự kiện bên ngoài như tín hiệu hoặc I/O hoạt động khởi xướng bởi các guest tồn tại, nó thoát tới user-mode.

    - Guest-mode: Đây là cấp độ phần cứng, nơi mà các lệnh mở rộng thiết lập của một CPU có khả năng ảo hóa được sử dụng để thực thi mã nguồn gốc, cho đến khi một lệnh được gọi như vậy cần sự hỗ trợ của KVM, một lỗi hoặc một gián đoạn từ bên ngoài.

Khi một máy ảo chạy, có rất nhiều chuyển đổi giữa các chế độ. Từ kernel-mode tới guest-mode và ngược lại rất nhanh, bởi vì chỉ có mã nguồn gốc được thực hiện trên phần cứng cơ bản. Khi I/O hoạt động diễn ra các luồng thực thi tới user-mode, rất nhiều thiết bị ảo I/O được tạo ra, do vậy rất nhiều I/O thoát ra và chuyển sang chế độ user-mode chờ. Hãy tưởng tượng mô phỏng một đĩa cứng và 1 guest đang đọc các block từ nó. Sau đó QEMU mô phỏng các hoạt động bằng cách giả lập các hoạt động bằng các mô phỏng hành vi của các ổ đĩa cứng và bộ điều khiển nó được kết nối. Để thực hiện các hoạt động đọc, nó đọc các khối tương ứng từ một tập tin lớp và trả về dữ liệu cho guest. Vì vậy, user-mode giả lập I/O có xu hướng xuất hiện một nút cổ chai làm chậm việc thực hiện máy ảo.

### 4. Cơ chế hoạt động

Để các máy ảo giao tiếp được với nhau, KVM sử dụng Linux Bridge và OpenVSwitch. Đây là 2 phần mềm cung cấp các giải pháp ảo hóa network. 

Trong bài này, chúng ta sẽ tìm hiểu về Linux Bridge.

- Linux bridge là một phần mềm đươc tích hợp vào trong nhân Linux để giải quyết vấn đề ảo hóa phần network trong các máy vật lý. Về mặt logic, Linux bridge sẽ tạo ra một con switch ảo để cho các VM kết nối được vào và có thể nói chuyện được với nhau cũng như sử dụng để ra mạng ngoài.

- Cấu trúc của Linux Bridge khi kết hợp với KVM – QEMU

            ![Imgur](https://imgur.com/EOH0zQf.png)

- Chú thích:

    - Bridge: tương đương với switch layer 2

    - Port: tương đương với port của switch thật

    - Tap (tap interface): có thể hiểu là giao diện mạng để các VM kết nối với bridge cho linux bridge tạo ra

    - fd (forward data): chuyển tiếp dữ liệu từ máy ảo tới bridge
- Các tính năng chính:

    - STP: Spanning Tree Protocol – giao thức chống lặp gói tin trong mạng

    - VLAN: chia switch (do linux bridge tạo ra) thành các mạng LAN ảo, cô lập traffic giữa các VM trên các VLAN khác nhau của cùng một switch.

    - FDB (forwarding database): chuyển tiếp các gói tin theo database để nâng cao hiệu năng switch. Database lưu các địa chỉ MAC mà nó học được. Khi gói tin Ethernet đến, bridge sẽ tìm kiếm trong database có chứa MAC address không. Nếu không, nó sẽ gửi gói tin đến tất cả các cổng.
