/*
Viết function tìm ra số lớn thứ nhì trong mảng các số. Ví dụ: max2Numbers([2, 1, 3, 4]) => 3
*/
package pkg

// Tìm số lớn nhất 
func Findmaxnumber01(number []int) int {
	var max int = number[0]
	for i := 0; i < len(number); i++ {
		if number[i] > max {
			max = number[i]
		}
	}
	return max
}

// Tìm số lớn thứ 2
func Findmaxnumber02(number []int) int {
	var max = Findmaxnumber01(number)
	max2 := 0
	for i := 0; i < len(number); i++ {
		if number[i] > max2 && number[i] < max {
			max2 = number[i]
		}
	}
	return max2

}
