# Tổng quan về SSL

## Mục lục

### 1. SSL là gì?

- SSL là viết tắt của từ Secure Sockets Layer. SSL là tiêu chuẩn của công nghệ bảo mật, truyền thông mã hoá giữa máy chủ Web server và trình duyệt. Tiêu chuẩn này hoạt động và đảm bảo rằng các dữ liệu truyền tải giữa máy chủ và trình duyệt của người dùng đều riêng tư và toàn vẹn. SSL hiện tại cũng là tiêu chuẩn bảo mật cho hàng triệu website trên toàn thế giới, nó bảo vệ dữ liệu truyền đi trên môi trường internet được an toàn.

- SSL đảm bảo rằng tất cả các dữ liệu được truyền giữa các máy chủ web và các trình duyệt được mang tính riêng tư, tách rời.

- Mục tiêu của SSL là bảo mật các thông tin nhạy cảm trong quá trình truyền trên internet như: thông tin cá nhân, thông tin thanh toán, thông tin đăng nhập.

- Nó là giải pháp thay thế cho phướng pháp truyền thông tin văn bản dạng plain text, văn bản loại này khi truyền trên internet sẽ không được mã hóa, nên việc áp dụng mã hóa vào sẽ khiến cho các bên thứ 3 không xâm nhập được vào thông tin của bạn, không đánh cắp hay chỉnh sửa được các thông tin đó.

- Hầu hết mọi người đều quen thuộc với các chứng chỉ SSL/TLS, đang được dùng bởi các website lớn và các webmaster nghiêm túc trong việc bảo vệ các giao dịch người dùng.

- Bạn có thể biết được website có đang dùng chứng chỉ bảo mật SSL/TLS hay không đơn giản bằng cách nhìn vào icon trong URL ngay trong thanh địa chỉ.

### 2. Chứng chỉ SSL hoạt động như thế nào?

- Chứng chỉ SSL/TLS hoạt động bằng cách tích hợp key mã hóa vào thông tin định danh công ty. Nó sẽ giúp công ty mã hóa mọi thông tin được truyền mà không bị ảnh hưởng hoặc chỉnh sửa bởi các bên thứ 3.

<img src="https://imgur.com/EzrI4HV.png">

- SSL/TLS hoạt động bằng cách sử dụng public và private key, đồng thời các khóa duy nhất của mỗi phiên giao dịch. Mỗi khi khách truy cập điền vào thanh địa chỉ SSL thông tin web browser hoặc chuyển hướng tới trang web được bảo mật, trình duyệt và web server đã thiết lập kết nối.

- Trong phiên kết nối ban đầu, public và private key được dùng để tạo session key, vốn được dùng để mã hóa và giải mã dữ liệu được truyền đưa. Session key sẽ được sử dụng trong một khoảng thời gian nhất định và chỉ có thể dùng cho phiên giao dịch này.

- Nếu có khóa màu xanh ngay đầu địa chỉ web thì tức là website đã thiết lập đúng SSL/TLS. Bạn có thể nhấn vào nút màu xanh đó để xem ai là người giữ chứng chỉ này.

### 3. Các loại chứng chỉ ssl

### Domain Validation

- Đối với loại chứng chỉ SSL này, cách xác nhực tên miền được thực hiện bằng email hoặc qua hồ sơ DNS.

- Nói một cách đơn giản, bạn cần phải xác nhận quyền sở hữu tên miền đó là của mình.

- Loại chứng chỉ này lấy được trong một vài phút (hoặc đôi khi một vài giờ).

- Điều này khá lý tưởng cho những cá nhân không thuộc tổ chức và không quan tâm đến vấn đề bảo mật.

- Đây là loại chứng chỉ SSL rẻ nhất và được khuyến nghị khi vấn đề an ninh không phải là yếu tố quan trọng nhất (ví dụ blog).

- Đánh giá:

    - Rẻ
    
    - Lấy được trong thời gian lâu nhất là vài giờ
    
    - Tốt nhất cho các blog cá nhân và các trang web không đặt nặng vấn đề bảo mật
### Organization Validation

- Đây là chứng chỉ tối thiểu được đề nghị cho các cổng thông tin thương mại điện tử. hay các eCommerce

- Certificate Authority (CA) xác nhận quyền sở hữu tên miền & các thông tin khác thông qua việc sử dụng các cơ sở dữ liệu công cộng.

- Sự khác biệt chủ yếu giữa DV & OV là việc xác thực công ty được thực hiện bởi các nhà cung cấp chứng chỉ. Nó không lớn như EV (chi tiết bên dưới), nhưng có khả năng tốt hơn DV.

- Đánh giá:

    - Yêu cầu 2-3 ngày làm việc để kích hoạt

    - Không thực sự hơn các chứng chỉ DV khác
    
    - Tối thiểu cho các cổng thông tin thương mại điện tử
### Extended Validation

- Đây là loại chứng chỉ được đánh giá cao cho các trang web với hoạt động giao dịch

- Rất dễ dàng để lấy được chứng chỉ DV & OC trong khi chứng chỉ EV đòi hỏi một quy trình xác thực nghiêm ngặt.

- Đây là loại chứng chỉ hiển thị các tổ chức mà chứng chỉ được cấp cho trong trình duyệt.

- Hầu hết các trang ngân hàng, tài chính và thương mại điện tử đều sử dụng chứng chỉ EV vì nó cung cấp thanh địa chỉ HTTPS màu xanh lá cây mà các bạn có thể thấy ở những trang uy tín.

