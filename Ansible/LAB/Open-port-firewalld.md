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