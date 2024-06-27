package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &Node{Value: value}
		return
	}

	bst.Root.Insert(value)
}

func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
			return
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
			return
		} else {
			n.Right.Insert(value)
		}
	}
}

func (bst *BST) Search(value int) *Node {
	if bst.Root == nil {
		return nil
	}
	return bst.Root.Search(value)
}

func (n *Node) Search(value int) *Node {
	if n == nil {
		return nil
	}
	if n.Value == value {
		return n
	}
	if n.Value > value {
		return n.Left.Search(value)
	}
	return n.Right.Search(value)
}

func (n *Node) findMin() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *BST) Delete(value int) {
	if bst.Root == nil {
		fmt.Println("BST is empty")
		return
	}
	bst.Root.Delete(value)

}

func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	if n.Value > value {
		n.Left = n.Left.Delete(value)
	} else if n.Value < value {
		n.Right = n.Right.Delete(value)
	} else {

		if (n.Left == nil) && (n.Right == nil) {
			return nil
		} else if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		minRight := n.Right.findMin()
		n.Value = minRight.Value
		n.Right = n.Right.Delete(minRight.Value)
	}
	return n
}

func (bst *BST) Visualize() {
	lines := buildTreeString(bst.Root, 0, true, make([]string, 0))
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Вспомогательная функция для построения строк представления дерева.
func buildTreeString(n *Node, indent int, last bool, result []string) []string {
	if n == nil {
		return result
	}

	// Формируем отступ.
	prefix := strings.Repeat("    ", indent)
	if indent > 0 {
		if last {
			prefix = prefix[:len(prefix)-4] + "└───"
		} else {
			prefix = prefix[:len(prefix)-4] + "├───"
		}
	}

	// Добавляем строку для текущего узла.
	line := fmt.Sprintf("%s%d", prefix, n.Value)
	result = append(result, line)

	// Рекурсивно обрабатываем левое и правое поддеревья.
	if n.Left != nil || n.Right != nil {
		if n.Left != nil {
			result = buildTreeString(n.Left, indent+1, n.Right == nil, result)
		} else {
			result = buildTreeString(nil, indent+1, false, result)
		}
		if n.Right != nil {
			result = buildTreeString(n.Right, indent+1, true, result)
		} else {
			result = buildTreeString(nil, indent+1, true, result)
		}
	}
	return result
}

func (bst *BST) PreOrder() {
	PreOrder(bst.Root)
	fmt.Println()
}

func PreOrder(n *Node) {
	if n == nil {
		return
	}
	result := n.Value

	PreOrder(n.Left)
	fmt.Print(result, " ")
	PreOrder(n.Right)
}

func main() {
	tree := &BST{}

	// Вставляем значения
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(40)
	tree.Insert(35)
	tree.Insert(45)
	tree.Insert(55)
	tree.Insert(52)
	tree.Insert(52)
	tree.Insert(60)

	tree.Visualize()

	tree.PreOrder()
}
