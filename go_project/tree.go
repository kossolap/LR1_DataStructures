package main

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height int
}

func newNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

func fixHeight(node *Node) {
	hL := getHeight(node.Left)
	hR := getHeight(node.Right)
	if hL > hR {
		node.Height = hL + 1
	} else {
		node.Height = hR + 1
	}
}

func rotateRight(node *Node) *Node {
	tmp := node.Left
	node.Left = tmp.Right
	tmp.Right = node
	fixHeight(node)
	fixHeight(tmp)
	return tmp
}

func rotateLeft(node *Node) *Node {
	tmp := node.Right
	node.Right = tmp.Left
	tmp.Left = node
	fixHeight(node)
	fixHeight(tmp)
	return tmp
}

func TInsert(node *Node, key int) *Node {
	if node == nil {
		return newNode(key)
	}
	if key < node.Key {
		node.Left = TInsert(node.Left, key)
	} else if key > node.Key {
		node.Right = TInsert(node.Right, key)
	} else {
		return node
	}
	fixHeight(node)
	bf := getBalance(node)
	if bf == 2 {
		if getBalance(node.Left) >= 0 {
			return rotateRight(node)
		}
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}
	if bf == -2 {
		if getBalance(node.Right) <= 0 {
			return rotateLeft(node)
		}
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}
	return node
}

func TFind(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if key == node.Key {
		return true
	}
	if key < node.Key {
		return TFind(node.Left, key)
	}
	return TFind(node.Right, key)
}

func findMin(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node
	}
	return findMin(node.Left)
}

func TRemove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = TRemove(node.Left, key)
	} else if key > node.Key {
		node.Right = TRemove(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		min := findMin(node.Right)
		node.Key = min.Key
		node.Right = TRemove(node.Right, min.Key)
	}
	fixHeight(node)
	bf := getBalance(node)
	if bf == 2 {
		if getBalance(node.Left) >= 0 {
			return rotateRight(node)
		}
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}
	if bf == -2 {
		if getBalance(node.Right) <= 0 {
			return rotateLeft(node)
		}
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}
	return node
}

func Inorder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	Inorder(node.Left, result)
	*result = append(*result, node.Key)
	Inorder(node.Right, result)
}

func TClear(node *Node) {
	_ = node
}
