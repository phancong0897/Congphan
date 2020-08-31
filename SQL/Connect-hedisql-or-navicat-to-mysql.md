# Tài liệu hướng dẫn kết nối HediSQL và Navicat đến MySQL

## Giới thiệu

- HeidiSQL là một công cụ quản trị nguồn mở và miễn phí cho MySQL, MariaDB và các nhánh của nó. Ngoài ra còn hỗ trợ Microsoft SQL Server và PostgreQuery. Sau đây là phần giới thiệu và hướng dẫn cài đặt quản lý cơ sở dữ liệu với phần mềm HeidiSQL.
- Một số tính năng nổi bật HeidiSQL:

    - HeidiSQL hỗ trợ nền tảng Windows XP, Vista, 7, 8, 10, Wine và Linux.

    - Kết nối với máy chủ MariaDB, MySQL, MS SQL và PostgreSQL.

    - Thiết lập kết nối SSH tới MariaDB hoặc MySQL.

    - Sử dụng Command Line hoặc GUI.

    - Các thao tác với bảng, khung nhìn, thủ tục, kích hoạt và sự kiện.
    
    - Khởi chạy các câu truy vấn SQL.
    
    - Nhập và xuất các tập tin SQL.
    
    - Phiên bản HeidiSQL Portable.

- Navicat Premium là một công cụ quản trị cơ sở dữ liệu đa kết nối nâng cao cho phép bạn kết nối đồng thời với tất cả các loại cơ sở dữ liệu một cách dễ dàng. Navicat cho phép bạn kết nối với cơ sở dữ liệu MySQL, MariaDB, Oracle, PostgreSQL, SQLite và SQL Server từ một ứng dụng duy nhất, giúp quản trị cơ sở dữ liệu trở nên dễ dàng.

- Các Tính Năng:

    - Tạo kết nối dễ dàng đến gần như tất cả các loại Database.

    - Tạo cơ sở dữ liệu.

    - Mô hình dữ liệu hiện ngay trước mặt, thêm khóa, xóa khóa sử dụng kéo chuột.

    -  Tư động format biểu đồ ER Diagram.

    - Viết câu lệnh SQL.

    - Thiết kế các Model.

## Kết nối HediSQL tới MySQL

- Dowload phần mềm HediSQL tạ trang chủ https://www.heidisql.com/download.php.

- Sau khi tải phần mềm HeidiSQL về máy bạn cần khởi chạy tập tin EXE

    Nhấn vào tùy chọn I accept the agreeement và nhấn Next để qua màn hình kế tiếp.

<img src="https://imgur.com/ThuWlyy.png">

- Tiếp theo bạn cần chọn đường dẫn thư mục lưu trữ trong Browse… hoặc có thể để mặc định.

<img src="https://imgur.com/LzSwMEZ.png">

- Bạn có thể chọn một số tùy chọn sau đó nhấn Next hoặc có thể để như mặc định.

<img src="https://imgur.com/FmiETB9.png">

- Bây giờ chỉ cần nhấn Install để bắt đầu cài đặt phần mềm HeidiSQL.

<img src="https://imgur.com/m81ADKq.png">

- Sau khi cài đặt xong thì bạn có thể sử dụng để kết nối với máy chủ cơ sở dữ liệu như MySQL, MariaDB, Microsoft SQL Server hoặc PostgreQuery.

<img src="https://imgur.com/Lea7enx.png">

## Kết nối Navicat tới MySQL

- Dowload phần mềm Navicat trên trang chủ: https://www.navicat.com/en/products

- Sau khi dowload xong, bạn càn chạy tệp exe và cài đặt như bình thường

<img src="https://imgur.com/iy8Mpik.png">

<img src="https://imgur.com/9lUmmms.png">

<img src="https://imgur.com/nLjBCnh.png">

<img src="https://imgur.com/tVkHqhR.png">

<img src="https://imgur.com/0gCNPDH.png">

- Sau khi cài đặt xong, bạn mở phần mềm Navicat Prenium để két nối tới MySQL

- Bạn hãy điền chính xác các thông tin host, port, username, password để có thể connect tới MySQL. Lưu ý nhớ phân quyền cho tài khoản databases.

<img src="https://imgur.com/qkcfzWd.png">

- Lưu ý: Trước khi bạn kết nối đến MySQL thì bạn hãy phần quyền cho nó bằng câu lệnh :

    - ` GRANT ALL PRIVILEGES ON wordpress.* TO user@IP IDENTIFIED BY '1234'; `

