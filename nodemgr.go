package genworld

import (
	"os"

	"github.com/zhs007/goutils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type NodeMgr struct {
	MapNodes map[string]*Node `yaml:"-"`
	Nodes    []*Node          `yaml:"nodes"`
}

func (mgr *NodeMgr) RebuildChildren() {
	for _, n := range mgr.MapNodes {
		n.Children = nil
	}

	for _, n := range mgr.MapNodes {
		for _, pid := range n.ParentID {
			pn := mgr.MapNodes[pid]
			if pn != nil {
				pn.Children = append(pn.Children, n)
			}
		}
	}
}

// onLoad - 这个接口只能在load以后调用
func (mgr *NodeMgr) onLoad() {
	mgr.MapNodes = make(map[string]*Node)

	for _, v := range mgr.Nodes {
		mgr.addNodeInMap(v)

		v.Each(func(cn *Node) {
			mgr.addNodeInMap(cn)
		})
	}

	// 全部加载完以后，再建立父子关系
	mgr.RebuildChildren()

	mgr.Nodes = nil
}

func (mgr *NodeMgr) addNodeInMap(node *Node) {
	mgr.MapNodes[node.CodeID] = node
}

func (mgr *NodeMgr) Merge(mgr1 *NodeMgr) {
	for _, v := range mgr1.MapNodes {
		mgr.addNodeInMap(v)
	}

	// 全部加载完以后，再建立父子关系
	mgr.RebuildChildren()
}

func LoadNodeMgr(fn string) (*NodeMgr, error) {
	data, err := os.ReadFile(fn)
	if err != nil {
		goutils.Error("LoadNodeMgr:ReadFile",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	mgr := &NodeMgr{}
	err = yaml.Unmarshal(data, mgr)
	if err != nil {
		goutils.Error("LoadNodeMgr:Unmarshal",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	mgr.onLoad()

	return mgr, nil
}
