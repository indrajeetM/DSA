package main

import (
	"fmt"
	"os"
)

// following is the structure for linked list
// this has two key-value pairs, which are
// data : which store the data of type string
// nextNode : which stores/points the address of the next node
type linkedlist struct {
	data     string
	nextNode *linkedlist
}

// This method helps to add the new node at the end of the list
func (list *linkedlist) insertNodeAtEnd(nodeData string) {

	//We will create a temporary node with the given data and address as nil/null/empty
	//As we are adding the node at the end of list, the new node should act as tail
	//And tail represent no further node in list, hence address is empty

	tempNode := &linkedlist{data: nodeData, nextNode: nil}

	//firstly will check if the immediate next node in the list is poiting to a nil/null/empty
	//that means its just head node present

	// before adding new node current list would look like this
	//[ Head | <nil> ]

	if list.nextNode == nil {
		list.nextNode = tempNode

		//after adding new node, list would look like this
		//[ Head | <0xc00012c028> ] => [ Head | <nil> ]

		return
	} else {
		//If the immediate next node is not ponting to empty address
		//then we have read throught the each
		tempList := list
		for tempList.nextNode != nil {
			tempList = tempList.nextNode
		}
		tempList.nextNode = tempNode
		return

	}
}

func (list *linkedlist) display() {
	if list == nil {
		fmt.Println("Empty list")
		return
	} else {
		for list != nil {
			if list.nextNode != nil {
				fmt.Printf("[ %v | %v ] ==> ", list.data, &list.nextNode)
			} else {
				fmt.Printf("[ %v | <nil> ] ==> ", list.data)
			}
			list = list.nextNode
		}
		fmt.Println("")
	}
}

func (list *linkedlist) insertNodeAtStart(nodeData string) *linkedlist {
	tempNode := &linkedlist{data: nodeData, nextNode: list.nextNode}
	list.nextNode = tempNode
	return list
}

func (list *linkedlist) removeNodeAtEnd() {

	for list.nextNode.nextNode != nil {
		list = list.nextNode
	}
	list.nextNode = nil
}

func (list *linkedlist) removeNodeAtStart() {
	if list.nextNode == nil {
		fmt.Println("This is head node")
	} else {
		list.nextNode = list.nextNode.nextNode
	}
}

func (list *linkedlist) insertNodeAtPosition(nodeData string, position int) {
	pre := list
	tempNode := &linkedlist{data: nodeData}
	for i := 0; i < position; i++ {
		if pre.nextNode == nil {
			pre.nextNode = tempNode
			return
		}
		pre = pre.nextNode
	}

	tempNode.nextNode = pre.nextNode
	pre.nextNode = tempNode
}

func (list *linkedlist) removeNodeAtPosition(position int) {
	pre := list
	for i := 0; i < position; i++ {
		if pre.nextNode == nil {
			fmt.Println("reached end")
			return
		}
		pre = pre.nextNode
	}
	pre.nextNode = pre.nextNode.nextNode
}

func main() {

	fmt.Println("Welcome to linked-list")
	fmt.Println("Select from below options")

	head := &linkedlist{data: "head"}
	var userChoice string
	for {
		fmt.Println("#################################")
		fmt.Println("1. Add node at the end")
		fmt.Println("2. Add node at the start")
		fmt.Println("3. Remove node at the end")
		fmt.Println("4. Remove node at the start")
		fmt.Println("5. Add node at the position")
		fmt.Println("6. Remove node at the position")
		fmt.Println("7. Display")
		fmt.Println("8. Exit")
		fmt.Scan(&userChoice)

		switch userChoice {
		case "1":
			var nodeData string
			fmt.Println("Enter your data of the node")
			fmt.Scan(&nodeData)
			head.insertNodeAtEnd(nodeData)
			head.display()
		case "2":
			var nodeData string
			fmt.Println("Enter your data of the node")
			fmt.Scan(&nodeData)
			head = head.insertNodeAtStart(nodeData)
			head.display()
		case "3":
			head.removeNodeAtEnd()
			head.display()
		case "4":
			head.removeNodeAtStart()
			head.display()
		case "5":
			var nodeData string
			fmt.Println("Enter your data of the node")
			fmt.Scan(&nodeData)
			var nodePosition int
			fmt.Println("Enter node position to insert")
			fmt.Scan(&nodePosition)
			head.insertNodeAtPosition(nodeData, nodePosition)
			head.display()
		case "6":
			var nodePosition int
			fmt.Println("Enter node position to remove")
			fmt.Scan(&nodePosition)
			head.removeNodeAtPosition(nodePosition)
			head.display()
		case "7":
			head.display()
		case "8":
			fmt.Println("Thank you")
			os.Exit(0)
		}

	}

}
