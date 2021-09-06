# Mạng network bridge trong Docker kết nối các container với nhau

### Network trong Docker, liên kết mạng các container

- Docker cho phép bạn tạo ra một network (giao tiếp mạng), sau đó các container kết nối vào network. Khi các container cùng một network thì chúng có thể liên lạc với nhau nhanh chóng qua tên của container và cổng (port) được lắng nghe của container trên mạng đó.

- Để liệt kê các network đang có:

    ` docker network ls `

- Các network được tạo ra theo một driver nào đó như bridge, none, overlay, macvlan. Trong phần này sẽ sử dụng đến bridge network: nó cho phép các container cùng network này liên lạc với nhau, cho phép gửi gói tin ra ngoài. Tạo một bridge network với tên network là name-network.

    ` docker network create --driver bridge name-network `

- Khi tạo một container, ấn định nó nối vào network có tên là name-network thì thêm vào tham số --network name-netowrk trong lệnh docker run. Khi một container có tên name-container đã được tạo, để thiết lập nó nối vào network có tên name-network.

    ` docker network connect name-network name-container `

    ![Imgur](https://imgur.com/TiLZ9Bg.png)

- Để tạo container gắn với network nào sử dụng câu lệnh"

    ` docker run -it --network name-network --name "container_name" imagesID `

    ![Imgur](https://imgur.com/386lzif.png)

### Cổng - Port

- Các kết nối mạng thông qua các cộng, để thiết lập cấu hình các cổng của container chú ý: có cổng bên trong container, có cổng mở ra bên ngoài (public), có giao thức của cổng (tpc, udp). Khi chạy container (tạo) cần thiết lập cổng thì đưa vào lệnh docker run tham số dạng sau:


    ` docker run -p public-port:target-port/protocol ...  `

    - public-port : cổng public ra ngoài (ví dụ 80, 8080 ...), các kết nối không cùng network đến container 
    phải thông qua cổng này.
    
    - target-port : cổng bên trong container, cổng public-port sẽ ánh xạ vào cổng này. Nếu các container cùng network có thể kết nối với nhau thông qua cổng này.

    Ví dụ: ` docker run -it --network br0 --name B1 -p 8888:80 busybox `

    - Ta sử dụng images busybox để kiểm tra public port. 

    ![Imgur](https://imgur.com/O6MPOEZ.png)

    Sau đó, ta sử dụng trình duyêt ở local, truy cập địa chỉ IP chạy docker engine thêm port của container public để kiểm tra nhé: IP:8080 , kết quả đã ra hiển thị file index.html

    ![Imgur](https://imgur.com/fUcpct3.png)

