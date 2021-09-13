# Sử dụng Dockerfile để tự động tạo các image trong Docker

# Sử dụng Dockerfile

- Dockerfile là một file text, trong đó chứa các dòng chỉ thị để Docker đọc và chạy theo chỉ thị đó để cuối cùng bạn có một image mới theo nhu cầu. Khi đang có một Dockerfile giả sử có tên là Dockerfile để ra lệnh cho Docker chạy nó bạn có thể gõ:

    ` docker build -t nameimage:version --force-rm -f Dockerfile . `

    - Bạn chú ý dấu . ở cuối lệnh docker build ở trên, có nghĩa tìm file có tên Dockerfile ở thư mục hiện tại.

    - -t nameimage:version là đặt tên và tag được gán cho image mới tạo ra.

### Các chỉ thị Dockerfile

- FROM : mọi Docker file đều có chỉ thị này, chỉ định image cơ sở

    Chỉ thị này chỉ ra image cơ sở để xây dựng nên image mới. Để xây dựng từ image nào đó thì bạn cần đọc document của Image đó để biết trong đó đang chứa gì, có thể chạy các lệnh gì trong đó ... Ví dụ, nếu bạn chọn xây dựng từ image centos:laste thì bạn bắt đầu bằng hệ điều hành CentOS và bạn có thể cài đặt, cập nhật các gói với yum, ngược lại nếu bạn chọn ubuntu:latest thì trình quản lý gói của nó là APT ...

- COPY ADD : sao chép dữ liệu

    COPY và ADD Trong Dockerfile được dùng để thêm thư mục, file vào Image. Cú pháp viết đó là:

    ` ADD thư_mục_nguồn thư_mục_đích `

    Trong đó thư_mục_nguồn là thư mục ở máy chạy Dockerfile, chứa dữ liệu cần thêm vào. thư_mục_đích là nơi dữ liệu được thêm vào ở container.

- ENV Trong Dockerfile

    Chỉ thị này dùng để thiết lập biến môi trường, như biến môi trường PATH ..., tùy hệ thống hay ứng dụng yêu cầu biến môi trường nào thì căn cứ vào đó để thiết lập.

    ` ENV biến giá_trị `

- RUN : chạy các lệnh.

    RUN Trong Dockerfile thi hành các lệnh, tương tự bạn chạy lệnh shell trên OS từ terminal.

    ` RUN lệnh-và-tham-số-cần-chạy `

    ` RUN ["lệnh", "tham số1", "tham số 2" ...] `

- VOLUME : gắn ổ đĩa, thư mục

    VOLUME Trong Dockerfile  chỉ thi tạo một ổ đĩa chia sẻ được giữa các container.

    ` VOLUME /dir_vol `

    Xem thêm Chia sẻ dữ liệu giữa các Container


- USER : user

    USER Trong Dockerfile
    
    Bạn thêm user được dùng khi chạy các lệnh ở chỉ thị RUN CMD WORKDIR.

    USER private

- WORKDIR : thư mục làm việc

    Thiết lập thư mục làm việc hiện tại chi các chỉ thị CMD, ENTRYPOINT, ADD thi hành.

    ` WORKDIR path_current_dir `

- EXPOSE : thiết lập cổng

    Để thiết lập cổng mà container lắng nghe, cho phép các container khác trên cùng mạng liên lạc qua cổng này hoặc đỉ ánh xạ cổng host vào cổng này.

    ` EXPOSE port `

- ENTRYPOINT, CMD Trong Dockerfile

    - Chạy lệnh trong chỉ thị này khi container được chạy.

    ENTRYPOINT commnad_script

    ENTRYPOINT ["command", "tham-số", ...]

    - CMD ý nghĩa tương tự như ENTRYPOINT, khác là lệnh chạy bằng shell của container.

    CMD command param1 param2

    - Chú ý ở dạng sau của CMD thì nó lại là thiết lập tham số cho ENTRYPOINT

    CMD ["tham-số1", "tham-số2"]

Đó là các chỉ thị cơ bản, đủ để bạn tự viết các Dockerfile xây dựng ra các image chứa các thành phần theo nhu cầu của mình.