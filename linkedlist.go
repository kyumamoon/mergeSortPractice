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

func main() {
	newList := linkedlist{}
	newNode01 := node{data: 1, next: nil}
	newNode02 := node{data: 2, next: nil}
	newNode03 := node{data: 3, next: nil}
	newNode04 := node{data: 4, next: nil}
	newNode05 := node{data: 5, next: nil}
	newNode06 := node{data: 6, next: nil}

	newList.append(&newNode01)
	newList.append(&newNode02)
	newList.append(&newNode03)
	newList.append(&newNode04)
	newList.append(&newNode05)
	newList.append(&newNode06)
	split, check := halfSplit(&newList)
	var left linkedlist
	var right linkedlist

	if check {
		left = split[0]
		right = split[1]
		left.printList()
		fmt.Printf("\n")
		right.printList()
		fmt.Println(check)
	} else {
		split[0].printList()
		fmt.Println(check)
	}

	split2, check := halfSplit(&left)
	var left2 linkedlist
	var right2 linkedlist
	if check {
		left2 = split2[0]
		right2 = split2[1]
		left2.printList()
		fmt.Printf("\n")
		right2.printList()
		fmt.Println(check)
	} else {
		split[0].printList()
		fmt.Println(check)
	}
	//fmt.Println(left.length)
	//fmt.Println(right.length)

}
