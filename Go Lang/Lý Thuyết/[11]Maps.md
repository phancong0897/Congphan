# Maps

### Map là gì?
Map (bản đồ) là một kiểu dữ liệu được dựng sẵn trong Go, một map là tập hợp các cặp key/value (khóa/giá trị) trong đó một value được liên kết với một key. Value chỉ được truy xuất bởi key tương ứng.

### Khởi tạo map
Một map được khởi tạo bằng cách truyền kiểu dữ liệu của key và value thông qua hàm make. Cú pháp như sau: make(map[type of key]type of value)

`personSalary := make(map[string]int)`

### Thêm phần tử vào map
Cú pháp để thêm phần tử mới vào map cũng tương tự như với array. Chương trình dưới đây thêm một số phần tử mới vào map personSalary.

```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := make(map[string]int)
    personSalary["steve"] = 12000
    personSalary["jamie"] = 15000
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}
```
- Ta cũng có thể vừa khởi tạo map kèm thêm khai báo phần tử vào luôn:

```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}
```
### Truy cập các phần tử của map

- Ta đã thêm một số phần tử vào map, giờ hãy tìm hiểu cách truy xuất chúng. Cú pháp để truy xuất phần tử của map: map[key]

```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    employee := "jamie"
    fmt.Println("Salary of", employee, "is", personSalary[employee])
}
```
- Nếu một phần tử không tồn tại thì sao? Map sẽ trả về giá trị 0 tương ứng với kiểu dữ liệu của phần tử đó. Trong trường hợp map personSalary này, nếu ta truy cập đến một phẩn tử không tồn tại, thì giá trị 0 của kiểu int là 0 sẽ được trả về.

```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    employee := "jamie"
    fmt.Println("Salary of", employee, "is", personSalary[employee])
    fmt.Println("Salary of joe is", personSalary["joe"])
}
```
- Nếu chúng ta muốn biết một key có tốn tại trong map hay không thì phải làm thế nào?  ` value, ok := map[key] `
```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    newEmp := "joe"
    value, ok := personSalary[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
    } else {
        fmt.Println(newEmp,"not found")
    }

}
```
- Dạng range của vòng lặp for được sử dụng để duyệt qua tất cả các phần tử của map.
```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("All items of a map")
    for key, value := range personSalary {
        fmt.Printf("personSalary[%s] = %d\n", key, value)
    }

}
```
### Xóa phần tử

- `delete(map, key) ` là cú pháp để xóa 1 key khỏi map. Hàm delete này không trả về bất kỳ giá trị nào.
```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("map before deletion", personSalary)
    delete(personSalary, "steve")
    fmt.Println("map after deletion", personSalary)

}
```
### Chiều dài của map

- Chiều dài của một map có thể được xác định bằng hàm len.
```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("length is", len(personSalary))

}
```

### Map là kiểu tham chiếu

- Tương tự như slice, map là kiểu tham chiếu. Khi một map được gán cho một biến mới, cả hai đều trỏ đến cùng một cấu trúc dữ liệu nội bộ. Do đó những thay đổi được thực hiện ở một biến sẽ ánh xạ đến biến kia.

```
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("Original person salary", personSalary)
    newPersonSalary := personSalary
    newPersonSalary["mike"] = 18000
    fmt.Println("Person salary changed", personSalary)

}
```