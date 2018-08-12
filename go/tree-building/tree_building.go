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

func Build(records []Record) (*Node, error) {
    if len(records) == 0 {
        return nil, nil
    }

    sort.Slice(records, func(i, j int) bool {
                            return records[i].ID < records[j].ID
                        })

    if !IsRootRecord(&records[0]) {
        return nil, errors.New("Root record not found")
    }

    rootNode := CreateNode(records[0].ID)

    nodeList := make([]*Node, len(records))
    nodeList[0] = rootNode
    nodeCount := 1

    for i := 1; i < len(records); i++ { 
        current:= &records[i]
        if current.ID != i || current.Parent >= current.ID {
            return nil, errors.New("Invalid record.")
        }

        childNode := CreateNode(current.ID)
        nodeList[i] = childNode
        nodeCount++

        parentNode := nodeList[current.Parent]
        if parentNode == nil {
            return nil, errors.New("Unexpected error")
        }
        InsertChild(parentNode, childNode)
    }

    return rootNode, nil
}

func IsRootRecord(record *Record) bool {
    return record.ID == 0 && record.Parent == 0
}

func CreateNode(ID int) *Node {
    node := new(Node)
    node.ID = ID
    return node
}

func InsertChild(parent, child *Node) {
    if parent.Children == nil {
        parent.Children = make([]*Node, 0)
    }
    parent.Children = append(parent.Children, child)
}
