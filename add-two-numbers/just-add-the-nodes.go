package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Nodes := countNodes(l1)
	l2Nodes := countNodes(l2)

	if l2Nodes > l1Nodes {
		l1, l2 = l2, l1
	}

	for left, right := l1, l2; left != nil; left = left.Next {
		if right != nil {
			left.Val += right.Val
			right = right.Next
		}

		if left.Val >= 10 {
			left.Val -= 10

			if left.Next != nil {
				left.Next.Val += 1
			} else {
				left.Next = &ListNode{
					1,
					nil,
				}
			}


		}

	}

	return l1
}

func countNodes(list *ListNode) int {
	count := 0
	for current := list; current != nil; current = current.Next  {
		count++
	}

	return count
}

func buildLinkedList(val uint64) *ListNode  {
	n := numberOfDigits(val)

	var list *ListNode
	for level := n; level > 0; level-- {
		divisor := power10(level)
		digit := val / divisor
		val -= digit * divisor

		list = &ListNode{
			int(digit),
			list,
		}

	}

	return list
}

func power10(n uint64) uint64 {
	f := uint64(1)
	for ; n > 1; n-- {
		f *= 10
	}
	return f
}

func numberOfDigits(number uint64) uint64 {
	i := 0
	if number == 0 {
		return uint64(1)
	}

	for number != 0 {
		i++
		number /= 10
	}
	return uint64(i)
}


func main()  {

	l1 := buildLinkedList(998)
	l2 := buildLinkedList(2)
	result := addTwoNumbers(l1, l2)
	fmt.Println(result)

}

