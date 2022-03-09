#Playbook creat user bằng ansible

```
- hosts: localhost
  tasks:
  - name: Create a login user
    user:
      name: congpv
      password: "{{ 'yourpasswd' | password_hash('sha512', 'mysecretsalt') }}"
      state: present
      createhome: yes 
```

```

- name: user module demo
  hosts: all
  become: true
  tasks:
    - name: user example present
      ansible.builtin.user:
        name: example
        password: "{{ 'yourpasswd' | password_hash('sha512', 'mysecretsalt') }}"
        groups:
          - wheel
          - adm
        state: "present"
        shell: "/bin/bash"
        system: false
        create_home: true
        home: "/home/example"
        comment: "Ansible example"
        generate_ssh_key: true

```