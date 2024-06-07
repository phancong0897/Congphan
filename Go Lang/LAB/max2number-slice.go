package max2number

func max01number(number []int) int {
	var max int = number(0)
	for i := 0; i < len(number); i++ {
		if number(i) > max {
			max = number[i]
		}
	}
}
