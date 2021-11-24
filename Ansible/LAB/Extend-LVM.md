# Hướng dẫn sử dụng ansible để extend lvm

- Tạo file playbook.yml với nội dung như sau:

```
- hosts: client02
  tasks:

  - name: create partion
    parted:
      device: /dev/sdb
      number: 1
      flags: [lvm]
      state: present
      part_end: 20GB

  - name: format disk
    filesystem:
      fstype: xfs
      dev: /dev/sdb1

  - name: extend vgs
    lvg:
      vg: rhel_vm02
      pvs: /dev/sda2,/dev/sdb1

  - name: extend lvm
    lvol:
      vg: rhel_vm02
      lv: home
      size: 100%PVS
      resizefs: true


```