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

func (mgr *NodeMgr) onLoad() {
	mgr.MapNodes = make(map[string]*Node)

	for _, v := range mgr.Nodes {
		mgr.addNodeInMap(v)

		v.Each(mgr, func(cn *Node) {
			mgr.addNodeInMap(cn)
		})
	}

	mgr.Nodes = nil
}

func (mgr *NodeMgr) addNodeInMap(node *Node) {
	mgr.MapNodes[node.CodeID] = node
}

func (mgr *NodeMgr) Merge(mgr1 *NodeMgr) {
	for _, v := range mgr1.MapNodes {
		mgr.addNodeInMap(v)
	}
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
