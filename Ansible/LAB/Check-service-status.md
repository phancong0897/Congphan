# Detect status service và thực thi playbook

- hosts: localhost
  tasks:
  - name: Populate service facts
    service_facts:

  - name: firewall
    firewalld:
      port: 8081/tcp
      zone: public
      permanent: true
      state: enabled
      immediate: true
    when:
      - ansible_facts.services['firewalld.service'].state == 'running'