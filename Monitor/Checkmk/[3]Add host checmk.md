# Hướng dẫn add host để checkmk giám sát

Sau khi cài đặt xong check_mk thì tất nhiên chúng ta sẽ phải thêm host vào để chúng ta giám sát chúng. Ở bài này tôi sẽ hướng dẫn các bạn làm sao để thêm một host để check_mk giám sát nó

Sau khi bạn đã làm xong [bài trước](https://github.com/phancong0897/Congphan/blob/master/Monitor/Checkmk/Add%20agent%20linux.md)  chúng ta cùng đi vào bài này nào. Sau khi đăng nhập vào web của check_mk thì bạn hãy làm theo các hướng dẫn bên dưới này nha!

### Bước 1: Chọn setup , chọn host

![Imgur](https://imgur.com/GQAKVDJ.png)

### Bước 2: Chọn add host để add host ceheckmk giám sát

![Imgur](https://imgur.com/gNMSZCi.png)

### Bước 3: Tiếp theo hãy điền thông tin của host mà bạn muốn giám sát bao gồm tên và địa chỉ IP của chúng. Sau đó Save & Save & go to connection tests để kiểm tra

![Imgur](https://imgur.com/aoSFNEX.png)

Sau khi kiểm tra phần ping và agent đã hoạt động --> Save & go to host properties

![Imgur](https://imgur.com/t7dBlUx.png)

### Bước 4: Ta chọn change để lưu cấu hình giám sát

![Imgur](https://imgur.com/EEUXsol.png)

Ta chọn Activate on selected all sites

![Imgur](https://imgur.com/D42GcQR.png)

Sau khi hoàn tất các bước , ta chọn monitor, all host và click host vừa adđ để kiểm tra lại, Do trước đó tôi đã add agent để giám sát Mysql, bài viết sau tôi sẽ hướng dẫn các bạn giám sát service Mysql.

![Imgur](https://imgur.com/OD1RIoo.png)


