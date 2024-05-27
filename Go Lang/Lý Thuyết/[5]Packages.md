# Packages

### Packages là gì? Tại sao chúng ta phải dùng packages?

- Trong các bài hướng dẫn trước, chúng ta thấy rằng một chương trình chạy bằng một file duy nhất có một hàm chính (main function) và một vài hàm khác. Trong các dự án thực tế, chúng ta không viết tất cả mã nguồn vào một file đơn lẻ như vậy. Chúng ta không thể tái sử dụng code hay bảo trì code theo cách viết này được. Để giải quyết vấn đề đó chúng ta phải cần đến packages (gói).

- Packages được sử dụng để tổ chức mã nguồn sao cho việc đọc và tái sử dụng mã nguồn dễ dàng hơn. Packages giúp phân chia mã nguồn thành nhiều phần, do đó việc bảo trì ứng dụng cũng thuận tiện hơn.

- Ví dụ, chúng ta cần tạo một ứng dụng xử lý hình ảnh với các tính năng như image cropping (cắt ảnh), sharpening (làm sắc nét), blurring (làm mờ) và color enhancement (tinh chỉnh màu sắc). Cách để xây dựng ứng dụng này là nhóm tất cả các code liên quan đến một tính năng vào trong một package riêng của nó. Ví dụ tính năng image cropping ở trong một package riêng, tính năng sharpening thì ở một package khác. Ưu điểm của việc này là, tính năng color enhancement có thể cần sử dụng một số hàm xử lý của tính năng sharpening. Khi đó ở mã nguồn của tính năng color enhancement ta có thể import (thêm vào) package sharpening và sử dụng các hàm của nó (chúng ta sẽ bàn về import sau). Bằng cách này code của chúng ta sẽ dễ dàng tái sử dụng.

### Tạo module

- Để tạo một module chạy lệnh. Và go.mod sẽ được tạo ra.

` go mod init <module name>  `

- Trong mỗi file go.mod thì có ba phần chính. Phần đâu khai báo tên module, thông thường tên module giống với một url chỉ tới git responstory của module đó và phiên bản go đang sử dụng. Phần thứ 2 là phần tuỳ chỉnh URL của package, đa số là rút gọn , mỗi khi bạn import một package thay vì phải viết dài như “import “github.com/dongnguyenltqb/demo/untils” giờ nếu bạn thay thế phần “github.com/dongnguyenltqb/demo” bằng “demo/utils” thì sẽ đỡ tốn công . Thực hiện bằng cách thêm dòng chữ sau vào, phần thứ 2 trên “require”. Và khi build go sẽ get các package đó bằng url gốc.

- Phần thứ 3 đây là phân quan trọng nhất, chỉ định rõ những Package mà bạn dùng trong module, bao gồm phiên bản, chú thích // indirect chỉ ra rằng module của bạn không dùng trực tiếp Package này mà một trong số những package bạn đang sử dụng dùng package đó, chú thích +incompatible chỉ ra việc Go đã bỏ qua những quy định về quản lí phiên bản theo tiêu chuẩn Semver

Thông thường khi cài đặt một package bằng cách sử dụng “go get” thì mặc định go sẽ tải xuống phiên bản mới nhất của package, để tránh điều này bạn có thể chỉ định rõ version, thậm chí là cả commit và branch của package.