- Đánh giá:

    - Mất khoảng 7-10 ngày để kích hoạt

    - Được đề xuất cho các tổ chức, có thanh địa chỉ HTTPS màu xanh lá cây phổ biến
    
### Các loại chứng chỉ SSL dựa trên số lượng tên miền và sub-domain bạn muốn
    
- Đây sẽ là một quyết định quan trọng cho những ai đang có kế hoạch mua một chứng chỉ SSL.

- Ngoài ba chứng chỉ trên, bạn cũng cần phải chọn các loại chứng chỉ dựa trên các lĩnh vực và các sub-domain mà mình có.

- Hãy cùng tìm hiểu vào các loại chứng chỉ SSL khác nhau với các số lượng tên miền hỗ trợ.

- Single Name SSL Certificate

    - Chỉ có một tên miền được bảo đảm bằng loại SSL này. Xem hướng dẫn cách dùng WordPress Thêm SSL Và HTTPS Trong WordPress nếu bạn chưa biết nhé !

    - Ví dụ, nếu bạn mua một chứng chỉ cho domain.com, nó sẽ không đảm bảo docs.domain.com hoặc shop.domain.com.

    - Lưu ý: Một Chứng Chỉ SSL Single Name cũng có thể được dùng để bảo đảm chỉ một sub-domain.

    - Ví dụ: Bạn có thể đảm bảo shop.domain.com và không domain.com.

- Chứng chỉ Wildcard SSL

    - Điều này đảm bảo sự không giới hạn các sub-domain và một tên miền duy nhất.

    - Ví dụ: Mua một chứng chỉ Wildcard SSL cho dieuhau.com cũng sẽ cho phép tôi sử dụng với các subdomain:

    - vip.dieuhau.com
    
    - forum.dieuhau.com

    - Tuy nhiên, nó sẽ không tác dụng với dạng abc.pro.dieuhau.com.

- Chứng Chỉ SSL Multi-domain

    - Một chứng Chỉ SSL Multi-domain hỗ trợ tất cả các loại tên miền và subdomain khác nhau.

    - Điều này được đề xuất cho những người có nhiều tên miền và subdomain.

- Chứng Chỉ Unified Communications (UCC)
    
    - Khi nghiên cứu để viết bài tôi đã biết thêm về loại chứng chỉ này.

    - UCCs cho phép khách hàng bảo vệ lên đến 100 tên miền bằng cách sử dụng cùng một chứng chỉ.

    - Các Chứng Chỉ Unified Communications được thiết kế đặc biệt để bảo đảm Microsoft® Exchange và các Office Communication Server.

    - Hãy nhớ rằng: Sử dụng một chứng chỉ cho nhiều tên miền không ảnh hưởng đến bạn dù bằng cách nào.

    - Các trang web phổ biến để mua chứng chỉ SSL:

        - Geotrust: Certificate Authority
        
        - Digicert
        
        - Let’s Encrypt: Certificate Authority (các chứng chỉ xác thực tên miền miễn phí)
        
        - CheapSSL Security
        
        - Trustwave
        
        - SSLs
### 5.Một số loại chứng chỉ SSL miễn phí.

### Let's Encrypt

- Let's Encrypt là nhà cung cấp chứng chỉ SSL miễn phí, tự động, hoạt động vì lợi ích của cộng đồng. Nó được quản lý bởi Internet Security Research Group (ISRG).

- Let's Encrypt cung cấp cho những người quản trị website một chứng nhận số cần thiết để kích hoạt HTTPS (SSL hoặc TLS) cho website của mình, hoàn toàn miễn phí, và theo cách thân thiện nhất có thể. Tất cả dựa trên mục tiêu tạo ra một môi trường Web an toàn, riêng tư và tôn trọng người dùng hơn.

- Let's Encrypt cung cấp chứng chỉ SSL loại Domain Validation, tức là sau khi cài đặt, trang web của bạn sẽ có một ổ khóa màu xanh trên thanh địa chỉ của trình duyệt, khi người dùng truy cập vào.

- Các nguyên tắc chính phía sau Let’s Encrypt là:

    - Miễn phí: Bất cứ ai sở hữu một tên miền có thể dùng Let’s Encrypt để có được một chứng chỉ tin cậy với chi phí bằng không.

    - Tự động: Phần mềm chạy trên máy chủ web có thể tương tác với Let’s Encrypt để lấy chứng chỉ một cách dễ dàng, cấu hình an toàn để sử dụng và tự động gia hạn.

    - An toàn: Let’s Encrypt sẽ phục vụ như một nền tảng để thúc đẩy các thực hành tốt nhất về bảo mật TLS, cả về phía CA và bằng cách giúp những người vận hành trang web bảo mật đúng cách các máy chủ của họ.

    - Trong suốt: Tất cả các chứng chỉ được cấp hoặc thu hồi sẽ được ghi lại công khai và có sẵn cho bất kỳ ai muốn kiểm tra.
    
    - Mở: Giao thức phát hành và gia hạn tự động sẽ được công bố như một tiêu chuẩn mở mà những người khác có thể áp dụng.
    
    - Hợp tác: Giống như các giao thức Internet cơ bản, Let’s Encrypt là một nỗ lực chung nhằm mang lại lợi ích cho cộng đồng, vượt ra ngoài tầm kiểm soát của bất kỳ một tổ chức nào.

### OpenSSL

- OpenSSL là một thư viện phần mềm cho các ứng dụng bảo mật truyền thông qua mạng máy tính chống nghe trộm hoặc cần phải xác định phe truyền thông ở bên đầu kia. Nó được sử dụng rộng rãi trong các máy chủ web internet, phục vụ phần lớn tất cả các trang web.

- 
