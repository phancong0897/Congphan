# Structures - Cấu trúc

### Structures là gì?

- Structures (cấu trúc) được người dùng định nghĩa cho một tập hợp các trường. Nó có thể được sử dụng để nhóm dữ liệu vào một cấu trúc thay vì viết từng loại riêng biệt dưới các dạng riêng biệt.

- Ví dụ một nhân viên (Employee) có một firstName, lastName và age. Structures Employee có ý nghĩa để nhóm ba thuộc tính này thành một cấu trúc duy nhất.

- Khai báo cấu trúc (Declaring a structure)

```
type Employee struct {  
    firstName string
    lastName  string
    age       int
}
```

- Đoạn mã trên khai báo một kiểu cấu trúc Employee có các trường firstName, lastName và age. Cấu trúc này cũng có thể được thực hiện nhỏ gọn hơn bằng cách khai báo các trường thuộc cùng một loại trong một dòng duy nhất theo sau là tên kiểu. Trong cấu trúc trên firstName và lastName thuộc về cùng một kiểu string và do đó cấu trúc có thể được viết lại như sau:

```
type Employee struct {  
    firstName, lastName string
    age, salary         int
}
```

- Employee được gọi là tên cấu trúc, nó tạo ra một kiểu mới có tên Employee có thể được sử dụng để tạo ra các cấu trúc kiểu Employee.

- Có thể khai báo cấu trúc mà không khai báo một kiểu mới và kiểu cấu trúc này được gọi là cấu trúc ẩn danh.

```
var employee struct {  
        firstName, lastName string
        age int
}
```

### Tên cấu trúc (Creating named structures)

- Ví dụ một cấu trúc được đặt tên Employee sử dụng trong một chương trình đơn giản.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {

    //creating structure using field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}
```

- Trong chương trình trên, chúng ta tạo một struct Employee. Cấu trúc emp1 được xác định bằng cách xác định giá trị cho mỗi tên trường. Không cần thiết rằng thứ tự của các tên trường phải giống như thứ tự trong khi khai báo kiểu cấu trúc. Ở đây chúng tôi đã thay đổi vị trí của lastName và di chuyển nó đến cuối cùng, emp2 được định nghĩa bằng cách bỏ qua các tên trường. Trong trường hợp này, cần duy trì thứ tự các trường giống như được chỉ định trong khai báo cấu trúc.

### Cấu trúc ẩn danh (Creating anonymous structures​)

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    var emp4 Employee //zero valued structure
    fmt.Println("Employee 4", emp4)
}
```
- Chương trình trên emp4 được định nghĩa nhưng nó không được khởi tạo với bất kỳ giá trị nào. Do đó firstName và lastName được gán giá trị rỗng "", age, salary được gán giá trị bằng 0.

### Cấu trúc 0 (Zero value of a structure)

- Cấu trúc ẩn danh cũng có thể xác định giá trị cho một số trường và bỏ qua phần còn lại. Trong trường hợp này, các tên trường bị bỏ qua và được gán giá trị bằng không.

```package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp5 := Employee{
        firstName: "John",
        lastName:  "Paul",
    }
    fmt.Println("Employee 5", emp5)
}
```
- Trong chương trình trên firstName và lastName được khởi tạo trong khi age và salary thì không. Do đó age và salary được gán giá trị bằng không.


### Truy cập các trường riêng lẻ của một cấu trúc (Accessing individual fields of a struct)

- Dấu chấm . toán tử được sử dụng để truy cập vào các trường riêng lẻ của một cấu trúc.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d", emp6.salary)
}
```

### Trỏ đến một cấu trúc (Pointers to a struct)

- Structs cũng có thể tạo ra con trỏ trỏ đến một cấu trúc.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", (*emp8).firstName)
    fmt.Println("Age:", (*emp8).age)
}
```

- Ngôn ngữ Go cho phép chúng ta tùy chọn sử dụng emp8.firstName thay vì (* emp8) .firstName để truy cập trường firstName.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp8 := &Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp8.firstName)
    fmt.Println("Age:", emp8.age)
}
```

### Trường ẩn danh (Anonymous fields)

- Có thể tạo các structs với các trường chỉ chứa một loại không có tên trường. Các loại trường này được gọi là các trường ẩn danh.

- Đoạn mã dưới đây tạo ra một cấu trúc Person có hai trường vô danh là string và int
```
type Person struct {  
    string
    int
}
```
```
package main

import (  
    "fmt"
)

type Person struct {  
    string
    int
}

