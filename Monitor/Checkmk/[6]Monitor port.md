# Giám sát port TCP

“Port” là một thuật ngữ chuyên ngành được sử dụng phổ thông trong network, system nói riêng và hệ thống nói chung. Đối với một dịch vụ được thiết kế chạy trên một port có số hiệu khác nhau. Phân loại port thành 2 loại phổ dựa trên giao thức mà port đó sử dụng là TCP port hay UDP port.

Mỗi dịch vụ sẽ chạy trên một port khác nhau. Chính vì vậy muốn biết dịch vụ có chạy ổn định hay không, có đang hoạt động chính xác cho phép mở kết nối tới client hay không đơn giản là giám sát trạng thái của các port. Bài viết này sẽ hướng dẫn bạn các bước để thiết lập giám sát TCP port sử dụng checkmk.

## Cài đặt

Trước tiên bạn cần cài agent trên máy muốn giám sát port. Bạn có thể tham khảo các cài đặt agent [tại đây](https://github.com/phancong0897/Congphan/blob/master/Monitor/Checkmk/%5B2%5DAdd%20agent%20linux.md#h%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-c%C3%A0i-%C4%91%E1%BA%B7t-agent-c%E1%BB%A7a-checkmk-tr%C3%AAn-centos-7)

- Truy cập vào Web UI để thực hiện các bước tiếp: 

    - Truy cập setup --> services

    ![Imgur](https://imgur.com/qmPFgCE.png)

    - Check TCP port connection

    ![Imgur](https://imgur.com/BWB8fPl.png)

    - Tạo rule monitor port

    [Imgur](https://imgur.com/c37sMpI.png)

    - Tham số khi tạo rule
    
    Mục 1: Mô tả rule

    Mục 2: Comment

    Mục 3: Port muốn check (ở đây tôi check port của mysql 3306)

    Mục 4: Một số tham số tùy chọn bổ sung
    
    Mục 5: Điều kiện áp dụng rule. Bạn có thể xác định rule này được áp dùng cho thư mục nào, những host có tag xác định, những host xác định hoặc áp dụng cho thư mục nhưng bỏ qua một số host xác định. Ở đây tôi để rule này chỉ áp dụng cho Host_01

    Sau đó lưu

    Một số tham số tùy chọn bổ sung ở mục 4 như sau:

    Service description: mô tả dịch vụ của port này
    
    DNS hostname: sử dụng domain thay vì IP theo mặc định
    
    Expected response time: xác đinh thời gian response để xác ddingj trạng thái OK hay Warning hoặc Critical
    
    Seconds before connection times out: Xác định thời gian trước khi kết nối times out.
    
    State for connection refusal: Trạng thái nếu kết nối bị từ chối
    
    Strings to expect in response: Chuỗi mong muốn trong kết quả trả về
    
    Maximum number of bytes to receive: Số byte lớn nhất có thể nhận
    
    Use SSL for the connection: Sử dụng SSL cho kết nối
    
    ![Imgur](https://imgur.com/LIi1mbc.png)

    ![Imgur](https://imgur.com/Jwc2lCo.png)

    - Cập nhật thay đổi

    ![Imgur](https://imgur.com/twCgpRd.png)

- Quay trở lại Host ta thấy port 80 đã được thêm vào giám sát

![Imgur](https://imgur.com/2K78aR9.png)

- Cảnh báo telegram

![Imgur](https://imgur.com/IZ54uK4.png)

