# Variables - Biến

### Biến là gì

- Biến là tên gọi của một vùng nhớ nào đó dùng để lưu trữ giá trị của một kiểu dữ liệu cụ thể. Có nhiều cú pháp khác nhau để khai báo biến trong go.

### Khai báo một biến

- Cú pháp: `var name type`

```
package main

import "fmt"

func main() {  
    var age int // khai báo biến
    fmt.Println("my age is", age)
}

```

- Câu lệnh var age int khai báo một biến có tên là age kiểu int. Ta chưa gán giá trị cho biến này. Nếu một biến không được gán giá trị, go sẽ tự động khởi tạo giá trị zero tương ứng với kiểu của biến. Trong trường hợp này age được gán giá trị là 0.

- Một biến có thể được gán bất kỳ giá trị nào cùng kiểu với nó. Ở ví dụ trên, ta có thể gán bất kỳ giá trị nguyên nào cho biến age.

```

package main

import "fmt"

func main() {  
    var age int // khai báo biến
    fmt.Println("my age is ", age)
    age = 27 //gán giá trị
    fmt.Println("my age is", age)
    age = 56 //gán giá trị
    fmt.Println("my new age is", age)
}

```

### Khai báo biến có giá trị khởi tạo

- Chúng ta có thể khởi tạo giá trị khi khai báo biến.

- Cú pháp: `var name type = initialvalue`

```
package main

import "fmt"

func main() {  
    var age int = 27 // khai báo biến có giá trị khởi tạo
    fmt.Println("my age is", age)
}

```

### Type inference (Suy luận kiểu)

- Nếu một biến có khai báo giá trị khởi tạo, Go có thể tự động suy ra kiểu của biến này từ giá trị khởi tạo đó. Do đó nếu một biến có giá trị khởi tạo, ta có thể bỏ qua việc khai báo kiểu của biến.

- Nếu biến được khai báo bằng cách sử dụng cú pháp var name = initialvalue, Go sẽ tự động suy ra kiểu biến đó từ giá trị khởi tạo.

### Khai báo nhiều biến

- Ta có thể khai báo nhiều biến một lúc trên cùng một câu lệnh.

- Cú pháp: ` var name1, name2 type = initialvalue1, initialvalue2 `

```
package main

import "fmt"

func main() {  
    var width, height = 100, 50 //"int" được bỏ qua
    fmt.Println("width is", width, "height is", height)
}
```

### Short hand declaration (Khai báo nhanh)

- Bên cạnh các cách khai báo biến như trên, Go cung cấp một cách khai báo khác nhanh và ngắn gọn hơn, có thể gọi là khai báo nhanh, bằng cách sử dụng toán tử :=

- Cú pháp: ` name := initialvalue `

```
package main

import "fmt"

func main() {  
    name, age := "naveen", 29 // khai báo nhanh
    fmt.Println("my name is", name, "age is", age)
}
```

- Cách khai báo nhanh được sử dụng khi có ít nhất một biến mới được khai báo phía bên trái của toán tử :=. Theo dõi ví dụ dưới đây:

```
package main

import "fmt"

func main() {  
    a, b := 20, 30 // khai báo nhanh biến a và b
    fmt.Println("a is", a, "b is", b)
    b, c := 40, 50 // biến b đã được khai báo rồi còn biến c giờ mới được khai báo
    fmt.Println("b is", b, "c is", c)
    b, c = 80, 90 // gán giá trị mới cho 2 biến đã được khai báo b và c
    fmt.Println("changed b is", b, "c is", c)
}
```