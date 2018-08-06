package tree

import (
    "errors"
)

const ROOT_ID = 0

type Record struct {
    ID, Parent int
}

type Node struct {
    ID       int
    Children []*Node
}

type NodeMap map[int]*Node

func Build(records []Record) (*Node, error) {
    if len(records) == 0 {
        return nil, nil
    }

    nodes := make(NodeMap)
    nodeCount := 0
    for i := 0; i < len(records); i++ { 
        currentRecord := &records[i]

        if !IsValidRecord(records, currentRecord) {
            return nil, errors.New("Invalid record.")
        }

        node, isNew := CreateOrGetNode(nodes, currentRecord.ID);
        if isNew {
            nodeCount++
        }

        if IsRootRecord(currentRecord) {
            continue
        }

        parentNode, isNew := CreateOrGetNode(nodes, currentRecord.Parent);
        if isNew {
            nodeCount++
        }

        if err := CheckParentChildNodesAreValid(parentNode, node);
        err != nil {
            return nil, err
        }

        if node != parentNode {
            InsertChildInOrder(parentNode, node)
        }
    }


    if len(records) != nodeCount {
        return nil, errors.New("Number of records doesn't match nodes created.")
    }

    if root := GetRootNode(nodes); root != nil {
        return root, nil
    }
    return nil, errors.New("Root node doesn't exist.")
}

func IsValidRecord(records []Record, record *Record) bool {
    return IsValidID(len(records), record.ID) &&
    IsValidID(len(records), record.Parent) &&
    (record.Parent < record.ID || IsRootRecord(record))
}

func IsValidID(LIMIT, ID int) bool {
    return 0 <= ID && ID < LIMIT
}

func IsRootRecord(record *Record) bool {
    return record.ID == ROOT_ID && record.Parent == ROOT_ID
}

func CreateOrGetNode(nodes NodeMap, ID int) (*Node, bool) {
    node := nodes[ID]
    if node != nil {
        return node, false
    }

    node = new(Node)
    node.ID = ID
    nodes[ID] = node
    return node, true;
}

func CheckParentChildNodesAreValid(parent, child *Node) (error) {
    if child == parent {
        return errors.New("Parent and child are the same node.")
    }

    if child.ID < parent.ID {
        return errors.New("Parent ID is greater than child ID.")
    }

    if child.ID == parent.ID {
        return errors.New("Parent and Child IDs are the same.")
    }
    return nil
}

func IsRootNode(node *Node) bool {
    return node.ID == 0
}

func GetRootNode(nodes NodeMap) *Node {
    return nodes[0]
}

func InsertChildInOrder(parent, child *Node) {
    if parent.Children == nil {
        parent.Children = make([]*Node, 0)
    }
    insertIndex := GetInsertIndexForChild(parent, child)

    if insertIndex == len(parent.Children) {
        parent.Children = append(parent.Children, child)
    } else {
        parent.Children = append(parent.Children, nil)
        copy(parent.Children[insertIndex+1:], parent.Children[insertIndex:])
        parent.Children[insertIndex] = child
    }
}

func GetInsertIndexForChild(parent, child *Node) int {
    insertIndex := -1
    for j := 0; j < len(parent.Children); j++ {
        parentChild := parent.Children[j]
        if parentChild.ID > child.ID {
            insertIndex = j
            break
        }
    }

    if insertIndex == -1 {
        insertIndex = len(parent.Children)
    }
    return insertIndex
}
