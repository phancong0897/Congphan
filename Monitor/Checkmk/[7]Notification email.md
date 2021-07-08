# Cấu hình cảnh báo qua email

### 1. Cấu hình postfix

- Remove Sendmail

    Trước tiên cần kiểm tra xem sendmail đã được cài đặt chưa bằng câu lệnh

    ` rpm -qa | grep sendmail `

    Nếu có kết quả trả về chứng tỏ sendmail đã được cài đặt. Ta cần remove nó

    ` yum remove sendmail* `

- Install postfix

    - Cài đặt postfix và một số gói liên quan

    ` yum -y install postfix cyrus-sasl-plain mailx `

    - Đặt postfix như MTA mặc định của hệ thống

    ` alternatives --set mta /usr/sbin/postfix `

    Nếu câu lệnh bị lỗi và trả về output /usr/sbin/postfix has not been configured as an alternative for mta thì thực hiện lệnh sau:

    ` alternatives --set mta /usr/sbin/sendmail.postfix `

    - Restart và enable postfix

    ```

    systemctl restart postfix
    
    systemctl enable postfix

    ```

- Configure Postfix

    - Mở file main.cf để chỉnh sửa

    ` vi /etc/postfix/main.cf `

    Thêm vào cuối file những dòng sau

    ```

    myhostname = hostname.example.com
    relayhost = [smtp.gmail.com]:587
    smtp_use_tls = yes
    smtp_sasl_auth_enable = yes
    smtp_sasl_password_maps = hash:/etc/postfix/sasl_passwd
    smtp_tls_CAfile = /etc/ssl/certs/ca-bundle.crt
    smtp_sasl_security_options = noanonymous
    smtp_sasl_tls_security_options = noanonymous

    ```

    - Configure Postfix SASL Credentials

    Tạo file /etc/postfix/sasl_passwd và thêm vào dòng sau

    ` [smtp.gmail.com]:587 username:password `

    - Trong đó:

    username: là địa chỉ email dùng để gửi mail

    password: là password của email dùng để gửi mail

    - Phân quyền cho file vừa tạo


    ```

    postmap /etc/postfix/sasl_passwd
    chown root:postfix /etc/postfix/sasl_passwd*
    chmod 640 /etc/postfix/sasl_passwd*
    systemctl reload postfix

    ```

### 2. Cấu hình cảnh báo trên site checkmk

- Tạo group chứa những user được nhận cảnh báo

![Imgur](https://imgur.com/QvudM6x.png)

![Imgur](https://imgur.com/bzeVpI4.png)

    Khai báo thông tin của group

    1: Khai báo tên của group
    2: Tên alias của group
    
![Imgur](https://imgur.com/V6XQLOr.png)

- Tạo user cho group để nhận thông báo
    
    Chọn 1 để vào users
    Chọn 2 để tạo user mới

![Imgur](https://imgur.com/xLE851g.png)

![Imgur](https://imgur.com/7VGH0eP.png)

- Cấu hình cảnh báo

Chọn event --> notifications --> create rule

![Imgur](https://imgur.com/Zp7w2Lz.png)

- Khai báo description cho rule và chọn method là Email

![Imgur](https://imgur.com/jyiKzHv.png)

![Imgur](https://imgur.com/IgRdml8.png)

[Imgur](https://imgur.com/XJ6SQu1.png)

- Chọn changes và lưu lại cấu hình

![Imgur](https://imgur.com/GJ8bnTP.png)

- Chọn group nhận cảnh báo

![Imgur](https://imgur.com/IgRdml8.png)

- Điều kiện cảnh báo khi server chuyển từ UP -> DOWN

![Imgur](https://imgur.com/XJ6SQu1.png)



