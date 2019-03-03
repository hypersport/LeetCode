/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	for slow, fast := head, head.Next; slow != fast; slow, fast = slow.Next, fast.Next.Next {
		if fast == nil || fast.Next == nil {
			return false
		}
	}
	return true
}
