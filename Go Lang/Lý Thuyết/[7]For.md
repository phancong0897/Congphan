# For - Vòng Lặp

- for là vòng lặp duy nhất có sẵn trong Go. Go không có vòng lặp while hay do ... while như ở các ngôn ngữ khác.

```
for initialisation; condition; post {  
}
```
- Câu lệnh khởi tạo initialisation sẽ chỉ được thực hiện một lần. Sau khi khởi tạo, điều kiện condition sẽ được kiểm tra, nếu điều kiện là true thì phần thân của vòng lặp trong dấu {} sẽ được thực thi. Post ở đây được hiểu như biến chạy, biến chạy sẽ được thực hiện sau mỗi lần lặp thành công của vòng lặp. Sau khi câu lệnh được thực hiện, điều kiện được kiểm tra lại, nếu điều kiện vẫn đúng, vòng lặp sẽ tiếp tục được thực hiện, nếu điều kiện sai, vòng lặp for sẽ kết thúc.

### break

- Lệnh break được sử dụng để dừng vòng lặp đột ngột và chuyển đến câu lệnh ngay sau vòng lặp for trước khi vòng lặp kết thúc theo bình thường.

```
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i > 5 {
            break //vòng lặp dừng nếu i > 5
        }
        fmt.Printf("%d ", i)
    }
    fmt.Printf("\n Câu lệnh ngay sau vòng lặp")
}
```

- Trong chương trình trên, giá trị của i được kiểm tra suốt mỗi lần lặp. Nếu i > 5 thì lệnh break được thực thi và vòng lặp dừng lại, câu lệnh ngay sau vòng lặp được chạy.

### Continue

- Lệnh continue được sử dụng để bỏ qua lệnh lặp hiện tại của vòng lặp for. Tất cả đoạn mã trong vòng lặp for sau khi gặp lệnh continue thì sẽ không được thực hiện cho phép lặp hiện tại. Vòng lặp sẽ chuyển sang bước lặp tiếp theo.

```
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%d ", i)
    }
}
```

- Trong chương trình trên if i%2 == 0 là kiểm tra phần dư của i chia 2 có bằng 0 hay không. Nếu bằng 0 thì lệnh continue được thực thi và điều khiển chuyển sang lần lặp tiếp theo của vòng lặp. Do đó, lệnh sau continue sẽ không được gọi và vòng lặp tiếp tục đến lần lặp tiếp theo. Nếu kiểm tra i%2 khác 0 thì lệnh in giá trị của i được thực thi.

### Biến thể của vòng lặp

```
package main

import (  
    "fmt"
)

func main() {  
    i := 0
    for ;i <= 10; { // biến khởi tạo và biến chạy có thể để trống
        fmt.Printf("%d ", i)
        i += 2
    }
}
```
- Như chúng ta đã biết, 3 thành phần của vòng lặp for là biến khởi tạo, điều kiện và biến chạy. Trong chương trình trên, biến khởi tạo và biến chạy được bỏ qua, i được khởi tạo ngoài vòng lặp for. Vòng lặp sẽ được thực hiện với i <= 10, i được tăng với bước nhảy là 2 bên trong vòng lặp for. 

- Ngoài ra, các dấu chấm phẩy của vòng lặp for trong chương trình trên cũng có thể được bỏ qua. Định dạng này có thể được coi là một thay thế cho vòng lặp while, khi bỏ dấu chấm phẩy, chương trình được viết lại như sau và kết quả không thay đổi.

### Vòng lặp vô hạn

- Cú pháp tạo một vòng lặp vô hạn là:

```
for {
}
```
