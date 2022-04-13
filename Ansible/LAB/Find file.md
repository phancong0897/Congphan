# Playbook tìm các file .sh 

```

- hosts: localhost
  gather_facts: false
  vars:
  tasks:
  - name: parse /tmp directory
    find:
      paths: /tmp
      patterns: '*.sh'
    register: list_of_files

  - debug:
      var: item.path
    with_items: "{{ list_of_files.files }}"

  - name: Before Ansible 2.3, option 'dest', 'destfile' or 'name' was used instead of 'path'
    replace:
      path: "{{ item.path }}"
      regexp: 'vhosts/cert_wildcard_vnpaytest_2021/cert_wildcard_vnpaytest_2021.cer'
      replace: 'aloalo'
    with_items: "{{ list_of_files.files }}"


```