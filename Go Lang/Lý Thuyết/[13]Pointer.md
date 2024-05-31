# Con trỏ - Pointer

### Con trỏ là gì?

- Một con trỏ là một biến chứa địa chỉ bộ nhớ của biến khác.

### Khai báo con trỏ

- *T là kiểu biến con trỏ mà trỏ đến một giá trị của kiểu T.

```
package main

import (  
    "fmt"
)

func main() {  
    b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("address of b is", a)
}
```
- Toán tử & được sử dụng để lấy địa chỉ của một biến. Tại dòng 9 của chương trình trên, chúng ta đang gán địa chỉ của biến b cho a có kiểu *int. Như vậy có thể nói a được trỏ đến b.

### Giá trị ZERO của con trỏ

- Giá trị zero của một con trỏ là nil.

```
package main

import (  
    "fmt"
)

func main() {  
    a := 25
    var b *int
    if b == nil {
        fmt.Println("b is", b)
        b = &a
        fmt.Println("b after initialization is", b)
    }
}
```
### Tham chiếu ngược (dereferencing a pointer)

- Tham chiếu ngược một con trỏ nghĩa là truy cập vào giá trị của biến mà con trỏ trỏ tới. *a là cú pháp tham chiếu ngược đến a.

```
package main  
import (  
    "fmt"
)

func main() {  
    b := 255
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
}
```
### Sử dụng con trỏ trong hàm

```
package main

import (  
    "fmt"
)

func change(val *int) {  
    *val = 55
}
func main() {  
    a := 58
    fmt.Println("value of a before function call is",a)
    b := &a
    change(b)
    fmt.Println("value of a after function call is", a)
}
```

### Hãy sử dụng con trỏ cho slice thay vì array trong hàm

- Giả định rằng chúng ta cần sửa đổi một array trong hàm và những thay đổi đó hiển thị với caller. Để làm điều này, ta sử dụng con trỏ cho array như một đối số trong hàm:

```
package main

import (  
    "fmt"
)

func modify(arr *[3]int) {  
    (*arr)[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}
```

- a[x] là cách viết gọn của (*a)[x]. Nên (*a)[0] trong ví dụ trên có thể được thay thế bởi arr[0]. Hãy thử viết lại ví dụ trên bằng cú pháp ngắn gọn:

```
package main

import (  
    "fmt"
)

func modify(arr *[3]int) {  
    arr[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(&a)
    fmt.Println(a)
}
```

- Mặc dù cách này sử dụng con trỏ cho array như một đối số trong hàm và thực hiện thay đổi nó, tuy nhiên vẫn không được coi là phương pháp "chuẩn" trong Go. Chúng ta hãy dùng slice thay vì array.

- Thử viết lại chương trình trên sử dụng slice.

```
package main

import (  
    "fmt"
)

func modify(sls []int) {  
    sls[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(a[:])
    fmt.Println(a)
}
```

### Go không hỗ trợ phép toán số học trên con trỏ

- Go không hỗ trợ phép toán số học trên con trỏ giống nhiều ngôn ngữ lập trình khác như C.