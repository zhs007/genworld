package genworld

type Node struct {
	Name     string   `yaml:"name"`
	CodeID   string   `yaml:"codeID"`
	TagID    []string `yaml:"tagID"` // 这里属于标记，一个节点，可以有多个标记，至于标记的具体用途，可以自由定义
	Info     string   `yaml:"info"`
	ParentID []string `yaml:"parentID"` // 因为这里可能是多个父节点的，所以下面children也应该用codeid来标识
	Children []string `yaml:"children"`
}

func (node *Node) Each(mgr *NodeMgr, oneach func(*Node)) {
	for _, c := range node.Children {
		cn := mgr.MapNodes[c]
		if cn != nil {
			cn.Each(mgr, oneach)
		}
	}
}
