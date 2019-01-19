package link

type LinkNode struct {
	Data         interface{}
	PreLinkNode  *LinkNode
	NextLinkNode *LinkNode
}

func (linkNode *LinkNode) Append(item interface{}) {
	index := linkNode
	for index.NextLinkNode != nil {
		index = index.NextLinkNode
	}
	node := &LinkNode{
		Data:        item,
		PreLinkNode: index,
	}
	index.NextLinkNode = node
}

func (linkNode *LinkNode) Insert(pos int, item interface{}) {
	node := linkNode
	for index := 0; index < pos; index++ {
		node = node.NextLinkNode
	}
	insertLinkNode := &LinkNode{
		Data:         item,
		PreLinkNode:  node,
		NextLinkNode: node.NextLinkNode,
	}
	node.NextLinkNode = insertLinkNode

}

func (linkNode *LinkNode) Del(pos int) *LinkNode {
	node := linkNode
	for index := 0; index < pos && node != nil; index++ {
		node = node.NextLinkNode
	}
	delLinkNode := &LinkNode{
		Data:         node.Data,
		PreLinkNode:  node.PreLinkNode,
		NextLinkNode: node.NextLinkNode,
	}
	node.PreLinkNode.NextLinkNode = node.NextLinkNode
	node.NextLinkNode.PreLinkNode = node.PreLinkNode
	return delLinkNode
}

func (linkNode *LinkNode) Reverse() {
	addr := &linkNode
	node := linkNode
	var temp *LinkNode
	for node.NextLinkNode != nil {
		temp = node.PreLinkNode
		node.PreLinkNode = node.NextLinkNode
		node.NextLinkNode = temp
		node = node.PreLinkNode
	}
	if temp != nil {
		*addr = node
	}
}

func (linkNode *LinkNode) SingleReverse() {
	if linkNode.NextLinkNode == nil {
		return
	}
	node := linkNode
	for node.NextLinkNode != nil {
		temp := node
		node = node.NextLinkNode
		flag := node.NextLinkNode
		node.NextLinkNode = temp
		node = flag
	}
	linkNode = node
}
