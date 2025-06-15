package sfa

import (
	"github.com/y-yu/sfa-go/common"
	"maps"
)

type StateMap map[common.State]common.State

func (lhs StateMap) Equal(rhs StateMap) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, rv := range rhs {
		if lv, ok := lhs[k]; !ok || lv != rv {
			return false
		}
	}

	return true
}

type StateStateMap map[common.State]StateMap

func (lhs StateStateMap) FindState(target StateMap) (common.State, bool) {
	for k, v := range maps.All(lhs) {
		if v.Equal(target) {
			return k, true
		}
	}

	return common.NewState(-1), false
}
