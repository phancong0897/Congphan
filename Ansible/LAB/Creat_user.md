---
- hosts: localhost
  tasks:
  - name: Create a login user
    user:
      name: congpv
      password: "{{ 'password' | password_hash('sha512', 'mysecretsalt') }}"
      state: present
      createhome: yes 