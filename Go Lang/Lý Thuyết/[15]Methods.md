# Methods - Phương thức

### Phương thức là gì?

- Phương thức (Method) là một hàm (function) được khai báo cho riêng một kiểu dữ liệu đặc biệt, kiểu dữ liệu này được gọi là vật nhận (receiver) nó được đặt giữa từ khóa func và tên của phương thức. Vật nhận này có thể là kiểu struct (cấu trúc) hoặc non-struct (phi cấu trúc). Vật nhận phải có sẵn để truy cập bên trong phương thức.

- Cú pháp để tạo một phương thức như sau:

```
func (t Type) methodName(parameter list) {

}
```

- Câu lệnh trên tạo một phương thức có tên là methodName với vật nhận là Type.

- Chúng ta thử viết một chương trình nhỏ tạo một phương thức trên một kiểu dữ liệu struct.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}

/*
 phương thức displaySalary() có vật nhận là Employee
*/
func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {  
    emp1 := Employee {
        name:     "Sam Adolf",
        salary:   5000,
        currency: "$",
    }
    emp1.displaySalary() //gọi phương thức displaySalary() của kiểu Employee
}
```

- Ở chương trình trên, chúng ta tạo một phương thức tên là displaySalary trên một struct có tên là Employee. Phương thức displaySalary() có quyền truy cập vào vật nhận e Employee bên trong nó. Bên trong phương thức, chúng ta sử dụng vật nhận e để in ra tên, tiền lương của một nhân viên.

- Chúng ta gọi đến phương thức bằng câu lệnh:

    ` emp1.displaySalary() `

### Tại sao phải sử dụng phương thức khi chúng ta đã có hàm?

```
package main

