package main

import "fmt"

type linkedlist struct {
	head   *node
	length int
}

type node struct {
	data int
	next *node
}

func (l *linkedlist) append(n *node) {
	if l.head == nil {
		l.head = n
		l.length++
	} else {
		current := l.head

		for {
			if current.next == nil {
				current.next = n
				l.length++
				break
			} else {
				current = current.next
			}
		}

	}
}

func halfSplit(l *linkedlist) ([]linkedlist, bool) {
	if l.length < 2 {
		return []linkedlist{*l}, false
	} else {
		mid := l.length / 2

		listLeft := linkedlist{}
		listRight := linkedlist{}

		first := l.head
		var second *node

		current := l.head
		{
			for i := 0; i < l.length; i++ {
				if i == l.length-mid {
					second = current
					break
				} else {
					current = current.next
				}
			}
		}

		{
			current := first
			for i := 0; i < l.length-mid; i++ {
				listLeft.append(&node{data: current.data})
				current = current.next
			}
		}
		{
			current := second
			for i := 0; i < mid; i++ {
				listRight.append(&node{data: current.data})
				current = current.next
			}
		}
		return []linkedlist{listLeft, listRight}, true
	}

}

func (l *linkedlist) printList() {
	if l.head == nil {
		fmt.Println("EMPTY")
	} else {
		current := l.head
		for {
			fmt.Println(current.data)
			if current.next == nil {
				break
			} else {
				current = current.next
			}
		}
	}
}

func (l *linkedlist) indexList(index int) int {
	if l.length == 0 || index > l.length {
		return 0
	}

	count := 0
	current := l.head

	for {
		if index == count {
			return current.data
		} else {
			count++
			current = current.next
		}
	}
}

func mergeSort(l linkedlist) linkedlist {
	if l.length < 2 {
		return l
	}

	var left, right linkedlist

	split, check := halfSplit(&l)
	if check {
		left = mergeSort(split[0])
		right = mergeSort(split[1])
	}

	return merge(left, right)
}

func merge(left, right linkedlist) linkedlist {
	answer := linkedlist{}
	leftCount := 0
	rightCount := 0

	for leftCount < left.length && rightCount < right.length {
		if left.indexList(leftCount) < right.indexList(rightCount) {
			answer.append(&node{data: left.indexList(leftCount)})
			leftCount++
		} else {
			answer.append(&node{data: right.indexList(rightCount)})
			rightCount++
		}
	}

	for ; leftCount < left.length; leftCount++ {
		answer.append(&node{data: left.indexList(leftCount)})
	}

	for ; rightCount < right.length; rightCount++ {
		answer.append(&node{data: right.indexList(rightCount)})
	}

	return answer
}

func main() {
	test := []int{111, 1, 7, 2, 1252151, 23, 23, 2, 6, 8, 9, 51}
	x := linkedlist{}

	for _, e := range test {
		x.append(&node{data: e})
	}

	y := mergeSort(x)
	y.printList()
}
