# Thành phần cấu tạo nên File cấu hình HAProxy

Cấu hình của HAProxy thường được tạo từ 4 thành phần bao gồm global, defaults, frontend, backend. 4 thành phần sẽ định nghĩa cách HAProxy nhận, xử lý các request, điều phối các request tới các Backend phía sau.

### Cấu trúc

Đầu tiên, sẽ có sự khác nhau giữa vị trí file cấu hình giữa 2 phiên bản EE - Enterprise Edition và CE - Community Edition. Trong bài viết tôi sẽ tập trung vào phiên bản CE tức phiên bản mở do công đồng đóng góp (Community Edition).

Đường dẫn file cấu hình HAProxy được lưu tại /etc/haproxy/haproxy.cfg với cấu trúc:

    ```
    global
    # Các thiết lập tổng quan

    defaults
    # Các thiết lập mặc định

    frontend
    # Thiết lập điều phối các request

    backend
    # Định nghĩa các server xử lý request

    ```

Hãy tưởng tưởng, khi các bạn khởi tạo dịch vụ haproxy, các thiết lập tại global sẽ được sử dụng để định nghĩa cách HAProxy được khởi tạo như số lượng kết nối tối đa, đường dẫn ghi file log, số process v.v. Sau đó các thiết lập tại mục defaults sẽ được áp dụng cho tất cả mục frontend, backend nằm phía sau (các bạn hoàn toàn có thể định nghĩa lại các giá tri mặc định tại frontend và backend). Có thể có nhiều mục frontend, backend được định nghĩa trong file cấu hình. Mục frontend được định nghĩa để điều hướng các request nhận được tới các backend. Mục backend sử dụng để định nghĩa các danh sách máy chủ dịch vụ (có Web server, Database, …) đây là nơi request được xử lý.

#### Global

- Được đặt trên cùng file cấu hình HAProxy, được định danh bằng từ khóa global và mục này luôn đứng riêng 1 dòng và được định nghĩa 1 lần duy nhất trong file cấu hình. Các thiết lập bên dưới global định nghĩa các thiết lập bảo mật, các điều chỉnh về hiệu năng áp dụng trên toàn HAProxy (áp dụng tại mức tiến trình HAProxy hoạt động)

- Ví dụ:

    ```

    global
    log         127.0.0.1 local2
    chroot      /var/lib/haproxy
    pidfile     /var/run/haproxy.pid
    maxconn     4000
    user        haproxy
    group       haproxy
    daemon
    stats socket /var/lib/haproxy/stats

    ```

- Ý nghĩa các cấu hình

    - maxconn: Chỉ định giới hạn số kết nối mà HAProxy có thể thiết lập. Sử dụng với mục đích bảo vệ load balancer khởi vấn đề tràn ram sử dụng.
    
    - log: Bảo đảm các cảnh báo phát sinh tại HAProxy trong quá trình khởi động, vận hành sẽ được gửi tới syslog
    
    - stats socket Định nghĩa runtime api, có thể sử dụng để disable server hoặc health checks, thấy đổi load balancing weights của server. Đọc thêm
    
    - user / group chỉ định quyền sử dụng để khởi tạo tiến trình HAProxy. Linux yêu cầu xử lý bằng quyền root cho nhưng port nhở hơn 1024. Nếu không định nghĩa user và group, HAProxy sẽ tự động sử dụng quyền root khi thực thi tiến trình.

#### Defaults

- Khi cấu hình tăng dần, phức tạp, khó đọc, các thiết lập cấu hình tại mục defaults giúp giảm các trùng lặp. Thiết lập tại mục defaults sẽ áp dụng cho tất cả mục frontend backend nằm sau nó. Các bạn hoàn toàn có thể thiết lập lại trong từng mục backend, frontend.

- Có thể có nhiều mục defaults. Chúng sẽ ghi đè lên nhau dựa theo vị trí (tức các mục defaults nằm sau sẽ ghi đè lên các mục defaults nằm trước).

