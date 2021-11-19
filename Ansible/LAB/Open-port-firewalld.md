# Để mở port firewall dùng ta dùng module firewalld

```
- hosts: centos7
  tasks:
  - name: firewall
    firewalld:
      port: 8081/tcp
      zone: public
      permanent: true
      state: enabled
      immediate: true

  - name: restart serice firewalld
    service:
      name: firewalld
      state: restarted

```

### Để mở service như http, https, ssh ta sử dụng module như sau:

```
- hosts: centos7
  tasks:
  - name: firewall
    firewalld:
      service: http
      zone: public
      permanent: true
      state: enabled
      immediate: true

  - name: restart serice firewalld
    service:
      name: firewalld
      state: restarted

```