func main() {  
    p := Person{"Naveen", 50}
    fmt.Println(p)
}
```

- Trong chương trình trên, Person là một struct với hai trường ẩn danh. p: = Person {"Naveen", 50} định nghĩa một biến kiểu Person. 

- Mặc dù một trường ẩn danh không có tên, theo mặc định, tên của trường ẩn danh là tên loại của nó. Ví dụ trong trường hợp của structs Person ở trên, mặc dù các trường là vô danh, theo mặc định chúng lấy tên của kiểu trường. Vì vậy, structs Person có 2 trường có tên string và int.

```
package main

import (  
    "fmt"
)

type Person struct {  
    string
    int
}

func main() {  
    var p1 Person
    p1.string = "naveen"
    p1.int = 50
    fmt.Println(p1)
}
```
### Cấu trúc lồng nhau (Nested structs​)

- Có thể là một cấu trúc chứa một trường mà lần lượt là một cấu trúc. Những loại cấu trúc này được gọi là cấu trúc lồng nhau.

```
package main

import (  
    "fmt"
)

type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age int
    address Address
}

func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.address = Address {
        city: "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:",p.age)
    fmt.Println("City:",p.address.city)
    fmt.Println("State:",p.address.state)
}
```

### Các trường được khuyến khích (Promoted fields​)

```
package main

import (  
    "fmt"
)

type Address struct {  
    city, state string
}
type Person struct {  
    name string
    age  int
    Address
}

func main() {  
    var p Person
    p.name = "Naveen"
    p.age = 50
    p.Address = Address{
        city:  "Chicago",
        state: "Illinois",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city) //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field
}
```

### Xuất cấu trúc (Exported Structs and Fields​)

- Nếu một kiểu cấu trúc bắt đầu bằng một chữ hoa, thì đó là một loại được xuất và nó có thể được truy cập từ các pakage khác. Tương tự, nếu các trường của cấu trúc bắt đầu bằng caps, chúng cũng có thể được truy cập từ các gói khác.

- Chúng ta tạo một thư mục có tên structs bên trong thư mục src của vùng làm việc với Go của bạn. Tạo một thư mục computer khác bên trong các cấu trúc.

- Cấu trúc gói trông giống như sau:

```
src  
   structs
          computer
                  spec.go
          main.go
```
- Bên trong thư mục computer, lưu chương trình bên dưới với tên tập tin spec.go
```
package computer

type Spec struct { //exported struct  
    Maker string //exported field
    model string //unexported field
    Price int //exported field
}
```

- Đoạn mã trên tạo ra một pakage computer có chứa một kiểu cấu trúc Spec đã xuất với hai trường Maker và Price và trường model chưa được xuất. Cho phép nhập pakage này từ pakage main và sử dụng cấu trúc Spec.

- Tạo một tệp có tên main.go bên trong thư mục structs và viết chương trình sau bên trong main.go

```
package main

import "structs/computer"  
import "fmt"

func main() {  
    var spec computer.Spec
    spec.Maker = "apple"
    spec.Price = 50000
    fmt.Println("Spec:", spec)
}
```

- Trong chương trình trên, chúng ta import "structs/computer" và truy cập vào hai trường xuất Maker và Price của struct Spec. Chương trình này có thể được chạy bằng cách thực hiện các lệnh go install structs.

### Cấu trúc bằng nhau (Structs Equality)

- Structs là các kiểu giá trị và có thể so sánh được nếu mỗi trường của chúng có thể so sánh được. Hai biến cấu trúc được coi là bằng nhau nếu các trường tương ứng của chúng bằng nhau.

```
package main

import (  
    "fmt"
)

type name struct {  
    firstName string
    lastName string
}


func main() {  
    name1 := name{"Steve", "Jobs"}
    name2 := name{"Steve", "Jobs"}
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{firstName:"Steve", lastName:"Jobs"}
    name4 := name{}
    name4.firstName = "Steve"
    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}
```

- Trong chương trình trên, struct name chứa hai trường string. Vì các string có thể so sánh được, nên có thể so sánh hai biến cấu trúc của tên kiểu, name1 và name2 là như nhau trong khi name3 và name4 thì không.

- Các biến cấu trúc không thể so sánh nếu chúng chứa các trường không thể so sánh được.

```
package main

import (  
    "fmt"
)

type image struct {  
    data map[int]int
}

func main() {  
    image1 := image{data: map[int]int{
        0: 155,
    }}
    image2 := image{data: map[int]int{
        0: 155,
    }}
    if image1 == image2 {
        fmt.Println("image1 and image2 are equal")
    }
}
```

- Trong chương trình trên, loại struct image chứa một trường có kiểu dữ liệu map, map không thể so sánh được, do đó không thể so sánh được image1 và image2. Nếu bạn chạy chương trình này, trình biên dịch sẽ lỗi như kết quả trên.