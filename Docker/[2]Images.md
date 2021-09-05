# Docker image

### Docker image là gì?

- Docker image là một file bất biến - không thay đổi, chứa các source code, libraries, dependencies, tools và các files khác cần thiết cho một ứng dụng để chạy.

- Do tính chất read-only của chúng, những images này đôi khi được gọi là snapshots. Chúng đại diện cho một application và virtual environment của nó tại một thời điểm cụ thể. Tính nhất quán này là một trong những tính năng tuyệt vời của Docker. Nó cho phép các developers test và thử nghiệm phần mềm trong điều kiện ổn định, thống nhất.

- Bạn có thể tạo số lượng Docker image không giới hạn từ một image base. Mỗi khi bạn thay đổi trạng thái ban đầu của một image và lưu trạng thái hiện có, bạn tạo một image mới với một layer mới ở trên nó.

- Do đó, Docker image có thể bao gồm một loạt các layer, mỗi layer khác nhau nhưng cũng có nguồn gốc từ layer trước. Các image layer đại diện cho các readonly-file mà chứa thêm các container layer mỗi khi bạn sử dụng nó để khởi động một virtual environment.

### Docker Container là gì?

- Docker container là một run-time environment mà ở đó người dùng có thể chạy một ứng dụng độc lập. Những container này rất gọn nhẹ và cho phép bạn chạy ứng dụng trong đó rất nhanh chóng và dễ dàng.

- Một tính năng quan trọng của container là tính chuẩn xác cho việc chạy các ứng dụng trong container. Không chỉ đảm bảo cho ứng dụng hoạt động như nhau trong các môi trường giống nhau, nó còn làm đơn giản việc cài đặt và chia sẻ cài đặt này cho các thành viên trong team.

- Vì container hoạt động độc lập, nó đảm bảo không làm ảnh hưởng xấu đến các container khác, cũng như server mà nó đang chạy trong đó. Docker được cho là "tạo ra sự độc lập tuyệt vời". Vì vậy, bạn sẽ không cần lo lắng việc máy tính của bạn bị xung đột do ứng dụng đang được phát triển được chạy trong container.

### Casc lệnh cơ bản

- liệt kê các image

    ` docker images -a `

    ![Imgur](https://imgur.com/8Y6a9xj.png)

- xóa một image (phải không container nào đang dùng)

    ` docker image rm imageid `

- tải về một image (imagename) từ [hub.docker.com](https://hub.docker.com/)

    ` docker pull imagename `

    ![Imgur](https://imgur.com/uB5Vrsa.png)

- liệt kê các container

    ` docker ps -a `

- Liệt kê các container đang chạy

    ` docker ps `

    ![Imgur](https://imgur.com/vl2J6ty.png)

- xóa container

    ` docker container rm containerid `

    ![Imgur](https://imgur.com/RpmsV9F.png)

 - Xoá container đang chạy

    `docker rm -f containerid`  

    ![Imgur](https://imgur.com/KOs7nok.png) 

- Tạo mới một container với các tham số --name để đặt tên container và -h để đặt hostname và truy cập container

    ` docker run -it --name "Container Name" -h hostname imageid  `

    ![Imgur](https://imgur.com/e097Sie.png)

- thoát termial vẫn giữ container đang chạy
    
    ` CTRL +P, CTRL + Q `

- Vào termial container đang chạy
    
    ` docker container attach containerid `

    ![Imgur](https://imgur.com/2ysnywQ.png)

- Chạy container đang dừng

    ` docker container start -i containerid `

    ![Imgur](https://imgur.com/7VecpTX.png)

- Chạy một lệnh trên container đang chạy
    
    ` docker exec -it containerid command `

    ![Imgur](https://imgur.com/OI3KX1K.png)