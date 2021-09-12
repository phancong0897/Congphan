# Thông tin docker, images và giám sát các container trong docker

### Các thành phần tạo nên Container

- Một số khái niệm mô tả hoạt động của Docker

    - Container: đó là một hộp kín để ứng dụng hoạt động. Mỗi container đều được tạo dựa trên một image (image đó chứa đủ cấu hình, thành phần dữ liệu). Khi bạn chạy một container xuất phát sừ một image, có một lớp (layer) được phép ghi thêm vào trên đỉnh của image (vậy container = image + layer được phép ghi). Sau đó, khi bạn lưu container này thành image mới (lệnh docker commit), một lớp image mới được thêm vào vào.

    - Image: như là ảnh chụp lại các cấu hình của container. Một image luôn ở trạng thái chỉ đọc, mọi thứ muốn thay đổi đều phải lưu ở tầng trên cùng (được phép ghi) của container, và tầng này có thể lưu lại để tạo image mới (thêm layer vào image cũ) và nó lại thành trạng thái chỉ đọc. Với quy trình như vậy, mỗi image đều phụ thuộc vào một hoặc nhiều image cha.

    - Platform Image : là một image mà nó không có image cha. Những image loại này chứa các biến môi trường, các gói, tiện ích để ưng dụng ảo hóa chạy, nó cũng chỉ đọc.

    - Registry là kho chứa các image, nơi chi sẻ, tải về các image.

    - Dockerfile một file cấu hình với cấu trúc để sinh ra images. Sử dụng file Dockerfile là cách tự động hóa việc tạo, chạy, sử dụng các container.

### Truy vấn thông tin về Image và Container

#### Lệnh docker history - lịch sử

- Lệnh này để truy vấn thông tin lịch sử các thao tác để hình thành nên một image. Cú pháp như sau

    ` docker image history name_or_id_of_image `

    ![Imgur](https://imgur.com/QuAyK30.png)

#### Lệnh docker inspect
    
- Để có được thông tin chi tiết về một image, container nào đó sử dụng đến lệnh docker inspect, lệnh này trả về thông tin về đối tượng cần truy vấn dưới dạng JSON. Cú pháp như sau:

    ` docker inspect name_or_id_of_image_container `

    ![Imgur](https://imgur.com/kUJTLxa.png)

#### Lệnh docker diff

- Lệnh docker diff để kiểm tra xem một container từ lúc nó được tạo ra đến giờ hệ thống file/thư mục thay đổi như thế nào. Cú pháp lệnh

    ` docker diff container-name-or-id `

    Kết quả liệt kê ra danh sách trên từng dòng những thư mục, file có sự thay đổi. Tiền tố đầu dòng có thể là A (thêm vào add), D (bị xóa đi delete) hoặc C (được được cập nhật - change).

    ![Imgur](https://imgur.com/HvMA6mJ.png)

#### Lệnh docker logs

- Để đọc thông tin log của container (mà không phải vào terminal của container) thì dùng tới lệnh docker logs, cú pháp cơ bản như sau:

    ` docker logs container-name-or-id `

    ![Imgur](https://imgur.com/ZKNC6wV.png)

    - Bạn có thể đưa vào một số tùy chọn sau của lệnh:

        - --tail n chỉ hiện thị n dòng cuối

        - -f hoặc --follow với tham số này, nếu container đang chạy nó sẽ tự động hiện thị thêm log mới nếu container phát sinh log. Ngắt giám sát log nhấn CTRL+C

    ![Imgur](https://imgur.com/j9X6bPN.png)


#### Giám sát container bằng docker stats

- Lệnh docker stats giám sát theo thời gian thực một số lại lượng sử dụng bởi container gồm: CPU, bộ nhớ, và lưu lượng mạng, số tiến trình. Cú pháp lệnh như sau:

    ` docker stats container1, container2 ... `

    ![Imgur](https://imgur.com/fkI8Pow.png)



