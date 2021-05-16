# Tổng quan về ảo hóa



### 1.Ảo hóa là gì?

- Ảo hóa là công nghệ cho phép khai thác triệt để khả năng hoạt động của các phần cứng trong hệ thống máy chủ bằng cách chạy đồng thời nhiều OS trên cùng lớp vật lý.

- Cùng chia sẻ tài nguyên phần cứng và được quản lý bởi lớp ảo hóa (Hypervisor).

- Lớp ảo hóa nằm giữa như một tầng trung gian giữa phần cứng (hardware) và phần mềm hệ điều hành (OS) giúp quản lý, phân phát tài nguyên phần cứng cho lớp OS ảo hoạt động ở trên.

Nó hoạt động như một tầng trung gian giữa phần cứng máy tính và phần mềm chạy trên nó. Ý tưởng của công nghệ máy chủ ảo hóa là từ một máy PC đơn lẻ có thể tạo thành nhiều máy ảo riêng biệt. Ảo hóa cho phép tạo nhiều máy ảo trên một máy chủ vật lý, mỗi một máy ảo cũng được cấp phát tài nguyên phần cứng như máy thật gồm có Ram, CPU, ổ cứng, Card mạng, các tài nguyên khác và hệ điều hành riêng. Khi chạy ứng dụng, người sử dụng không nhận biết được ứng dụng đó chạy trên lớp phần cứng ảo.

### 2. Các thành phần của một hệ thống ảo hóa

[Imgur](https://imgur.com/pmibevj)

- Tài nguyên vật lý chính (Host machine / Host hardwave): Máy chủ vật lý, CPU, RAM, ổ đĩa cứng, card mạng… Nhiệm vụ là chia tài nguyên cấp cho các máy ảo.

- Phần mềm ảo hóa (Hypervisor): cung cấp truy cập cho mỗi máy chủ ảo đến tài nguyên của máy chủ vật lý, lập kế hoạch và phân chia tài nguyên vật lý cho các máy chủ ảo, cung cấp giao diện quản lý cho các máy chủ ảo

- Hệ điều hành khách (Guest Operating System): được cài đặt trên một máy chủ ảo, thao tác như ở trên hệ điều hành thông thường.

- Mảy ảo (Virtual Machine): nó hoạt động như một máy chủ vật lý thông thường với tài nguyên riêng, giao diện riêng, hệ điều hành riêng.

- Một hệ thống ảo hóa bắt buộc phải có đầy đủ các thành phần: tài nguyên vật lý, phần mềm ảo hóa, máy chủ ảo và hệ điều hành khách. Khi có đầy đủ 4 thành phần của hệ thống ảo hóa, người dùng có thể dễ dàng xây dựng cho mình một hệ thống ứng dụng ảo hóa hoàn chỉnh.



