# If - Else

### if là câu lệnh điều kiện
- Cú pháp của nó như sau:
```
if condition {  
}
```
- Nếu điều kiện là đúng (true), đoạn mã bên trong cặp dấu ngoặc móc { } sẽ được thực thi.

- Không giống những ngôn ngữ khác như C, cặp dấu { } là bắt buộc ngay cả khi chỉ có một đoạn mã bên trong { }.

- Câu lệnh if có thể có tùy chọn thêm các thành phần else if và else.

```
if condition {  
} else if condition {
} else {
}
```
- Trong khối lệnh có thể có một hoặc nhiều nhánh else if được sử dụng. Câu lệnh điều kiện sẽ thực thi từ trên xuống dưới đến khi thỏa mãn điều kiện đúng. Có nghĩa là, khi điều kiện trong câu lệnh if hay else if là đúng, khối lệnh tương ứng sẽ được thực thi, trường hợp không có điều kiện nào đúng, khối lệnh trong else sẽ được thực thi.

- Có một biến thể khác của lệnh if, trong đó có thành phần câu lệnh (tùy chọn) được thực thi trước khi kiểm tra điều kiện. Cú pháp của nó như sau:
```
if statement; condition {  
}
```
```
package main

import (  
   "fmt"
)

func main() {  
   if num := 10; num % 2 == 0 { // Kiểm tra xem có phải số chẵn hay không
       fmt.Println(num,"is even")
   } else {
       fmt.Println(num,"is odd")
   }
}
```
- Trong ví dụ trên, biến num được khởi tạo trong câu lệnh if. Một điều cần lưu ý là biến num chỉ truy cập được bên trong câu lệnh if và else. Cụ thể, phạm vi của biến num bị giới hạn trong khối lệnh if else. Nếu ta cố truy cập từ bên ngoài khối lệnh if else, trình biên dịch sẽ báo lỗi.

- Câu lệnh else bắt buộc đứng ngay sau dấu móc ngoặc đóng } của lệnh điều kiện, nếu không, trình biên dịch sẽ báo lỗi.

- Theo nguyên tắc này, dấu ; sẽ được chèn ngay sau dấu }, nếu nó nằm ở cuối dòng. Do vậy dấu ; sẽ tự động được chèn ngay sau dấu } của lệnh điều kiện if.