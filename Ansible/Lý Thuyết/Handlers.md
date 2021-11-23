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

### 2. Handler Listen

Từ Ansible version 2.2 trở đi thì Ansible đã hỗ trợ một cấu hình mang tên ‘Handler Listen‘. Với keyword ‘listen‘ sử dụng khi cấu hình một tác vụ ‘handler‘, ta có thể hiểu nó giống như việc gom nhóm các tác vụ ‘handler‘ thành một group theo tên ‘listen‘. Và bạn chỉ cần gọi tên của ‘listen‘ được khai báo là có thể chạy toàn bộ các tác vụ ‘handler‘ con của nhóm ‘listen‘ đó.

```
- hosts: all

  tasks:
    - name: configure nginx
      template: src=nginx.conf.j2 dest=/etc/nginx.conf
      notify: restart_nginx_cluster

  handlers:
     - name: restart_uwsgi
       service: name=uwsgi state=restarted
       listen: restart_nginx_cluster            

     - name: restart_nginx
       service: name=nginx state=restarted
       listen: restart_nginx_cluster

```

- Giải thích:

    - Ta có một playbook đơn giản, là cần copy cấu hình file nginx.conf lên con máy chủ remote. Sau đó cần restart toàn bộ service liên quan webserver nginx, ở đây là uwsgi và nginx . Thì thay vì phải gọi (notify) từng tên tác vụ handler như ở phần đầu, thì giờ mình có thể gọi tên nhóm ‘listen‘ là ‘restart_nginx_cluster‘ tức nhóm các tác vụ ‘handler‘ đã được khai báo và thực thi . Việc khai báo ‘listen‘ này nhìn chung rất gọn gàng nếu bạn có nhiều ‘handler’ cần phải sử dụng cùng lúc.

    - Lưu ý là các tác vụ ‘handler‘ trong một nhóm ‘listen‘ vẫn chạy theo thứ tự đã khai báo trong mục danh sách tác vụ ‘handler‘ nhé.