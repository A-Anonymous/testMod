package testMod

func Reverse(x int) int {
	re := 0
	for x != 0{
		t := x % 10
		re = re * 10 + t
		x = x / 10
	}
	if re < -2147483648 || re > 2147483647 {
		return 0
	}
	return re
}