import (  
    "fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}

/*
 phương thức displaySalary() được chuyển thành hàm displaySalary với Employee là tham số truyền vào
*/
func displaySalary(e Employee) {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {  
    emp1 := Employee{
        name:     "Sam Adolf",
        salary:   5000,
        currency: "$",
    }
    displaySalary(emp1)
}
```

- Ở chương trình trên, phương thức displaySalary được chuyển thành hàm displaySalary và struct Employee là tham số truyền vào hàm đó. Chương trình này cũng cho output tương tự như trên Salary of Sam Adolf is $5000.

- Vậy tại sao chúng ta phải dùng phương thức khi mà ta có thể viết một chương trình tương tự với hàm. Có một vài lý do cho việc này.

    - Go không phải là một ngôn ngữ lập trình hướng đối tượng thuần túy và nó không hỗ trợ các lớp (class). Do đó viết các phương thức trên các kiểu dữ liệu là một cách để có được hành vi tương tự như các lớp.
    
    - Các phương thức có thể trùng tên nếu được xác định trên các kiểu dữ liệu khác nhau trong khi các hàm thì không được phép trùng tên. Giả sử chúng ta có 2 struct là Square và Circle. Chúng ta có thể định nghĩa 2 phương thức cùng có tên là Area trên cả Square và Circle. Hãy xem chương trình sau.

```
package main

import (  
    "fmt"
    "math"
)

type Rectangle struct {  
    length int
    width  int
}

type Circle struct {  
    radius float64
}

func (r Rectangle) Area() int {  
    return r.length * r.width
}

func (c Circle) Area() float64 {  
    return math.Pi * c.radius * c.radius
}

func main() {  
    r := Rectangle{
        length: 10,
        width:  5,
    }
    fmt.Printf("Area of rectangle %d\n", r.Area())
    c := Circle{
        radius: 12,
    }
    fmt.Printf("Area of circle %f", c.Area())
}
```

### Vật nhận là con trỏ so với vật nhận là giá trị

- Chúng ta thấy các phương thức trên có vật nhận là giá trị. Ta cũng có thể tạo các phương thức với vật nhận là con trỏ. Sự khác biệt giữa vật nhận là giá trị và vật nhận là con trỏ đó là, những thay đổi được thực hiện bên trong một phương thức đối với vật nhận là con trỏ sẽ thực sự thay đổi vật nhận đó, còn đối với vật nhận là giá trị thì không. Để hiểu rõ hơn hay theo dõi chương trình sau.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    name string
    age  int
}

/*
Phương thức với vật nhận là giá trị 
*/
func (e Employee) changeName(newName string) {  
    e.name = newName
}

/*
Phương thức với vật nhận là con trỏ  
*/
func (e *Employee) changeAge(newAge int) {  
    e.age = newAge
}

func main() {  
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    (&e).changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e.age)
}
```

- Ở chương trình trên, phương thức changeName có vật nhận là giá trị (e Employee) trong khi phương thức changeAge có vật nhận là con trỏ (e *Employee). Sự thay đổi được thực hiện đối với trường name của struct Employee xảy ra bên trong phương thức changeName sẽ không được hiển thị, chương trình in ra cùng một tên trước và sau khi phương thức e.changeName("Michael Andrew") được gọi. Đối với phương thức changeAge hoạt động trên vật nhận là con trỏ (e *Employee), sự thay đổi đối với trường age sau khi gọi phương thức (&e).changeAge(51) lại được hiển thị

- Viết lại chương trình trên và sử dụng e.changeAge(51) thay vì (&e).changeAge(51), ta được output giống nhau.

```
package main

import (  
    "fmt"
)

type Employee struct {  
    name string
    age  int
}

/*
Phương thức với vật nhận là giá trị  
*/
func (e Employee) changeName(newName string) {  
    e.name = newName
}

/*
Phương thức với vật nhận là con trỏ 
*/
func (e *Employee) changeAge(newAge int) {  
    e.age = newAge
}

func main() {  
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    e.changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e.age)
}
```

### Khi nào sử dụng vật nhận là con trỏ và khi nào thì sử dụng vật nhận là giá trị

- Thông thường vật nhận là con trỏ sẽ được sử dụng khi những sự thay đổi đối với vật nhận bên trong phương thức cần được hiển thị cho người gọi.

- Vật nhận là con trỏ cũng có thể được sử dụng ở những nơi mà cần tốn bộ nhớ để sao chép cấu trúc dữ liệu. Giả sử ta có một struct có rất nhiều trường, sử dụng struct này như một vật nhận là giá trị trong một phương thức sẽ cần toàn bộ struct đó được sao chép, như vậy sẽ rất tốn bộ nhớ. Trong trường hợp này, nếu vật nhận là con trỏ thì struct đó sẽ không phải sao chép nữa mà chỉ có con trỏ trỏ đến nó được sử dụng trong phương thức.

- Những trường hợp còn lại thì vật nhận là giá trị đều có thể được sử dụng.

### Phương thức của trường ẩn danh

- Các phương thức thuộc về một trường ẩn danh nằm trong một struct nào đó có thể được gọi như thể chúng là của chính struct đó.

```
package main

import (  
    "fmt"
)

type address struct {  
    city  string
    state string
}

func (a address) fullAddress() {  
    fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {  
    firstName string
    lastName  string
    address
}

func main() {  
    p := person{
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city:  "Los Angeles",
            state: "California",
        },
    }

    p.fullAddress() //truy cập phương thức fullAddress của struct address

}
```

### Vật nhận là giá trị trong một phương thức so với đối số là giá trị trong một hàm

- Chủ đề này khá mới đối với các go newbie. Tôi sẽ cố gắng làm cho nó càng rõ ràng càng tốt.

- Khi một hàm có đối số là giá trị, nó sẽ chỉ chấp nhận đối số là giá trị.

- Khi một phương thức có vật nhận là giá trị, nó sẽ chấp nhận cả vật nhận là con trỏ lẫn vật nhận là giá trị.

```
package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func area(r rectangle) {  
    fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {  
    fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    area(r)
    r.area()

    p := &r
    /*
       lỗi biên dịch, cannot use p (type *rectangle) as type rectangle 
       in argument to area  
    */
    //area(p)

    p.area()//dùng con trỏ p gọi đến phương thức area() có vật nhận là giá trị
}
```

- Ta thấy hàm func area(r rectangle) nhận đối số là giá trị và phương thức func (r rectangle) area() nhận vật nhận là giá trị.

- Chúng ta gọi đến hàm area với đối số là giá trị: area(r) và nó hoạt động bình thường. Tương tự, chúng ta gọi đến phương thức area sử dụng vật nhận là giá trị: r.area() và nó cũng hoạt động bình thường.

- Sau đó, chúng ta khởi tạo 1 con trỏ p trỏ tới r: p := &r. Nếu chúng ta cố truyền con trỏ này vào hàm area (hàm area chỉ chấp nhận giá trị), thì trình biên dịch sẽ báo lỗi. Tôi đã comment dòng code thực hiện điều đó //area(p). Nếu bạn uncomment dòng này, trình biên dịch sẽ ném ra lỗi sau compilation error, cannot use p (type *rectangle) as type rectangle in argument to area.

- Giờ chúng ta sẽ qua đến phần quan trọng, dòng code p.area() gọi đến phương thức area, phương thức area vốn được khai báo nhận vật nhận là giá trị giờ lại sử dụng vật nhận là con trỏ p. Và điều này hoàn toàn hợp lệ. Lý do là dòng code p.area() sẽ được Go ngầm hiểu là (*p).area().

### Vật nhận là con trỏ trong một phương thức so với đối số là con trỏ một trong hàm

- Tương tự như với đối số là giá trị, một hàm có đối số là con trỏ sẽ chỉ chấp nhận con trỏ trong khi phương thức có vật nhận là con trỏ sẽ chấp nhận cả vật nhận là giá trị lẫn vật nhận là con trỏ.

```
package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func perimeter(r *rectangle) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()//dùng giá trị r gọi đến phương thức perimeter() có vật nhận là con trỏ

}
```

- Ở chương trình trên, chúng ta định nghĩa hàm perimeter nhận con trỏ làm đối số func perimeter(r *rectangle) và phương thức perimeter có vật nhận là con trỏ func (r *rectangle) perimeter().

- Sau đó, chúng ta gọi đến hàm perimeter truyền đối số là con trỏ p perimeter(p) và gọi đến phương thức thức perimeter với vật nhận là con trỏ p p.perimeter(). Tất cả đều hoạt động tốt.

- Tiếp theo, chúng ta thử gọi đến hàm perimeter với đối số truyền vào là giá trị r: //perimeter(r). Điều này là không được phép vì một hàm được khai báo nhận đối số là con trỏ sẽ không nhận các đối số là giá trị. Nếu bạn uncomment dòng này và chạy thì chương trình sẽ ném ra lỗi cannot use r (type rectangle) as type *rectangle in argument to perimeter.

- Giờ chúng ta gọi đến phương thức perimeter với vật nhận là giá trị r, dù phương thức perimeter này vốn được khai báo có vật nhận là con trỏ. Điều này hoàn toàn hợp lệ và dòng code r.perimeter() sẽ được Go ngầm hiểu là (&r).perimeter()

### Phương thức đối với kiểu dữ liệu non-struct

- Từ đầu tới giờ, chúng ta đã định nghĩa các phương thức trên kiểu dữ liệu struct (cấu trúc). Chúng ta cũng có thể định nghĩa các phương thức trên kiểu dữ liệu non-struct (phi cấu trúc) tuy nhiên có một vấn đề. Để định nghĩa một phương thức trên một kiểu dữ liệu, thì phương thức và kiểu dữ liệu của vật nhận của phương thức đó phải được định nghĩa trên cùng một package. Ở các phần bên trên, tất cả các struct và phương thức của các struct đó đều được chúng ta định nghĩa trên cùng package main và do đó chúng hoạt động bình thường.

```
package main

func (a int) add(b int) {  
}

func main() {

}
```

- Ở chương trình trên, tôi thử thêm một phương thức có tên là add cho một kiểu dữ liệu được dựng sẵn trong Go là kiểu int. Điều này là không hợp lệ, vì phương thức add và kiểu int không được định nghĩa trên cùng một package. Chương trình ném ra lỗi biên dịch cannot define new methods on non-local type int.

- Cách để giải quyết vấn đề này là tạo một kiểu dữ liệu là bí danh cho kiểu dựng sẵn int, sau đó tạo một phương thức mà kiểu dữ liệu bí danh này chính là vật nhận.

```
package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {  
    return a + b
}

func main() {  
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
```
- Ở chương trình trên, chúng ta đã tạo một kiểu dữ liệu bí danh là myInt cho kiểu int. Sau đó, ta định nghĩa phương thức add với myInt chính là vật nhận.