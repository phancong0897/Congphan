# Tạo scripts check bad disk

## Thực hiện trên host cần giám sát

- Chạy lệnh để cài đặt các gói smartmontool:
`yum install smartmontools -y`

- Cài đặt plugins :

    `  wget https://raw.githubusercontent.com/uncelvel/tutorial-ceph/master/docs/monitor/check_mk/plugins/smart `

    - Phần quyền cho file .smart
    
    ` chmod +x smart`

    - Run scripts

    ` ./smart `

- Truy cập checkmk, chọn đúng host và discovery lại.