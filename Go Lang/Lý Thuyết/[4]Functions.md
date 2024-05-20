# Function

### Function là gì?

- Ta có thể hiểu Function là một khuôn chứa các dòng code thực hiện một công việc nhất định. Function tiếp nhận đầu vào, sau đó thực hiện một vài tính toán và trả về kết quả đầu ra.

### Khai báo Function

Cú pháp chung để khai báo một function như sau

```
func functionname(parametername type) returntype {  
 //function body
}
```
- Việc khai báo Function bắt đầu với từ khóa func, tiếp sau đó là functionname. Các tham số được đặt ở giữa, sau cùng là returntype. Cú pháp để xác định một tham số là tên tham số rồi đến loại của tham số đó. Thứ tự sắp xếp của tham số: (parameter1 type, parameter2 type). Sau đó code sẽ được đặt ở giữa "{" và "}" và đây chính là phần thân của một function.

- Các tham số và kiểu trả về (return) được tùy chọn trong một function. Do đó cú pháp bên dưới cũng hợp lệ 

```
func functionname() {  
}
```

### Function mẫu 

```
package main

import (  
    "fmt"
)

func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}
func main() {  
    price, no := 90, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is", totalPrice)
}
```

### Nhiều giá trị trả về

```
package main

import (  
    "fmt"
)

func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

func main() {  
     area, perimeter := rectProps(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f", area, perimeter) 
}
```

### Các giá trị trả về được đặt tên

- Có thể trả về các giá trị đã được đặt tên từ một function. Nếu một giá trị trả về được đặt tên, nó có thể xem như là được khai báo là một biến trong dòng đầu tiên của function. 

- RectProps ở trên có thể viết lại bằng cách sử dụng các giả trị trả về có tên là

```
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}
```

- area và perimeter là các giá trị trả về có tên trong hàm trên. Lưu ý câu lệnh return trong function không trả về bất kì giá trị nào một cách rõ ràng. Khi area và perimeter  được xác định khi khai báo function như các giá trị trả về, chúng sẽ tự động được trả về từ function khi gặp câu lệnh return.

### Định danh trống

- "_" được hiểu là định danh trống trong Go. Nó có thể sử dụng cho bất kì giá trị nào của bất kì loại nào. Chúng ta cùng tìm hiểu việc sử dụng định danh trống này.

- Function rectProps trả về chu vi và diện tích của hình chữ nhật. Nếu chúng ta chỉ cần area và loại bỏ đi perimeter. Đây là nơi "_" sẽ được sử dụng.

- Chương trình bên dưới chỉ có area được trả về từ Function rectProps

```
package main

import (  
    "fmt"
)

func rectProps(length, width float64) (float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
func main() {  
    area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area)
}
```
- Ở dòng thứ 13, chúng ta chỉ có area và "_" được sử dụng để loại bỏ tham số perimeter.
