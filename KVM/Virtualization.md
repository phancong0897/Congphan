# Tổng quan về ảo hóa



### 1.Ảo hóa là gì?

- Ảo hóa là công nghệ cho phép khai thác triệt để khả năng hoạt động của các phần cứng trong hệ thống máy chủ bằng cách chạy đồng thời nhiều OS trên cùng lớp vật lý.

- Cùng chia sẻ tài nguyên phần cứng và được quản lý bởi lớp ảo hóa (Hypervisor).

- Lớp ảo hóa nằm giữa như một tầng trung gian giữa phần cứng (hardware) và phần mềm hệ điều hành (OS) giúp quản lý, phân phát tài nguyên phần cứng cho lớp OS ảo hoạt động ở trên.

Nó hoạt động như một tầng trung gian giữa phần cứng máy tính và phần mềm chạy trên nó. Ý tưởng của công nghệ máy chủ ảo hóa là từ một máy PC đơn lẻ có thể tạo thành nhiều máy ảo riêng biệt. Ảo hóa cho phép tạo nhiều máy ảo trên một máy chủ vật lý, mỗi một máy ảo cũng được cấp phát tài nguyên phần cứng như máy thật gồm có Ram, CPU, ổ cứng, Card mạng, các tài nguyên khác và hệ điều hành riêng. Khi chạy ứng dụng, người sử dụng không nhận biết được ứng dụng đó chạy trên lớp phần cứng ảo.

### 2. Các thành phần của một hệ thống ảo hóa

![Imgur](https://imgur.com/pmibevj.png)

- Tài nguyên vật lý chính (Host machine / Host hardwave): Máy chủ vật lý, CPU, RAM, ổ đĩa cứng, card mạng… Nhiệm vụ là chia tài nguyên cấp cho các máy ảo.

- Phần mềm ảo hóa (Hypervisor): cung cấp truy cập cho mỗi máy chủ ảo đến tài nguyên của máy chủ vật lý, lập kế hoạch và phân chia tài nguyên vật lý cho các máy chủ ảo, cung cấp giao diện quản lý cho các máy chủ ảo

- Hệ điều hành khách (Guest Operating System): được cài đặt trên một máy chủ ảo, thao tác như ở trên hệ điều hành thông thường.

- Mảy ảo (Virtual Machine): nó hoạt động như một máy chủ vật lý thông thường với tài nguyên riêng, giao diện riêng, hệ điều hành riêng.

- Một hệ thống ảo hóa bắt buộc phải có đầy đủ các thành phần: tài nguyên vật lý, phần mềm ảo hóa, máy chủ ảo và hệ điều hành khách. Khi có đầy đủ 4 thành phần của hệ thống ảo hóa, người dùng có thể dễ dàng xây dựng cho mình một hệ thống ứng dụng ảo hóa hoàn chỉnh.

### 3. Ảo hóa hoạt động như thế nào?

Ảo hóa được xây dựng dựa trên giải pháp chia một máy vật lý thành nhiều máy con. Giải pháp này được biết đến với cái tên là Virtual Machine Monitor (VMM) hay thường được gọi là Hypervisor. VMM cho phép tạo tách rời các máy ảo và điều phối truy cập của các máy ảo này đến tài nguyên phần cứng và cấp phát tài nguyên tự động theo nhu cầu sử dụng. Cấu trúc này giúp cân bằng khả năng điện toán để mang lại:

- Nhiều ứng dụng chạy trên cùng một server, mỗi máy ảo được lập trình trên máy chủ, do đó nhiều ứng dụng và các hệ điều hành có thể cùng lúc chạy trên một host.

- Tối đa hóa công suất sử dụng và tối thiếu hóa server: Mỗi máy chủ vật lý được sử dụng với đầy đủ công suất, cho phép giảm đáng kể chi phí nhờ sử dụng tối đa server.

- Cấp phát tài nguyên và ứng dụng nhanh chóng, dễ dàng. Máy ảo được triển khai từ một file chứa đầy đủ phần mềm với cơ chế đơn giản là copy và Điều này mang đến sự đơn giản, nhanh chóng và linh hoạt chưa từng có
cho việc quản lý và cung cấp hạ tầng Công nghệ thông tin. Máy ảo thậm chí cóthể di chuyển sang một server vật lý khác trong khi vẫn đang chạy, hoạt động bình thường. Doanh nghiệp có thể ảo hóa những ứng dụng quan trọng của doanh nghiệp để nâng cao hiệu suất, sự ổn định, khả năng mở rộng và giảm thiểu chi phí.

### 4. Mục tiêu của ảo hóa

Ảo hóa xoay quanh 4 mục tiêu chính: Availability, Scalability, Optimization, Management.

- Availability: giúp các ứng dụng hoạt động liên tục bằng cách giảm thiểu (bỏ qua) thời gian chết (downtime) khi phần cứng gặp sự cố, khi nâng cấp hoặc di chuyển.

- Scalability:  khả năng tùy biến, thu hẹp hay mở rộng mô hình server dễ dàng mà không làm gián đoạn ứng dụng.

- Optimization: sử dụng triệt để nguồn tài nguyên phần cứng và tránh lãng phí bằng cách giảm số lượng thiết bị vật lý cần thiết (giảm số lượng server, switch, cáp, v.v. )

- Management: khả năng quản lý tập trung, giúp việc quản lý trở nên dễ dàng hơn bao giờ hết.

## Lợi ích và nhược điểm của ảo hóa

### Lợi ích:

- Tiết kiệm năng lượng tiêu thụ, giảm chi phí duy trì server (tiền điện để chạy và làm mát server)

- Giảm số lượng thiết bị vật lý cần thiết (giảm số lượng server, switch, cáp, phí gia công)

- Tận dụng tối đa nguồn tài nguyên, tránh lãng phí.

- Quản lý tập trung, liên tục, nâng cao hiệu quả làm việc của quản trị viên.

- Khả năng mở rộng dể dàng

### Nhược điểm:

- Nếu máy chủ có cấu hình phần cứng thấp nhưng lại có một máy ảo sử dụng quá nhiều tài nguyên hoặc chạy quá nhiều máy ảo sẽ làm chậm toàn bộ hệ thống bao gồm các máy ảo và các ứng dụng chạy trên máy ảo. Đồng thời do một hoặc vài máy chủ phải đảm nhận nhiều máy ảo chạy trên nó nên máy chủ gặp trục trặc, sự cố thì các máy ảo cũng sẽ bị ảnh hưởng theo.

- Còn ở góc độ bảo mật, nếu hacker nắm quyền điều khiển một máy chủ vật lý chứa các máy ảo thì hacker có thể kiểm soát được tất cả các máy ảo trong nó.



