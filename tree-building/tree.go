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

func (n *Node) findChildren(records []Record) error {
	for _, r := range records {
		if r.ID == n.ID {
			continue
		}
		if r.Parent > r.ID {
			return errors.New("wrong parent")
		}
		if r.Parent == r.ID {
			return errors.New("cycle directly")
		}
		if r.Parent == n.ID {
			if n.Children == nil {
				n.Children = make([]*Node, 0)
			}
			if len(n.Children) > 0 && n.Children[len(n.Children)-1].ID == r.ID {
				return errors.New("duplicate node")
			}
			child := &Node{ID: r.ID}
			n.Children = append(n.Children, child)
			if err := child.findChildren(records); err != nil {
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
		return nil, errors.New("no root node")
	}

	node := &Node{}
	records = records[1:]

	prevID := 0
	for i, r := range records {
		if r.ID == prevID {
			return nil, errors.New("duplicate node")
		}
		if r.ID != i+1 {
			return nil, errors.New("non-continuous")
		}
		prevID = r.ID
	}

	err := node.findChildren(records)

	return node, err
}
