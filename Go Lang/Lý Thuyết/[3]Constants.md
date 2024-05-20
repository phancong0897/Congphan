# Constants - Hằng Số

### Định nghĩa

- Khái niệm constant (hằng số) được sử dụng trong Go để thể hiện những giá trị cố định như 5, -89, "I love Go", 67.89,...

- Hằng số tức là những giá trị không thể thay đổi, do vậy đoạn code dưới đây sẽ dưới đây sẽ báo lỗi không thể gán lại giá trị cho a.

```
package main

func main() {
    const a = 55 
    a = 89 //cannot assign to a
}
```

- Giá trị của 1 hằng số cần được biết ngay tại thời gian biên dịch. Do vậy nó không thể được gán 1 giá trị được trả bởi 1 function, bởi vì function chỉ trả về kết quả khi chạy chương trình.

```
package main

import (  
    "fmt"
    "math"
)

func main() {  
    fmt.Println("Hello, playground")
    var a = math.Sqrt(4) //chạy bình thường với từ khóa var
    const b = math.Sqrt(4)//lỗi , const initializer math.Sqrt(4) is not a constant
}
```

- Ở chương trình trên, a chỉ là 1 biến bình thường được khai báo bằng từ khóa var, do vậy nó có thể được gán giá trị trả về của function math.Sqrt(4).

- Còn b là 1 hằng số và giá trị của nó cần được tính ngay tại thời gian biên dịch. Nhưng funcion math.Sqrt(4) chỉ chạy và trả về kết quả khi chạy chương trình, do vậy đoạn code trên báo lỗi.

### String constant (Hằng số chuỗi)

- Mọi giá trị nằm trong cặp dấu ngoặc kép đều là hằng số chuỗi trong Go. Ví dụ, các chuỗi như "Hello World" hay "Golang" đều là hằng số trong Go.

- Vậy kiểu của 1 hằng số chuỗi là gì? Câu trả lời là untyped (không có kiểu).

- 1 hằng số chuỗi như "Hello World" không có kiểu gì.

### Hằng số boolean

Hằng số boolean có 2 giá trị là true và false. Những quy tắc của hằng số chuỗi cũng được áp dụng cho boolean.

```
package main

func main() {  
    const trueConst = true
    type myBool bool
    var defaultBool = trueConst //chạy bình thường
    var customBool myBool = trueConst //chạy bình thường
    defaultBool = customBool //lỗi do 2 biến có kiểu khác nhau
}
```
### Hằng số Số (numeric constant)

- Numeric constant bao gồm hằng số tự nhiên (integer), thập phân (float) và các hằng số phức hợp (complex). Ngoài ra còn có thêm 1 số biến thể khác nữa.

### Biểu thức số học

- Numeric constant có thể được sử dụng thoaỉ mái trong các biểu thức, và kiểu của nó chỉ cần khi nó được gán cho 1 biến khác, hoặc sử dụng ở những đoạn code yêu cầu trả về kiểu dữ liệu của nó.

```
package main

import (  
    "fmt"
)

func main() {  
    var a = 5.9/8
    fmt.Printf("a's type %T value %v",a, a)
}
```

- Ở chương trình trên, 5.9 là 1 số thập phân và 8 là 1 số nguyên. Tuy nhiên biểu thức 5.9/8 vẫn chạy được bình thường. Kết quả của biểu thức là 0.7375 là 1 số thập phân, do vậy biến a có kiểu là float. Chương trình sẽ trả về kết quả là a's type float64 value 0.7375.