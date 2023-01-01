package main

import (
	"fmt"

	"github.com/jalal-akbar/leet-code/leetcode"
)

func main() {
	// 1. Two Nums
	// var nums = []int{2, 10, 7, 15}
	// var target = 9
	// n := twosums.TwoSums(nums, target)
	// fmt.Println( "Two Nums", n)
	// 2. Add Two Numbers
	nod := leetcode.AddTwoNumbers(&leetcode.ListNode{Value: 243}, &leetcode.ListNode{Value: 564})
	fmt.Println(nod)
	// 9. Palidnrome Number
	// x := 101
	// res := palindromenumber.IsPalindrome(x)
	// fmt.Println("9. Palindrome Number", res)

}
