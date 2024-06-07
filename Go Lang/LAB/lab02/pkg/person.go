package pkg

import (
	"sort"
)

type Person struct {
	Name        string
	SalaryRatio int
	BonusMoney  int
}

// Hàm sắp xếp tên nhân viên tăng dần theo bảng chữ cái
func SortByName(list []Person) (result []Person) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}

// Hàm Tính lương
func (p Person) getsalary() int {
	return p.SalaryRatio*1500000 + p.BonusMoney
}

// Sắp xếp lương theo thứ tự giảm dần
func SortBySalary(list []Person) (result []Person) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].getsalary() > list[j].getsalary()
	})

	return list
}

// Lấy ra số lương cao nhất
func FindMaxSalary(list []Person) (result int) {
	personSortBySalary := SortBySalary(list)
	for i := 0; i < len(personSortBySalary); i++ {
		if personSortBySalary[i].getsalary() < personSortBySalary[0].getsalary() {
			result = personSortBySalary[i].getsalary()
			break
		}
	}
	return result
}

// In ra thông tin nhân viên có mức lương cao nhất
func FindMaxSalaryPerson1(list []Person) (result []Person) {
	max3 := FindMaxSalary(list)
	for _, e := range list {
		if e.getsalary() == max3 {
			result = append(result, e)
		}
	}
	return result
}

// Lấy ra số lương của nhân viên cao thứ hai
func FindMax02Salary(list []Person) (result int) {
	var max = FindMaxSalary(list)
	max2 := 0
	for i := 0; i < len(list); i++ {
		if list[i].getsalary() > max2 && list[i].getsalary() < max {
			max2 = list[i].getsalary()
		}
	}
	return max2
}
// In ra thông tin nhân viên có mức lương cao thứ hai
func FindMaxSalaryPerson2(list []Person) (result []Person){
	max4 := FindMax02Salary(list)
	for _, e := range list{
		if e.getsalary() = max4{
			result = append(result, e)
		}
	}
}