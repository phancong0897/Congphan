# Tổng quan về DNS

## Mục lục

### [1. DNS là gì?](https://github.com/phancong0897/Congphan/blob/master/DNS/DNS-overview.md#1-dns-l%C3%A0-g%C3%AC-)

### [2. Nguyên tắc làm việc của DNS](https://github.com/phancong0897/Congphan/blob/master/DNS/DNS-overview.md#2-nguy%C3%AAn-t%E1%BA%AFc-l%C3%A0m-vi%E1%BB%87c-c%E1%BB%A7a-dns)

### [3. Các khái niệm trong DNS](https://github.com/phancong0897/Congphan/blob/master/DNS/DNS-overview.md#3-c%C3%A1c-kh%C3%A1i-ni%E1%BB%87m-trong-dns)

### [Nguồn tham khảo](https://github.com/phancong0897/Congphan/blob/master/DNS/DNS-overview.md#ngu%E1%BB%93n-tham-kh%E1%BA%A3oSSSS)

## 1. DNS là gì ?

- DNS là từ viết tắt trong tiếng Anh của Domain Name System, là Hệ thống phân giải tên được phát minh vào năm 1984 cho Internet, chỉ một hệ thống cho phép thiết lập tương ứng giữa địa chỉ IP và tên miền. Hệ thống tên miền (DNS) là một hệ thống đặt tên theo thứ tự cho máy vi tính, dịch vụ, hoặc bất kỳ nguồn lực tham gia vào Internet. Nó liên kết nhiều thông tin đa dạng với tên miền được gán cho những người tham gia. Quan trọng nhất là, nó chuyển tên miền có ý nghĩa cho con người vào số định danh (nhị phân), liên kết với các trang thiết bị mạng cho các mục đích định vị và địa chỉ hóa các thiết bị khắp thế giới.

## 2. Nguyên tắc làm việc của DNS

- Mỗi nhà cung cấp dịch vụ vận hành và duy trì DNS server riêng của mình, gồm các máy bên trong phần riêng của mỗi nhà cung cấp dịch vụ đó trong Internet. Tức là, nếu một trình duyệt tìm kiếm địa chỉ của một website thì DNS server phân giải tên website này phải là DNS server của chính tổ chức quản lý website đó chứ không phải là của một tổ chức (nhà cung cấp dịch vụ) nào khác.

- INTERNIC (Internet Network Information Center) chịu trách nhiệm theo dõi các tên miền và các DNS server tương ứng. INTERNIC là một tổ chức được thành lập bởi NFS (National Science Foundation), AT&T và Network Solution, chịu trách nhiệm đăng ký các tên miền của Internet. INTERNIC chỉ có nhiệm vụ quản lý tất cả các DNS server trên Internet chứ không có nhiệm vụ phân giải tên cho từng địa chỉ.

- DNS có khả năng truy vấn các DNS server khác để có được 1 cái tên đã được phân giải. DNS server của mỗi tên miền thường có hai việc khác biệt. Thứ nhất, chịu trách nhiệm phân giải tên từ các máy bên trong miền về các địa chỉ Internet, cả bên trong lẫn bên ngoài miền nó quản lý. Thứ hai, chúng trả lời các DNS server bên ngoài đang cố gắng phân giải những cái tên bên trong miền nó quản lý.

- DNS server có khả năng ghi nhớ lại những tên vừa phân giải. Để dùng cho những yêu cầu phân giải lần sau. Số lượng những tên phân giải được lưu lại tùy thuộc vào quy mô của từng DNS.

## 3. Các khái niệm trong DNS 

- Không gian tên miền (Domain name space) : Không gian tên miền là một kiến trúc dạng cây (hình), có chứa nhiều nốt (node). Mỗi nốt trên cây sẽ có một nhãn và có không hoặc nhiều resource record (RR), chúng giữ thông tin liên quan tới tên miền. Nốt root không có nhãn.

