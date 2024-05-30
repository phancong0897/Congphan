# Kiểu dữ liệu string

- String là một kiểu dữ liệu thường gặp nhất trong ngôn ngữ lập trình. Do có một số khác biệt trong việc sử dụng so với các ngôn ngữ khác, chúng tôi dành riêng phần này để nói về string.

### String là gì?

- String trong Go là một slice chứa các byte. String có thể được tạo ra bằng cách bao quanh nội dung của nó trong dấu “ “. Hãy xem một ví dụ đơn giản dưới đây, tạo một string và in ra màn hình:

```
package main

import (  
    "fmt"
)

func main() {  
    name := "Hello World"
    fmt.Println(name)
}
```

### Truy cập các byte trong một string

- Do string là một slice của byte, chúng ta có thể truy cập từng phần tử byte của một string.
```
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}

func main() {  
    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
}
```

### Rune

- Một rune là một kiểu dữ liệu có sẵn trong Go, một cách gọi khác của kiểu dữ liệu int32. Rune đại diện cho một điểm mã trong Go. Không quan trọng điểm mã chứa bao nhiêu byte, nó có thể đại diện bởi một rune. Chúng ta hãy sửa ví dụ trên để in ra ký tự sử dụng rune:

```
package main

import (  
   "fmt"
)

func printBytes(s string) {  
   for i:= 0; i < len(s); i++ {
       fmt.Printf("%x ", s[i])
   }
}

func printChars(s string) {  
   runes := []rune(s)
   for i:= 0; i < len(runes); i++ {
       fmt.Printf("%c ",runes[i])
   }
}

func main() {  
   name := "Hello World"
   printBytes(name)
   fmt.Printf("\n")
   printChars(name)
   fmt.Printf("\n\n")
   name = "Señor"
   printBytes(name)
   fmt.Printf("\n")
   printChars(name)
}
```

### Vòng lặp for range của một string

- Ví dụ phần trên là một cách hoàn hảo để lặp từng phần tử rune của một string. Tuy nhiên, Go cung cấp một cách dễ dàng hơn để làm được điều đó, bằng cách sử dụng vòng lặp for ... range.

```
package main

import (  
   "fmt"
)

func printCharsAndBytes(s string) {  
   for index, rune := range s {
       fmt.Printf("%c starts at byte %d\n", rune, index)
   }
}

func main() {  
   name := "Señor"
   printCharsAndBytes(name)
}
```

### Thiết lập string từ slice chứa các byte

```
package main

import (  
   "fmt"
)

func main() {  
   byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
   str := string(byteSlice)
   fmt.Println(str)
}
```

### Tạo string từ slice chứa các rune

```
package main

import (  
   "fmt"
)

func main() {  
   runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
   str := string(runeSlice)
   fmt.Println(str)
}
```
- Trong ví dụ trên, biến runeSlice chứa các điểm mã Unicode của từ Señor dạng hexadecimal. Chương trình sẽ in ra Señor.

### Độ dài của string

```
package main

import (  
   "fmt"
   "unicode/utf8"
)

func length(s string) {  
   fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}
func main() {  
   word1 := "Señor" 
   length(word1)
   word2 := "Pets"
   length(word2)
}
```
- Hàm RuneCountInString(s string) (n int) của gói UTF8 thường được sử dụng để tìm độ dài của string. Phương thức này coi string như một đối số và trả về số lượng rune trong đó.

### Kiểu dữ liệu string là bất biến

- Kiểu dữ liệu string là bất biến, có nghĩa là, một khi string được tạo ra, nó không thể bị thay đổi. Hãy cùng xem ví dụ dưới đây:

```
package main

import (  
   "fmt"
)

func mutate(s string)string {  
   s[0] = 'a' // bất kỳ ký tự unicode hợp lệ trong ngoặc nháy đơn đều là một rune 
   return s
}
func main() {  
   h := "hello"
   fmt.Println(mutate(h))
}
```