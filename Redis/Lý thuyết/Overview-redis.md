# Tìm hiểu tổng quan về Redis

### 1. Redis là gì?

Redis là tên viết tắt của Remote Dictionary Server (Máy chủ từ điển từ xa), là kho dữ liệu key-value, trong bộ nhớ, mã nguồn mở và có tốc độ truy cập nhanh để dùng làm cơ sở dữ liệu, cache, message broker và queue. Redis có thể được sử dụng như một database được lưu trữ trong ram để tăng tốc độ xử lí.

Redis hiện cung cấp thời gian phản hồi ở tốc độ chưa đến một mili giây, vì Redis cho phép bạn lưu trữ trên RAM của mình nhanh hơn so với khi lưu trữ trên các ổ cứng, khoảng 150.000 lần so với truy cập ổ HDD và nhanh hơn 500 lần so với truy cập SSD.

Thông thường khi ta sử dụng cơ sở dữ liệu thì vẫn lưu trữ trên ổ đĩa cứng. Ví dụ khi query một cơ sở dữ liệu để đọc 10.000 bản ghi. Nếu dữ liệu được lưu trữ trên đĩa, thì nó sẽ mất trung bình 30 giây, trong khi chỉ mất khoảng 0,0002 giây để đọc từ RAM, đó chính là lí do mà Redis dùng RAM để lưu trữ. Redis thường được tạo trên một máy chủ riêng hoặc set giới hạn bộ nhớ nhất định được sử dụng trên máy chủ dùng chung cho nên sẽ không bị tình trạng thiếu RAM khi chạy đồng thời các ứng dụng khác.

Tuy nhiên, khi lưu trữ trên RAM thì nhanh thật nhưng lại gặp 1 vấn đề là bị mất điện thì dữ liệu cũng mất tiêu. Để ngăn chặn việc mất dữ liệu xảy ra, có một mô-đun được tích hợp sẵn để ghi trạng thái trong bộ nhớ vào file trên đĩa trong những trường hợp nhất định. Các file này được tải lại khi khởi động lại redis. Vì vậy, dữ liệu sẽ không bị mất. Ngoài ra, để tăng tính sẵn sàng và khả năng chịu lỗi của hệ thống. Redis có thể cấu hình theo dạng Cluster với kỹ thuật Master-Slave giúp hệ thống redis luôn sẵn sàng đáp ứng sao lưu dữ liệu trên đĩa cứng và phục hồi dữ liệu khi gặp sự cố.

### 2. Lợi ích của Redis

#### In-Memory Data Store

Toàn bộ dữ liệu Redis nằm trong bộ nhớ chính của máy chủ, khác với các cơ sở dữ liệu thường lưu phần lớn dữ liệu trên ổ đĩa hoặc ổ SSD. So với cơ sở dữ liệu trên ổ đĩa truyền thống trong đó phần lớn các tác vụ đều yêu cầu truy cập lặp lại tới ổ đĩa, còn Redis thì chỉ việc lấy ra kết quả đã tính toán. Do đó hiệu suất nhanh thấy rõ rệt với các tác vụ đọc hoặc ghi thông thường mất chưa đầy một mili giây và hỗ trợ hàng triệu tác vụ mỗi giây.

#### Cấu trúc dữ liệu linh hoạt

- Khác với những kho dữ liệu key-value đơn giản có cấu trúc dữ liệu giới hạn, Redis có nhiều cấu trúc dữ liệu khác nhau nên đáp ứng được nhu cầu ứng dụng của bạn. Kiểu dữ liệu Redis gồm có:

    - String – văn bản hoặc dữ liệu nhị phân có kích thước lên tới 512MB

        ```
        127.0.0.1:6379> auth redis@123
        OK
        127.0.0.1:6379> set myname congpv
        OK
        127.0.0.1:6379> get myname
        "congpv"

        ```

    - List – một tập hợp các chuỗi được sắp xếp theo thứ tự như khi được thêm vào
    
    - Set – một tập hợp chưa được sắp xếp các chuỗi, có khả năng giao cắt, liên kết và khác với các kiểu tập khác
    
    - Sorted Set – Tập được sắp xếp theo giá trị
    
    - Hash – một cấu trúc dữ liệu dùng để lưu trữ danh sách các trường và giá trị
    
    - Bitmap – một kiểu dữ liệu cho phép thực hiện các tác vụ quy mô bit

    - HyperLogLogs – một cấu trúc dữ liệu xác suất để ước tính các thành phần duy nhất trong một tập dữ liệu

#### Đơn giản và dễ sử dụng

Redis đơn giản hóa code của bạn bằng cách cho phép bạn viết ít dòng code hơn để lưu trữ, truy cập và sử dụng dữ liệu trong các ứng dụng của mình. Ví dụ: nếu ứng dụng của bạn có dữ liệu được lưu dạng hashmap và bạn muốn lưu trữ dữ liệu đó trong DB thì bạn chỉ cần sử dụng cấu trúc dữ liệu hashes để lưu chúng. Thao tác tương tự trên DB không có cấu trúc dữ liệu hashes sẽ yêu cầu nhiều dòng code để chuyển đổi từ định dạng này sang định dạng khác. Với mỗi kiểu dữ liệu thì Redis thêm nhiều tùy chọn để thao tác và tương tác với dữ liệu của bạn.

