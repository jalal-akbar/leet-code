package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	// if x > 0 {
	// 	return true
	// }
	// if x < 0 || x%10 == 0 {
	// 	return false
	// }
	// reverseInt := 0
	// if x > reverseInt {
	// 	pop := x % 10
	// 	x = x / 10
	// 	reverseInt = reverseInt*10 + pop
	// }
	// if x == reverseInt || x == reverseInt%10 {
	// 	return true
	// }
	// return false
	xString := strconv.Itoa(x)
	//temp := ""
	for i := 0; i < len(xString)/2; i++ {
		if xString[i] != xString[len(xString)-1-i] {
			return false
		}

	}
	return true
}
