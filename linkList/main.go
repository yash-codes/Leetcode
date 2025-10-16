/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
2->4->3
5->6->4
--------
7->0->8

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//BY ME After

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	finalLinkList := &ListNode{Val: 0}
	current := finalLinkList
	var carry int = 0

	for l1 != nil || l2 != nil || carry > 0 {
		x := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		y := 0
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

	}
	return finalLinkList.Next
}

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//BY ME Before
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var retLL = &ListNode{Val: 0}
	current := retLL
	var carry, sum int
	count := 0

	for l1 != nil || l2 != nil || carry > 0 {
		fmt.Println("itiration: ", count)
		count++
		sum = 0
		if l1 == nil && l2 == nil {
			fmt.Println("l1 and l2 are nil, carry: ", carry)
			current.Next = &ListNode{Val: carry}
			current = current.Next
			carry = 0
		} else if l1 == nil {

			fmt.Println("l1 is nil, carry: ", carry)
			sum = l2.Val + carry
			fmt.Println(sum)
			if sum > 9 {
				carry = sum / 10
				sum = sum % 10
			} else {
				carry = 0
			}

			fmt.Println("digits: ", sum, carry)
			//retLL = l2
			//retLL.Val = sum
			current.Next = &ListNode{Val: sum}
			l2 = l2.Next
			//retLL = retLL.Next
			current = current.Next

		} else if l2 == nil {
			fmt.Println("l2 is nil, carry: ", carry)
			sum = l1.Val + carry
			fmt.Println(sum)
			if sum > 9 {
				carry = sum / 10
				sum = sum % 10
			} else {
				carry = 0
			}

			fmt.Println("digits: ", sum, carry)
			//retLL.Val = sum
			current.Next = &ListNode{Val: sum}
			l1 = l1.Next
			//retLL = retLL.Next
			current = current.Next
		} else {
			fmt.Println("l1 and l2 both not nil, carry: ", carry)
			sum = l1.Val + l2.Val + carry
			fmt.Println(sum)
			if sum > 9 {
				carry = sum / 10
				sum = sum % 10
			} else {
				carry = 0
			}

			fmt.Println("digits: ", sum, carry)
			//retLL.Val = sum
			current.Next = &ListNode{Val: sum}

			l1 = l1.Next
			l2 = l2.Next
			//retLL = retLL.Next
			current = current.Next
		}
	}
	return retLL.Next
}
*/

// BY CHATGPT
// addTwoNumbers adds two numbers represented by linked lists.
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var retLL = &ListNode{Val: 0}
	current := retLL
	var carry int

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		digit := sum % 10

		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return retLL.Next
}
*/
// buildList creates a linked list from a slice of integers.
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, val := range nums {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

// printList prints the linked list in a readable format.
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

// main function to test addTwoNumbers
func main() {
	// Initialize l1 = [2, 4, 3], l2 = [5, 6, 4]
	fmt.Printf("\nCase1\n")
	l1 := buildList([]int{2, 4, 3}) // represents 342
	l2 := buildList([]int{5, 6, 4}) // represents 465

	// Add the two numbers
	result := addTwoNumbers(l1, l2) // Should return 807 -> [7, 0, 8]

	// Print the result

	//printList(result)
	fmt.Println("Input L1: ")
	printList(l1)
	fmt.Println("Input L2: ")
	printList(l2)
	fmt.Println("Result: ")
	printList(result)

	fmt.Printf("\n\nCase2\n")
	l3 := buildList([]int{2, 4, 3, 9}) // represents 342
	l4 := buildList([]int{5, 6, 4, 9}) // represents 465

	// Add the two numbers
	result2 := addTwoNumbers(l3, l4) // Should return 807 -> [7, 0, 8]

	// Print the result

	//printList(result)
	fmt.Println("Input L1: ")
	printList(l3)
	fmt.Println("Input L2: ")
	printList(l4)
	fmt.Println("Result: ")
	printList(result2)

	fmt.Printf("\n\nCase2\n")
	l5 := buildList([]int{2, 4, 3, 9, 9, 2, 4}) // represents 342
	l6 := buildList([]int{5, 6, 4, 9})          // represents 465

	// Add the two numbers
	result3 := addTwoNumbers(l5, l6) // Should return 807 -> [7, 0, 8]

	// Print the result

	//printList(result)
	fmt.Println("Input L1: ")
	printList(l5)
	fmt.Println("Input L2: ")
	printList(l6)
	fmt.Println("Result: ")
	printList(result3)
}
