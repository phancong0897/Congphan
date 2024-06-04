# Interfaces - I

### Interface là gì?

- Định nghĩa chung về interface trong thế giới Hướng đối tượng đó là ‘‘Interface xác định hành vi của một đối tượng’’. Nó chỉ xác định những gì đối tượng phải làm. Cách để thực hiện các hành vi này tùy thuộc vào đối tượng đó.

- Trong Go, một interface là tập hợp các khai báo phương thức. Khi một kiểu dữ liệu định nghĩa tất cả các phương thức trong một interface, thì nó được gọi là implement interface đó (triển khai interface). Điều này tương tự như trong thế giới OOP. Interface xác định một kiểu dữ liệu nên có những phương thức nào và cách nó implement những phương thức đó.

- Ví dụ, WashingMachine là một interface chứa các khai báo phương thức như Cleaning() và Drying(). Bất kỳ kiểu dữ liệu nào định nghĩa ra phương thức Cleaning() và Drying() đều được gọi là implement interface WashingMachine.

### Khai báo và triển khai interface

```
package main

import (  
    "fmt"
)

//định nghĩa 1 interface
type VowelsFinder interface {  
    FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {  
    var vowels []rune
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

func main() {  
    name := MyString("Sam Anderson")
    var v VowelsFinder
    v = name // Điều này được chấp nhận vì MyString implement interface VowelsFinder
    fmt.Printf("Vowels are %c", v.FindVowels())
}
```

- Ở chương trình trên ta khởi tạo một interface tên là VowelsFinder có một phương thức là FindVowels() []rune. Khởi tạo một kiểu dữ liệu tên là MyString.

- Sau đó, chúng ta khai báo phương thức FindVowels() []rune cho vật nhận là kiểu MyString. Lúc này, MyString được gọi là implement interface VowelsFinder. Điều này có chút khác biệt so với các ngôn ngữ khác như Java, trong đó một class phải khai báo rõ ràng khi nó implement một interface bằng cách sử dụng từ khóa implements. Điều này không cần thiết trong Go, các interface trong Go ngầm hiểu là được implement nếu có một kiểu dữ liệu chứa tất cả các phương thức được khai báo trong interface đó.

- Trong hàm main() bên trên, chúng ta gán biến name có kiểu MyString cho biến v kiểu VowelsFinder. Điều này được chấp nhận vì MyString implement interface VowelsFinder. Lệnh v.FindVowels() ở dòng tiếp theo gọi đến phương thức FindVowels trên kiểu MyString và in ra tất cả các nguyên âm trong chuỗi Sam Anderson. Chương trình in ra Vowels are [a e o]

- Chúng ta đã khởi tạo và implement thành công một interface!

### Công dụng thực tế của interface

- Ví dụ trên hướng dẫn chúng ta cách tạo và implement một interface, nhưng nó không thực sự cho thấy công dụng thực tế của một interface. Thay vì v.FindVowels() nếu chúng ta sử dụng name.FindVowels() trong chương trình trên, nó vẫn sẽ làm việc và chúng ta không sử dụng gì đến interface được tạo ra.

- Giờ chúng ta sẽ xem xét một ứng dụng thực tế của interface.

- Chúng ta viết một chương trình đơn giản tính tổng chi phí của một công ty dựa trên mức lương của từng nhân viên. Giả định tất cả các chi phí tính bằng USD.

```
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId  int
    basicpay int
}

//tiền lương của nhân viên permanent bằng tổng của basic pay và pf
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//tiền lương của nhân viên contract chỉ là basic pay
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
tổng chi phí được tính bằng cách duyệt qua từng phần tử của slice SalaryCalculator
và tính tổng mức lương của từng nhân viên
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

}
```

- Ở chương trình trên chúng ta khai báo một interface tên là SalaryCalculator có một phương thức là CalculateSalary() int.

- Chúng ta có 2 loại nhân viên trong công ty là Permanent (nhân viên chính thức) và Contract (nhân viên hợp đồng) được định nghĩa bằng kiểu struct. Mức lương của nhân viên Permanent là tổng của basicpay và pf còn đối với nhân viên Contract thì chỉ là basicpay. Điều này được thể hiện trong các phương thức CalculateSalary tương ứng. Bằng cách khai báo phương thức này, cả 2 struct Permanent và Contract đều đang implement interface SalaryCalculator.

- Hàm totalExpense được khai báo bên dưới thể hiện sự tiện ích của việc sử dụng interface. Hàm này nhận một slice các interface SalaryCalculator []SalaryCalculator làm tham số. Trong hàm main() chúng ta truyền một slice với các phần tử gồm cả 2 kiểu Permanent và Contract vào hàm totalExpense. Hàm totalExpense tính toán chi phí bằng cách gọi đến phương thức CalculateSalary của kiểu tương ứng, điều này được thực hiện ở câu lệnh expense = expense + v.CalculateSalary().

