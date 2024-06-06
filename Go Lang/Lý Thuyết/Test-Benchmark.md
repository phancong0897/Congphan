# Test và Benchmark trong Go

- Go lang cung cấp 2 kỹ thuật:

    - 1. Unit Test kiểm thử logic chạy có theo ý đồ của lập trình viên không
    - 2. Benchmark kiểm thử tốc độ thực thi

Chúng ta có thể viết file *test.go nằm trong cùng package hoặc thư mục riêng

File test hay benchmark luôn phải kết thúc bằng `test.go` và `import testing`

Hàm benchmark luôn phải bắt đầu bằng `fun Benchmark`

- Ví dụ :
    
    - Lưu ý nhớ tạo module

file main.go

```go
package main

func removeItemSliceNotKeepOrder(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

// Xóa phần tử giữ nguyên thứ tự
// cách 1
func removeItemSliceKeepOrder(a []string, i int) []string {
	copy(a[1:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

// cách 2
func removeItemSliceKeepOrder2(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

```

file remove_slice_benchmark_test.go
```go
package main

import (
	"testing"
)

// Xóa phần tử ưu tiên tốc độ, không theo thứ tự
func Benchmark_removeItemSliceNotKeepOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeItemSliceKeepOrder([]string{"a", "b", "c", "d", "e"}, 2)
	}
}

func Benchmark_removeItemSliceKeepOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeItemSliceKeepOrder([]string{"a", "b", "c", "d", "e"}, 2)
	}
}

func Benchmark_removeItemSliceKeepOrder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeItemSliceKeepOrder2([]string{"a", "b", "c", "d", "e"}, 2)
	}
}

```

- Chạy lệnh sau : `go test -bench .` sẽ trả ra kết quả benchmark

```
PS E:\Go Lang\Arrays-loop> go test -bench .
goos: windows
goarch: amd64
pkg: arrays-loop
cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
Benchmark_removeItemSliceNotKeepOrder-12        134349687                8.852 ns/op
Benchmark_removeItemSliceKeepOrder-12           135429922                8.861 ns/op
Benchmark_removeItemSliceKeepOrder2-12          139383878                8.600 ns/op
PASS
ok      arrays-loop     6.381s
PS E:\Go Lang\Arrays-loop> 
```