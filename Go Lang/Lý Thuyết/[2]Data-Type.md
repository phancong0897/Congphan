# Data Type - Kiểu dữ liệu

- Các kiểu dữ liệu có sẵn trong Go:
    
    - Kiểu bool 

    - Kiểu chuỗi (string)

    - Kiểu dữ liệu số (numeric type) 

​        - int8, int16, int32, int64, int

        - uint8, uint16, uint32, uint64, uint
        
        - float32, float64
        
        - complex64, complex128
        
        - byte
        
        - rune


### bool

- Bool là kiểu dữ liệu chỉ nhận 2 giá trị hoặc true hoặc false (hoặc đúng hoặc sai).


```
package main

import "fmt"

func main() {
	a := true //a được gán bằng true
	b := false // b được gán bằng false
	fmt.Println("a:", a, "b:", b) 
	c := a && b // c được gán bằng a&&b 
	fmt.Println("c:", c) 
	d := a || b // d được gán bằng a||b
	fmt.Println("d:", d) 
}

```

- Trong chương trình trên, a được gán bằng true, b được gán bằng false.

- c được gán bằng giá trị của a && b. Phép và ( && ) trả về giá trị đúng (true) khi và chỉ chi cả a và b là true, ngược lại giá trị trả về là sai (false). 

- d được gán bằng giá trị của a || b. Phép hoặc ( || ) trả về giá trị là true khi hoặc a hoặc b là true hoặc cả a và b là true, ngược lại giá trị trả về là false khi cả a và b đều là false. 

### Số nguyên

- Đây là các giá trị số nguyên giống hoàn toàn như trong toán, tuy nhiên số nguyên trong máy tính có giới hạn (tức không có giá trị nào là vô cùng ∞ cả).

- Các kiểu số nguyên trong Go là uint8, uint16, uint32, uint64, int8, int16, int32, int64, int. 

- Các con số 8, 16, 32, 64 có nghĩa là máy tính cần dùng bao nhiêu bit để biểu diễn số nguyên đó.

- uint tức là unsigned int – là kiểu số nguyên không âm. Bảng dưới đây cho biết giới hạn của từng loại kiểu số nguyên:

```
KIỂU	GIỚI HẠN
uint8	0 – 255
uint16	0 – 65535
uint32	0 – 4294967295
uint64	0 – 18446744073709551615
int8	-128 – 127
int16	-32768 – 32767
int32	-2147483648 – 2147483647
int64	-9223372036854775808 – 9223372036854775807

```

- Ví dụ:

```
package main

import "fmt"

func main() {
	var num1 int16 = 20132
	var num2 int16 = 23244
    fmt.Println("Tong 2 so num1 và num2 là:", num1 + num2)

	var num3 int32 = 20132
	var num4 int32 = 23244
    fmt.Println("Tong 2 so num3 và num4 là:", num3 + num4)

	var num5 int = 20132
	var num6 int = 23244	
    fmt.Println("Tong 2 so num5 và num6 là:", num5 + num6) 
}
```

- Cộng hai số kiểu int16: num1 + num2 = 43376. Tuy nhiên, kiểu dữ liệu int16 chỉ có thể lưu trữ các giá trị trong khoảng từ -32768 đến 32767. Vì 43376 vượt quá giới hạn dương của int16, xảy ra tràn số. Cụ thể, tổng bị tràn và kết quả cuối cùng được tính như sau:  43376 - 65536 = -22160 (65536 là tổng số giá trị có thể được lưu trữ trong int16).

- Cộng hai số kiểu int32: num3 + num4 = 43376. Vì kiểu int32 có thể lưu trữ giá trị từ -2147483648 đến 2147483647, tổng 43376 nằm trong phạm vi hợp lệ của int32. Do đó, kết quả bạn nhận được là 43376.

- Cộng hai số kiểu int: num5 + num6 = 43376 . Trong Go, kiểu int thường là int32 hoặc int64 tùy thuộc vào kiến trúc hệ thống (32-bit hoặc 64-bit). Tuy nhiên, dù là int32 hay int64, cả hai đều có khả năng lưu trữ giá trị 43376 mà không bị tràn số. Vì vậy, kết quả bạn nhận được là 43376.

- Đặc biệt trong Go còn có 3 kiểu số nguyên phụ thuộc hệ điều hành là uint, int và uintptr, 3 kiểu dữ liệu này có giới hạn giống như kiến trúc của hệ điều hành mà bạn đang sử dụng. Ví dụ nếu bạn đang dùng Windows 64 bit thì kiểu int sẽ có giới hạn giống như kiểu uint64. Thông thường khi sử dụng số nguyên bạn dùng int là đủ rồi.

