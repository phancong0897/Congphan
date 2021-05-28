# Tìm hiểu mô hình mạng NAT trong KVM

## [1. Tổng quan]()

## [2. Tạo mô hình NAT trên KVM]()

### 1. Tổng quan

- KVM được biết đến là một cơ sở hạ tầng ảo hóa cho nhân Linux. KVM cũng cung cấp các mô hình mạng trong việc ảo hóa network. Các mô hình bao gồm:

    - NAT

    - Host-only

    - Linux-bridge

- Ở [bài trước](https://github.com/phancong0897/Congphan/blob/master/KVM/Linux-Bridge.md#4-c%C3%A0i-%C4%91%E1%BA%B7t-v%C3%A0-qu%E1%BA%A3n-l%C3%AD-linux-bridge) tôi đã giới thiệu về mô hình mạng linux-bridge. Ở bài viết này tôi sẽ giới thiệu về mô hình NAT trên KVM.

- Nếu như với mô hình linux bridge KVM tạo ra một virtual switch thì ta cũng có thể hình dung với mô hình mạng NAT này KVM sẽ tạo ra một thiết bị là virtual router. Khi ta tạo một dải mạng với mô hình NAT thì lúc này virtual router sẽ NAT từ dải mạng mà ta tạo ra ra địa chỉ của card mạng vật lý trên KVM host để đi ra ngoài internet.

- Khi một dải mạng tạo ra ta sẽ thấy trên KVM host xuất hiện một thêm một card mạng. Card mạng này đóng vai trò là gateway cho dải mạng mà ta tạo ra.

- Như vậy VM có thể biết được địa chỉ ở bên ngoài internet nhưng các máy ở bên ngoài sẽ không thể thấy được VM. Như vậy địa chỉ của KVM host sẽ đại diện cho các VM khi giao tiếp với bên ngoài.

### 2. Tạo mô hình NAT trên KVM

- Ở đây tôi sử dụng công cụ virt-manager để tạo và quản lý mô hình mạng này.

    - Trước tiên bạn vào công cụ virt-manager trên thanh công cụ click vào Edit chọn Connection Details chọn Virtual Networks và click vào biểu tượng dấu cộng như bên dưới.

    ![Imgur](https://imgur.com/lyDfavZ.png)

    ![Imgur](https://imgur.com/WBpgM6N.png)

    ![Imgur](https://imgur.com/XE5VJ10.png)

    ![Imgur](https://imgur.com/BQFuddn.png)

    ![Imgur](https://imgur.com/O02BcHR.png)

    ![Imgur](https://imgur.com/RgPzpAh.png)

- Ta có thể kiểm tra VM đã nhận đúng dải mạng

    ![Imgur](https://imgur.com/YcEtoLI.png)

- Ping thử để kiểm tra xem mạng có hoạt động

    ![Imgur](https://imgur.com/wr3cOW3.png)

- Sử dụng lệnh tracert để kiểm tra đường đi gói tin

    ![Imgur](https://imgur.com/vNLg1Ri.png)

Đầu tiên, gói tin sẽ đi ra gateway của card mạng mình vừa tạo( tức 192.168.100.1) , sau đó sẽ đi ra gateway của card mạng vật lý ( tức 10.10.34.1), sau đó sẽ đi ra ngoài.