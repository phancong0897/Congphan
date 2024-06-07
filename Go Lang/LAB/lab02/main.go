/*
Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:

Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
*/
package main

import (
	"fmt"
	"lab02/pkg"
)

func main() {
	number := []int{1, 2, 3, 5, 4, 6, 9, 8}
	fmt.Println(pkg.Findmaxnumber01(number))
	fmt.Println("Số lớn thứ hai trong dãy là:", pkg.Findmaxnumber02(number))
	persons := []pkg.Person{
		{"Hang", 5, 1000000},
		{"An", 7, 1500000},
		{"Cong", 10, 1300000},
		{"Ngoc", 7, 2000000},
		{"Oanh", 8, 2500000},
		{"Tam", 8, 2600000},
	}

	
	fmt.Println(pkg.SortByName(persons))   // Sắp xếp theo tên tăng dần theo bảng chữ cái
	fmt.Println(pkg.SortBySalary(persons))  // Sắp xếp lương nhân viên theo chiều giảm dần
	fmt.Println(pkg.FindMaxSalary(persons))  // Tìm ra số lương cao nhất
	fmt.Println(pkg.FindMax02Salary(persons)) // Tìm ra số lương cao thứ hai
	fmt.Println(pkg.FindMaxSalaryPerson1(persons)) // In ra thông tin của người có lương cao nhất
	fmt.Println(pkg.FindMaxSalaryPerson2(persons))  // In ra thông tin của người có lương cao nhất
}