- Ưu điểm lớn nhất của hàm totalExpense này là nó có thể được mở rộng đến bất kỳ loại nhân viên mới nào mà không cần phải thay đổi code. Giả sử công ty bổ sung một loại nhân viên mới là Freelancer với cách tính lương khác. Freelancer này chỉ việc truyền vào đối số slice của hàm totalExpense mà không phải thay đổi bất kỳ 1 dòng code nào trong hàm totalExpense. Freelancer cũng implement interface SalaryCalculator.

- Output của chương trình trên là: Total Expense Per Month $14050

### Biểu diễn nội dung của interface

- Một interface có thể được biểu diễn như một bộ (type, value). Type là kiểu dữ liệu cụ thể (concrete type) của interface và value lưu trữ giá trị của kiểu dữ liệu đó.

- Thử viết một chương trình nhỏ để hiểu hơn về điều này.

```
package main

import (  
    "fmt"
)

type Tester interface {  
    Test()
}

type MyFloat float64

func (m MyFloat) Test() {  
    fmt.Println(m)
}

func describe(t Tester) {  
    fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {  
    var t Tester
    f := MyFloat(89.7)
    t = f
    describe(t)
    t.Test()
}
```
- Interface Tester có một phương thức là Test() và kiểu dữ liệu MyFloat implement interface này. Trong hàm main(), chúng ta gán biến f có kiểu MyFloat vào biến t kiểu Tester. Giờ thì kiểu cụ thể của t là MyFloat và giá trị của t là 89.7. Hàm describe in ra giá trị và kiểu cụ thể của interface Tester

### Kiểu interface rỗng

- Một interface không có phương thức nào thì được gọi là interface rỗng. Nó được biểu diễn dưới dạng interface{}. Vì interface rỗng không có thương thức, nên tất cả các kiểu dữ liệu đều có thể implement interface rỗng.

```
package main

import (  
    "fmt"
)

func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    s := "Hello World"
    describe(s)
    i := 55
    describe(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe(strt)
}
```

- Ở chương trình trên, hàm describe(i interface{}) lấy một interface rỗng làm đối số do đó nó có thể truyền vào bất kỳ kiểu dữ liệu nào.

- Chúng ta thử truyền string, int và struct vào hàm describe, chương trình in ra:

```
Type = string, value = Hello World  
Type = int, value = 55  
Type = struct { name string }, value = {Naveen R} 
```

### Type Assertion

- Type assertion được sử dụng để ép kiểu giá trị cơ bản của interface.

- i.(T) là cú pháp để ép kiểu giá trị cơ bản của interface i về một kiểu cụ thể là T.

- Thử viết một chương trình nhỏ về type assertion.

```
package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}
func main() {  
    var s interface{} = 56
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)
}
```

- Khi Steven Paul được truyền vào hàm assert, ok sẽ là false vì kiểu cụ thể của i không phải là int và v nhận giá trị 0 là zero value của kiểu int. Chương trình in ra.

### Sử dụng Switch với kiểu dữ liệu

- Trường hợp switch với biểu thức tùy chọn là kiểu dữ liệu được dùng để so sánh kiểu dữ liệu cụ thể của một interface với nhiều kiểu dữ liệu được chỉ định trong các  trường hơp khác nhau. Nó tương tự như trường hợp switch case.

- Cú pháp cho switch với kiểu dữ liệu cũng tương tự như cú pháp i.(T) trong Type assertion. Kiểu T ở đây chính là các từ khóa của các kiểu dữ liệu mà chúng ta muốn so sánh. Hãy theo dõi chương trình dưới đây.

```
package main

import (  
    "fmt"
)

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func main() {  
    findType("Naveen")
    findType(77)
    findType(89.98)
}
```

- Ở chương trình trên, câu lệnh switch i.(type) để khai báo sử dụng switch với kiểu dữ liệu. Mỗi một case sẽ so sánh kiểu dữ liệu cụ thể của i với một kiểu dữ liệu được chỉ định để so sánh. Nếu case nào phù hợp, thì in ra câu lệnh tương ứng. Output của chương trình là.
```
I am a string and my value is Naveen  
I am an int and my value is 77  
Unknown type  
```

- Trong hàm main(), câu lệnh cuối là findType(89.98), giá trị 89.98 ở đây là kiểu dữ liệu float64 và nó không phù hợp với case nào, vì thế Unknown type được in ra ở dòng cuối cùng.

- Chúng ta cũng có thể so sánh một kiểu dữ liệu với interface. Nếu chúng ta có một kiểu dữ liệu và nếu kiểu đó implement một interface, ta có thể so sánh chính kiểu dữ liệu này với interface mà nó implement.

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

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}

func main() {  
    findType("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)
}
```

- Ở chương trình trên, struct Person implement interface Describer. Trong hàm findType, chúng ta khai báo switch với kiểu dữ liệu, biến v so sánh với kiểu interface Describer. Vì biến p là kiểu Person implement interface Describer do đó khi chương trình chạy đến dòng findType(p), case này đạt yêu cầu và phương thức Describer() được gọi.

```
unknown type  
Naveen R is 25 years old  
```