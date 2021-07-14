# Tìm hiểu về ELK stack

### 1. Tổng quan

- ELK là viết tắt tên của 3 thành phần chính của hệ thống này: E – Elasticsearch , L – Logstash, K – Kibana, phục vụ cho việc giám sát hệ thống. ELK được phát triển từ đầu những năm 2000 và cho đến nay đã có hơn 250 triệu lượt tải xuống và sử dụng. Hiện tại phiên bản mới nhất của ELK là 7.8. Đậy là công cụ giám sát tập trung mã nguồn mở rất mạnh, có thể xử lý rất nhiều bài toán quản lý hệ thống mạng nên rất được các công ty, tổ chức tin dùng.

- ELK stack là tập hợp 3 phần mềm đi chung với nhau, phục vụ cho công việc logging. 3 phần mềm này lần lượt là:

    - Elasticsearch: Cơ sở dữ liệu để lưu trữ, tìm kiếm và query log.

    - Logstash: Tiếp nhận log từ nhiều nguồn, sau đó xử lý log và ghi dư liệu vào Elasticsearch

    - Kibana: Giao diện để quản lý, thống kê log. Đọc thông tin từ Elasticsearch

- Điểm mạnh của ELK là khả năng thu thập, hiển thị, truy vấn theo thời gian thực. Có thể đáp ứng truy vấn một lượng dữ liệu cực lớn.

- Hình dưới đây là luồng hoạt động của ELK stack:

    ![Imgur](https://imgur.com/c8MsbUO.png)

    Đầu tiên, log sẽ được đưa đến Logstash. (Thông qua nhiều con đường, ví dụ như server gửi UDP request chứa log tới URL của Logstash, hoặc Beat đọc file log và gửi lên Logstash).

    Logstash sẽ đọc những log này, thêm những thông tin như thời gian, IP, parse dữ liệu từ log (server nào, độ nghiêm trọng, nội dung log) ra, sau đó ghi xuống Elasticsearch.

    Khi muốn xem log, người dùng vào URL của Kibana. Kibana sẽ đọc thông tin log trong Elasticsearch, hiển thị lên giao diện cho người dùng query và xử lý.

### 2. Các thành phần và nguyên lí hoạt động ELK Stack

### 2.1 Các thành phần

-	Elasticsearch: Dùng để tìm kiếm và query log.

-	Logstash: Tiếp nhận log từ nhiều nguồn, sau đó xử lý log và ghi vào cơ sở dữ liệu.

-	Kibana: Giao diện để quản lý, thống kê log. Đọc thông tin từ Elasticsearch.

-	Beats: Một tập các công cụ chuyên dùng để thu thập dữ liệu cực mạnh.

-	Elastalert: Dùng để tạo các alert cảnh báo cho hệ thống.
 
### 2.2 Nguyên lí hoạt động

### 2.2.1 ELK thu thập, phân tích và cảnh báo sự kiện.
-	Đầu tiền, thông tin cần giám sát sẽ được đưa đến máy chủ ELK thông qua nhiều con đường, ví dụ như server gửi UDP request chứa log tới URL của logstash, hoặc Beats thu thập các thông tin từ các bộ công cụ chuyên dụng cài trên các server và gửi lên Logstash hoặc Elasticsearch.

-	Logstash sẽ đọc những log này, thêm những thông tin như thời gian, IP, parse dữ liệu từ log ra, sau đó ghi xuống database là Elasticsearch. Ngoài ra, Logstash còn có thể được sử dụng như 1 Threat Intelligent. Nó sẽ đối chiếu các trường thông tin nó nhận được với cơ sở dữ liệu các thông tin độc hại/ vi phạm chính sách,…

-	Khi muốn xem log, người dùng vào URL của Kibana, Kibana sẽ đọc thông tin log trong Elasticsearch, hiển thị giao diện cho người dùng query và xử lý. Kibana hiển thị thông tin log cho người dùng.

-	Elastalert sẽ đảm nhận nhiệm vụ kiểm tra toàn bộ các sự kiện được lưu trong Elasticsearch, nếu phát hiện ra sự kiện nào bất thường dựa vào các tập rules được thiết lập, nó sẽ gửi cảnh báo cho người quản trị các phương thức mà người quản trị mong muốn.

### 2.2.2 ELK giao tiếp giữa các thành phần thông qua kênh mã hoá

- ELK Stack hỗ trợ tính năng mã hoá đường truyền giao tiếp giữa các thành phần khi chuyển tiếp dữ liệu về cơ sở dữ liệu cũng như truy vấn dữ liệu. Nó sử dụng phương thức mã hoá SSL/TLS
Khi tính năng năng bảo mật của Elasticsearch được sử dụng thì mọi kết nối tới elasticsearch đều phải được xác thực bao gồm cả các kết nối từ Kibana hay các máy chủ ứng dụng khác.

- Elasticsearch có 2 level kết nối:

    -	Transport Communications: lớp này được sử dụng cho các kết nối nội bộ giữa các Elasticsearch nodes.
    
    -	HTTP Communications: giao thức HTTP này được sử dụng cho các kết nối từ các máy khách kết nối tới Elasticsearch Cluster.

    ![Imgur](https://imgur.com/S4fsh7L.png)

    (1): Phía Client gửi một yêu cầu kết nối tới cho Server

    (2): Phía Server phản hồi lại kết nối của phía Client, đồng thời gửi kèm theo chứng thư và các thông tin liên qua của Server. Ngoài ra, tại bước này, phía Server cũng có thể yêu cầu phía client gửi lại chứng thư của Client để định danh/ xác thực

    (3): Phía Client sẽ kiểm tra, xác thực/ tin tưởng các thông tin định danh của Server được xác định trong tệp chứng thư. 

    (4)(5)(6): Phía Client sau khi đã xác thực được thông tin của server, client sẽ tạo một khóa đối xứng và mã hóa khóa đó bằng khóa công khai của Server và gửi trả lại về phía Server. Server sẽ sử dụng khóa bí mật của mình để giải mã khóa đối xứng đã được mã hóa. Ngoài ra, phía Server kiểm tra, xác thực/ tin tưởng các thông tin định danh của Client để tạo kết nối với Client.

    (7)(8): Sau khi thực hiện hoàn tất các bước xác thực và tạo khóa mã hóa, phía Client và Server sẽ hoàn tất việc tạo một kênh kết nối mã hóa với nhau.
    
    (9): Khóa đối xứng sẽ được sử dụng cho việc mã hoá, giải mã các thông tin được truyền tải giữa phía Client và Server giúp cho mọi thông tin được bảo mật.


