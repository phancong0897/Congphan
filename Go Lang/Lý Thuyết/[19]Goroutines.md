# Goroutines

### Goroutines là gì?

- Goroutines là các hàm hoặc phương thức chạy đồng thời với các hàm/ phương thức khác. Goroutines có thể được coi là những luồng gọn nhẹ. Chi phí tạo một Goroutine tương đối thấp so với một luồng. Do vậy, những ứng dụng Go có hàng ngàn 

- Goroutines chạy đồng thời là điều hết sức bình thường.

- Ưu điểm của Goroutines so với luồng

    - Goroutines có chi phí cực thấp so với luồng. Chúng chỉ chiếm khoảng vài kb trong stack và stack có thể "phình to" và "thu nhỏ" tùy vào nhu cầu của chương trình, trong khi đó các luồng trong stack phải được xác định và không thể thay đổi.
    
    - Goroutines được ghép kênh với một số ít hơn các luồng của HĐH. Có thể chỉ có một luồng trong một chương trình với hàng ngàn Goroutines. Nếu bất kỳ Goroutine trong những khối luồng yêu cầu chờ người dùng nhập dữ liệu, thì một luồng khác của HĐH được tạo ra và những Goroutines còn lại được di chuyển đến luồng mới. Tất cả việc này được giám sát bởi trình thực thi chạy (runtime) và chúng ta - những lập trình viên - đang trừu tượng hóa những chi tiết phức tạp, và đưa ra một API tường minh để có thể làm việc đồng thời.
    
    - Goroutines trao đổi thông qua các kênh. Các kênh được thiết kế để ngăn ngừa các khả năng xung đột xảy ra khi truy cập bộ nhớ chia sẻ sử dụng Goroutines. Các kênh có thể hình dung như một đường ống sử dụng các Goroutines trao đổi thông tin. Chúng ta sẽ bàn về các kênh chi tiết hơn trong bài tới.

### Làm thế nào để bắt đầu một Goroutine?

- Nhập từ khóa go vào trước hàm hoặc phương thức, chúng ta sẽ có một Goroutine chạy đồng thời.
```
package main

import (  
    "fmt"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()
    fmt.Println("main function")
}
```
- Tại dòng 11, go hello() bắt đầu một Goroutine mới. Bây giờ hàm hello() sẽ chạy đồng thời cùng với hàm main(). Hàm main() chạy trong Goroutine của riêng nó và được gọi là main Goroutine.

- Chạy thử chương trình và bạn sẽ thấy một sự ngạc nhiên!

- Chương trình sẽ output ra dòng chữ main function. Điều gì dã xảy ra khi Goroutine khởi chạy? Chúng ta cần nắm được hai thuộc tính chủ yếu của Goroutine để hiểu tại sao như vậy:

- Khi một Goroutine mới khởi chạy, tính năng goroutine sẽ gọi trở lại ngay lập tức. Không giống hàm, hệ thống không chờ Goroutine chạy xong, mà trả lại ngay lập tức dòng code tiếp theo ngay sau khi Goroutine gọi và bất kỳ giá trị trả lại nào từ Goroutine sẽ bị bỏ qua.
Goroutine chính cần chạy để các Goroutine khác chạy được. Nếu Goroutine chính dừng lại thì chương trình cũng dừng và không Goroutine nào chạy nữa.
Tôi nghĩ giờ bạn có thể hiểu tại sao Goroutine không chạy. Sau khi gọi go hello() tại dòng 11, hệ thống trả về ngay lập tức dòng tiếp theo của code mà không đợi hello goroutine kết thúc và in ra main function. Sau đó Goroutine chính dừng khi không còn đoạn code nào để thực thi và do đó, hello Goroutine không còn cơ hội để chạy.

- Chúng ta sẽ fix lại lỗi như sau:
```
package main

import (  
    "fmt"
    "time"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}
```
- Tại dòng 13 của chương trình trên, chúng ta gọi phương thức Sleep của gói time, dùng để "ru ngủ" goroutine khi được khởi chạy. Trong tình huống này, goroutine chính sẽ đi vào trạng thái ngủ trong 1 giây. Bây giờ gọi đến go hello() có đủ thời gian thực thi trước khi Goroutine chính dừng lại. Chương trình này sẽ output ra Hello world goroutine trước, sau đó 1 giây sẽ output tiếp main function.

- Phương pháp sử dụng sleep trong Goroutine chính để đợi những Goroutine khác kết thúc việc thực thi là một mẹo nhỏ (hack) chúng ta sử dụng để hiểu cách làm việc của các Goroutines. Những kênh có thể được sử dụng để tạo khối cho Goroutine chính đến khi tất cả những Goroutines khác kết thúc thực thi. Chúng ta sẽ bàn về các kênh trong bài tiếp theo.

### Chạy nhiều Goroutines

- Thử viết thêm một chương trình chạy nhiều Goroutines để hiểu rõ hơn về Goroutine.

```
package main

import (  
    "fmt"
    "time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}
func main() {  
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main terminated")
}
```
- Chương trình trên bắt đầu hai Goroutine tại dòng 21 và 22. Hai Goroutine này chạy đồng thời với nhau. Thời gian Goroutine "ngủ" khởi tạo là 250 ms và sau đó in ra 1, sau đó tiếp tục "ngủ" và in ra 2, cứ thế tiếp tục đến khi in ra 5. Tương tự Goroutine alphabets in ra các chữ cái từ a đến e và có 400ms để "ngủ". Goroutine chính khởi chạy numbers và alphabets Goroutine, "ngủ" trong 3000 ms và sau đó dừng lại. 

- Chương trình output ra:

```
1 a 2 3 b 4 c 5 d e main terminated  
```