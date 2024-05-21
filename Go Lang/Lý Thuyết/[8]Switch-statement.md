# Cấu trúc rẽ nhánh Switch ... case

- Một switch là một câu lệnh điều kiện đánh giá một biểu thức và so sánh nó với một danh sách các kết quả phù hợp và thực thi các khối mã.

- Cú pháp:
```
switch (expression)
{
    case constant_1:
    {
        Statements;
        break;
    }
    case constant_2:
    {
        Statements;
        break;
    }
    .
    .
    .
    case constant_n:
    {
        Statements;
        break;
    }
    default:
    {
        Statements;
    }
}
```

### Default

Default sẽ được thực thi khi không có case nào đúng.

### Mutilple expressions in case

- Trong switch có thể bao gồm nhiều biểu thức trong một case bằng việc tách chúng bởi dấu phẩy.

```
package main

import (
	"fmt"
)

func main() {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("Nguyên âm")
	default:
		fmt.Println("Không phải một nguyên âm")
	}
}
```
- Chương trình trên kiểm tra xem letter có phải là một nguyên âm hay không và các trường hợp của nguyên âm được viết trong cùng một case và cách nhau bởi dấu phẩy.

### Expressionless Switch

- Biểu thức trong Switch là tuỳ chọn và nó có thể được bỏ qua. Nếu nó được bỏ qua thì switch được coi là đúng và mỗi biểu thức trong case được đánh giá và thực thi câu lệnh tương ứng.

```
package main

import (
	"fmt"
)

func main() {
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("num lớn hơn 0 và nhỏ hơn 50")
	case num >= 51 && num <= 100:
		fmt.Println("num lớn hơn 51 và nhỏ hơn 100")
	case num >= 101:
		fmt.Println("num lớn hơn 100")
	}

}
```

- Trong chương trình trên, biểu thức không có trong switch và do đó nó được coi là đúng, mỗi biểu thức trong case đều được đánh giá. Loại này của switch có thể được coi là một thay thế cho nhiều mệnh đề if ... else.

### Fallthrough

- Trong Go, việc đi ra khỏi câu lệnh switch là ngay lập tức sau khi một case đúng được thực thi. Một lệnh fallthrough được sử dụng để chuyển quyền kiểm soát đến câu lệnh của case ngay sau nó. Điều này sẽ giúp tất cả các case đều được kiểm tra.

```
package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {

	switch num := number(); { //num không phải là một hằng số
	case num < 50:
		fmt.Printf("%d thì nhỏ hơn 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d thì nhỏ hơn 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d thì nhỏ hơn 200", num)
	}

}
```

- Biểu thức switch và case không chỉ là hằng số. Trong chương trình trên, num được khởi tạo bằng giá trị trả về của hàm number(). 

- Fallthrough phải là lệnh cuối cùng trong case, nếu đặt đâu đó ở giữa trình biên dịch sẽ báo lỗi. 

- Đây là phần cuối cùng trong bài hướng dẫn này. Có một loại khác của switch là type switch, điều này sẽ được tìm hiểu khi chúng ta nghiên cứu về interfaces. 