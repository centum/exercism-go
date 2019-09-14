// Package tree implement parser tree struct
package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// AppendChild append new child to Node
func (n *Node) AppendChild(child *Node) {
	i := sort.Search(len(n.Children), func(i int) bool { return n.Children[i].ID > child.ID })
	n.Children = append(n.Children, &Node{})
	copy(n.Children[i+1:], n.Children[i:])
	n.Children[i] = child
}

var (
	ErrNonContinuous   = errors.New("non-continuous")
	ErrWrongParent     = errors.New("wrong parent")
	ErrCycleDependency = errors.New("cycle dependency")
	ErrNoRoot          = errors.New("no root node")
)

// Build build Node struct from list Record
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	mapNodes := make(map[int]*Node, len(records))
	maxID := 0 // needed to detect non-continuous
	for _, r := range records {
		if _, ok := mapNodes[r.ID]; !ok {
			mapNodes[r.ID] = &Node{ID: r.ID}
		}

		if r.ID == 0 {
			if r.Parent == 0 {
				continue
			}
			return nil, ErrWrongParent
		}

		if r.Parent >= r.ID {
			return nil, ErrCycleDependency
		}

		if _, ok := mapNodes[r.Parent]; !ok {
			mapNodes[r.Parent] = &Node{ID: r.Parent}
		}

		mapNodes[r.Parent].AppendChild(mapNodes[r.ID])

		if r.ID > maxID {
			maxID = r.ID
		}
	}
	if len(records) != maxID+1 {
		return nil, ErrNonContinuous
	}

	node, ok := mapNodes[0]
	if !ok {
		return nil, ErrNoRoot
	}

	return node, nil
}
