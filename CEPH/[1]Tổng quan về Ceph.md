# Tổng Quan Về Ceph

### Giới Thiệu

- Ceph là dự án mã nguồn mở, cung cấp giải pháp lưu trữ dữ liệu. Ceph cung cấp giải pháp lưu trữ phân tán mạnh mẽ, đáp ứng các yêu cầu về khả năng mở rộng, hiệu năng, khả năng chịu lỗi cao. Xuất phát từ mục tiêu ban đầu, Ceph được thiết kết với khả năng mở rộng không giới hạn, hỗ trợ lưu trữ tới mức exabyte, cùng với khả năng tương thích cao với các phần cứng có sẵn. Và từ đó, Ceph trở nên nổi bật trong ngành công nghiệp lưu trữ đang phát triển và mở rộng. Hiện nay, các nền tảng hạ tầng đám mây công cộng, riêng và lai dần trở nên phổ biến và to lớn. Bên cạnh đó, phần cứng - thành phần quyết định hạ tầng đám mây đang dần gặp các vấn đề về hiệu năng, khả năng mở rộng. Ceph đã và đang giải quyết được các vấn đề doanh nghiệp đang gặp phải, cung cấp hệ thống, hạ tầng lưu trữ mạnh mẽ, độ tin cậy cao.

- Nguyên tắc cơ bản của Ceph:

    - Khả năng mở rộng tất cả thành phần;
    
    - Khả năng chịu lỗi cao;
    
    - Giải pháp dựa trên phần mềm, hoàn toàn mở, tính thích nghi cao;
    
    - Chạy tương thích với mọi phần cứng;

- Ceph xây dựng kiến trúc mạnh mẽ, khả năng mở rộng không giới hạn, hiệu năng cao, cung cấp giải pháp thống nhất, nền tảng mạnh mẽ cho doanh nghiệp, giảm bớt sự phụ thuộc vào phần cứng đắt tiền. Hệ sinh thái Ceph cung cấp giải pháp lưu trữ dựa trên block, file, object, và cho phép tùy chỉnh theo ý muốn. Nền tảng Ceph xây dựng dựa trên các object, tổ chức trên các block. Tất cả các kiểu dữ liệu, block, file đều được lưu dưới dạng object, quản trị bởi Ceph cluster. Object storage hiện đã trở thành giải pháp cho hệ thống lưu trữ truyền thống, cho phép xây dựng kiến trúc hạ tầng độc lập với phần cứng.

Ceph xây dựng giải pháp dựa trên Object, Các object được tổ chức, nhân bản trên toàn cluster, nâng cao tính bảo đảm dữ liệu. Tại Ceph, Object sẽ không tồn tại đường dẫn vật lý, toàn bộ Object được quản trị dưới dạng Key Object, tạo nền tảng mở với khả năng lưu trữ tới hàng petabyte-exabyte.

<h3 align="center"><img src="D:\Github\Congphan\CEPH\Images\Ceph storage cluster.png"></h3>