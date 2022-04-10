package _141_Linked_List_Cycle

import "github.com/halfrost/LeetCode-Go/structures"

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	mapCycle := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := mapCycle[head]; ok {
			return true
		}
		mapCycle[head] = struct{}{}
	}
	return false
}
