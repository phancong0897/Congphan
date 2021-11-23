# Tìm hiểu cấu hình Handler trong Ansible Playbook

### 1. Handler Ansible là gì ?

- Handler trong Ansible là gì ? Handler là một tác vụ y như ‘task’ trong playbook nhưng khác ở chỗ ‘task‘ được tự động khởi chạy theo flow cấu hình hoặc theo ‘tag‘ còn ‘handler‘ thì không thể khởi chạy nếu không có tín hiệu (notify) triệu gọi sử dụng.

- Tóm gọn

    - handler : là danh sách một hoặc nhiều tác vụ chờ được triệu gọi sử dụng từ task. Một tác vụ handler được khai báo cấu hình y hệt một task, nhưng yêu cầu phải có tên tác vụ rõ ràng để dùng khi triệu gọi.
    
    - notify : là keyword dùng để triệu gọi một hoặc nhiều handler trong một task.

- Lưu ý chung:

    - Khi notify các tác vụ ‘handler‘ , thì các tác vụ sẽ luôn chạy theo thứ tự từ trên xuống đã được định nghĩa.

    - Nếu có hai tác vụ ‘handler’ cùng tên, thì chỉ chạy một cái duy nhất theo thứ tự đầu tiên.

    - Không sử dụng [tag] cho tác vụ handler được.

        ```

        ---
        - hosts: vb
        ...

        tasks:
        ...
        notify:

        handlers:
        ...

        ```

### Ví dụ

Bạn có thể gọi sử dụng nhiều tác vụ handler trong một task.

    ```

    ---
    - hosts: web
      tasks:
    - copy:
        src: main.conf
        dest: /etc/nginx/nginx.conf
      notify: 
        - restart nginx
        - stop nginx

    handlers:
    - name: restart nginx
      service: name=nginx state=restarted

    - name: stop nginx
      service: name=nginx state=stop
    
    ```

