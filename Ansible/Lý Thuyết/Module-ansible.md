# Tìm hiểu về các module thường được sử dụng trong ansible

### Item

- Đặt trường hợp , ta cần cài đặt các gói httpd, mariadb-server, php-fpm... thay vì viết module yum cho từng gói cài đặt. Giờ ta sẽ nhóm vào Item để chạy 1 lần luôn. Hay ta cần thao tác giống nhau trên các service ta có thể sử dụng item

```
- hosts: centos7
  sudo: yes

  tasks:
    - name: Install Apache.
      yum:
        name: "{{ item }}"
        state: present
      with_items:
        - apache2
        - mysql-server
        - php
        - php-mysql
    - name: Restart Apache and Mysql
      service:
        name: "{{item}}"
        state:  running
      with_items:
          - apache2
          - mysql
```

### Handlers

Handlers giúp chúng ta gọi lại hành động thực thi nhiều lần (notify) mà không cần phải viết lại.

```
- hosts: centos7
  sudo: yes



  tasks:
    - name: Install Apache.
      apt:
        name: "{{ item }}"
        state: present
      with_items:
        - apache2
        - mysql-server
        - php
        - php-mysql

    - name: deploy html file
      template:
        src: /tmp/index.html
        dest: /var/www/html/index.html
      notify: restart web

  handlers:
    - name: restart web
      service:
        name: "{{ item }}"
        state:  restart
      with_items:
          - apache2
          - mysql
```

### Variables và Template

Đặt giá trị cho biến cố định

```
- hosts: centos7
  sudo: yes


  vars:
     - domain_name: "tuanduong.com"
     - index_file: "index.html"
     - config_file: "tuanduong.conf"


  tasks:
    - name: Install Apache.
      apt:
        name: "{{ item }}"
        state: present
      with_items:
        - apache2
        - mysql-server
        - php
        - php-mysql

    - name: deploy html file
      template:
        src: /tmp/{{ index_file }}
        dest: /var/www/html/index.html
      notify: restart web

  handlers:
    - name: restart web
      service:
        name: "{{ item }}"
        state:  running
      with_items:
          - apache2
          - mysql
```

#### FLIE

Trong Ansible, có nhiều các module làm việc với tệp tin, thư mục, links trên các node đích (node client) như copy, template, file,… thường được sử dụng. Trước tiên, chúng ta sẽ cùng tìm hiểu về file module. File module giúp quản lý tập tin và các thuộc tính của nó. Ngoài việc taọ, xóa, xác định vị trí của tệp tin file module cũng đặt các quyền và quyền sở hữu hay thiết lập symlinks cho tệp tin.

```
  hosts: centos7
  tasks:
    - name: Change file ownership, group and permissions
      file:
        path: /etc/thuntt.config
        owner: thuntt
        group: thuntt
        mode: '0644'
 
    - name: Create an insecure file
      file:
        path: /viblo
        owner: root
        group: root
        mode: '1777'
 
    - name: Create a symbolic link
      file:
        src: /file/to/link/to
        dest: /path/to/symlink
        owner: thuntt
        group: thuntt
        state: link
```

#### COPY

Copy module là module thường được sử dụng khi chúng ta muốn sao chép một tệp tin từ Ansible server (Management node) đến các node đích (client node).

```
- name: copy file from local to remote with owner, group and file permissions (octal)
  copy:
    src: test_file
    dest: $HOME/test_file
    owner: thuntt
    group: thuntt
    mode: 0644
 
- name: copy file from local to remote with owner, group and file permissions (octal)
  copy:
    src: test_file
    dest: $HOME/test_file
    owner: thuntt
    group: thuntt
    mode: 0644
 
- name: copy file from local to remote with root as the owner (become required)
  copy:
    src: test_file
    dest: "/home/{{ ansible_user }}/test_file"
    owner: root
    group: root
    mode: 0644
  become: true
```