- Ví dụ đơn giản, các bạn có thế thiết lập mode http tại mục defaults, khi đó toàn bộ các mục frontend backend listen sẽ đều dùng mode http làm mặc định.

- Ví dụ

    ```
    mode                    http
    log                     global
    option                  httplog
    option                  dontlognull
    option                  http-server-close
    option forwardfor       except 127.0.0.0/8
    option                  redispatch
    retries                 3
    option  http-server-close
    timeout http-request    10s
    timeout queue           1m
    timeout connect         10s
    timeout client          1m
    timeout server          1m
    timeout http-keep-alive 10s
    timeout check           10s
    maxconn                 3000

    ```

timeout connect chỉ định thời gian HAProxy đợi thiết lập kết nối TCP tới backend server. Hậu tố s tại 10s thể hiện khoảng thời gian 10 giây, nếu bạn không có hậu tố s, khoảng thời gian sẽ tính bằng milisecond. xem thêm.

timeout server chỉ định thời gian chờ kết nối tới backend server.

- Lưu ý:

    - Khi thiết lập mode tcp thời gian timeout server phải bằng timeout client

    - log global: Chỉ định ‘frontend’ sẽ sử dụng log settings mặc định (trong mục global).
    
    - mode: Thiết lập mode định nghĩa HAProxy sẽ sử dụng TCP proxy hay HTTP proxy. Cấu hình sẽ áp dụng với toàn frontend và backend khi bạn chỉ mong muốn sử dụng 1 mode mặc định trên toàn backend (Có thể thiết lập lại giá trị tại backend)

    - maxconn: Thiết lập chỉ định số kết nối tối đa, mặc định bằng 2000.
    
    - option httplog: Bổ sung format log dành riêng cho các request http bao gồm (connection timers, session status, connections numbers, header v.v). Nếu sử dụng cấu hình mặc định các tham số sẽ chỉ bao gồm địa chỉ nguồn và địa chỉ đích.
    
    - option http-server-close: Khi sử dụng kết nối dạng keep-alive, tùy chọn cho phép sử dụng lại các đường ống kết nối tới máy chủ (có thể kết nối đã đóng) nhưng đường ống kết nối vẫn còn tồn tại, thiết lập sẽ giảm độ trễ khi mở lại kết nối từ phía client tới server.
    
    - option dontlognull: Bỏ qua các log format không chứa dữ liệu
    
    - option forwardfor: Sử dụng khi mong muốn backend server nhận được IP thực của người dùng kết nối tới. Mặc định backend server sẽ chỉ nhận được IP của HAProxy khi nhận được request. Header của request sẽ bổ sung thêm trường X-Forwarded-For khi sử dụng tùy chọn.
    
    - option redispatch: Trong mode HTTP, khi sử dụng kỹ thuật stick session, client sẽ luôn kết nối tới 1 backend server duy nhất, tuy nhiên khi backend server xảy ra sự cố, có thể client không thể kết nối tới backend server khác (Trong bài toán load balancer). Sử dụng kỹ thuật cho phép HAProxy phá vỡ kết nối giữa client với backend server đã xảy ra sự cố. Đồng thời, client có thể khôi phục lại kết nối tới backend server ban đầu khi dịch vụ tại backend server đó trở lại hoạt động bình thường.
    
    - retries: Số lần thử kết nối lại backend server trước khi HAProxy đánh giá backend server xảy ra sự cố.
    
    - timeout check: Kiểm tra thời gian đóng kết nối (chỉ khi kết nối đã được thiết lập)
    
    - timeout http-request: Thời gian chờ trước khi đóng kết nối HTTP
    
    - timeout queue: Khi số lượng kết nối giữa client và haproxy đạt tối đã (maxconn), các kết nối tiếp sẽ đưa vào hàng đợi. Tùy chọn sẽ làm sạch hàng chờ kết nối.