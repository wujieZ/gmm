package gmm

import "strings"

type node struct {
	pattern  string  // 待匹配的路由， 例如 /a/:name
	part     string  // 路由中的一部分， 例如 :name
	children []*node // 子节点， 例如[doc, introl]
	isWild   bool    // 是否精确匹配, part 含有 : 活 * 时为true
}

// 第一个匹配成功的节点， 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点， 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 节点插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 节点查询
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.pattern, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
