# Sự khác biệt giữa module command và shell trong ansible

Sẽ có những lúc bạn cần thực hiện các câu lệnh đơn giản thay vì xài module khác phải không nào ? Việc thực hiện các câu lệnh có thể được hỗ trợ bởi module Shell hoặc module Command . Vậy đâu là sự khác biệt giữa hai module này mà ta cần biết để lựa chọn sử dụng.

### Module Command

- Với module Command, thì các câu lệnh được truyền vào sẽ được thực thi không thông qua shell linux. Vì vậy một số thứ như biến môi trường $HOME sẽ không tồn tại, các tính năng chuyển hướng input/output ( < , > ) hay pipes ( | ) sẽ không hoạt động. Module Command an toàn hơn vì sẽ không ảnh hưởng đến môi trường của user thực thi.

- Cú pháp cơ bản

    ```
    - name: return motd to registered var
      command: cat /etc/motd
      register: mymotd
    
    ```

### Module Shell

- Với module Shell, thì Ansible sẽ mở một shell linux trên máy chủ remote (mặc định là /bin/sh) . Sau đó câu lệnh tham số sẽ được thực thi trên shell remote. Vì vậy các tính năng như chuyển hướng input/output, pipe sẽ hoạt động được.

- Cú pháp cơ bản

    ```
    - name: Execute the command in remote shell
      shell: somescript.sh >> somelog.txt
    
    ```
