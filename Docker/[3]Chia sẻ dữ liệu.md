# Chia sẻ dữ liệu giữa Docker Host và Container

Máy Host là hệ thống bạn đang chạy Docker Engine. Một thư mục của máy Host có thể chia sẻ để các Container đọc, lưu dữ liệu.

### Chia sẻ dữ liệu Host/Container

- Để chia sẻ dữ liệu giữa máy host và container ta sử dụng lệnh sau:

    ` docker run -it -v path_host:path_container --name imageID `

    ![Imgur](https://imgur.com/069Nfbl.png)

Khi container được tao ra, thư mục /home/dulieu sẽ ánh xá đến thư mục /home/congpv ở máy host.

### Chia sẻ dữ liệu giữa các Container

- Giờ chạy, tạo container khác cũng nhận thư mục chia sẻ dữ liệu như container đầu tiên( ở đây ví dụ container tên congpv). Để làm điều đó thêm vào tham số --volumes-from congpv

    ` docker run -it --volumes-from container_vừatạo iamgesID `

    ![Imgur](https://imgur.com/HZ9wKjG.png)

### Quản lý các ổ đĩa với docker volume

- Liệt kê danh sách các ổ đĩa:

    ` docker volume ls `

- Tạo một ổ đĩa:

    ` docker volume create name_volume `

- Xem thông tin chi tiết về đĩa:

    ` docker volume inspect name_volume `

- Xóa một ổ đĩa:

    ` docker volume rm name_volume `

    ![Imgur](https://imgur.com/Pc41rVp.png)

- Mount một ổ đĩa vào container (--mount)

    ``` 

    # Tạo ổ đĩa có tên disk1
    docker volume create disk1

    # Mount ổ đĩa vào container
    
    # container truy cập tại /home/disk1

    docker run -it --mount source=firstdisk,target=/home/firstdisk  ubuntu

    ```

    `docker run -it --mount source=volume,target=path_container --name "name" imagesID `

    ![Imgur](https://imgur.com/cJxtSEY.png)

- Gán ổ đĩa vào container khi tạo container (-v)

    - Nếu muốn ổ đĩa bind dữ liệu đến một thư mục cụ thể của máy HOST thì tạo ổ đĩa với tham số như sau:

    ` docker volume create --opt device=path_in_host --opt type=none --opt o=bind  volumename `

    - Sau đó ổ đĩa này gán vào container với tham số -v (không dùng --mount)

    ![Imgur](https://imgur.com/ksBMAPX.png)
