# Tổng quan về lệnh chattr trong linux

### Mục lục

### [1. Chattr?]()

### [2. Cách thêm thuộc tính trên file để bảo veek file khỏi bị xóa]()

### [3. Cách dùng chattr để bảo vệ thư mục]()

### [4. Áp dụng lệnh chattr để bảo vệ /etc/passwd và /etc/shadow]

## 1. Chattr?

- chattr là viết tắt của Change Attribute. Đây là câu lệnh cho phép bạn thay đổi thuộc tính của file giúp bảo vệ file khỏi bị xóa hoặc ghi đè nội dung, dù cho bạn có đang là user root đi nữa.

- Cú pháp lệnh:

    - ` chattr [operator] [flags] [filename] `

    - Các operator:

        - +: Thêm thuộc tính cho file

        - -: Gỡ bỏ thuộc tính khỏi file

        - =: Giữ nguyên thuộc tính của file

        - Các flag (hay các thuộc tính của file):

    - Có nhiều flag, ở đây tôi sẽ giới thiệu các flag thường dùng là:

        + i: Flag này khiến file không thể rename, không thể tạo symlink, không thể thực thi, không thể write. Chỉ có user root mới set và unset được thuộc tính này.
    
        + a: Flag này khiến file không thể rename, không thể tạo symlink, không thể thực thi, chỉ có thể nối thêm nội dung vào file. Chỉ có user root mới set và unset được thuộc tính này.
        
        + d: file có thuộc tính này sẽ không được backup khi tiến trình dump chạy
        
        + S: Nếu một file có thuộc tính này bị sửa, thay đổi sẽ được cập nhật đồng bộ trên ổ cứng.
        + A: Khi file có thuộc tính được truy cập, giá trị atime của file sẽ không thay đổi

## 2. Cách thêm thuộc tính trên file để bảo vệ file khỏi bị xóa

- Chúng ta thêm thuộc tính i (immutable) cho file.

    - ` chattr +i /opt/test/xinchao.txt `

- Bây giờ ta thử xóa file trên:

    ```
    [root@localhost ~]# rm -rf /opt/test/xinchao.txt
    rm: cannot remove ‘/opt/test/xinchao.txt’: Operation not permitted

    ```
- Cách để unset thuộc tính đã thêm cho file

    - Ta sử dụng operator -

    - VÍ dụ tôi sẽ unset thuộc tính i trên file heelo.txt

    ` [root@localhost ~]# chattr -i /opt/test/xinchao.txt `
    
    - Sau đó bạn hãy thử xóa file:

    ` [root@localhost ~]# rm -rf /opt/test/xinchao.txt `

## 3. Cách dùng chattr để bảo vệ thư mục

- Để bảo vệ cả thư mục và các file bên trong thư mục đó, ta dùng flag -R (recursively) và +i với đường dẫn của thư mục đó.

    - ` [root@localhost ~]# chattr -R +i /opt/test `

- Thử xóa thư mục này đi

```
[root@localhost ~]# rm -rf /opt/test
rm: cannot remove ‘/opt/test’: Operation not permitted
```

- Có thể thấy ta không thể xóa thư mục vidu và các file bên trong nó.

- Để unset quyền trên, ta sử dụng flag -R (recursively) và -i với đường dẫn của thư mục đó.

    - ` [root@localhost ~]# chattr -R -i /opt/test `

- Sau khi chạy lệnh này, ta có thể sửa, xóa thư mục và file bên trong như bình thường.

## 4. Áp dụng lệnh chattr để bảo vệ file /etc/passwd và /etc/shadow

- Thêm thuộc tính i cho các file này để tránh bị xóa nhầm. Lưu ý rằng ta không thể tạo thêm user trên khi sử dụng cách này.

    ``` 

    [root@localhost ~]# chattr +i /etc/passwd
    [root@localhost ~]# chattr +i /etc/shadow

    ```

- Thử thêm user

    ``` 
    [root@localhost ~]# useradd congphan
    useradd: cannot open /etc/passwd
    ```

### Nguồn tham khảo:

https://news.cloud365.vn/cli-3-gioi-thieu-lenh-chattr-de-bao-ve-su-toan-ven-cua-file/