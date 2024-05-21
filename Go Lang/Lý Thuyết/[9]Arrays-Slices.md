# Arrays and Slices - Mảng avf slice

### Mảng

- Mảng là tập hợp của các phần tử có cùng một loại. Ví dụ: tập hợp các số nguyên 5, 8, 9, 79, 76 tạo thành một mảng. Nếu các phần tử trong mảng có type (loại) khác nhau Go sẽ không thực thi được, ví dụ một mảng chứa cả chuỗi và số nguyên sẽ không được coi là hợp lệ trong Go.

- Khai báo

- Một mảng thuộc loại n[T]. n sẽ biểu thị số phần tử có trong mảng và T biểu thị loại của mỗi phần tử. Số lượng phần tử n cũng được coi là một phần của type (Chúng ta sẽ thảo luận chi tiết hơn trong phần này).

- Có nhiều cách khác nhau để khai báo mảng. Hãy xem từng phần một.

```
package main

import (  
    "fmt"
)


func main() {  
    var a [3]int //int array with length 3
    a[0] = 12 // array index starts at 0
    a[1] = 78
    a[2] = 50
    fmt.Println(a)
```

- Chúng ta cùng tạo một mảng bàng cách sử dụng short hand declaration (khai báo ngắn).
```
package main 

import (  
    "fmt"
)

func main() {  
    a := [3]int{12, 78, 50} // short hand declaration to create array
    fmt.Println(a)
}
```

### Mảng là một kiểu giá trị

- Mảng trong Go là các loại giá trị và không phải là loại tham chiếu. Điều này có nghĩa là khi chúng được gán cho một biến mới, một bản copy của bản gốc sẽ được gán cho biến mới này. Nếu thay đổi được thực hiện trên biến mới, nó sẽ không xét lại trong mảng ban đầu.
```
package main

import "fmt"

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}
```

- Dòng 7 trong chương trình trên, một bản sao của a được gán cho b. Ở dòng 8, phần tử đầu tiên của b được đổi thành Singapore. Điều này sẽ không xét lại trong mảng ban đầu a. Chương trình vẫn sẽ output. 

- Tương tự, khi mảng được truyền cho các hàm như các tham số, chúng được truyền theo giá trị và mảng ban đầu không thay đổi.

```
package main

import "fmt"

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func main() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}
```

### Độ dài của mảng

- Độ dài của mảng được xác định bằng cách truyền mảng như một tham số tới function len 

```
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of a is",len(a))
}
```

### Vòng lặp mảng sử dụng range (phạm vi)

- Vòng lặp for có thể sử dụng để lặp qua các phần tử trong một mảng

```
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
```
### Mảng đa chiều

```
package main

import (  
    "fmt"
)

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {  
    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
}
```

### Tạo một slice 

```
package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)
}
```

- Có một cách khác để tạo một slide.
```
package main

import (  
    "fmt"
)

func main() {  
    c := []int{6, 7, 8} //creates and array and returns a slice reference
    fmt.Println(c)
}
```
