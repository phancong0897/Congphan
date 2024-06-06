// index loop
```go
func indexLoop() {
	cars := []string{"Toyota", "Mercedes", "BMW"}
	for i := 0; i < len(cars); i++ {
		fmt.Println(i, cars[i])
	}
}
```go
/*
Go lang không có while
*/
```go
func whileLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	i := 0
	for i < len(cars) {
		fmt.Println(i, cars[i])
	}
}
```
```go
func rangeLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	for index, car := range cars {
		fmt.Println(index, car)
	}
}
```

/*
reverseLoop sẽ trả kết quả từ phần tử cuối cùng đến phần tử đầu tiên ngược lại với các hàm trên trả từ phần tử đầu tiên đến phần tử cuối cùng.
Defer : Đưa một lệnh vào một ngăn xếp đặc biệt. Trước khi hàm thoát, thì sẽ thực hiện tuần tự các lệnh trong ngăn xếp này theo cơ chế Last in First Out
*/
```go
func reverseLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	for index, car := range cars {
		defer fmt.Println(index, car)
	}
}
```

/*
Mảng 2 chiều
*/
```go
func demoArray2D(){
	langs := [][]string{
		{"C", "C#", "python"},
		{"java", "perl"},
		{"C++", "go", "Rust", "Crystal"}
	}
	for _, v := range langs{
		for _, lang := range v{
			fmt.Print(lang, "")
		}
	}
}
```

/*
Slice
array trong Go là static aray. Số phần tử mảng không thay đổi sau khi khởi tạo.
Muốn thay đổi phần tử, thêm, sửa, xóa phải dùng slice.
Phần tử trong mảng array có thể là value có thể là reference (con trỏ đến vùng nhớ) tùy thuộc vào kiểu phần tử mảng
Chỉ mục truy xuất đến mảng luôn bắt đầu từ 0.
Chú ý chỉ số sau ký tự ":"(hai chấm) là exclusive chứ không phải inclusive, có nghĩa là không bao gồm
*/

```go
func demoSlice(){
	a := []string{"a", "b", "c", "d", "e", "f"}
	a = append(a, "e")
	fmt.Println(len(a))
	fmt.Println(a[:2])  // Lấy 2 phần tử đầu tiên
	fmt.Println(a[2:])  // Bỏ qua 2 phần tử đầu tiên
	fmt.Println(a[len(a)-2:]) //Lấy 2 phần tử cuối cùng
	fmt.Println(a[1:3])
}
```