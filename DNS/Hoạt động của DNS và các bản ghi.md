# Cơ chết hoạt động của DNS và các loại bản ghi DNS

## Mục lục

### [1. Cơ chết hoạt động của DNS](https://github.com/phancong0897/Congphan/blob/master/DNS/Ho%E1%BA%A1t%20%C4%91%E1%BB%99ng%20c%E1%BB%A7a%20DNS%20v%C3%A0%20c%C3%A1c%20b%E1%BA%A3n%20ghi.md#1-c%C6%A1-ch%E1%BA%BFt-ho%E1%BA%A1t-%C4%91%E1%BB%99ng-c%E1%BB%A7a-dns-1)

### [2. Các loại bản ghi DNS](https://github.com/phancong0897/Congphan/blob/master/DNS/Ho%E1%BA%A1t%20%C4%91%E1%BB%99ng%20c%E1%BB%A7a%20DNS%20v%C3%A0%20c%C3%A1c%20b%E1%BA%A3n%20ghi.md#2-c%C3%A1c-lo%E1%BA%A1i-b%E1%BA%A3n-ghi)

### [ Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/DNS/Ho%E1%BA%A1t%20%C4%91%E1%BB%99ng%20c%E1%BB%A7a%20DNS%20v%C3%A0%20c%C3%A1c%20b%E1%BA%A3n%20ghi.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3o)

## 1. Cơ chết hoạt động của DNS

<img src="https://imgur.com/iE8EAtR.png">

- Trước hết chương trình trên máy người sử dụng gửi yêu cầu tìm kiếm địa chỉ IP ứng với tên miền beta.example.com tới máy chủ quản lý tên miền cục bộ(Local DNS Server hay Local Name Server) thuộc mạng của nó.

- Máy chủ tên miền cục bộ này kiểm tra trong cơ sở dữ liệu của nó có chứa cơ sở dữ liệu chuyển đổi từ tên miền sang địa chỉ IP của tên miền mà người sử dụng yêu cầu không. Trong trường hợp máy chủ tên miền cục bộ có cơ sở dữ liệu này, nó sẽ gửi trả lại địa chỉ IP của máy có tên miền nói trên.

- Trong trường hợp máy chủ tên miền cục bộ không có cơ sở dữ liệu về tên miền này nó sẽ hỏi lên các máy chủ tên miền ở mức cao nhất ( máy chủ tên miền làm việc ở mức ROOT). Máy chủ tên miền ở mức ROOT này sẽ chỉ cho máy chủ tên miền cục bộ địa chỉ của máy chủ tên miền quản lý các tên miền có đuôi .COM.

- Máy chủ tên miền cục bộ gửi yêu cầu đến máy chủ quản lý tên miền có đuôi .com tìm tên miền beta.example.com. Máy chủ tên miền quản lý các tên miền .com sẽ gửi lại địa chỉ của máy chủ quản lý tên miền example.com.

- Máy chủ tên miền cục bộ sẽ hỏi máy chủ quản lý tên miền  example.com  này địa chỉ IP của tên miền beta.example.com. Do máy chủ quản lý tên miền example.com có cơ sở dữ liệu về tên miền beta.example.com nên địa chỉ IP của tên miền này sẽ được gửi trả lại cho máy chủ tên miền cục bộ.

- Máy chủ tên miền cục bộ chuyển thông tin tìm được đến máy của người sử dụng. 
Người sử dụng dùng địa chỉ IP này để kết nối đến server chứa trang web có địa chỉ beta.example.com.

- Các thành phần của DN

    - DNS Cache: Thuật ngữ này thường bị nhầm lẫn do nó có ít nhất 2 ý nghĩa. Đầu tiên DNS Cache có thể là danh sách tên và địa chỉ IP mà bạn đã truy vấn và đã được giải quyết và được lưu vào bộ nhớ cache để không có lưu lượng truy cập mạng được tạo ra và truy cập nhanh hơn. Ý nghĩa thứ hai liên quan đến một DNS Server chỉ đơn giản là thực hiện các truy vấn đệ quy và bộ đệm ẩn mà không thực sự là một máy chủ có thẩm quyền.

    - Resolvers :Là bất kỳ host nào trên Internet cần dùng để tra cứu thông tin tên miền, giống như thiết bị bạn đang sử dụng để đọc trang web này.

    - Name Servers: Những server này chứa cơ sở dữ liệu về tên và địa chỉ IP và phục vụ các yêu cầu DNS cho client.

    - Name Space: Là cơ sở dữ liệu về địa chỉ IP và các tên liên quan của chúng.

    
## 2. Các loại bản ghi

- Start of Authority (SOA) resource record: định nghĩa các tham số toàn cục cho zone hoặc tên miền. Một tệp tin zone chỉ được phép chứa một mẩu tin SOA và phải nằm ở vị trí đầu tiên trước các mẩu tin khác.

- Name server (NS) resource record: chỉ ra Máy chủ tên miền (Name server) của zone đó.

- A Resource Records (mẩu tin địa chỉ): mẩu tin cho biết địa chỉ IP tương ứng của một tên miền, có dạng như “example IN A 172.16.48.1”

- PTR Records (mẩu tin con trỏ): ngược lại với A record, PTR chỉ ra tên miền tương ứng của một địa chỉ IP, có dạng như “1.48.16.172.in-addr.arpa. IN PTR example.com.”

- CNAME Resource Records: một dạng record giúp tạo ra biệt hiệu cho một tên miền, ví dụ mẩu tin CNAME “ftp.example.com. IN CNAME ftp1.example.com.” cho phép trỏ tên miền ftp.example.com sang ftp1.example.com

- MX Resource Records (mẩu tin Mail exchange): chỉ ra máy chủ mail của tên miền.

- TXT Resource Records (mẩu tin text): chứa thông tin dạng văn bản không định dạng, thường dùng để chứa các thông tin bổ sung.-Nameserver and ZonesCác chương trình lưu trữu toàn bộ thông tin về domain namespace gọi là nameserver. Nameserver thông thường sẽ có thông tin hoàn chỉnh về một phần nào đó của domain namespace gọi là zone, zone này load từ file hoặc từ nameserver khác.

## Nguồn tham khảo 

https://news.cloud365.vn/hoat-dong-cua-he-thong-dns/