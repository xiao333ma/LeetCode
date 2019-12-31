package other

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
     Val int
     Next *ListNode
 }



func main() {
	head := &ListNode{Val:1}
	head.Next = &ListNode{Val:2}
	head.Next.Next = &ListNode{Val:3}
	head.Next.Next.Next = &ListNode{Val:4}
	head.Next.Next.Next.Next = &ListNode{Val:5}
	head.Next.Next.Next.Next.Next = &ListNode{Val:6}
	middleNode(head)

}

func middleNode(head *ListNode) *ListNode {
	array := make([]*ListNode, 0)
	for {
		if head != nil {
			array = append(array, head)
			head = head.Next
		} else {
			break
		}
	}
	return array[len(array) / 2]
}

//876.Middle_of_the_Linked_List
func middleNode(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}