# Tổng quan về Rsync

### Mục lục

### [1. Giới thiệu về Rsync]()

### [2. Cách sử dụng Rsync]()

### [3. Sử dụng Rsync kết hợp với SSH để Tự động hóa việc đồng bộ ]()

### [4. Cấu hình tự động hóa Rsync]()
### 1. Giới thiệu về Rsync

- Rsync(remote sync) là công cụ đồng bộ file, thư mục của hệ điều hành Linux. Nó sử dụng thuật toán khi copy dữ liệu sao cho dữ liệu phải copy là nhỏ nhất (chỉ copy những gì thay đổi giữa nguồn và gốc), khi đồng bộ nó giữ nguyên mọi thuộc tính của file, thư mục (từ chủ sở hữu, quyền truy cập file ...).

- Các tính năng quan trọng của rsync:

    - Tốc độ: Lần đầu tiên rsync sao chép toàn bộ nội dung giữa các thư mục nguồn và các thư mục đích. Trong lần tiếp theo rsync chỉ chuyển các khối hoặc byte đã thay đổi đến vị trí đích, điều này làm cho việc chuyển rất nhanh.

    - Bảo mật: rsync cho phép mã hóa dữ liệu bằng giao thức ssh trong khi truyền.

    - Ít băng thông: rsync sử dụng nén và giải nén khối dữ liệu theo từng khối ở đầu gửi và nhận tương ứng. Vì vậy, băng thông được sử dụng bởi rsync sẽ luôn ít hơn so với các giao thức truyền tệp khác.

    - Đặc quyền : Không yêu cầu đặc quyền đặc biệt để cài đặt và thực thi rsync.
- Mặc định hầu hết các bản phân phối Linux có sẵn công cụ này, nếu chưa có thì chúng ta chạy lệnh sau:

    - Đối với bản phân phối Red Hat/CentOS

        ` yum install rsync `

    - Đối với bản phân phối Debian/Ubuntu

        ` apt-get install rsysnc `

### 2. Cách sử dụng rsync

- Cú pháp của rsync:

    ` rsync option source destination `

- Trong đó:

    - source: đường dẫn thư mục chứa dữ liệu gốc muốn đồng bộ, nơi truyền dữ liệu.

    - destination: đường dẫn nơi chứa dữ liệu đồng bộ đến, nơi nhận dữ liệu.
    
    - option: các tham số tùy chọn.

- Các tùy chọn thường được sử dụng:

    - -a: Tùy chọn này sẽ gọi toàn user, group, permission của dữ liệu truyền đi.

    - -v: Hiển thị trạng thái truyền tải file để chúng ta có thể theo dõi.

    - -h: Kết hợp với tùy chọn -v để định dạng dữ liệu show ra dễ nhìn hơn.

    - -z: Nén dữ liệu trước khi truyền đi giúp tăng tốc quá trình đồng bộ file.
    
    - -e: Sử dụng giao thức SSH để mã hóa dữ liệu.
    
    - -P: Tùy chọn này dùng khi đường truyền không ổn định, nó sẽ gửi tiếp các file chưa được gửi đi khi có kết nối trở lại.
    
    - --delete: Xóa dữ liệu ở destination nếu source không tồn tại dữ liệu đó.
    
    - --exclude: Loại trừ dữ liệu không muốn truyền đi.
- Lưu ý: Chỉ dùng --delete khi bạn chắc chắn rằng "source" và "destination" đều đúng, nhất là về vị trí, nếu bạn để sai vị trí hoặc sai tên thư mục thì có thể dẫn đến mất dữ liệu toàn bộ khi dùng rsync.

- rsync trên cùng máy chủ: 

    - Dùng câu lệnh: rsync -<tùy chọn> <thư mục nguồn> <thư mục đích>

- rsync giữa các máy chủ

    - Dùng lệnh: rsync -<tùy chọn> <thư mục nguồn> [<user>@<IP đích]:<thư mục đích>



### 3. Sử dụng Rsync kết hợp với SSH để Tự động hóa việc đồng bộ

### 3.1. Download từ Server

- Trước tiên ta thử thực hiện việc đồng bộ bằng lệnh sau:

    - ` rsync -avzh ssh root@172.16.2.230:/root/test2/ /root/backup `

- Trong đó:

    - rsync -avzhe: Câu lệnh rsync cùng với các tùy chọn.

    - ssh : Là việc sử dụng rsync qua SSH.

    - root@172.16.2.230:/root/test2/: Thư mục của máy chủ đích mà ta muốn đồng bộ.

```

[root@localhost ~]# rsync -avzhe ssh root@172.16.2.230:/root/test2/ /root/backup/
The authenticity of host '172.16.2.230 (172.16.2.230)' can't be established.
ECDSA key fingerprint is SHA256:ehUsHPwoKbhzLtDU2Jttwvd14n4aFUmgQZV92u6Qyko.
ECDSA key fingerprint is MD5:bb:1d:5a:88:c6:9d:31:df:c9:59:6c:66:cb:30:df:6e.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.16.2.230' (ECDSA) to the list of known hosts.
root@172.16.2.230's password:
receiving incremental file list
./
alo.txt
aloalo.txt
test.txt
root@/
root@/alo.txt
root@/aloalo.txt
root@/test.txt

```

### 3.2 Upload lên Server 

- SSH Port tiêu chuẩn

    - ` rsync -avH /root/test/hello.txt -e ssh root@172.16.2.230:/root/test2/ ` 

- Trong đó:

    -  /root/test/hello.txt: local path nơi tệp sẽ được tải lên máy chủ đích (từ xa), trong đó file là tên tệp.

    - root: username của người dùng từ xa mà bạn sẽ đăng nhập vào máy chủ mục tiêu (từ xa).

    - 172.16.2.230: Tên máy chủ hoặc địa chỉ IP của máy chủ đích (từ xa).

    - /root/tets2: Đường dẫn từ xa cho tệp sẽ được tải lên máy chủ đích (từ xa), trong đó file là tên tệp.


### 4. Cấu hình tự đồng hóa rsync

- Bước 1: Tạo tệp Shell Scripts

    - ` vi /backups/rsync.sh `

- Ta thêm dòng lệnh rsync ở trên vào Scripts sau đó lưu lại. Vì đã có SSH Key pair nên chúng ta chỉ cần 1 dòng này trong scripts:

    - ` rsync -avzhe "ssh -p port" user@IP:/backup/ /backups/ `

- Bước 2: Lưu lại và phần quyền thực thi cho tệp Shell Scripts vừa tạo

    - ` chmod +x /backups/rsync.sh `

- Bước 3: Đưa Script vào Crontab để tự động hóa

    - ` crontab -e `

    - Thêm nội dung sau:

    ` 0 8 * * * /backups/rsync.sh `

### Nguồn hướng dẫn

https://hocvps.com/rsync/

https://bizflycloud.vn/tin-tuc/huong-dan-cai-dat-cong-cu-truyen-tai-du-lieu-rsync-tren-linux-20180309115710949.htm

Tại đây có nghĩa là việc đồng bộ sẽ được diễn ra vào lúc 8 giờ sáng hàng ngày. Ta lưu lại là thành công.


