package main

import "fmt"

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	for i := len(nodes) / 2 - 1;  i >= 0; i--{
		down(nodes, i, len(nodes))
	}
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	child := 2 * i + 1
	if child >= n{
		return
	}
	temp := nodes[i].Value
	for{
		if child + 1 < n && nodes[child].Value > nodes[child + 1].Value{
			child += 1
		}
		if temp <= nodes[child].Value{
			break
		}
		nodes[i].Value = nodes[child].Value
		i = child
		child  = 2 * i + 1
		if child >= n{
			break
		}
	}
	nodes[i].Value = temp
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	if j == 0{
		return
	}
	temp := nodes[j].Value
	parent := (j - 1) / 2
	for nodes[parent].Value > temp{
		nodes[j].Value = nodes[parent].Value
		j = parent
		if j == 0{
			break
		}
		parent = (j - 1) / 2
	} 
	nodes[j].Value = temp
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	minNode := nodes[0]
	nodes[0].Value = nodes[len(nodes) - 1].Value
	nodes = nodes[ : len(nodes) - 1]
	down(nodes, 0, len(nodes) )
	return minNode, nodes

}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes) - 1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	for i := 0; i < len(nodes); i++{
		if nodes[i].Value == node.Value{
			nodes[i] = nodes[len(nodes) - 1] //将最后一个结点替换掉该结点
			nodes = nodes[ : len(nodes) - 1] 
			up(nodes, i) //如果最后一个节点比原来的小，那么可能会需要up
			down(nodes, i, len(nodes)) //反之，可能会需要down
			i-- //i--是为了防止拿上来的最后结点的值也是要删除的，因此重新检查该位置
		}
	} 
	fmt.Println(nodes)
	return nodes
}

func main() {
	testNodes := []Node{Node{7}, Node{4}, Node{6}, Node{2}, Node{7}, Node{0}, Node{9}, Node{4}}
	Init(testNodes)
	
	fmt.Println(testNodes)

	testNodes = Push(Node{5}, testNodes)

	fmt.Println(testNodes)

	testNodes = Push(Node{1}, testNodes)
	
	fmt.Println(testNodes)
	
	testNodes = Remove(testNodes, Node{4})

	var minNode Node
	for len(testNodes) > 0{
		minNode, testNodes = Pop(testNodes)
		fmt.Printf("%d ", minNode.Value)
	}


}