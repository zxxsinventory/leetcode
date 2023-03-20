package main

import "fmt"
import "time"

func main() {
	// 1. Two Sum
	// fmt.Println(twoSum([]int{1,2,5,4,7}, 8))

	// 2. Add Two Numbers
	// l12 := &ListNode{Val: 5, Next: nil}
	// l11 := &ListNode{Val: 5, Next: l12}
	// l22 := &ListNode{Val: 5, Next: nil}
	// l21 := &ListNode{Val: 5, Next: l22}
	// // addTwoNumbers(l11, l21)
	// res := addTwoNumbers(l11, l21)
	// for res != nil {
	// 	fmt.Println(res.Val)
	// 	res = res.Next
	// }

	// 3. Longest Substring Without Repeating Characters	
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring("bbbbb"))
	// fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	start := time.Now()
	time.Sleep(time.Second*1)
	fmt.Println(time.Now() - start)

}

func twoSum(nums []int, target int) []int{
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		remain := target - nums[i]
		if _, ok := m[remain]; ok {
			return []int{i, m[remain]}
		}
		m[nums[i]] = i
	}
	return nil
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil 
	}
  	head := &ListNode{Val: 0, Next: nil}
  	current := head
  	carry := 0
  	for l1 != nil || l2 != nil {
  		// fmt.Println("l1: ", l1.Val)
  		// fmt.Println("l1: ", l2.Val)
  		var x, y int
  		if l1 == nil {
  			x = 0
  		} else {
  			x = l1.Val
  		}
  		if l2 == nil {
  			y = 0
  		} else {
  			y = l2.Val
  		}
  		next := &ListNode{Val: (x + y + carry)%10, Next: nil}
  		// fmt.Println("v: ", next.Val)
  		carry = (x + y + carry)/10
  		// fmt.Println("v:", (x + y + carry) % 10)
  		current.Next = next
  		current = next
		if l1 != nil {
			l1 = l1.Next 
		}
		if l2 != nil {
		    l2 = l2.Next
		}
  	}
  	// fmt.Println("c:", carry)
  	if carry > 0 {
  		current.Next = &ListNode{Val: 1, Next: nil}
  	}
  	return head.Next
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	l := len(s)

	var freq [256]int
	left, right, result := 0, -1, 0

	for left < len(s) {
		if right + 1 < l && freq[s[right + 1] - 'a'] == 0 {
			freq[s[right + 1] - 'a']++
			right++
			if right - left + 1 > result {
				result = right - left + 1
			}
		} else {
			freq[s[left] - 'a']--
			left++
		}
	}
	return result
}





