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

var (
	ErrNonContinuous = errors.New("non-continuous")
	ErrDuplicateNode = errors.New("duplicate node")
	ErrWrongParent   = errors.New("wrong parent")
	ErrCycleDirectly = errors.New("cycle directly")
	ErrNoRoot        = errors.New("no root node")
)

func (n *Node) findChildren(records *[]Record, pos int) error {
	for i, r := range *records {
		pos++
		if r.ID != pos {
			return ErrNonContinuous
		}
		if r.ID == n.ID {
			return ErrDuplicateNode
		}
		if r.Parent > r.ID {
			return ErrWrongParent
		}
		if r.Parent == r.ID {
			return ErrCycleDirectly
		}
		if r.Parent == n.ID {
			child := &Node{ID: r.ID}
			n.Children = append(n.Children, child)
			recs := (*records)[i+1:]
			if err := child.findChildren(&recs, pos); err != nil {
				return err
			}
		}
	}
	return nil
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, ErrNoRoot
	}

	records = records[1:]

	node := &Node{}

	err := node.findChildren(&records, 0)

	return node, err
}
