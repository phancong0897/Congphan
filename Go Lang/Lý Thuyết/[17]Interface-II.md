# Interfaces – II

### Triển khai Interfaces sử dụng con trỏ và giá trị

- Tất cả các Interfaces ví dụ mà chúng ta đã thảo luận trong phần 1 đã được thực hiện bằng cách sử dụng giá trị. Nó cũng có thể thực hiện các giao diện bằng cách sử dụng con trỏ. Có một sự tinh tế cần lưu ý trong khi thực hiện các Interfaces bằng cách sử dụng con trỏ. Hãy tìm hiểu rằng nó trong chương trình sau đây:

```
package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() { //implemented using value receiver  
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {  
    state   string
    country string
}

func (a *Address) Describe() { //implemented using pointer receiver  
    fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {  
    var d1 Describer
    p1 := Person{"Sam", 25}
    d1 = p1
    d1.Describe()
    p2 := Person{"James", 32}
    d1 = &p2
    d1.Describe()

    var d2 Describer
    a := Address{"Washington", "USA"}

    /* compilation error if the following line is
       uncommented
       cannot use a (type Address) as type Describer
       in assignment: Address does not implement
       Describer (Describe method has pointer
       receiver)
    */
    //d2 = a

    d2 = &a //This works since Describer interface
    //is implemented by Address pointer in line 22
    d2.Describe()

}
```
- Trong chương trình trên, cấu trúc Person triển khai Interfaces Describer bằng cách sử dụng giá trị.

- Như chúng ta đã học được trong quá trình thảo luận về các phương thức, các phương thức với bộ nhận giá trị chấp nhận cả con trỏ và giá trị.

- p1 là một giá trị của kiểu Person và nó được gán cho d1, tương tự, d1 được gán cho &p2.

- Kết quả:
```
Sam is 25 years old  
James is 32 years old  
State Washington Country USA  
```

### Triển khai nhiều Interfaces

```
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    DisplaySalary()
}

type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}

func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var s SalaryCalculator = e
    s.DisplaySalary()
    var l LeaveCalculator = e
    fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}
```
- Chương trình trên có hai Interfaces là SalaryCalculator và LeaveCalculator được khai báo.

- Cấu trúc Employee được định nghĩa cung cấp các triển khai cho phương thức DisplaySalary của Interfaces SalaryCalculator và phương thức CalculateLeavesLeft của Interfaces LeaveCalculator.

- Chúng ta gán e cho một biến kiểu Interfaces SalaryCalculator, sau đó gán cùng một biến e cho một biến kiểu LeaveCalculator. Điều này là có thể vì e của loại Employee thực hiện cả hai Interfaces SalaryCalculator và LeaveCalculator.

- Kết quả:
```
Naveen Ramanathan has salary $5200
Leaves left = 25
```

### Nhúng  Interfaces

```
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    DisplaySalary()
}

type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

type EmployeeOperations interface {  
    SalaryCalculator
    LeaveCalculator
}

type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}

func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var empOp EmployeeOperations = e
    empOp.DisplaySalary()
    fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}
```

- Interfaces EmployeeOperations của chương trình trên được tạo ra bằng cách nhúng các Interfaces SalaryCalculator và LeaveCalculator.

- Bất kỳ loại nào được cho là để triển khai Interfaces EmployeeOperations nếu nó cung cấp các định nghĩa phương thức cho các phương thức có trong cả hai Interfaces SalaryCalculator và LeaveCalculator.

- Employee struct triển khai thực hiện giao diện EmployeeOperations vì nó cung cấp định nghĩa cho cả hai phương thức DisplaySalary và CalculateLeavesLeft.

- Biến e của kiểu Employee được gán cho empOp của kiểu EmployeeOperations. Tiếp theo, các phương thức DisplaySalary () và CalculateLeavesLeft () được gọi trên empOp.

- Kết quả:
```
Naveen Ramanathan has salary $5200  
Leaves left = 25 
```

### Giá trị 0 của Interfaces

- Giá trị bằng không của một Interfaces là  nil. Một Interfaces nil có cả giá trị cơ bản và cụ thể là nil.

```
package main

import "fmt"

type Describer interface {  
    Describe()
}

func main() {  
    var d1 Describer
    if d1 == nil {
        fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
    }
}
```

- Nếu bạn cố gắng thử phương thức nil trong Interfaces, chương trình sẽ lỗi vì nil trong Interfaces không có giá trị cơ bản cũng như một loại cụ thể.