# Hướng dẫn sử dụng extend disk, tạo pvs, vg ,lv theo disk 

- Mô hình sẽ là khi ta thêm một disk vào server, ta không thể biết được đó là disk có tên như nào để tiến hành extend. Tôi sẽ hướng dẫn các bạn theo mô hình này nhé.

### Viết playbook cho phần extend như sau:

- Tạo folder trong roles để sử dụng roles ansible

    `mkdir /etc/ansible/roles/extendisk `

- Tạo task trong folder checkdisk:
    ` mkdir /etc/ansible/roles/extendisk/tasks `

- Viết playbook cho phần extend với nối dụng như sau:

    ` vi /etc/ansible/roles/extendisk/tasks/main.yml `

```

- name: create partion
  parted:
    device: "{{ pvs_name }}"
    number: 1
    flags: [lvm]
    state: present
    part_end: "{{ Disk }}"

- name: format disk
  filesystem:
    fstype: "{{ file_system }}"
    dev: "{{ pvs_name }}1"

- name: extend vgs
  lvg:
    vg: "{{vg_name}}"
    pvs: "{{ pvs_name }}1"

- name: creat lvm
  lvol:
    vg: "{{vg_name}}"
    lv: "{{lv_name}}"
    size: 100%PVS

```

### Viết playbook checkdisk

- Tạo folder trong roles để sử dụng roles ansible

  `mkdir /etc/ansible/roles/checkdisk `

- Tạo task trong folder checkdisk:

    ` mkdir /etc/ansible/roles/checkdisk/tasks `

- Viết playbook cho phần extend với nối dụng như sau:

    ` vi /etc/ansible/roles/checkdisk/tasks/main.yml `

```
- name: check disk
  shell: lsblk
  register: lsblk_status

  - debug:
      var: lsblk_status

```

### Viết playbook chạy role

` vi /etc/ansbile/playbook.yml `

```
- hosts: centos7
  roles:
    - checkdisk

- hosts: centos7
  vars_prompt:

  - name: "pvs_name"
    prompt: "name pvs"
    private: no

  - name: "vg_name"
    prompt: "Nhap ten VG"
    private: no

  - name: "lv_name"
    prompt: "Nhap ten LV"
    private: no

  - name: "file_system"
    prompt: "Nhap dinh dang"
    private: no

  - name: "Disk"
    prompt: "Disk"
    private: no

  roles:
    - extendisk

- hosts: centos7
  vars_prompt:
  - name: "mount_point"
    prompt: "Nhap folder mount"
    private: no
  tasks:

  - name: mkdir data
    shell: mkdir -p /vnpay

  - name: mount disk
    mount:
      path: /vnpay
      src: "/dev/mapper/{{mount_point}}"
      fstype: fs
      state: present

```


- Chạy playbook

    ` ansible-playbook -i /etc/ansible/hosts /etc/ansible/playbook.yml `

- Hoạt động của file như sau:

    - Trước tiên mình sẽ sử dụng lệnh lsblk để kiểm tra disk vừa thêm vào
    - Sau đó sẽ nhập các pvs,vg,lv để gán vào biến
    - Chạy task để tạo pvs, vg,lv
    - Sau đó tiến hành mount vg vừa tạo vào folder vnpay