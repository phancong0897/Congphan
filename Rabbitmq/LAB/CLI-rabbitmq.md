# Các câu lệnh quản trị rabbitmq

- Tạo user quản trị và phân quyền user làm admin

    ```

    rabbitmqctl add_user [user name] [Password]
    rabbitmqctl set_user_tags [user name] administrator

    ```

- Tạo vhost và phân quyền

    ```

    rabbitmqctl add_vhost [name vhost]
    rabbitmqctl set_permissions -p [name vhoste] [user name] "^admin-.*" ".*" ".*"
    rabbitmqctl set_permissions -p [name vhoste] [user name] ".*" ".*" ".*"

    ```

- List user và tag

    ` rabbitmqctl list_users `

- Kiểm tra quyền của user

    ` rabbitmqctl list_user_permissions [user name] `

- Kiểm tra trạng thái cluster

    ` rabbitmqctl cluster_status `

- Stop service rabbitmq

    ` systemctl stop rabbitmq-server `

- Start service rabbitmq

    ` systemctl start rabbitmq-server `

- Restart vhost

    ` rabbitmqctl restart_vhost --vhost=[name vhost] `

- Loại bỏ node khỏi cluster

    ` rabbitmqctl reset ` 

- Stop application

    ` rabbitmqctl stop_app `

- Start application

    ` rabbitmqctl start_app `

- Joind node vào cluster rabbitmq

    ` rabbitmqctl join [cluster] `

- Delete queue trong vhost

    ` rabbitmqctl --vhost=[name vhost] delete_queue [name queue] `

- set_vm_memory_high_watermark: Theo mặc định, khi máy chủ RabbitMQ sử dụng trên 40% RAM khả dụng, nó sẽ báo động bộ nhớ và chặn tất cả các kết nối đang xuất bản thông báo. Sau khi cảnh báo bộ nhớ đã xóa (ví dụ: do máy chủ phân trang các thông báo vào đĩa hoặc gửi chúng cho các máy khách sử dụng và xác nhận việc phân phối ) dịch vụ bình thường sẽ tiếp tục.

    ` rabbitmqctl set_vm_memory_high_watermark 0.5 (Ngưỡng 50%) `

- set_disk_free_limit :Định cấu hình giới hạn dung lượng trống trên đĩa : Giới hạn dung lượng trống của ổ đĩa được định cấu hình với cài đặt disk_free_limit . Theo mặc định, 50MB được yêu cầu để trống trên phân vùng cơ sở dữ liệu (xem mô tả vị trí tệp cho vị trí cơ sở dữ liệu mặc định). Tệp cấu hình này đặt giới hạn dung lượng trống của ổ đĩa thành 1GB

    ` rabbitmqctl set_disk_free_limit 1G `

- set_vhost_limits: Trong hầu hết các shell, bạn rất có thể cần phải trích dẫn điều này.

	- Các giới hạn được công nhận là:

		kết nối tối đa
		hàng đợi tối đa

	- Sử dụng giá trị âm để chỉ định "không có giới hạn".

	- lệnh này giới hạn số lượng tối đa các kết nối đồng thời trong vhost "vnpay" là 64:

	` Rabbitmqctl set_vhost_limits -p vnpay '{"max-results": 64}' `
	
	- Lệnh này giới hạn số lượng hàng đợi tối đa trong vhost "vnpay" là 256:

	` Rabbitmqctl set_vhost_limits -p vnpay '{"max-queues": 256}' `
	
	- Lệnh này xóa giới hạn số lượng kết nối tối đa trong vhost "vnpay":

	` Rabbitmqctl set_vhost_limits -p vnpay '{"max-results": -1}' `
	
	- Lệnh này vô hiệu hóa các kết nối máy khách trong vhost "vnpay":

	` Rabbitmqctl set_vhost_limits -p vnpay '{"max-results": 0}' `

#### Link tham khảo

https://www.rabbitmq.com/rabbitmqctl.8.html


