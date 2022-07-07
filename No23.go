package main

import (
	"fmt"
	"sort"
)

/*

 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type nodeSorter struct {
	vals []*ListNode
}

func (ns nodeSorter) Len() int {
	return len(ns.vals)
}

func (ns nodeSorter) Less(i, j int) bool {
	return ns.vals[i].Val < ns.vals[j].Val
}

func (ns *nodeSorter) Swap(i, j int) {
	t := ns.vals[i]
	ns.vals[i] = ns.vals[j]
	ns.vals[j] = t
}

func mergeKLists(lists []*ListNode) *ListNode {
	vals := make([]*ListNode, 0)
	for _, list := range lists {
		for itr := list; itr != nil; itr = itr.Next {
			vals = append(vals, itr)
		}
	}
	if len(vals) <= 0 {
		return nil
	}
	sort.Sort(&nodeSorter{vals})
	for i := 1; i < len(vals); i++ {
		vals[i-1].Next = vals[i]
	}
	vals[len(vals)-1].Next = nil
	return vals[0]
}

func main() {
	var li1 *ListNode
	var li2 *ListNode

	litest := []*ListNode{li1, li2}
	test := mergeKLists(litest)
	fmt.Print(test.Val)
}
