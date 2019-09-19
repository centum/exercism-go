// Package tree implement parser tree struct
package tree

import (
	"errors"
	"sort"
)

// Record represents input dataset
type Record struct {
	ID, Parent int
}

// Node represents node of tree
type Node struct {
	ID       int
	Children []*Node
}

var (
	// ErrNonContinuous error broken sequence
	ErrNonContinuous = errors.New("non-continuous")
	// ErrWrongParent error wrong parent
	ErrWrongParent = errors.New("wrong parent")
	// ErrCycleDependency error detected cycle dependency
	ErrCycleDependency = errors.New("cycle dependency")
	// ErrNoRoot error raised if not fount root node
	ErrNoRoot = errors.New("no root node")
)

// Build build Node struct from list Record
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	mapNodes := make(map[int]*Node, len(records))
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

		mapNodes[r.Parent].Children = append(mapNodes[r.Parent].Children, mapNodes[r.ID])
	}

	if i := len(records) - 1; i != records[i].ID {
		return nil, ErrNonContinuous
	}

	node, ok := mapNodes[0]
	if !ok {
		return nil, ErrNoRoot
	}

	return node, nil
}
