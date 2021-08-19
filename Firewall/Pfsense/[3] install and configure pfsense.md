# Cài đặt và cấu hình cơ bản pfsense

### 1. Chuẩn bị

- Tài file iso của pfsense

    Tải file iso từ trang https://www.pfsense.org/download/

- Chuẩn bị VM để cài đặt Pfsense:

    - Chuẩn bị 2 Card mạng: 1 đường LAN và 1 đường WAN (kết nối Internet).

    - Đường WAN sẽ sử dụng để Client kết nối tới.

    - Đường LAN để quản trị, có quyền SSH tới các VM thông qua LAN.

### 2. Cài đặt Pfsense

![Imgur](https://imgur.com/YupcTjW.png)

![Imgur](https://imgur.com/oO92RXV.png)

![Imgur](https://imgur.com/tyNoiVH.png)

![Imgur](https://imgur.com/IS6Ak5q.png)

![Imgur](https://imgur.com/h8I9D75.png)

![Imgur](https://imgur.com/Z44zMEk.png)

Sau khi hoàn thành các bước trên, đợi 1-2 phút để VM khởi động lại.

### Cấu hình cơ bản cho pfsense

- Cấu hình IP tĩnh cho đường WAN

Sau khi cấu hình xong sẽ có IP cho đường WAN, bạn có thể thay đổi IP theo mô hình của các bạn nhé.

![Imgur](https://imgur.com/7tOf5mi.png)

- Giao diện Pfsense sau khi hoàn thành:

Tài khoản mặc định để truy cập là admin / pfsense

![Imgur](https://imgur.com/9wkfnON.png)

- Thiết lập cấu hình Pfsense qua giao diện.

![Imgur](https://imgur.com/RgLAisE.png)

![Imgur](https://imgur.com/deIkfEc.png)

![Imgur](https://imgur.com/NgxhqtC.png)

![Imgur](https://imgur.com/C7EHjz1.png)

![Imgur](https://imgur.com/Wze3peL.png)

![Imgur](https://imgur.com/T4DRhcU.png)

![Imgur](https://imgur.com/UIpZerR.png)

![Imgur](https://imgur.com/yLVgDwH.png)

![Imgur](https://imgur.com/GJR0KxB.png)

- Tắt Hardware Checksum Offloading

![Imgur](https://imgur.com/nQLVrNZ.png)

![Imgur](https://imgur.com/X707mU0.png)

![Imgur](https://imgur.com/FX2Ivw7png)

- Bổ sung Rule Firewall cho phép kết nối giao diện Pfsense qua IP Public

Cho phép kết nối qua HTTP:

![Imgur](https://imgur.com/iOTTvUN.png)

![Imgur](https://imgur.com/TkdAWEU,png)

Cho phép kết nối qua HTTPS:

![Imgur](https://imgur.com/Kgj99k9.png)

![Imgur](https://imgur.com/w6A29m1.png)

- Chuyển kết nối dashboard mặc định Pfsense tới HTTPS

![Imgur](https://imgur.com/t3RPzhx.png)

- Enable card mạng LAN

![Imgur](https://imgur.com/o8pcLqj.png)

![Imgur](https://imgur.com/jFIat0f.png)

Đặt IPv4 cho card mạng LAN: 192.168.50.10 sau đó Save lại:

![Imgur](https://imgur.com/zZjFuwy.png)

Apply Changes để xác nhận cấu hình.

- Quay lại màn hình Dashboard đã nhận IP của 2 card mạng:

![Imgur](https://imgur.com/TMPw8jj.png)