- Tên miền (Domain name): Tên miền được tạo thành từ các nhãn và phân cách nhau bằng dấu chấm (.)Mỗi node trong cây có một nhãn riêng. Một nhãn rỗng đại diện cho root. Một domain name đầy đủ của bất kỳ một node nào trong “cây” là thứ tự các nhãn trong đường dẫn từ node đó đến root, tên mỗi node đc phân chia bằng các dấu chấm (.) Việc đánh tên này sẽ xác định được vị trí của node đó trong “cây”. Mỗi node trong cây đc xác định bằng một FQDN – tên đầy đủ chỉ đường dẫn đến node trong cây., ví dụ example.com. Tên miền còn được chia theo cấp độ như tên miền top level, tên miền cấp 1, cấp 2….  các loại top-level-domain .com: các tổ chức công ty thương mại.org: các tổ chức phi lợi nhuận.net: các trung tâm hỗ trợ về mạng.edu: các tổ chức giáo dục .gov: các tổ chức chính phủ.

- Primary DNS Server (PDS)

    - Primary DNS Server (PDS) là nguồn xác thực thông tin chính thức cho các tên miền mà nó được phép quản lý. Thông tin về một tên miền do PDS được phân cấp quản lý thì được lưu trữ tại đây và sau đó có thể được chuyển sang các Secondary DNS Server (SDS).

    - Các tên miền do PDS quản lý thì được tạo, và sửa đổi tại PDS và sau đó được cập nhật đến các SDS .

- Secondary DNS Server (SDS)

    - DNS được khuyến nghị nên sử dụng ít nhất là hai DNS server để lưu địa chỉ cho mỗi một vùng (zone). PDS quản lý các vùng và SDS được sử dụng để lưu trữ dự phòng cho vùng, và cho cả PDS. SDS không nhất thiết phải có nhưng khuyến khích hãy sử dụng . SDS được phép quản lý tên miền nhưng dữ liệu về tên miền không phải được tạo ra từ SDS mà được lấy về từ PDS.

    - SDS có thể cung cấp các hoạt động ở chế độ không tải trên mạng. Khi lượng truy vấn vùng (zone) tăng cao, PDS sẽ chuyển bớt tải sang SDS (quá trình này còn được gọi là cân bằng tải), hoặc khi PDS bị sự cố thì SDS hoạt động thay thế cho đến khi PDS hoạt động trở lại .

    - SDS thường được sử dụng tại nơi gần với các máy trạm (client) để có thể phục vụ cho các truy vấn một cách dễ dàng. Tuy nhiên, cài đặt SDS trên cùng một subnet hoặc cùng một kết nối với PDS là không nên. Điều đó sẽ là một giải pháp tốt để dự phòng cho PDS, vì khi kết nối đến PDS bị hỏng thì cũng không ảnh hưởng gì tới đến SDS.

    - Ngoài ra, PDS luôn duy trì một lượng lớn dữ liệu và thường xuyên thay đổi hoặc thêm các địa chỉ mới vào các vùng. Do đó, DNS server sử dụng một cơ chế cho phép chuyển các thông tin từ PDS sang SDS và lưu giữ trên đĩa. Khi cần phục hồi dữ liệuvề các vùng, chúng ta có thể sử dụng giải pháp lấy toàn bộ ( full ) hoặc chỉ lấy phần thay đổi (incrememtal).

- Các khái niệm trong Zone : 

    - primary zone: cho phép đọc và ghi cơ sở dữ liệu và có toàn quyền trong việc update dữ liệu của DNS

    - secondary zone: cho phép đọc và ghi bản sao của sơ sở dữ liệu và muốn được cập nhật zone thì phải đồng bộ với Primary zone

    - forwarder: là kỹ thuật cho phép name server nội bộ gửi yêu cầu truy vấn đến server khác để phân giải những tên miền bên ngoài hệ thống

    - Delegation (sự ủy quyền) : 1 miền có thể tổ chức thành miền con, mỗi miền con có thể ủy quyền cho 1 tổ chức khác, và tổ chức này phải chịu trách nhiệm duy trì thông tin trong miền này

Một trong những mục đích của DNS đó là quản trị phân tán. Ta có thể chia nhỏ việc quản lý thành nhiều phần khác nhau. Một domain có thể có nhiều subdomains.

## Nguồn tham khảo

https://itforvn.com/tu-hoc-mcsa-mcse-2016-lab-3-cau-hinh-dns-server-tren-windows-server-2016/
