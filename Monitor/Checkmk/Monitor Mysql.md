# Giám sát MySQL

## Cài đặt

- Bước 1: Cài đặt checkmk agent

Trước tiên bạn cần cài đặt checkmk agent trên máy chủ đang chạy dịch vụ MySQL

- Bước 2: Cấu hình MySQL

Bạn cần tạo một user mysql dùng cho việc giám sát

` GRANT SELECT, SHOW DATABASES ON *.* TO 'mysqlmonitor'@'localhost' IDENTIFIED BY 'mysqlmonitor'; `

Khai báo thông tin của user dùng để giám sát vào file mysql.cfg

```

cat <<EOF > /etc/check_mk/mysql.cfg
[client]
user=mysqlmonitor
password=mysqlmonitor
EOF

```

Thay đổi quyền cho file vừa tạo để chắc chắn rằng nếu không phải user root thì sẽ không được phép sửa file này

` chmod 400 /etc/check_mk/mysql.cfg `

- Bước 3: Copy file plugin từ checkmk server sang máy chủ mysql

Truy cập vào checkmk server để copy file plugin giám sát mysql sang bên máy đang chạy mysql

` scp /opt/omd/versions/2.0.0p5.cfe/share/check_mk/agents/plugins/mk_mysql root@10.10.10.77:/usr/lib/check_mk_agent/plugins `

Trong đó 10.10.10.77 là địa chỉ của máy chạy mysql

- Bước 4: Thực hiện thêm service trên WATO

Bây giờ truy cập vào Web UI để thực hiện discovery các dịch vụ