#### Replication - Persistence

Redis sử dụng kiến trúc primary-replica và hỗ trợ sao chép không đồng bộ, dữ liệu có thể được sao chép sang nhiều máy chủ khác nhau. Điều này giúp hiệu suất đọc được cải thiện (vì các yêu cầu có thể được phân chia giữa các máy chủ) và phục hồi nhanh hơn khi máy chủ chính gặp sự cố ngừng hoạt động. Để duy trì, Redis hỗ trợ sao lưu point-in-time backups (sao chép dữ liệu Redis vào đĩa).

#### Độ khả dụng cao và quy mô linh hoạt

Redis có kiến trúc primary-replica theo cấu trúc liên kết dạng một nút chính hoặc cụm. Kiến trúc này cho phép bạn xây dựng những giải pháp có độ khả dụng cao, đảm bảo hiệu suất ổn định và tin cậy. Khi cần điều chỉnh kích thước cụm, bạn có nhiều tùy chọn khác nhau để thay đổi quy mô theo chiều dọc và thay đổi quy mô theo chiều ngang. Việc này cho phép tăng cụm theo nhu cầu của bạn.

#### Khả năng mở rộng

Redis là dự án mã nguồn mở đã có cộng đồng lớn tin dùng. Không có giới hạn về nhà cung cấp hoặc công nghệ vì Redis được có tính tiêu chuẩn mở, hỗ trợ các định dạng dữ liệu mở và có tập hợp máy khách phong phú.

### 3. Một số ứng dụng của Redis

#### CACHING

Redis là một lựa chọn tuyệt vời để triển khai in-memory cache để giảm độ trễ truy cập dữ liệu, tăng thông lượng và giảm tải khỏi cơ sở dữ liệu quan hệ hoặc NoSQL trong ứng dụng của bạn. Redis có thể phục vụ các dữ liệu được yêu cầu thường xuyên ở thời gian phản hồi dưới một mili giây. Và cho phép bạn dễ dàng mở rộng quy mô cho các tải cao hơn mà không cần tăng phụ trợ tốn kém. Persistent session caching, web page caching và caching của các đối tượng được sử dụng thường xuyên như hình ảnh, tệp và siêu dữ liệu là các ví dụ phổ biến về caching với Redis.

#### SESSION STORE

Redis như một kho lưu trữ dữ liệu trong bộ nhớ với tính sẵn sàng cao và bền bỉ là lựa chọn phổ biến của các nhà phát triển ứng dụng để lưu trữ và quản lý dữ liệu session cho các ứng dụng của mình. Redis cung cấp độ trễ dưới một phần nghìn giây và khả năng phục hồi cần thiết để quản lý dữ liệu session như hồ sơ người dùng, thông tin đăng nhập, trạng thái phiên và cá nhân hóa cụ thể của người dùng.

#### REAL-TIME ANALYTICS

Redis có thể được sử dụng với các giải pháp streaming như Apache Kafka và Amazon Kinesis như một in-memory data store để nhập, xử lý và phân tích dữ liệu real-time với độ trễ nhỏ hơn một mili giây. Redis là một lựa chọn lý tưởng cho các hệ thống real-time analytics như truyền thông xã hội, quảng cáo và IoT,...

### 4. Sử dụng Redis ở đâu và khi nào?

Redis chủ yếu được sử dụng cho mục đích lưu vào bộ nhớ đệm. Đây là thế mạnh lớn nhất của Redis vì nó cung cấp nhiều loại dữ liệu khác nhau với cấu trúc từ điên (key-value) của nó.

#### Full Page Cache (FPC)

Điều đó có nghĩa là, chúng ta có thể sử dụng Redis để lưu vào bộ đệm đầy đủ các trang có lưu lượng truy cập cao và nội dung tĩnh. Nếu bản thân trang có xu hướng thay đổi hoặc nội dung được tạo động, thì nó không nên được lưu vào bộ nhớ cache trong mọi trường hợp bất kể đó là Redis hay Memcached. Nếu trang chủ yếu chứa nội dung tĩnh hoặc nội dung được làm mới sau mỗi x khoảng thời gian, thì chúng ta có thể lưu vào bộ nhớ cache của trang đó và nó sẽ được phục vụ ngay lập tức!

#### Session Cache

Ngoài cache file tĩnh, Redis còn có thể dùng để lưu session (phiên đăng nhập). Không giống như Memcached không có tính bền bỉ (Persistence), Redis có cho phép lưu trữ session. Xử lý các phiên bằng Redis giúp tăng trải nghiệm người dùng khi tương tác với trang web.

Tính năng Persistence rất quan trọng để lưu trữ các phiên, để tránh mất dữ liệu trong các phần quan trọng của tương tác với người dùng. Ví dụ: xử lý thanh toán, thêm một mặt hàng vào giỏ hàng hoặc yêu cầu bất kỳ hành động nào với tư cách là người dùng đã xác thực.


### Link tham khảo

- https://viblo.asia/p/tim-hieu-ve-redis-LzD5dXEW5jY

- https://redis.io/commands#list
