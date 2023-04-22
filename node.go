package genworld

type Node struct {
	Name     string   `yaml:"name"`
	CodeID   string   `yaml:"codeID"`
	Info     string   `yaml:"info"`
	ParentID []string `yaml:"parentID"`
	Children []*Node  `yaml:"children"`
}

func (node *Node) Each(oneach func(*Node)) {
	for _, c := range node.Children {
		oneach(c)

		c.Each(oneach)
	}
}
