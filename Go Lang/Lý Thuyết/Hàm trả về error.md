# Hàm trả về error

- Do go không có try/catch nên luôn phải kiểm trả lỗi trả về từ mỗi func.

- Quy ước tham số trả về cuối cùng luôn là error

- Go lang có cú pháp đặc biệt gọi là if assignment gồm 2 vế:

    - 1. Asignment : `result, err := Sqrt(4)`

    - 2. Condition: `err != nil`

- Ví dụ:
```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("so khong hop le %g")
	}
	return math.Sqrt(f), nil
}

func main() {
	if a, err := Sqrt(4); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Khai căn số vừa nhập là: %f", a)
	}
}

```