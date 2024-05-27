# Variadic Functions

### Variadic function là gì?

- Variadic function (Hàm bất định) là hàm chấp nhận một đối số có biến kiểu số (variable number of arguments)

### Cú pháp

- Nếu tham số cuối cùng của một hàm được ký hiệu bằng ...T, hàm đó sẽ chấp nhận mọi đối số chứa số (number) thuộc kiểu T cho tham số cuối.

- Lưu ý rằng chỉ có tham số cuối cùng của hàm mới được coi là bất định

### Một vài ví dụ và giải thích về cách hoạt động của Variadic function

- Bạn đã bao giờ thắc mắc vì sao hàm append lại được dùng để nối dữ liệu vào một slice chứa number trong đối số chưa? Vì nó là một hàm bất định.

` func append(slice []Type, elems ...Type) []Type `

- Phía trên là định nghĩa của hàm append. Trong định nghĩa này elems là một tham số bất định. Do đó hàm append có thể nhận một đối số có biến là number.

Hãy cùng tạo một hàm bất định, chúng ta sẽ viết một chương trình đơn giản để tìm trong một dãy input có tồn tại giá trị interger nào không:

```
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}

```

Trong chương trình trên, func find(num int, nums... int) nhận một đối số có biến chứa số cho tham số nums. Trong hàm find, kiểu dữ liệu của nums tương đương với []int, vd: một slice kiểu interger.

Hàm bất định hoạt động bằng việc chuyển các đối số có biến kiểu số được truyền vào hàm cho một slice mới mang kiểu dữ liệu của tham số bất định. Ví dụ ở dòng trong chương trình trên, đối số của hàm find là 89, 90, 95. Hàm find chứa tham số bất định int. Do đó cả ba tham số này sẽ được bộ biên dịch chuyển đổi và đưa vào trong một slice có kiểu dữ liệu int []int{89, 90, 95} sau đó mới truyền vào hàm find.

Ở dòng  vòng lặp for chạy qua slice nums và in ra vị trí của num trong slice nếu có, nếu vị trí của nums không được in ra nghĩa là không tìm thấy giá trị nào cả.

Tại dòng hàm find chỉ có một đối số, chúng ta không truyền vào bất kỳ đối số nào cho tham số bất định nums ...int, trong trường hợp này là hoàn toàn hợp lệ khi nums là một nil slice với độ dài và sức chứa bằng 0.

### Truyền slice vào một hàm bất định

```
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    nums := []int{89, 90, 95}
    find(89, nums)
}
```

Ở dòng find ta sẽ truyền vào hàm một slice bên trong là một đối số có biến chứa số.

Cách làm này sẽ không có tác dụng, chương trình trên sẽ không hoạt động được và báo lỗi sau: main.go:23: cannot use nums (type []int) as type int in argument to find

Vì sao chương trình này lại không chạy, hãy cùng xem cấu trúc của hàm find dưới đây: 

`func find(num int, nums ...int)`

Theo định nghĩa về hàm bất định, nums ...int nghĩa là: hàm find sẽ chấp nhận mọi đối số có biến kiểu int.

Vẫn ở dòng find nums được truyền vào sẽ là một đối số bất định của hàm find. Như đã nói, những đối số bất định này sẽ được chuyển thành slice thuộc kiểu int cùng lúc hàm find cũng chứa những đối số kiểu int. Trong trường hợp này nums đã là một slice kiểu int rồi, slice []int mới được bộ biên dịch cố gắng tạo ra bằng cách sử dụng nums `find(89, []int{nums})` cũng sẽ không hoạt động được vì nums là []int chứ không phải là int

- Có một cú pháp đặc biệt để truyền một slice vào trong một hàm bất định, hãy thêm hậu tố ... vào sau slice, tiếp đó slice sẽ được truyền vào hàm mà không cần tạo slice mới.

Trong chương trình trên, nếu bạn thay find(89, nums) bằng find(89, nums...) vào dòng find chương trình sẽ được biên dịch

