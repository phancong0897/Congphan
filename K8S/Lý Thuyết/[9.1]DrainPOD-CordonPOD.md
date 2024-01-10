# Drain POD and Cordon  POD

- Cordon nodes: Để đảm bảo không có pod mới nào được lên lịch thực hiện trên node này nữa
    
    ` kubectl cordon [node-name] `

- Drain node: Sẽ thực hiện move hết những Pod đang chạy trên node này để giải phóng node
    
    ` kubectl drain --ignore-daemonsets --force [node-name] 