### Số thực

- Đây là các giá trị số có phần thập phân, ví dụ 1.234, 123.4… Việc lưu trữ cũng như thao tác với số thực trong máy tính khá phức tạp nên chúng ta cũng không đi sâu vào làm gì. Ở đây chúng ta chỉ có một số lưu ý như sau:

    - Số thực không bao giờ chính xác một cách tuyệt đối, rất khó để biểu diễn chính xác một số thực. Ví dụ như phép trừ 1.01 – 0.99 sẽ cho ra kết quả là 0.020000000000000018 chứ không phải là 0.02 như bạn vẫn nghĩ.

    - Cũng giống như số nguyên, số thực trong máy tính cũng có nhiều giới hạn khác nhau.

- Trong Go có 2 kiểu số thực là float32 và float64. Thông thường để biểu diễn số thực, bạn chỉ cần dùng float64 là đủ.

```
package main

import (  
    "fmt"
)

func main() {  
    a, b := 5.67, 8.97
    fmt.Printf("Kiểu dữ liệu của a là %T và của b là %T\n", a, b)
    sum := a + b
    sub := a - b
    fmt.Println("Tổng a và b:", sum)
    fmt.Println("Hiệu a và b:", sub)
}
```

### Số phức

- Trong Go, có 2 kiểu số phức là complex64 và complex128

- Complex64: Số phức có phần thực float32 và phần ảo

- Complex128: Số phức có phần thực float64 và phần ảo

- Số phức được tạo bằng một hàm Complex để xây dựng phần thực và phần ảo:

    `func complex(r, i FloatType) ComplexType `

    Hoặc số phức cũng có thể được tạo bằng cú pháp viết tắt:

    ` c := 6 + 7i `

- Phần thực và phần ảo như tham số và kết quả trả về là kiểu số phức. Khi khai báo số phức, chúng ta nên khai báo phần thực và phần ảo ở cùng kiểu dữ liệu. Ví dụ phần thực và phần ảo đều là float32 thì kết quả trả về sẽ có kiểu complex64, phần thực và phần ảo đều là float64 thì kết quả trả về sẽ có kiểu complex128.

```
package main

import (
	"fmt"
)

func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("Tổng 2 số phức c1 và c2:", cadd)
	cmul := c1 * c2
	fmt.Println("Tích 2 số phức c1 và c2:", cmul)
}

```

### Các kiểu dữ liệu số khác

byte là tên gọi khác của uint8

rune là tên gọi khác của int32

### Chuỗi

Trong Go, chuỗi là một tập hợp các byte, bây giờ chúng ta giả sử chuỗi là một tập hợp các kí tự. Ví dụ

```
package main

import (  
    "fmt"
)

func main() {  
    first := "Naveen"
    last := "Ramanathan"
    name := first +" "+ last
    fmt.Println("My name is",name)
}
```

Trong chương trình trên, các biến first, last được gán lần lượt bằng các chuỗi "Naveen", "Ramanathan". Các chuỗi được nối với nhau bởi dấu cộng, khoảng trắng cũng là một chuỗi, biến name được gán bằng việc cộng các chuỗi vừa khai báo ở trên và chuỗi khoảng trắng. Có một số thao tác khác với chuỗi, chúng ta sẽ tìm hiểu trong một hướng dẫn khác.

### Ép kiểu

- Ngôn ngữ Go rất nghiêm ngặt và chặt chẽ, nên chúng không cho phép tự động chuyển đổi (ép kiểu) kiểu dữ liệu.

```
package main

import (  
    "fmt"
)

func main() {  
    i := 55      //int
    j := 67.8    //float64
    sum := i + j //int + float64 not allowed
    fmt.Println(sum)
}
```

- Đoạn code trong ví dụ trên hoàn toàn chạy được trong C, nhưng trong Go, điều này sẽ không hoạt động và hệ thống báo 2 số i và j không có cùng kiểu dữ liệu.

- Để sửa lỗi, cả i và j phải cùng kiểu dữ liệu. Hãy chuyển j thành kiểu int,

```
package main

import (  
    "fmt"
)

func main() {  
    i := 55      //int
    j := 67.8    //float64
    sum := i + int(j) //j is được ép kiểu thành int
    fmt.Println(sum)
}
```