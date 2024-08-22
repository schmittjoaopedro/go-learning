package copy_list_with_random_pointer

// start 18:37
// end 18:57
// https://leetcode.com/problems/copy-list-with-random-pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	var newList *Node = nil
	var pairs map[*Node]*Node = make(map[*Node]*Node)

	var currOri *Node = head
	var currCop *Node = nil
	for currOri != nil {
		newNode := &Node{currOri.Val, nil, nil}
		if currCop == nil {
			newList = newNode
			currCop = newNode
		} else {
			currCop.Next = newNode
			currCop = newNode
		}
		pairs[currOri] = currCop
		currOri = currOri.Next
	}

	currOri = head
	for currOri != nil {
		if currOri.Random != nil {
			pairs[currOri].Random = pairs[currOri.Random]
		}
		currOri = currOri.Next
	}

	return newList
}
