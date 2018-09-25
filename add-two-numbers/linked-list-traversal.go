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

	l1val := parseLinkedList(l1)
	l2val := parseLinkedList(l2)

	result := l1val + l2val

	return buildLinkedList(result)
}

func parseLinkedList(list *ListNode) uint64 {
	val := uint64(0)

	level := uint64(1)
	for current := list; current != nil; current = current.Next  {
		multiplier := power10(level)
		val += uint64(current.Val) * multiplier

		level++
	}

	return val
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

func main()  {

	l1 := buildLinkedList(1000000000000000000000000000001)
	l2 := buildLinkedList(564)
	result := addTwoNumbers(l1, l2)
	fmt.Println(result)

}

