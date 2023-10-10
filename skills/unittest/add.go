package unittest

// 求和函数
func Sum(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}
